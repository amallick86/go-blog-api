package main

import (
	"fmt"
	"go-blog-api/auto"
	"go-blog-api/config"
	"go-blog-api/router"
	"log"
	"net/http"
)

func main() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	listen(config.PORT)
}
func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
