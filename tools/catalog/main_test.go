package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/rules/models"
)

// TestRenderMatchesCommittedCatalog ensures the committed catalog.go is in
// the exact format produced by the generator.
func TestRenderMatchesCommittedCatalog(t *testing.T) {
	want, err := os.ReadFile(filepath.Join("..", "..", "rules", "models", "catalog.go"))
	if err != nil {
		t.Fatal(err)
	}

	got, err := render(&catalog{
		InstanceServerTypes: models.InstanceServerTypes,
		RdbNodeTypes:        models.RdbNodeTypes,
		RedisNodeTypes:      models.RedisNodeTypes,
		LBTypes:             models.LBTypes,
	})
	if err != nil {
		t.Fatal(err)
	}

	if string(got) != string(want) {
		t.Errorf("rendered catalog differs from rules/models/catalog.go:\n%s", got)
	}
}

func TestFetchCatalog(t *testing.T) {
	mux := http.NewServeMux()

	// Instance types span two pages in fr-par-1 and differ in nl-ams-1.
	mux.HandleFunc("/instance/v1/zones/fr-par-1/products/servers", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth-Token") != "secret" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		switch r.URL.Query().Get("page") {
		case "1":
			fmt.Fprint(w, `{"servers":{"DEV1-S":{},"PRO2-S":{}},"total_count":3}`)
		case "2":
			fmt.Fprint(w, `{"servers":{"GP1-XS":{}},"total_count":3}`)
		default:
			t.Errorf("unexpected page %q", r.URL.Query().Get("page"))
		}
	})
	mux.HandleFunc("/instance/v1/zones/nl-ams-1/products/servers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"servers":{"DEV1-S":{},"L4-1-24G":{}},"total_count":2}`)
	})
	mux.HandleFunc("/rdb/v1/regions/fr-par/node-types", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"node_types":[{"name":"db-dev-s"},{"name":"db-pro2-xxs"}],"total_count":2}`)
	})
	// Redis is only available in fr-par-1: nl-ams-1 answers 404 and is skipped.
	mux.HandleFunc("/redis/v1/zones/fr-par-1/node-types", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"node_types":[{"name":"RED1-micro"}],"total_count":1}`)
	})
	mux.HandleFunc("/lb/v1/zones/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"lb_types":[{"name":"LB-S"},{"name":"LB-GP-M"}],"total_count":2}`)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	c := &client{baseURL: server.URL, token: "secret", http: server.Client()}

	got, err := fetchCatalog(context.Background(), c, []string{"fr-par-1", "nl-ams-1"}, []string{"fr-par"})
	if err != nil {
		t.Fatal(err)
	}

	want := &catalog{
		InstanceServerTypes: []string{"DEV1-S", "GP1-XS", "L4-1-24G", "PRO2-S"},
		RdbNodeTypes:        []string{"db-dev-s", "db-pro2-xxs"},
		RedisNodeTypes:      []string{"RED1-micro"},
		LBTypes:             []string{"LB-GP-M", "LB-S"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestFetchCatalogFailsWhenNoLocalityAnswers(t *testing.T) {
	server := httptest.NewServer(http.NotFoundHandler())
	defer server.Close()

	c := &client{baseURL: server.URL, token: "secret", http: server.Client()}

	if _, err := fetchCatalog(context.Background(), c, []string{"fr-par-1"}, []string{"fr-par"}); err == nil {
		t.Fatal("expected an error")
	}
}

func TestFetchCatalogFailsOnUnauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	c := &client{baseURL: server.URL, token: "wrong", http: server.Client()}

	if _, err := fetchCatalog(context.Background(), c, []string{"fr-par-1"}, []string{"fr-par"}); err == nil {
		t.Fatal("expected an error")
	}
}
