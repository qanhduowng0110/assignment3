package middleware

import (
	"assignment3/cmd/api"
	"assignment3/cmd/jwt"
	"assignment3/ent"
	"assignment3/ent/user"
	"assignment3/util"
	"context"
	"net/http"
)

type ContextKey struct {
	Name string
}

var userCtxKey = &ContextKey{"user"}

func MiddlewareAuthenticateJWT() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			// token := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(header), "Bearer"))

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			//Create user and check if username exists in the database
			users, err := api.DB.User.Query().Where(user.Username(username)).Only(util.Ctx)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, users)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *ent.User {
	raw, _ := ctx.Value(userCtxKey).(*ent.User)
	return raw
}

func ForAdminContext(ctx context.Context) *ent.User {
	raw, _ := ctx.Value("Admin").(*ent.User)
	return raw
}
