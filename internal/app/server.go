package app

import (
	"github.com/MlDenis/go-url-shortener/internal/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	host        string
	urlsArchive urlsarchive.URLsArchive
}

func NewServer(host string, urlsArchive urlsarchive.URLsArchive) *Server {

	return &Server{
		host:        host,
		urlsArchive: urlsArchive,
	}
}

func Run() error {

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	urlArchive := urlsarchive.NewURLArchive()
	s := NewServer(cfg.ServerAdress, urlArchive)

	e := echo.New()
	e.GET("/:id", s.Get())
	e.POST("/", s.Post())

	e.Logger.Fatal(e.Start(s.host))

	return nil
}
