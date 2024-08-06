package main

import (
	"context"
	firebase "firebase.google.com/go"
	"go-chat-firebase/handler"
	"go-chat-firebase/routers"
	"google.golang.org/api/option"
	"log"
	"os"
)

func main() {

	home, err := os.Getwd()
	if err != nil {
		log.Fatalf("error get path: %v\n", err)
	}
	ctx := context.Background()
	opt := option.WithCredentialsFile(home + "/go-chat-7343f-firebase-adminsdk-ax6ib-42ebb29264.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}
	defer client.Close()

	router := routers.NewRouter(handler.NewMessage(client))
	routers.EndPoints(router)

}
