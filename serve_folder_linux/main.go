package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) < 2 {
		log.Print("## not enough arguements \n>> using Default port  >> 0.0.0.0:9500\n")
		fmt.Println(">> serving \"" + dir + "\" on " + " 0.0.0.0:9500")
		http.Handle("/", http.FileServer(http.Dir(dir)))
		http.ListenAndServe("0.0.0.0:9500", nil)
	}

	fmt.Println("serving " + dir + "on " + os.Args[1])
	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe(os.Args[1], nil)
}
