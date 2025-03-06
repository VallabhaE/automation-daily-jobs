package messagehelper

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)


func bot(){
	url := "https://api.telegram.org/bot7177711386:AAHOpDuGxQ1KAb1V2sKGIZCPSJ8iR11irBg/getUpdates"
	resp,err := http.Get(url)
	if err!=nil{
		return
	}

	body, err := ioutil.ReadAll(resp.Body) // or io.ReadAll(resp.Body) in Go 1.16+
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Print the response body as a string
	fmt.Println(string(body))


}