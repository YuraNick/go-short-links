package auth

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/req"
	"go/adv-demo/pkg/resp"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("/auth/login")
		_, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		res := LoginResponse{
			Token: "jwt",
		}
		resp.SetJson(w, res, 201)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/auth/register")
		_, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		res := RegisterResponse{
			Token: "jwt",
		}
		resp.SetJson(w, res, 201)
	}
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}
