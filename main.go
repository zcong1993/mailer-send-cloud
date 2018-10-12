package main

import (
	sender2 "github.com/zcong1993/mailer-send-cloud/sender"
	"github.com/zcong1993/mailer/client"
	"github.com/zcong1993/mailer/service"
	"github.com/zcong1993/mailer/utils"
	"net/http"
	"os"
	"time"
)

func main() {
	rabbit := utils.EnvOrDefault("RABBIT", "amqp://guest:guest@localhost:5672/")
	exchangeName := utils.EnvOrDefault("EXCHANGE_NAME", "mail")
	retryExchangeName := utils.EnvOrDefault("RETRY_EXCHANGE_NAME", "mail_retry")
	qName := utils.EnvOrDefault("QUEUE_NAME", "mail")

	sender := &sender2.SendCloud{
		ApiKey:     os.Getenv("API_KEY"),
		ApiUser:    os.Getenv("API_USER"),
		Client:     http.Client{Timeout: time.Second * 5},
		ApiAddress: sender2.API_ADDRESS,
	}

	logger := client.NewDefaultLogger(10)

	service.RunService(rabbit, exchangeName, retryExchangeName, qName, sender, logger, utils.MustToInt(utils.EnvOrDefault("MAX_RETRY", "2")))
}
