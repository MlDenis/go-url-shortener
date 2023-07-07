package app

import (
	"github.com/MlDenis/go-url-shortener/internal/config"
	"io"
	"net/http"
	"path"
	"github.com/google/uuid"
)

func Run() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	http.HandleFunc("/", ServeTheRequest)
	http.ListenAndServe(cfg.BaseURL, nil)

	return nil
}

func ServeTheRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetHandler(w, r)
	case http.MethodPost:
		PostHandler(w, r)
	default:
		BadRequestHandler(w, r)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	_, shortURL := path.Split(r.URL.Path)
	URL, found := URLsArchive[shortURL]
	if != found {
		BadRequestHandler(w, r)
	}

	http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	shortURL := generateShortURL(string(body))
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))

}

func generateShortURL(URL string) string {
	if value, found := URLsArchive[URL]; found {
		return value
	}

	shortURL := uuid.New().String()
	URLsArchive[shortURL] = URL
	return shortURL
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusBadRequest)
}