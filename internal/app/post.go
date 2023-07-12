package app

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (s *Server) Post() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		body, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}
		if len(body) == 0 {
			return ctx.String(http.StatusBadRequest, "empty request body")
		}

		urlID := uuid.New().String()
		shortURL := s.host + "/" + urlID
		s.urlsArchive.Add(urlID, string(body))

		return ctx.String(http.StatusCreated, shortURL)
	}
}
