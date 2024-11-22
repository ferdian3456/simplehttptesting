package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type RequestData struct {
	Name string `json:"name"`
}

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, _ *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Hello world")
	})
	router.POST("/data", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		var data RequestData
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Fprintf(writer, "Failed to decode the request's json")
			panic(err)
		}

		log.Println(data)

		fmt.Fprintf(writer, "Post data: %s", data.Name)
	})

	router.PUT("/data/:PutID", func(writer http.ResponseWriter, _ *http.Request, params httprouter.Params) {
		PutIDParams := params.ByName("PutID")
		fmt.Fprintf(writer, "Testing Put Method with: %s", PutIDParams)
	})

	router.DELETE("/data/:DeleteID", func(writer http.ResponseWriter, _ *http.Request, params httprouter.Params) {
		DeleteIDParams := params.ByName("DeleteID")
		fmt.Fprintf(writer, "Testing Delete Method with: %s", DeleteIDParams)
	})

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: router,
	}

	log.Printf("Server is running on: %s", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
