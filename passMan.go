package passMan

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
)

func defaultErrorHandler(err error, c *Context) error {
	slog.Error("error", "err", err)
	return nil
}

type (
	ErrorHandler func(error, *Context) error
	thisones     interface {
		PassMan
	}
)

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	ctx      context.Context
}

func (c *Context) Render(component templ.Component) error {
	return component.Render(c.ctx, c.response)
}

type Handler func(c *Context) error

type PassMan struct {
	ErrorHandler ErrorHandler
	router       *httprouter.Router
}

func New() *PassMan {
	return &PassMan{
		router:       httprouter.New(),
		ErrorHandler: defaultErrorHandler,
	}
}

func (s *PassMan) Start(port string) error {
	return http.ListenAndServe(port, s.router)
}

func (s *PassMan) Get(path string, h Handler, plugs ...Handler) {
	s.router.GET(path, s.makeHTTPRouterHandle(h))
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
