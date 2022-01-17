package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/togdon/frogbot/pkg/swagger/server/restapi"
	"github.com/togdon/frogbot/pkg/swagger/server/restapi/operations"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewFrogbotAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			// error handle
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)

	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(GetHelloUser)

	api.GetFrogNameHandler = operations.GetFrogNameHandlerFunc(GetFrogByName)

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

//Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("Ribbit")
}

//GetHelloUser returns Hello + your name
func GetHelloUser(user operations.GetHelloUserParams) middleware.Responder {
	return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
}

//GetFrogByName returns a Frog in png
func GetFrogByName(Frog operations.GetFrogNameParams) middleware.Responder {

	var URL string
	if Frog.Name != "" {
		URL = "https://github.com/togdon/frogbot/raw/main/frogs/" + Frog.Name + ".png"
	} else {
		//by default we return Frogboss
		URL = "https://github.com/togdon/frogbot/raw/main/frogs/frogboss.png"
	}

	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("error")
	}

	return operations.NewGetFrogNameOK().WithPayload(response.Body)
}
