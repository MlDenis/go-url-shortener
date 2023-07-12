package app

import (
	"github.com/MlDenis/go-url-shortener/internal/config"
	"github.com/MlDenis/go-url-shortener/internal/urlsarchive"
	"testing"
)

type testEnv struct {
	urlsArchive urlsarchive.URLsArchive
	s           *Server
}

func newTestEnv(t *testing.T) *testEnv {
	cfg, _ := config.NewConfig()

	te := &testEnv{}

	te.urlsArchive = urlsarchive.NewURLArchive()
	te.s = NewServer(
		cfg.ServerAdress,
		te.urlsArchive,
	)
	return te
}
