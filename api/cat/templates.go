// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// TemplatesOption is a non-required Templates option that gets applied to an HTTP request.
type TemplatesOption func(r *transport.Request)

// WithTemplatesName - a pattern that returned template names must match.
func WithTemplatesName(name string) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesFormat - a short version of the Accept header, e.g. json, yaml.
func WithTemplatesFormat(format string) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesH - comma-separated list of column names to display.
func WithTemplatesH(h []string) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesHelp - return help information.
func WithTemplatesHelp(help bool) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesLocal - return local information, do not retrieve the state from master node (default: false).
func WithTemplatesLocal(local bool) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesMasterTimeout - explicit operation timeout for connection to master node.
func WithTemplatesMasterTimeout(masterTimeout time.Duration) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesS - comma-separated list of column names or column aliases to sort by.
func WithTemplatesS(s []string) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesV - verbose mode. Display column headers.
func WithTemplatesV(v bool) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesErrorTrace - include the stack trace of returned errors.
func WithTemplatesErrorTrace(errorTrace bool) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesFilterPath - a comma-separated list of filters used to reduce the respone.
func WithTemplatesFilterPath(filterPath []string) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesHuman - return human readable values for statistics.
func WithTemplatesHuman(human bool) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesIgnore - ignores the specified HTTP status codes.
func WithTemplatesIgnore(ignore []int) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesPretty - pretty format the returned JSON response.
func WithTemplatesPretty(pretty bool) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// WithTemplatesSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithTemplatesSourceParam(sourceParam string) TemplatesOption {
	return func(r *transport.Request) {
	}
}

// Templates - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-templates.html for more info.
//
// options: optional parameters.
func (c *Cat) Templates(options ...TemplatesOption) (*TemplatesResponse, error) {
	req := c.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := c.transport.Do(req)
	return &TemplatesResponse{resp}, err
}

// TemplatesResponse is the response for Templates.
type TemplatesResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *TemplatesResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}