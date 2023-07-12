package urlarchive

import (
	"fmt"
)

type UA struct {
	urlArchive map[string]string
}

func NewURLArchive() *UA {
	return &UA{
		urlArchive: make(map[string]string),
	}
}

func (d *UA) Add(id string, url string) {
	d.urlArchive[id] = url
}

func (d *UA) Get(id string) (string, error) {
	URL, found := d.urlArchive[id]
	if !found {
		return "", fmt.Errorf("URL not found")
	}
	return URL, nil
}
