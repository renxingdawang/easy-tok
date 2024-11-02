// Code generated by Kitex v0.11.3. DO NOT EDIT.

package relationsrv

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	relation "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/relation"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationAction(ctx context.Context, request *relation.TokRelationActionRequest, callOptions ...callopt.Option) (r *relation.TokRelationActionResponse, err error)
	RelationFollowList(ctx context.Context, request *relation.TokRelationFollowListRequest, callOptions ...callopt.Option) (r *relation.TokRelationFollowListResponse, err error)
	RelationFollowerList(ctx context.Context, request *relation.TokRelationFollowerListRequest, callOptions ...callopt.Option) (r *relation.TokRelationFollowerListResponse, err error)
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
	return &kRelationSrvClient{
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

type kRelationSrvClient struct {
	*kClient
}

func (p *kRelationSrvClient) RelationAction(ctx context.Context, request *relation.TokRelationActionRequest, callOptions ...callopt.Option) (r *relation.TokRelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, request)
}

func (p *kRelationSrvClient) RelationFollowList(ctx context.Context, request *relation.TokRelationFollowListRequest, callOptions ...callopt.Option) (r *relation.TokRelationFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFollowList(ctx, request)
}

func (p *kRelationSrvClient) RelationFollowerList(ctx context.Context, request *relation.TokRelationFollowerListRequest, callOptions ...callopt.Option) (r *relation.TokRelationFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFollowerList(ctx, request)
}