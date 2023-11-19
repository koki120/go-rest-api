package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/koki120/go-rest-api/domain/entity"
	"github.com/koki120/go-rest-api/log"
)

func Authentication(next http.Handler) http.HandlerFunc {
	logger := log.NewLogger()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("TODO:authentication")

		ctx := r.Context()
		ctx = SetUserToContext(ctx, entity.User{})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type ContextKey string

const userKey ContextKey = "userKey"

func SetUserToContext(ctx context.Context, user entity.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetUserFromContext(ctx context.Context) (entity.User, error) {
	v := ctx.Value(userKey)
	user, ok := v.(entity.User)
	if !ok {
		return entity.User{}, errors.New("no user found in context")
	}
	return user, nil
}
