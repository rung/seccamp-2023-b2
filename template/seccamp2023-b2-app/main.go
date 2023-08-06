package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

func main() {
	log.Println("Starting server...")

	h := requestHandler

	// Cloud Run injects the PORT environment variable into the container.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	if err := fasthttp.ListenAndServe(":"+port, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}

}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello World!")
}

// detected function by linter
func detected() {
	_, err := os.Open("test")
	_, err = os.Open("test2")
	if err != nil {
		return
	}
	return
}
