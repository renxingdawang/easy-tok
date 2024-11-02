package main

import (
	feed "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/feed/feedsrv"
	"log"
)

func main() {
	svr := feed.NewServer(new(FeedSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
