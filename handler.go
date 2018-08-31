package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Handlers enumerates all requests being handled
type Handlers map[string]ResourceInterface

// Ping is a hackable function for /ping:GET resource
var Ping = func(req Request) (Response, error) {
	var w Response

	w.Headers = map[string]string{}

	err := w.Stat(http.StatusOK)
	w.Headers["Allow"] = "GET"
	w.Headers["Accept"] = "application/json"

	return w, err
}

// NewHandler returns a callback handler for (AWS) APIGateway
func NewHandler(h Handlers) func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var err error

		r := Request(req)
		p := r.GetResourceKey()
		w := Response{
			StatusCode: http.StatusNotFound,
		}

		if r.Resource == "/ping" && r.HTTPMethod == http.MethodGet {
			w, err = Ping(r)
		} else if _, ok := h[p]; ok {
			w, err = triggerMethod(ctx, h[p], r)
		} else {
			w, err = defaultHandler(r)
			err = fmt.Errorf("api.NewHandler: undeclared handler `%s`", p)
		}

		return events.APIGatewayProxyResponse(w), err
	}
}

func defaultHandler(req Request) (Response, error) {
	var w Response

	err := w.Stat(http.StatusNotFound)

	return w, err
}
