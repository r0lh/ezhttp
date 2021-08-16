package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("[!] No port provided.\nUsage: %s <port>\ne.g. %s 9000\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}

	port := fmt.Sprintf(":%s", os.Args[1])

	fs := http.FileServer(http.Dir(""))

	log.Fatal(http.ListenAndServe(port, fs))
}
