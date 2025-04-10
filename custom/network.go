package custom

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func HttpGet(url string) (body string, erro error) {
	//var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36")
	client := &http.Client{}
	// Enviar a requisição
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	//	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		bodyString := string(bodyBytes)
		return bodyString, nil
	} else {
		err := errors.New("Status Code: " + string(resp.StatusCode))
		return "", err
	}
}

func StractDomain(url string) (string, error) {
	splited_url := strings.Split(url, "/")
	if len(splited_url) < 3 {
		erro := errors.New("A URL mencionada não é valida")
		return "", erro
	}
	if splited_url[1] != "" || splited_url[2] == "" {
		erro := errors.New("A URL mencionada não é valida")
		return "", erro
	}
	return splited_url[2], nil
}
