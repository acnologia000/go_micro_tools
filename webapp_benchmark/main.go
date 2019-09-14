package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"flag"
	"strconv"
)

func main() {

	number := flag.Int("n",1000000,"number hits to make on url for benchmark")
	url := flag.String("url","http://localhost:8080","provide port for testing url")
	start := time.Now()

	flag.Parse()

	for i := 0; i < *number; i++ {

		resp, err := http.Get(*url)
		if err != nil {
			log.Fatal(err)
		}
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		//log.Print(string(dat))
	}
	log.Print("total time take for "+strconv.Itoa(*number)+" reqs = " + start.Sub(time.Now()).String())
}
