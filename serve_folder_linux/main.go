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

	log.Printf(">> file name : %s | port :%s \n", *fileNamePtr , *port)

	if *fileNamePtr == "*" {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}
	

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

// func printIp(port *string) {

// 	localIp:= GetLocalIP()

// 		if localIp=="" {
// 			log.Println(">>could not display address please find it the long way")
// 			log.Println(">>check the ip of current device by 'ipconfig' or 'ifconfig' ")
// 			log.Println(">>than use the port number displayed below with ip to dail in")
// 		} else {
// 			log.Printf("\n>>go to : '%s"+*port+"' from other device in local network to acesss",localIp)
// 		}

// }

// func GetLocalIP() string{
//     addrs, err := net.InterfaceAddrs()
//     if err != nil {
//         return ""
//     }
//     for _, address := range addrs {
//         // check the address type and if it is not a loopback the display it
//         if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
//             if ipnet.IP.To4() != nil {
//                 return ipnet.IP.String()
//             }
//         }
//     }
//     return ""
// }
