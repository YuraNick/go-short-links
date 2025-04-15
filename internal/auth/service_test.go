package auth_test

import (
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"testing"
)

type MockdUserRepository struct {
}

func (repo *MockdUserRepository) Create(user *user.User) (*user.User, error) {
	return user, nil
}

func (repo *MockdUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "test.mail@email.ru"
	authService := auth.NewAuthService(&MockdUserRepository{})
	email, err := authService.Register(initialEmail, "123", "Вячеслав")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("Email expect %s got %s", initialEmail, email)
	}
}
