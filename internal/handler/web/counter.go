package web

import (
	"github.com/erry-az/templ-exmpl/view/layout"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const sessionCountKey = "count"

type Counter struct {
	global int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (h *Counter) Page(c echo.Context) error {
	sessionCount := h.sessionGetInt(c, sessionCountKey)

	return Render(c, layout.Counter(h.global, sessionCount))
}

func (h *Counter) Add(c echo.Context) error {
	var (
		addType      = c.FormValue("add")
		sessionCount = h.sessionGetInt(c, sessionCountKey)
	)

	if addType == "global" {
		h.global++
	}

	if addType == "session" {
		h.sessionSetInt(c, sessionCountKey, sessionCount+1)
	}

	return h.Page(c)
}

func (h *Counter) sessionSetInt(c echo.Context, key string, value int) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	sess.Values[key] = value

	sess.Save(c.Request(), c.Response())
}

func (h *Counter) sessionGetInt(c echo.Context, key string) int {
	sess, _ := session.Get("session", c)

	rawResult := sess.Values[key]
	if rawResult == nil {
		return 0
	}

	result, ok := rawResult.(int)
	if !ok {
		return 0
	}

	return result
}
