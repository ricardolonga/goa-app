//************************************************************************//
// API "clients": Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/ricardolonga/goa-app/design
// --out=$(GOPATH)/src/github.com/ricardolonga/goa-app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ClientController is the controller interface for the Client actions.
type ClientController interface {
	goa.Muxer
	Show(*ShowClientContext) error
}

// MountClientController "mounts" a Client resource controller on the given service.
func MountClientController(service *goa.Service, ctrl ClientController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowClientContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/clients/:clientId", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Client", "action", "Show", "route", "GET /clients/:clientId")
}
