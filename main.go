package main

import (
	"fmt"

	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	pusherClient := pusher.Client{
		AppID:   "",
		Key:     "",
		Secret:  "",
		Cluster: "eu",
		Secure:  true,
	}

	data := map[string]string{"message": "hello world"}
	err := pusherClient.Trigger("alerts-channel", "alert", data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
