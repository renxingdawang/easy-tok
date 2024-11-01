package main

import (
	"context"
	user "easy-tok/kitex_gen/easy/tok/user"
)

// UserSrvImpl implements the last service interface defined in the IDL.
type UserSrvImpl struct{}

// Register implements the UserSrvImpl interface.
func (s *UserSrvImpl) Register(ctx context.Context, req *user.TokUserRegisterRequest) (resp *user.TokUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserSrvImpl interface.
func (s *UserSrvImpl) Login(ctx context.Context, req *user.TokUserRegisterRequest) (resp *user.TokUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserById implements the UserSrvImpl interface.
func (s *UserSrvImpl) GetUserById(ctx context.Context, req *user.TokUserRequest) (resp *user.TokUserResponse, err error) {
	// TODO: Your code here...
	return
}
