package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	botToken := "5790417363:AAEVTbxssoejCIZhDDPi4ChnyHrZE7LZAGs"
	botApi := "https://api.telegram.org/bot"
	botURL := botApi + botToken

	offset := 0

	for {
		updates, err := GetUpdate(botURL, offset)
		if err != nil {
			log.Println("Something went wrong", err.Error())
		}
		for _, update := range updates {
			err = response(botURL, update)
			offset = update.UpdateId + 1
		}
		fmt.Println(updates)
	}
}

func GetUpdate(botURL string, offset int) ([]Update, error) {
	resp, err := http.Get(botURL + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var restResponse RestResponse

	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}

	return restResponse.Result, nil
}

func response(botURL string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatID
	botMessage.Text = update.Message.Text

	marshal, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}

	_, err = http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}
	return nil
}
