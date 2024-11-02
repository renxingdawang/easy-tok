package main

import (
	favorite "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/favorite/favoritesrv"
	"log"
)

func main() {
	svr := favorite.NewServer(new(FavoriteSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
