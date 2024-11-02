package main

import (
	publish "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/publish/publishsrv"
	"log"
)

func main() {
	svr := publish.NewServer(new(PublishSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
