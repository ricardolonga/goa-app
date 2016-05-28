//************************************************************************//
// API "clients": Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/ricardolonga/goa-app/design
// --out=$(GOPATH)/src/github.com/ricardolonga/goa-app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// ClientHref returns the resource href.
func ClientHref(clientID interface{}) string {
	return fmt.Sprintf("/clients/%v", clientID)
}
