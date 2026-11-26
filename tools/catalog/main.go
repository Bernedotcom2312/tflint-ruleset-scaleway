// Command catalog regenerates rules/models/catalog.go from the Scaleway API.
//
// Usage:
//
//	SCW_SECRET_KEY=... go run ./tools/catalog -output rules/models/catalog.go
//
// The commercial types of every zone (or region) are merged, so that a type
// available in a single locality is still considered valid.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"text/template"
	"time"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/rules/models"
)

const defaultAPIURL = "https://api.scaleway.com"

func main() {
	output := flag.String("output", "rules/models/catalog.go", "file to write")
	apiURL := flag.String("api-url", defaultAPIURL, "Scaleway API URL")
	timeout := flag.Duration("timeout", 2*time.Minute, "global timeout")
	flag.Parse()

	token := os.Getenv("SCW_SECRET_KEY")
	if token == "" {
		log.Fatal("SCW_SECRET_KEY must be set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	client := &client{
		baseURL: *apiURL,
		token:   token,
		http:    &http.Client{Timeout: 30 * time.Second},
	}

	catalog, err := fetchCatalog(ctx, client, models.Zones, models.Regions)
	if err != nil {
		log.Fatal(err)
	}

	src, err := render(catalog)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(*output, src, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %s", *output)
}

// catalog holds the sorted, de-duplicated lists written to catalog.go.
type catalog struct {
	InstanceServerTypes []string
	RdbNodeTypes        []string
	RedisNodeTypes      []string
	LBTypes             []string
}

func fetchCatalog(ctx context.Context, c *client, zones, regions []string) (*catalog, error) {
	instance, err := collect(ctx, c, zones, "/instance/v1/zones/%s/products/servers", "per_page", instanceServerNames)
	if err != nil {
		return nil, fmt.Errorf("instance types: %w", err)
	}
	rdb, err := collect(ctx, c, regions, "/rdb/v1/regions/%s/node-types", "page_size", namedItems("node_types"))
	if err != nil {
		return nil, fmt.Errorf("rdb node types: %w", err)
	}
	redis, err := collect(ctx, c, zones, "/redis/v1/zones/%s/node-types", "page_size", namedItems("node_types"))
	if err != nil {
		return nil, fmt.Errorf("redis node types: %w", err)
	}
	lb, err := collect(ctx, c, zones, "/lb/v1/zones/%s/lb-types", "page_size", namedItems("lb_types"))
	if err != nil {
		return nil, fmt.Errorf("lb types: %w", err)
	}

	return &catalog{
		InstanceServerTypes: instance,
		RdbNodeTypes:        rdb,
		RedisNodeTypes:      redis,
		LBTypes:             lb,
	}, nil
}

// pageParser extracts the names of a page and the total number of items.
type pageParser func(body []byte) (names []string, total int, err error)

// collect walks every page of path for each locality and returns the
// sorted union of names. A locality where the product is not available is
// skipped, but at least one locality must answer.
func collect(ctx context.Context, c *client, localities []string, path, pageSizeParam string, parse pageParser) ([]string, error) {
	const pageSize = 100

	names := map[string]struct{}{}
	answered := 0

	for _, locality := range localities {
		seen := 0
		for page := 1; ; page++ {
			body, err := c.get(ctx, fmt.Sprintf(path, locality), map[string]string{
				"page":        fmt.Sprint(page),
				pageSizeParam: fmt.Sprint(pageSize),
			})
			var unavailable *notAvailableError
			if errors.As(err, &unavailable) {
				log.Printf("skipping %s: %s", locality, err)
				break
			}
			if err != nil {
				return nil, err
			}

			pageNames, total, err := parse(body)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", locality, err)
			}
			if page == 1 {
				answered++
			}
			for _, name := range pageNames {
				names[name] = struct{}{}
			}

			seen += len(pageNames)
			if len(pageNames) == 0 || seen >= total {
				break
			}
		}
	}

	if answered == 0 {
		return nil, errors.New("no locality answered")
	}

	list := make([]string, 0, len(names))
	for name := range names {
		list = append(list, name)
	}
	sort.Strings(list)
	return list, nil
}

// instanceServerNames parses a ListServersTypes response, whose servers are
// keyed by commercial type.
func instanceServerNames(body []byte) ([]string, int, error) {
	var resp struct {
		Servers    map[string]json.RawMessage `json:"servers"`
		TotalCount int                        `json:"total_count"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, 0, err
	}
	names := make([]string, 0, len(resp.Servers))
	for name := range resp.Servers {
		names = append(names, name)
	}
	return names, resp.TotalCount, nil
}

// namedItems parses a list response whose items, under key, have a name.
func namedItems(key string) pageParser {
	return func(body []byte) ([]string, int, error) {
		var resp map[string]json.RawMessage
		if err := json.Unmarshal(body, &resp); err != nil {
			return nil, 0, err
		}

		var items []struct {
			Name string `json:"name"`
		}
		if raw, ok := resp[key]; ok {
			if err := json.Unmarshal(raw, &items); err != nil {
				return nil, 0, err
			}
		}

		var total int
		if raw, ok := resp["total_count"]; ok {
			if err := json.Unmarshal(raw, &total); err != nil {
				return nil, 0, err
			}
		}

		names := make([]string, 0, len(items))
		for _, item := range items {
			names = append(names, item.Name)
		}
		return names, total, nil
	}
}

// client is a minimal Scaleway API client, so that the ruleset does not
// depend on the Scaleway SDK.
type client struct {
	baseURL string
	token   string
	http    *http.Client
}

// notAvailableError is returned when a product is not available in a locality.
type notAvailableError struct {
	status int
	url    string
}

func (e *notAvailableError) Error() string {
	return fmt.Sprintf("GET %s: HTTP %d", e.url, e.status)
}

func (c *client) get(ctx context.Context, path string, query map[string]string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range query {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Set("X-Auth-Token", c.token)
	req.Header.Set("User-Agent", "tflint-ruleset-scaleway/catalog")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch {
	case resp.StatusCode == http.StatusOK:
		return body, nil
	case resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusBadRequest:
		return nil, &notAvailableError{status: resp.StatusCode, url: req.URL.Path}
	default:
		return nil, fmt.Errorf("GET %s: HTTP %d: %s", req.URL.Path, resp.StatusCode, bytes.TrimSpace(body))
	}
}

var catalogTemplate = template.Must(template.New("catalog").Parse(`// Code generated by tools/catalog. DO NOT EDIT.

package models

// InstanceServerTypes lists the commercial types of Instances (also used by Kubernetes pools).
var InstanceServerTypes = []string{
{{- range .InstanceServerTypes }}
	{{ printf "%q" . }},
{{- end }}
}

// RdbNodeTypes lists the node types of Managed Databases for PostgreSQL and MySQL.
var RdbNodeTypes = []string{
{{- range .RdbNodeTypes }}
	{{ printf "%q" . }},
{{- end }}
}

// RedisNodeTypes lists the node types of Managed Databases for Redis.
var RedisNodeTypes = []string{
{{- range .RedisNodeTypes }}
	{{ printf "%q" . }},
{{- end }}
}

// LBTypes lists the commercial types of Load Balancers.
var LBTypes = []string{
{{- range .LBTypes }}
	{{ printf "%q" . }},
{{- end }}
}
`))

func render(c *catalog) ([]byte, error) {
	var buf bytes.Buffer
	if err := catalogTemplate.Execute(&buf, c); err != nil {
		return nil, err
	}
	return format.Source(buf.Bytes())
}
