package main

import (
	"eat_box/config"
	"eat_box/internal/router"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	r := router.NewRouter()
	err = r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
