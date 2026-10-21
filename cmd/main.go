package main

import (
	"fmt"

	"github.com/eduardotashiro/cron-redeploy/internal/config"
)

func main() {
	config := config.Load()
	service := config.ServiceID
	token := config.Token

	fmt.Printf("service: %s \ntoken: %s\n", service, token)

}
