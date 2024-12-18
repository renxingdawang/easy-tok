// Code generated by Kitex v0.11.3. DO NOT EDIT.

package publishsrv

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	publish "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/publish"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"PublishAction": kitex.NewMethodInfo(
		publishActionHandler,
		newPublishSrvPublishActionArgs,
		newPublishSrvPublishActionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"PublishList": kitex.NewMethodInfo(
		publishListHandler,
		newPublishSrvPublishListArgs,
		newPublishSrvPublishListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeletePublish": kitex.NewMethodInfo(
		deletePublishHandler,
		newPublishSrvDeletePublishArgs,
		newPublishSrvDeletePublishResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	publishSrvServiceInfo                = NewServiceInfo()
	publishSrvServiceInfoForClient       = NewServiceInfoForClient()
	publishSrvServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return publishSrvServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return publishSrvServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return publishSrvServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PublishSrv"
	handlerType := (*publish.PublishSrv)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "publish",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func publishActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*publish.PublishSrvPublishActionArgs)
	realResult := result.(*publish.PublishSrvPublishActionResult)
	success, err := handler.(publish.PublishSrv).PublishAction(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPublishSrvPublishActionArgs() interface{} {
	return publish.NewPublishSrvPublishActionArgs()
}

func newPublishSrvPublishActionResult() interface{} {
	return publish.NewPublishSrvPublishActionResult()
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*publish.PublishSrvPublishListArgs)
	realResult := result.(*publish.PublishSrvPublishListResult)
	success, err := handler.(publish.PublishSrv).PublishList(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPublishSrvPublishListArgs() interface{} {
	return publish.NewPublishSrvPublishListArgs()
}

func newPublishSrvPublishListResult() interface{} {
	return publish.NewPublishSrvPublishListResult()
}

func deletePublishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*publish.PublishSrvDeletePublishArgs)
	realResult := result.(*publish.PublishSrvDeletePublishResult)
	success, err := handler.(publish.PublishSrv).DeletePublish(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPublishSrvDeletePublishArgs() interface{} {
	return publish.NewPublishSrvDeletePublishArgs()
}

func newPublishSrvDeletePublishResult() interface{} {
	return publish.NewPublishSrvDeletePublishResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PublishAction(ctx context.Context, request *publish.TokPublishActionRequest) (r *publish.TokPublishActionResponse, err error) {
	var _args publish.PublishSrvPublishActionArgs
	_args.Request = request
	var _result publish.PublishSrvPublishActionResult
	if err = p.c.Call(ctx, "PublishAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, request *publish.TokPublishListRequest) (r *publish.TokPublishListResponse, err error) {
	var _args publish.PublishSrvPublishListArgs
	_args.Request = request
	var _result publish.PublishSrvPublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeletePublish(ctx context.Context, request *publish.TokDeletePublishRequest) (r *publish.TokDeletePublishResponse, err error) {
	var _args publish.PublishSrvDeletePublishArgs
	_args.Request = request
	var _result publish.PublishSrvDeletePublishResult
	if err = p.c.Call(ctx, "DeletePublish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
