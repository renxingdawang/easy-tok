package main

import (
	"context"
	favorite "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/favorite"
)

// FavoriteSrvImpl implements the last service interface defined in the IDL.
type FavoriteSrvImpl struct{}

// FavoriteAction implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteAction(ctx context.Context, request *favorite.TokFavoriteActionRequest) (resp *favorite.TokFavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteList(ctx context.Context, request *favorite.TokFavoriteListRequest) (resp *favorite.TokFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}
