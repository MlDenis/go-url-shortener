package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) Get() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		URL, err := s.urlsArchive.Get(ctx.Param("id"))
		if err != nil {
			return ctx.String(http.StatusBadRequest, fmt.Errorf("get URL error: %s", err).Error())
		}

		ctx.Response().Header().Set("Location", URL)
		return ctx.NoContent(http.StatusTemporaryRedirect)
	}
}
