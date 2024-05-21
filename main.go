package main

import (
	"eat_box/config"
	"eat_box/internal/kafka"
	"eat_box/internal/router"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	go kafka.ListenScore() //启动一个协程来监听score
	r := router.NewRouter()
	err = r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
