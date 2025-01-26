package internal

import (
	"context"

	"github.com/laclipasa/la-clipasa/ent"
)

type ctxKeyUser struct{}

// GetUserFromContext retrieves the user from the context.
func GetUserFromContext(ctx context.Context) (*ent.User, bool) {
	u, ok := ctx.Value(ctxKeyUser{}).(*ent.User)
	if !ok {
		return nil, false
	}
	return u, true
}

// SetUserContext stores the user in the context.
func SetUserContext(ctx context.Context, u *ent.User) context.Context {
	return context.WithValue(ctx, ctxKeyUser{}, u)
}
