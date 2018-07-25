package api

import "github.com/aws/aws-lambda-go/events"

// Request represents APIGateway proxy request
type Request events.APIGatewayProxyRequest

// GetResourceKey specifies for the handler routing
func (r Request) GetResourceKey() string {
	return r.Resource
}
