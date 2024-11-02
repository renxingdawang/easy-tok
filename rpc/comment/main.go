package main

import (
	comment "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/comment/commentsrv"
	"log"
)

func main() {
	svr := comment.NewServer(new(CommentSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
