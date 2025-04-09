package middleware

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/jwt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, jwtData := jwt.NewJwt(config.Auth.Secret).Parse(token)
		// if !isValid {

		// }
		fmt.Println(isValid)
		fmt.Println(jwtData)
		next.ServeHTTP(w, r)
	})
}
