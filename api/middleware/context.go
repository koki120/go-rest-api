package middleware

import (
	"context"
	"errors"

	"github.com/koki120/go-rest-api/domain/entity"
)

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
