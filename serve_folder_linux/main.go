package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"flag"
)

func main() {

	fileNamePtr := flag.String("file","*","name of file , if you want to serve a single file")
	port := flag.String("port","9900","provide port for 0.0.0.0:<port>")
	
	flag.Parse()

	log.Println("file name : %s \nport :%s ", fileNamePtr , port)

	if *fileNamePtr == "*" {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}
	
	log.Print(">> using Default port  >> 0.0.0.0:"+*port+"\n")
	log.Println(">> serving \"" + dir + "\" on " + " 0.0.0.0:"+*port)
	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe("0.0.0.0:"+*port, nil)
	

	log.Println("serving " + dir + "on " + os.Args[1])
	http.Handle("/", http.FileServer(http.Dir(dir)))	
	} else {
		log.Println(*fileNamePtr)
		http.HandleFunc("/",func(w http.ResponseWriter,r* http.Request) {http.ServeFile(w,r,*fileNamePtr)})
		log.Println(">> serving \"" + *fileNamePtr + "\" on " + " 0.0.0.0:"+*port)
		http.ListenAndServe("0.0.0.0:"+*port, nil)
	}
	
}
