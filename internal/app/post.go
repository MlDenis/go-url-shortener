package httphandlers

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
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
