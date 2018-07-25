package api

import (
	"fmt"
	"net/http"
	"strings"
)

// ResourceInterface represents the interface for a generic API resource
type ResourceInterface interface {
	AllowedMethods() []string

	Get(Request, *Response) error

	Post(Request, *Response) error

	Patch(Request, *Response) error

	Put(Request, *Response) error

	Delete(Request, *Response) error
}

func triggerMethod(r ResourceInterface, req Request) (Response, error) {
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
				err = r.Get(req, &w)

			case http.MethodPost:
				err = r.Post(req, &w)

			case http.MethodPatch:
				err = r.Patch(req, &w)

			case http.MethodPut:
				err = r.Put(req, &w)

			case http.MethodDelete:
				err = r.Delete(req, &w)

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
func (r Resource) Get(req Request, w *Response) error {
	w.Stat(http.StatusMethodNotAllowed)
	return nil
}

// Post executes a POST http action
func (r Resource) Post(req Request, w *Response) error {
	w.Stat(http.StatusMethodNotAllowed)
	return nil
}

// Patch executes a PATCH http action
func (r Resource) Patch(req Request, w *Response) error {
	w.Stat(http.StatusMethodNotAllowed)
	return nil
}

// Put executes a PUT http action
func (r Resource) Put(req Request, w *Response) error {
	w.Stat(http.StatusMethodNotAllowed)
	return nil
}

// Delete executes a DELETE http action
func (r Resource) Delete(req Request, w *Response) error {
	w.Stat(http.StatusMethodNotAllowed)
	return nil
}
