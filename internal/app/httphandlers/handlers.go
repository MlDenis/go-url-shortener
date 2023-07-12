package app

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"path"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	_, shortURL := path.Split(r.URL.Path)

	if shortURL == "" {
		BadRequestHandler(w, r)
		return
	}

	URL, found := URLsArchive[shortURL]
	if !found {
		BadRequestHandler(w, r)
	}

	w.Header().Set("Location", URL)
	fmt.Printf("%s", w)
	//http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
	w.WriteHeader(http.StatusTemporaryRedirect)

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	shortURL := generateShortURL(string(body), "http://localhost:8080/")
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))

}

func generateShortURL(URL string, serverAdress string) string {
	if value, found := URLsArchive[URL]; found {
		return value
	}

	shortURLIdentifier := uuid.New().String()
	URLsArchive[shortURLIdentifier] = URL
	fmt.Printf("Debug: %s%s\n", serverAdress, shortURLIdentifier)
	return serverAdress + shortURLIdentifier
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusBadRequest)
}
