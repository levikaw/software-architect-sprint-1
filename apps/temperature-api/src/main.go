package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var port string
	args := os.Args
	if len(args) > 1 {
		num, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error converting '%s' to an integer: %v\n", args[1], err)
		}

		port = fmt.Sprintf("%d", num)
	} else {
		port = fmt.Sprintf("%d", 8081)
	}

	http.HandleFunc("/", handler)

	fmt.Printf("Starting server at port %s", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
