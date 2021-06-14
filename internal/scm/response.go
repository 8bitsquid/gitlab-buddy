package scm

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

type IResponse interface {
	GetBody() string
	GetStatusCode() int
	GetMetadata() interface{}
	parseBody()
}

type Response struct {
	*http.Response
	BodyString string
	Metadata   interface{}
}

func NewResponse(resp *http.Response) IResponse {
	r := &Response{Response: resp}
	r.parseBody()
	return r
}

// TODO: find better way - minor redundancy
func NewResponseWithMetadata(resp *http.Response, metadata interface{}) IResponse {
	r := &Response{Response: resp, Metadata: metadata}
	r.parseBody()
	return r
}

func (r *Response) GetBody() string {
	return r.BodyString
}

func (r *Response) GetStatusCode() int {
	return r.StatusCode
}

// Intended to be extended by clients with custom response data
func (r *Response) GetMetadata() interface{} {
	return r.Metadata
}

func (r *Response) parseBody() {
	defer r.Response.Body.Close()
	bodyString, err := io.ReadAll(r.Response.Body)
	if err != nil {
		zap.S().Errorw("Unable to parse response body")
	}
	r.BodyString = string(bodyString)
}
