package main

import (
	user "easy-tok/kitex_gen/easy/tok/user/usersrv"
	"log"
)

func main() {
	svr := user.NewServer(new(UserSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
