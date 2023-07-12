package httphandlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"path"
)

func (h* ) Get() echo.HandlerFunc  {
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
