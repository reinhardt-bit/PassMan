package passMan

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
)

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	ctx      context.Context
}

func defaultErrorHandler(err error, c *Context) error {
	slog.Error("error", "err", err)
	return nil
}

type (
	ErrorHandler func(err error, c *Context)
	thisones     interface {
		PassMan
	}
)

func (c *Context) Renderer(component templ.Component) error {
	return component.Render(c.response, c.request)
}

type Handler func(c *Context) error

type PassMan struct {
	ErrorHandler ErrorHandler
	router       *httprouter.Router
}

func (s *PassMan) makeHTTPRouterHandle(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ctx := &Context{
			response: w,
			request:  r,
			ctx:      context.Background(),
		}
		if err := h(ctx); err != nil {
			s.ErrorHandler(err, ctx)
		}
	}
}
