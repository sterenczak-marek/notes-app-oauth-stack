package internal

import (
	"context"
	"log"

	"github.com/sterenczak-marek/notes-app-oauth-stack/oauth_consumer/internal/models"
)

type key int

const (
	userContextKey key = 1
)

func NewUserContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}
func MustGetUserFromContext(ctx context.Context) *models.User {
	u, ok := ctx.Value(userContextKey).(*models.User)
	if !ok {
		log.Panicf("Unable to get user from context. Ensure if middleware order is correct.")
	}
	return u
}
