package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file, default config use")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "my.test@mail.ru",
		Password: "$2a$10$jqqElt6xt6xQE7lagt6YNuHNHRsrbREiHLS1d2eDinDA44Vu0rTlm", //123
		Name:     "Василий",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "my.test@mail.ru").
		Delete(&user.User{})
}

func TestLoggingSuccess(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	db := initDb()
	initData(db)

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "my.test@mail.ru",
		Password: "123",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatalf("Token empty")
	}
	removeData(db)
}

func TestLoggingFail(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	db := initDb()
	initData(db)

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "Hey-2@ya.ru",
		Password: "1",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, res.StatusCode)
	}

	removeData(db)
}
