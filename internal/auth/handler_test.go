package auth_test

import (
	"bytes"
	"encoding/json"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}
	userRepo := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "SecretKey",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}
	return &handler, mock, err
}

func TestLoginSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("test.mail@email.ru", "$2a$10$jqqElt6xt6xQE7lagt6YNuHNHRsrbREiHLS1d2eDinDA44Vu0rTlm") //password 123
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test.mail@email.ru",
		Password: "123",
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("got %d expected %d", w.Code, http.StatusOK)
	}
}

func TestRegisterHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	emptyRows := sqlmock.NewRows([]string{"email", "password"})
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("test.mail@email.ru", "$2a$10$jqqElt6xt6xQE7lagt6YNuHNHRsrbREiHLS1d2eDinDA44Vu0rTlm") //password 123
	mock.ExpectQuery("SELECT").WillReturnRows(emptyRows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(rows)
	mock.ExpectCommit()
	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "test.mail@email.ru",
		Password: "123",
		Name:     "Игорь",
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("got %d expected %d", w.Code, 201)
	}
}
