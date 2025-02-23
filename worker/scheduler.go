package worker

import (
	"errors"
	"fmt"
	"smsSender/client"
	"smsSender/db"
	"smsSender/model"
	"time"
)

var quit chan bool

func Start() chan bool {
	var ticker = time.NewTicker(2 * time.Minute)
	quit = make(chan bool, 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				sendMessages()
			case <-quit:
				ticker.Stop()
			}
		}
	}()

	return quit
}

func Stop() error {
	if quit == nil {
		return errors.New("Active worker not found!")
	}
	if isClose() {
		return errors.New("Active worker already closed!")
	}

	close(quit)
	return nil
}

func isClose() bool {
	select {
	case <-quit:
		return true
	default:
		return false
	}
}

func sendMessages() bool {
	messages, err := db.GetUnsentMessages()
	if err != nil {
		return true
	}

	for _, message := range messages {
		response, err := client.HandlePostRequest(message)
		if err != nil {
			return true
		}
		saveRedis(response)
		updateSendingStatus(message)
	}

	fmt.Println(messages)
	return false
}

func updateSendingStatus(message model.MessageDto) {
	err := db.UpdateSentMessages(message)
	if err != nil {
		return
	}
}

func saveRedis(smsResponse *model.SmsResponse) {
	err := db.SaveSms(smsResponse)
	if err != nil {
		return
	}
}
