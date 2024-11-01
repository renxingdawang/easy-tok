// Code generated by Kitex v0.11.3. DO NOT EDIT.

package usersrv

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/user"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newUserSrvRegisterArgs,
		newUserSrvRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newUserSrvLoginArgs,
		newUserSrvLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetUserById": kitex.NewMethodInfo(
		getUserByIdHandler,
		newUserSrvGetUserByIdArgs,
		newUserSrvGetUserByIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userSrvServiceInfo                = NewServiceInfo()
	userSrvServiceInfoForClient       = NewServiceInfoForClient()
	userSrvServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userSrvServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userSrvServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userSrvServiceInfoForClient
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
	serviceName := "UserSrv"
	handlerType := (*user.UserSrv)(nil)
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
		"PackageName": "user",
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

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserSrvRegisterArgs)
	realResult := result.(*user.UserSrvRegisterResult)
	success, err := handler.(user.UserSrv).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserSrvRegisterArgs() interface{} {
	return user.NewUserSrvRegisterArgs()
}

func newUserSrvRegisterResult() interface{} {
	return user.NewUserSrvRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserSrvLoginArgs)
	realResult := result.(*user.UserSrvLoginResult)
	success, err := handler.(user.UserSrv).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserSrvLoginArgs() interface{} {
	return user.NewUserSrvLoginArgs()
}

func newUserSrvLoginResult() interface{} {
	return user.NewUserSrvLoginResult()
}

func getUserByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserSrvGetUserByIdArgs)
	realResult := result.(*user.UserSrvGetUserByIdResult)
	success, err := handler.(user.UserSrv).GetUserById(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserSrvGetUserByIdArgs() interface{} {
	return user.NewUserSrvGetUserByIdArgs()
}

func newUserSrvGetUserByIdResult() interface{} {
	return user.NewUserSrvGetUserByIdResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *user.TokUserRegisterRequest) (r *user.TokUserRegisterResponse, err error) {
	var _args user.UserSrvRegisterArgs
	_args.Req = req
	var _result user.UserSrvRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *user.TokUserRegisterRequest) (r *user.TokUserRegisterResponse, err error) {
	var _args user.UserSrvLoginArgs
	_args.Req = req
	var _result user.UserSrvLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserById(ctx context.Context, req *user.TokUserRequest) (r *user.TokUserResponse, err error) {
	var _args user.UserSrvGetUserByIdArgs
	_args.Req = req
	var _result user.UserSrvGetUserByIdResult
	if err = p.c.Call(ctx, "GetUserById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
