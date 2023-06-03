package hermes

import (
	"context"
	"fmt"
	"time"

	"github.com/imroc/req/v3"
	"github.com/namsral/flag"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"jetshop/integration/common"
	"jetshop/integration/hermes/response"
)

const (
	listThreadURI      = "/api/im/threads"
	getDetailThreadURI = "/api/im/threads/detail"
	listMessageURI     = "/api/im/messages"
)

var (
	endpoint     = ""
	ClientId     = ""
	ClientSecret = ""
)

func init() {
	flag.StringVar(&endpoint, "hermes_endpoint", "localhost:4000", "hermes endpoint")
	flag.StringVar(&ClientId, "hermes_client_id", "client_id", "hermes client id")
	flag.StringVar(&ClientSecret, "hermes_client_secret", "client_secret", "hermes client secret")
	flag.Parse()
}

type Client interface {
	SetTracer(tracer trace.Tracer)
	ListThread(ctx context.Context, sellerId string, startTime int64, pageSize int) (*response.ListThread, error)
	GetThread(ctx context.Context, sellerId, threadId string) (*response.Thread, error)
	ListMessage(ctx context.Context, sellerId, threadId string, startTime int64, pageSize int) (*response.ListMessage, error)
}

type client struct {
	*req.Client
	clientId     string
	clientSecret string
	helper       *hermesHelper
}

func NewClient() Client {
	c := req.C().SetBaseURL(endpoint)

	if common.ShowLog {
		c = c.DevMode()
	}

	return &client{
		Client:       c,
		clientId:     ClientId,
		clientSecret: ClientSecret,
		helper:       NewHermesHelper(ClientId, ClientSecret),
	}
}

func (c *client) addSign(path string, data map[string]string) error {
	t := time.Now()

	data["client_id"] = c.clientId
	data["timestamp"] = fmt.Sprintf("%d", t.UnixMilli())
	data["sign_method"] = "sha256"
	data["sign"] = c.helper.Sign(data, path, c.clientSecret)

	return nil
}

type apiNameType int

const apiNameKey apiNameType = iota

func (c *client) SetTracer(tracer trace.Tracer) {
	c.WrapRoundTripFunc(func(rt req.RoundTripper) req.RoundTripFunc {
		return func(req *req.Request) (resp *req.Response, err error) {
			ctx := req.Context()
			apiName, ok := ctx.Value(apiNameKey).(string)
			if !ok {
				apiName = req.URL.Path
			}
			_, span := tracer.Start(req.Context(), apiName)
			defer span.End()
			span.SetAttributes(
				attribute.String("http.url", req.URL.String()),
				attribute.String("http.method", req.Method),
				attribute.String("http.req.header", req.HeaderToString()),
			)
			if len(req.Body) > 0 {
				span.SetAttributes(
					attribute.String("http.req.body", string(req.Body)),
				)
			}
			resp, err = rt.RoundTrip(req)
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			if resp.Response != nil {
				span.SetAttributes(
					attribute.Int("http.status_code", resp.StatusCode),
					attribute.String("http.resp.header", resp.HeaderToString()),
					attribute.String("http.resp.body", resp.String()),
				)
			}
			return
		}
	})
}

func withAPIName(ctx context.Context, name string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, apiNameKey, name)
}

func (c *client) ListThread(ctx context.Context, sellerId string, startTime int64, pageSize int) (*response.ListThread, error) {
	data := map[string]string{
		"seller_id":  sellerId,
		"start_time": fmt.Sprintf("%d", startTime),
		"page_size":  fmt.Sprintf("%d", pageSize),
	}

	if err := c.addSign(listThreadURI, data); err != nil {
		return nil, err
	}

	var res response.ListThread

	if err := c.Get(listThreadURI).
		SetContext(ctx).
		SetQueryParams(data).
		Do(withAPIName(ctx, "client.list_thread")).
		Into(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *client) GetThread(ctx context.Context, sellerId, threadId string) (*response.Thread, error) {
	data := map[string]string{
		"seller_id": sellerId,
		"thread_id": threadId,
	}

	if err := c.addSign(getDetailThreadURI, data); err != nil {
		return nil, err
	}

	var res response.Thread

	if err := c.Get(getDetailThreadURI).
		SetContext(ctx).
		SetQueryParams(data).
		Do(withAPIName(ctx, "client.detail_tread")).
		Into(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *client) ListMessage(ctx context.Context, sellerId, threadId string, startTime int64, pageSize int) (*response.ListMessage, error) {
	data := map[string]string{
		"seller_id":  sellerId,
		"thread_id":  threadId,
		"start_time": fmt.Sprintf("%d", startTime),
		"page_size":  fmt.Sprintf("%d", pageSize),
	}

	if err := c.addSign(listMessageURI, data); err != nil {
		return nil, err
	}

	var res response.ListMessage

	if err := c.Get(listMessageURI).
		SetContext(ctx).
		SetQueryParams(data).
		Do(withAPIName(ctx, "client.list_message")).
		Into(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
