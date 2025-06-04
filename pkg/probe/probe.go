package probe

import (
	"net/http"
	"time"
)

// HTTPProbe faz uma requisição GET e retorna o status code.
func HTTPProbe(url string) (int, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
