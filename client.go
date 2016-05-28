package main

import (
	"github.com/goadesign/goa"
	"github.com/ricardolonga/goa-app/app"
)

// ClientController implements the client resource.
type ClientController struct {
	*goa.Controller
}

// NewClientController creates a client controller.
func NewClientController(service *goa.Service) *ClientController {
	return &ClientController{Controller: service.NewController("ClientController")}
}

// Show runs the show action.
func (c *ClientController) Show(ctx *app.ShowClientContext) error {
	// TBD: implement
	res := &app.GoaExampleClient{}
	return ctx.OK(res)
}
