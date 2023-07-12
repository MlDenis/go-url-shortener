package urlsarchive

import (
	"fmt"
)

type UA struct {
	urlsArchive map[string]string
}

func NewURLArchive() *UA {
	return &UA{
		urlsArchive: make(map[string]string),
	}
}

func (d *UA) Add(id string, url string) {
	d.urlsArchive[id] = url
}

func (d *UA) Get(id string) (string, error) {
	URL, found := d.urlsArchive[id]
	if !found {
		return "", fmt.Errorf("URL not found")
	}
	return URL, nil
}
