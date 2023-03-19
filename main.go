package main

import (
	"log"
	"net/http"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {

	log.Println("starting bot")

	api := NewCodexApi(os.Getenv("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"))

	pref := tele.Settings{
		URL:         "",
		Token:       os.Getenv("5793305089:AAFdb-xyTAXkor7A3Mne4wdzlT82hAIvkBs"),
		Updates:     0,
		Poller:      &tele.LongPoller{Timeout: 10 * time.Second},
		Synchronous: false,
		Verbose:     false,
		ParseMode:   "",
		OnError: func(error, tele.Context) {
		},
		Client:  &http.Client{
			Transport: nil,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				// Add your logic here
				// ...
				// Return an appropriate error value
				return nil
			},
			Jar:     nil,
			Timeout: 0,
		},
		Offline: false,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/ask", func(c tele.Context) error {

		if !c.Message().FromGroup() {
			return c.Reply("Not Allowed, invite to group first")
		}

		var answer string

		question := c.Message().Payload

		result, err := api.GetCodexSuggestion(question)

		if len(result.Choices) > 0 {
			answer = result.Choices[0].Text
		} else {
			answer = "I don't know"
		}

		if err != nil {
			answer = err.Error()
		}

		answer += "\n\n generated using text-davinci-003 model"

		return c.Reply(answer)
	})

	b.Start()

}
