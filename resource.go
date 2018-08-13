package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// ResourceInterface represents the interface for a generic API resource
type ResourceInterface interface {
	AllowedMethods() []string

	Get(context.Context, Request, *Response) error

	Post(context.Context, Request, *Response) error

	Patch(context.Context, Request, *Response) error

	Put(context.Context, Request, *Response) error

	Delete(context.Context, Request, *Response) error
}

func triggerMethod(ctx context.Context, r ResourceInterface, req Request) (Response, error) {
	var (
		err error
		w   = Response{}
	)

	if w.Headers == nil {
		w.Headers = map[string]string{}
	}

	w.Headers["Allow"] = strings.Join(r.AllowedMethods(), ", ")
	w.Headers["Accept"] = "application/json"

	for _, i := range r.AllowedMethods() {
		if i == req.HTTPMethod {
			switch i {
			case http.MethodGet:
				err = r.Get(ctx, req, &w)

			case http.MethodPost:
				err = r.Post(ctx, req, &w)

			case http.MethodPatch:
				err = r.Patch(ctx, req, &w)

			case http.MethodPut:
				err = r.Put(ctx, req, &w)

			case http.MethodDelete:
				err = r.Delete(ctx, req, &w)

			default:
				w.Stat(http.StatusMethodNotAllowed)
				err = fmt.Errorf("api.triggerMethod: explicitly rejected request `%s`:`%s`", req.Resource, req.HTTPMethod)
			}

			return w, err
		}
	}

	err = w.Stat(http.StatusMethodNotAllowed)
	err = fmt.Errorf("api.triggerMethod: not in `AllowedMethods`. rejected request `%s`:`%s`", req.Resource, req.HTTPMethod)
	return w, err
}

// Resource represents a generic API resource
type Resource struct{}

// AllowedMethods returns the list of allowed methods for this resource
func (r Resource) AllowedMethods() []string {
	return []string{}
}

// Get executes a GET http action
func (r Resource) Get(ctx context.Context, req Request, w *Response) error {
	w.Stat(http.StatusMethodNotAllowed)
	return nil
}

// Post executes a POST http action
func (r Resource) Post(ctx context.Context, req Request, w *Response) error {
	w.Stat(http.StatusMethodNotAllowed)
	return nil
}

// Patch executes a PATCH http action
func (r Resource) Patch(ctx context.Context, req Request, w *Response) error {
	w.Stat(http.StatusMethodNotAllowed)
	return nil
}

// Put executes a PUT http action
func (r Resource) Put(ctx context.Context, req Request, w *Response) error {
	w.Stat(http.StatusMethodNotAllowed)
	return nil
}

// Delete executes a DELETE http action
func (r Resource) Delete(ctx context.Context, req Request, w *Response) error {
	w.Stat(http.StatusMethodNotAllowed)
	return nil
}
