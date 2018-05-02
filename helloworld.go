package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s\n", req.Proto, req.URL)
<<<<<<< HEAD
	fmt.Fprintln(w, "HELLO ! It's not working")
=======
	fmt.Fprintln(w, "HELLO ! It's jenkins autobuild testing")
>>>>>>> 5777ed777fc75905b41ae4a2f02814f1cdeeddb9
	fmt.Fprintln(w, "See the documentation at http://docs.deis.io/ for more information.")
}
