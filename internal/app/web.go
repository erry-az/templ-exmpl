package app

import (
	"github.com/erry-az/templ-exmpl/internal/handler/web"
	"github.com/erry-az/templ-exmpl/server"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func NewWeb() error {
	srv := server.NewHttp(&server.HttpConfig{
		Port: 8080,
	})

	webCounter := web.NewCounter()

	srv.AddRoute(func(e *echo.Echo) {
		e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
		
		e.GET("/", webCounter.Page)
		e.POST("/", webCounter.Add)
	})

	return srv.Start()
}
