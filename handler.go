package api

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Handlers enumerates all requests being handled
type Handlers map[string]ResourceInterface

// NewHandler returns a callback handler for (AWS) APIGateway
func NewHandler(h Handlers) func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var err error

		r := Request(req)
		p := r.GetResourceKey()
		w := Response{
			StatusCode: http.StatusNotFound,
		}

		if _, ok := h[p]; ok {
			w, err = triggerMethod(h[p], r)
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
