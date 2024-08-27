package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/belovetech/go-graphql/internal/users"
	"github.com/belovetech/go-graphql/pkg/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// allow unauthenticated user
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := strings.Split(header, " ")[1]
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
			}

			// create user and check if user exist in the db
			user := users.User{Username: username}
			id, err := users.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = strconv.Itoa(id)

			// put it in the context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// call next with the context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContect(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)

	fmt.Print(raw)
	return raw
}
