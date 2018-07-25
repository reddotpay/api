package api_test

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/reddotpay/api"
)

type Task struct {
	api.Resource
}

func (t Task) AllowedMethods() []string {
	return []string{http.MethodGet, http.MethodTrace}
}

func (t Task) Get(r api.Request, w *api.Response) error {
	w.StatusCode = 200
	return w.Output(map[string]string{"foo": "bar"})
}

func ExampleNewHandler() {
	var (
		task       = Task{}
		apiHandler = api.Handlers{
			"/task/{taskId}": task,
		}
	)

	// undeclared resource
	r, err := api.NewHandler(apiHandler)(
		events.APIGatewayProxyRequest{
			Resource:   "/undeclared",
			HTTPMethod: http.MethodGet,
		},
	)

	fmt.Printf("Response [%d] [allow: `%s`, accept: `%s`]: `%s`\n",
		r.StatusCode, r.Headers["Allow"], r.Headers["Accept"], r.Body)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	// if method was not set
	r, err = api.NewHandler(apiHandler)(
		events.APIGatewayProxyRequest{
			Resource:       "/task/{taskId}",
			PathParameters: map[string]string{"taskId": "1"},
			HTTPMethod:     http.MethodDelete,
		},
	)

	fmt.Printf("Response [%d] [allow: `%s`, accept: `%s`]: `%s`\n",
		r.StatusCode, r.Headers["Allow"], r.Headers["Accept"], r.Body)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	// if method was not allowed
	r, err = api.NewHandler(apiHandler)(
		events.APIGatewayProxyRequest{
			Resource:       "/task/{taskId}",
			PathParameters: map[string]string{"taskId": "1"},
			HTTPMethod:     http.MethodTrace,
		},
	)

	fmt.Printf("Response [%d] [allow: `%s`, accept: `%s`]: `%s`\n",
		r.StatusCode, r.Headers["Allow"], r.Headers["Accept"], r.Body)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	// happy scenario
	r, err = api.NewHandler(apiHandler)(
		events.APIGatewayProxyRequest{
			Resource:       "/task/{taskId}",
			PathParameters: map[string]string{"taskId": "1"},
			HTTPMethod:     http.MethodGet,
		},
	)

	fmt.Printf("Response [%d] [allow: `%s`, accept: `%s`]: `%s`\n",
		r.StatusCode, r.Headers["Allow"], r.Headers["Accept"], r.Body)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	// Output:
	// Response [404] [allow: ``, accept: ``]: `{"message":"Not Found"}`
	// error:  api.NewHandler: undeclared handler `/undeclared`
	// Response [405] [allow: `GET, TRACE`, accept: `application/json`]: `{"message":"Method Not Allowed"}`
	// error:  api.triggerMethod: not in `AllowedMethods`. rejected request `/task/{taskId}`:`DELETE`
	// Response [405] [allow: `GET, TRACE`, accept: `application/json`]: `{"message":"Method Not Allowed"}`
	// error:  api.triggerMethod: explicitly rejected request `/task/{taskId}`:`TRACE`
	// Response [200] [allow: `GET, TRACE`, accept: `application/json`]: `{"foo":"bar"}`
}
