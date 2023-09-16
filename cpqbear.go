package cpqbear

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func GetAccessToken(url string, username string, password string) (string, error) {
	payload := []byte("grant_type=password&username=" + username + "&password=" + password)

	request, err := http.NewRequest("GET", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var tokenResponse AccessTokenResponse

	err = json.Unmarshal(responseBody, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}
