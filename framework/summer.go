package framework

import (
	"github.com/Meduzz/summer/api"
	"github.com/Meduzz/summer/errors"
	"github.com/gin-gonic/gin"
)

type (
	Hook interface {
		Identity(*gin.Context) string
		Verify(string, *api.Request) bool
	}

	Summer struct {
		handlers map[string]api.Handler
		Auth     Hook
	}
)

func NewFramework() *Summer {
	return &Summer{
		handlers: make(map[string]api.Handler),
	}
}

func (s *Summer) Register(name string, handler api.Handler) {
	s.handlers[name] = handler
}

func (s *Summer) Remove(name string) {
	delete(s.handlers, name)
}

func (s *Summer) Handle(req *api.Request) *api.Response {
	handler, ok := s.handlers[req.Method]

	// drop unknown notifications...?
	if !ok && !req.IsNotification() {
		err := errors.MethodNotFoundError()
		return &api.Response{
			JsonRPC: api.JsonRPC,
			Error:   err,
			ID:      req.ID,
		}
	}

	return handler(req)
}
