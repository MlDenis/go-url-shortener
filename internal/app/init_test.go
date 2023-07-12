package httphandlers

type MockShortener struct {
	state map[string]string
}

func (ms *MockShortener) POSTRequest(baseURL string) (string, error) {
	return ms.state[baseURL], nil
}

func (ms *MockShortener) GETRequest(shortenURL string) (string, error) {
	return ms.state[shortenURL], nil
}
