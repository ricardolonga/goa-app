//************************************************************************//
// API "clients": Application Contexts
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
	"strconv"
)

// ShowClientContext provides the client show action context.
type ShowClientContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	ClientID int
}

// NewShowClientContext parses the incoming request URL and body, performs validations and creates the
// context used by the client controller show action.
func NewShowClientContext(ctx context.Context, service *goa.Service) (*ShowClientContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowClientContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramClientID := req.Params["clientId"]
	if len(paramClientID) > 0 {
		rawClientID := paramClientID[0]
		if clientID, err2 := strconv.Atoi(rawClientID); err2 == nil {
			rctx.ClientID = clientID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("clientId", rawClientID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowClientContext) OK(r *GoaExampleClient) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.client")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowClientContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
