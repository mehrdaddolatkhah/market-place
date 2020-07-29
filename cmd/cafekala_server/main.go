package main

import (
	"fmt"
	"log"
	"net/http"

	mongodb "github.com/mehrdaddolatkhah/cafekala_server/pkg/repository/mongo"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/transport/rest"
)

func main() {
	// connect to MongoDB
	client, err := mongodb.ConnectToDatabase()

	if err != nil {
		fmt.Println("Failed to connect MongoDB", err)
	}

	// repository := repository.Repository{
	// 	Client: client,
	// }

	// userRepo := repository.NewUserRepo(client)

	// h := business.NewBaseHandler(userRepo)

	// http.HandleFunc("/", h.HelloWorld)

	// s := &http.Server{
	// 	Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	// }

	// s.ListenAndServe()

	//http.ListenAndServe(":3000", rest.RouteHandler(client))
	router := rest.RouteHandler(client)
	log.Fatal(http.ListenAndServe(":8080", router))

}
