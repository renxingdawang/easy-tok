package main

import (
	relation "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/relation/relationsrv"
	"log"
)

func main() {
	svr := relation.NewServer(new(RelationSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
