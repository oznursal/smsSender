package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"smsSender/model"
)

func HandlePostRequest(dto model.MessageDto) (*model.SmsResponse, error) {
	url := "https://webhook.site/40850de7-286a-447e-abba-a760d20326cb"

	// create post body
	postBody, _ := json.Marshal(map[string]string{
		"to":      dto.PhoneNumber,
		"content": dto.Content,
	})

	body := bytes.NewBuffer(postBody)
	respBody, err := http.Post(url, "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(respBody.Body)

	var res model.SmsResponse
	err = json.NewDecoder(respBody.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	fmt.Println("Status:", respBody.Status)

	return &res, nil
}
