package main

import (
	"context"
	//"github.com/renxingdawang/easy-tok/dao/mysql"
	user "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/user"

	"sync"
)

// UserSrvImpl implements the last service interface defined in the IDL.
type UserSrvImpl struct{}

var UserSrvIns *UserSrvImpl
var UserSrvOnce sync.Once

func GetUserSrvImpl() *UserSrvImpl {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrvImpl{}
	})
	return UserSrvIns
}

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
