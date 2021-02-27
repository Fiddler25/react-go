package main

import (
	"react-go/backend/config/db"
	"react-go/backend/config/router"
)

func main() {
	db.Init()
	router.Init()
}
