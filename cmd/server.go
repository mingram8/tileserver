package main

import (
"flag"
"fmt"
"github.com/mingram/tileserver/middleware"
"log"
"net/http"

"github.com/rs/cors"

)

func main() {
	listenPort := flag.Int("listenPort", 8080, "Listen Port")
	rootDir := flag.String("directory", ".", "Directory that the files are in")
	flag.Parse()

	var httpHandler middleware.HttpHandler
	httpHandler.RootDirectory = *rootDir
	mux := http.NewServeMux()

	mux.HandleFunc("/health",httpHandler.Health)
	mux.HandleFunc("/tiles",httpHandler.Tiles)

	wrapper := httpHandler.NewLogger(mux)

	handler := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT"},
		// Enable Debugging for testing, consider disabling in production
	}).Handler(wrapper)

	listenAddress := fmt.Sprintf(":%d", *listenPort)
	log.Println("Running on ", listenAddress)

	log.Fatal(http.ListenAndServe(listenAddress, handler))

}