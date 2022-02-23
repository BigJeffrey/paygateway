package middlewares

import (
	"log"
	"net/http"
	"paygateway/controllers"

	"github.com/dgrijalva/jwt-go"
)

var mySignedKey = []byte("mySecredPhrase")

func (m *Middleware) IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		if r.RequestURI == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				controllers.ApiError(w, "Pleas log in first", http.StatusUnauthorized)
				return
			}
			controllers.ApiError(w, "Please log in first", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value

		tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return mySignedKey, nil
		})
		if err != nil {
			controllers.ApiError(w, "Please log in first", http.StatusUnauthorized)
			return
		}

		if tkn.Valid {
			next.ServeHTTP(w, r)
		} else {
			controllers.ApiError(w, "Please log in first", http.StatusUnauthorized)
			return
		}
	})
}
