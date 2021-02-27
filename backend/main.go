package main

import (
	"react-go/backend/router"
	"react-go/backend/config/db"
)

func main() {
	db.Init()
	router.Init()
}
