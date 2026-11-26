package models

// Zones lists the Scaleway Availability Zones.
//
// Kept in sync with scw.AllZones from github.com/scaleway/scaleway-sdk-go.
var Zones = []string{
	"fr-par-1",
	"fr-par-2",
	"fr-par-3",
	"nl-ams-1",
	"nl-ams-2",
	"nl-ams-3",
	"pl-waw-1",
	"pl-waw-2",
	"pl-waw-3",
	"it-mil-1",
	"it-mil-2",
	"it-mil-3",
}

// Regions lists the Scaleway regions.
//
// Kept in sync with scw.AllRegions from github.com/scaleway/scaleway-sdk-go.
var Regions = []string{
	"fr-par",
	"nl-ams",
	"pl-waw",
	"it-mil",
}
