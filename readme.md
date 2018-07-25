# api
--
    import "github.com/reddotpay/api"


## Usage

```go
var StatusMessage = map[int]string{

	100: "Continue",
	101: "Switching Protocols",
	102: "Processing",

	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-authoritative Information",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	207: "Multi-Status",
	208: "Already Reported",
	226: "IM Used",

	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Found",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	307: "Temporary Redirect",
	308: "Permanent Redirect",

	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Payload Too Large",
	414: "Request-URI Too Long",
	415: "Unsupported Media Type",
	416: "Requested Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a teapot",
	421: "Misdirected Request",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	444: "Connection Closed Without Response",
	451: "Unavailable For Legal Reasons",
	499: "Client Closed Request",

	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	508: "Loop Detected",
	510: "Not Extended",
	511: "Network Authentication Required",
	599: "Network Connect Timeout Error",
}
```
StatusMessage maps http satus code to the corresponding message
https://httpstatuses.com

#### func  NewHandler

```go
func NewHandler(h Handlers) func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
```
NewHandler returns a callback handler for (AWS) APIGateway

#### type Handlers

```go
type Handlers map[string]ResourceInterface
```

Handlers enumerates all requests being handled

#### type Request

```go
type Request events.APIGatewayProxyRequest
```

Request represents APIGateway proxy request

#### func (Request) GetResourceKey

```go
func (r Request) GetResourceKey() string
```
GetResourceKey specifies for the handler routing

#### type Resource

```go
type Resource struct{}
```

Resource represents a generic API resource

#### func (Resource) AllowedMethods

```go
func (r Resource) AllowedMethods() []string
```
AllowedMethods returns the list of allowed methods for this resource

#### func (Resource) Delete

```go
func (r Resource) Delete(req Request, w *Response) error
```
Delete executes a DELETE http action

#### func (Resource) Get

```go
func (r Resource) Get(req Request, w *Response) error
```
Get executes a GET http action

#### func (Resource) Patch

```go
func (r Resource) Patch(req Request, w *Response) error
```
Patch executes a PATCH http action

#### func (Resource) Post

```go
func (r Resource) Post(req Request, w *Response) error
```
Post executes a POST http action

#### func (Resource) Put

```go
func (r Resource) Put(req Request, w *Response) error
```
Put executes a PUT http action

#### type ResourceInterface

```go
type ResourceInterface interface {
	AllowedMethods() []string

	Get(Request, *Response) error

	Post(Request, *Response) error

	Patch(Request, *Response) error

	Put(Request, *Response) error

	Delete(Request, *Response) error
}
```

ResourceInterface represents the interface for a generic API resource

#### type Response

```go
type Response events.APIGatewayProxyResponse
```

Response represents APIGateway proxy response

#### func (*Response) Output

```go
func (r *Response) Output(v interface{}) error
```
Output sets the body to the default output type

#### func (*Response) Stat

```go
func (r *Response) Stat(i int) error
```
Stat sets the header and the message to a standard output status
