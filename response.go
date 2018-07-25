package api

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// Response represents APIGateway proxy response
type Response events.APIGatewayProxyResponse

// Output sets the body to the default output type
func (r *Response) Output(v interface{}) error {
	if r.Headers == nil {
		r.Headers = map[string]string{}
	}

	r.Headers["Content-Type"] = "application/json"

	b, err := json.Marshal(v)

	r.Body = string(b)

	return err
}

// Stat sets the header and the message to a standard output status
func (r *Response) Stat(i int) error {
	r.StatusCode = i

	return r.Output(map[string]string{
		"message": StatusMessage[i],
	})
}
