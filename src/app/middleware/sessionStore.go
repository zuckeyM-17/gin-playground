package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SessionStore() gin.HandlerFunc {
	store := cookie.NewStore([]byte("secret"))
	options := sessions.Options{
		HttpOnly: true,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 30,
	}
	store.Options(options)
	return sessions.Sessions("gin-playground", store)
}
