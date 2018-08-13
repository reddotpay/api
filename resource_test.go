package api_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/reddotpay/api"
	"github.com/stretchr/testify/assert"
)

type Sample struct {
	api.Resource
}

func (s Sample) AllowedMethods() []string {
	return []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	}
}

func testMethod(m string, t *testing.T) {
	r, _ := api.NewHandler(api.Handlers{"/sample/{sampleId}": (Sample{})})(
		context.TODO(),
		events.APIGatewayProxyRequest{
			Resource:       "/sample/{sampleId}",
			PathParameters: map[string]string{"sampleId": "1"},
			HTTPMethod:     m,
		},
	)

	assert := assert.New(t)

	assert.Equal(405, r.StatusCode)
	assert.Equal(`{"message":"Method Not Allowed"}`, r.Body)
	assert.Equal("GET, POST, PUT, PATCH, DELETE", r.Headers["Allow"])
}

func TestResource_Get(t *testing.T) {
	testMethod(http.MethodGet, t)
}

func TestResource_Put(t *testing.T) {
	testMethod(http.MethodPut, t)
}

func TestResource_Patch(t *testing.T) {
	testMethod(http.MethodPatch, t)
}

func TestResource_Post(t *testing.T) {
	testMethod(http.MethodPost, t)
}

func TestResource_Delete(t *testing.T) {
	testMethod(http.MethodDelete, t)
}

func TestResource_AllowMethods(t *testing.T) {
	type Foo struct {
		api.Resource
	}

	r, err := api.NewHandler(api.Handlers{"/foo/{fooId}": (Foo{})})(
		context.TODO(),
		events.APIGatewayProxyRequest{
			Resource:       "/foo/{fooId}",
			PathParameters: map[string]string{"fooId": "1"},
			HTTPMethod:     http.MethodGet,
		},
	)

	assert := assert.New(t)

	assert.EqualError(err, "api.triggerMethod: not in `AllowedMethods`. rejected request `/foo/{fooId}`:`GET`")
	assert.Equal(405, r.StatusCode)
	assert.Equal(`{"message":"Method Not Allowed"}`, r.Body)
	assert.Equal("", r.Headers["Allow"])
}
