//************************************************************************//
// API "clients": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/ricardolonga/goa-app/design
// --out=$(GOPATH)/src/github.com/ricardolonga/goa-app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// GoaExampleClient media type.
//
// Identifier: application/vnd.goa.example.client+json
type GoaExampleClient struct {
	// API href for making requests on the client
	Href string `json:"href" xml:"href"`
	// Unique client Id
	ID int `json:"id" xml:"id"`
	// Name of client
	Name string `json:"name" xml:"name"`
}

// Validate validates the GoaExampleClient media type instance.
func (mt *GoaExampleClient) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}
