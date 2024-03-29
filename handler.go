package handler

import (
	"net/http"

	"github.com/machinebox/graphql"
)

// Response of function call
type Response struct {

	// Body the body will be written back
	Body []byte

	// StatusCode needs to be populated with value such as http.StatusOK
	StatusCode int

	// Header is optional and contains any additional headers the function response should set
	Header http.Header
}

// Request of function call
type Request struct {
	Body        []byte
	Header      http.Header
	QueryString string
	Method      string
}

// Context of the function call
type Context struct {
	GraphqlClient *graphql.Client
	GetSecret     func(secretName string) (secretBytes []byte, err error)
	GetEnv        func(key string) (string, bool)
}

// FunctionHandler used for a serverless Go method invocation
type FunctionHandler interface {
	Handle(req Request, context Context) (Response, error)
}

func init() {

}
