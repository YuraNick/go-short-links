package main

// go run ./migrations/automigrate.go

import (
	"go/adv-demo/configs"
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	db.AutoMigrate(&link.Link{}, &user.User{})
}
