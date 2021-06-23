// Package solicitations provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package solicitations

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"github.com/jiangdi0924/spsdk/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("spsdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetSolicitationActionsForOrder request
	GetSolicitationActionsForOrder(ctx context.Context, amazonOrderId string, params *GetSolicitationActionsForOrderParams) (*http.Response, error)

	// CreateProductReviewAndSellerFeedbackSolicitation request
	CreateProductReviewAndSellerFeedbackSolicitation(ctx context.Context, amazonOrderId string, params *CreateProductReviewAndSellerFeedbackSolicitationParams) (*http.Response, error)
}

func (c *Client) GetSolicitationActionsForOrder(ctx context.Context, amazonOrderId string, params *GetSolicitationActionsForOrderParams) (*http.Response, error) {
	req, err := NewGetSolicitationActionsForOrderRequest(c.Endpoint, amazonOrderId, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) CreateProductReviewAndSellerFeedbackSolicitation(ctx context.Context, amazonOrderId string, params *CreateProductReviewAndSellerFeedbackSolicitationParams) (*http.Response, error) {
	req, err := NewCreateProductReviewAndSellerFeedbackSolicitationRequest(c.Endpoint, amazonOrderId, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewGetSolicitationActionsForOrderRequest generates requests for GetSolicitationActionsForOrder
func NewGetSolicitationActionsForOrderRequest(endpoint string, amazonOrderId string, params *GetSolicitationActionsForOrderParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "amazonOrderId", amazonOrderId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/solicitations/v1/orders/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "marketplaceIds", params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateProductReviewAndSellerFeedbackSolicitationRequest generates requests for CreateProductReviewAndSellerFeedbackSolicitation
func NewCreateProductReviewAndSellerFeedbackSolicitationRequest(endpoint string, amazonOrderId string, params *CreateProductReviewAndSellerFeedbackSolicitationParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "amazonOrderId", amazonOrderId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/solicitations/v1/orders/%s/solicitations/productReviewAndSellerFeedback", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "marketplaceIds", params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetSolicitationActionsForOrder request
	GetSolicitationActionsForOrderWithResponse(ctx context.Context, amazonOrderId string, params *GetSolicitationActionsForOrderParams) (*GetSolicitationActionsForOrderResp, error)

	// CreateProductReviewAndSellerFeedbackSolicitation request
	CreateProductReviewAndSellerFeedbackSolicitationWithResponse(ctx context.Context, amazonOrderId string, params *CreateProductReviewAndSellerFeedbackSolicitationParams) (*CreateProductReviewAndSellerFeedbackSolicitationResp, error)
}

type GetSolicitationActionsForOrderResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetSolicitationActionsForOrderResponse
	JSON400      *GetSolicitationActionsForOrderResponse
	JSON403      *GetSolicitationActionsForOrderResponse
	JSON404      *GetSolicitationActionsForOrderResponse
	JSON413      *GetSolicitationActionsForOrderResponse
	JSON415      *GetSolicitationActionsForOrderResponse
	JSON429      *GetSolicitationActionsForOrderResponse
	JSON500      *GetSolicitationActionsForOrderResponse
	JSON503      *GetSolicitationActionsForOrderResponse
}

// Status returns HTTPResponse.Status
func (r GetSolicitationActionsForOrderResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSolicitationActionsForOrderResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateProductReviewAndSellerFeedbackSolicitationResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *CreateProductReviewAndSellerFeedbackSolicitationResponse
	JSON400      *CreateProductReviewAndSellerFeedbackSolicitationResponse
	JSON403      *CreateProductReviewAndSellerFeedbackSolicitationResponse
	JSON404      *CreateProductReviewAndSellerFeedbackSolicitationResponse
	JSON413      *CreateProductReviewAndSellerFeedbackSolicitationResponse
	JSON415      *CreateProductReviewAndSellerFeedbackSolicitationResponse
	JSON429      *CreateProductReviewAndSellerFeedbackSolicitationResponse
	JSON500      *CreateProductReviewAndSellerFeedbackSolicitationResponse
	JSON503      *CreateProductReviewAndSellerFeedbackSolicitationResponse
}

// Status returns HTTPResponse.Status
func (r CreateProductReviewAndSellerFeedbackSolicitationResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateProductReviewAndSellerFeedbackSolicitationResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetSolicitationActionsForOrderWithResponse request returning *GetSolicitationActionsForOrderResponse
func (c *ClientWithResponses) GetSolicitationActionsForOrderWithResponse(ctx context.Context, amazonOrderId string, params *GetSolicitationActionsForOrderParams) (*GetSolicitationActionsForOrderResp, error) {
	rsp, err := c.GetSolicitationActionsForOrder(ctx, amazonOrderId, params)
	if err != nil {
		return nil, err
	}
	return ParseGetSolicitationActionsForOrderResp(rsp)
}

// CreateProductReviewAndSellerFeedbackSolicitationWithResponse request returning *CreateProductReviewAndSellerFeedbackSolicitationResponse
func (c *ClientWithResponses) CreateProductReviewAndSellerFeedbackSolicitationWithResponse(ctx context.Context, amazonOrderId string, params *CreateProductReviewAndSellerFeedbackSolicitationParams) (*CreateProductReviewAndSellerFeedbackSolicitationResp, error) {
	rsp, err := c.CreateProductReviewAndSellerFeedbackSolicitation(ctx, amazonOrderId, params)
	if err != nil {
		return nil, err
	}
	return ParseCreateProductReviewAndSellerFeedbackSolicitationResp(rsp)
}

// ParseGetSolicitationActionsForOrderResp parses an HTTP response from a GetSolicitationActionsForOrderWithResponse call
func ParseGetSolicitationActionsForOrderResp(rsp *http.Response) (*GetSolicitationActionsForOrderResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetSolicitationActionsForOrderResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetSolicitationActionsForOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest GetSolicitationActionsForOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest GetSolicitationActionsForOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest GetSolicitationActionsForOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest GetSolicitationActionsForOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest GetSolicitationActionsForOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest GetSolicitationActionsForOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest GetSolicitationActionsForOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest GetSolicitationActionsForOrderResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseCreateProductReviewAndSellerFeedbackSolicitationResp parses an HTTP response from a CreateProductReviewAndSellerFeedbackSolicitationWithResponse call
func ParseCreateProductReviewAndSellerFeedbackSolicitationResp(rsp *http.Response) (*CreateProductReviewAndSellerFeedbackSolicitationResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CreateProductReviewAndSellerFeedbackSolicitationResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest CreateProductReviewAndSellerFeedbackSolicitationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest CreateProductReviewAndSellerFeedbackSolicitationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest CreateProductReviewAndSellerFeedbackSolicitationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest CreateProductReviewAndSellerFeedbackSolicitationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest CreateProductReviewAndSellerFeedbackSolicitationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest CreateProductReviewAndSellerFeedbackSolicitationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest CreateProductReviewAndSellerFeedbackSolicitationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest CreateProductReviewAndSellerFeedbackSolicitationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest CreateProductReviewAndSellerFeedbackSolicitationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}
