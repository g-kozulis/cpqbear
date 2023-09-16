package cpqbear

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}

func GetAccessToken(url string, username string, password string) {
	payload := []byte("grant_type=password&username=" + username + "&password=" + password)

	request, error := http.NewRequest("GET", url, bytes.NewBuffer(payload))
	if error != nil {
		fmt.Println(error)
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println(error)
	}

	formattedData := formatJSON(responseBody)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Response body: ", formattedData)

	// clean up memory after execution
	defer response.Body.Close()
}
