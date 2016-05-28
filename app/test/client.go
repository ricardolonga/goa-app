package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/ricardolonga/goa-app/app"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// ShowClientNotFound test setup
func ShowClientNotFound(t *testing.T, ctrl app.ClientController, clientID int) {
	ShowClientNotFoundCtx(t, context.Background(), ctrl, clientID)
}

// ShowClientNotFoundCtx test setup
func ShowClientNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.ClientController, clientID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/clients/%v", clientID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["clientId"] = []string{fmt.Sprintf("%v", clientID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ClientTest"), rw, req, prms)
	showCtx, err := app.NewShowClientContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// ShowClientOK test setup
func ShowClientOK(t *testing.T, ctrl app.ClientController, clientID int) *app.GoaExampleClient {
	return ShowClientOKCtx(t, context.Background(), ctrl, clientID)
}

// ShowClientOKCtx test setup
func ShowClientOKCtx(t *testing.T, ctx context.Context, ctrl app.ClientController, clientID int) *app.GoaExampleClient {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/clients/%v", clientID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["clientId"] = []string{fmt.Sprintf("%v", clientID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "ClientTest"), rw, req, prms)
	showCtx, err := app.NewShowClientContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.GoaExampleClient)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.GoaExampleClient", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}
