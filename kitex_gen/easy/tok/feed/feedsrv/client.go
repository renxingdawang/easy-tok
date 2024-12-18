// Code generated by Kitex v0.11.3. DO NOT EDIT.

package feedsrv

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	feed "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/feed"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetUserFeed(ctx context.Context, request *feed.TokFeedRequest, callOptions ...callopt.Option) (r *feed.TokFeedResponse, err error)
	GetVideoById(ctx context.Context, request *feed.TokFeedRequest, callOptions ...callopt.Option) (r *feed.TokVideo, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kFeedSrvClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFeedSrvClient struct {
	*kClient
}

func (p *kFeedSrvClient) GetUserFeed(ctx context.Context, request *feed.TokFeedRequest, callOptions ...callopt.Option) (r *feed.TokFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserFeed(ctx, request)
}

func (p *kFeedSrvClient) GetVideoById(ctx context.Context, request *feed.TokFeedRequest, callOptions ...callopt.Option) (r *feed.TokVideo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoById(ctx, request)
}
