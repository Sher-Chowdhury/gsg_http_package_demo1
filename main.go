package main

import (
	"net/http"
)

func httpResponse(w http.ResponseWriter, r *http.Request) {

	// this works
	//fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])

	// this also works.
	w.Write([]byte("Hello world. I am a web server"))
}

func main() {

	// the HandleFunc calls the httpResponse and provides the 2 arguements for us.
	http.HandleFunc("/homepage", httpResponse)
	http.HandleFunc("/homepage/", httpResponse)
	println("About to start listening on http://localhost:8080, try going to http://localhost:8080/homepage")
	http.ListenAndServe(":8080", nil)

	// var webresponse http.ResponseWriter([]byte("Hello world. I am a web server"))
}

/*
See also:

https://yourbasic.org/golang/http-server-example/

*/
