package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"log"
)

func main() {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	message := &messaging.Message{
		Data: map[string]string{
			"key": "value",
		},
		Android: &messaging.AndroidConfig{
			Priority: "HIGH",
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					ContentAvailable: true,
				},
			},
		},
		// Token dispositivo destino
		Token: "",
	}

	result, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalf("error send Messaging client: %v\n", err)
	}

	log.Printf("resul as send messagin: %v\n", result)
}
