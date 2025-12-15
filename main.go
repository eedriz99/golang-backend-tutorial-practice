package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"user/user"
)

type server struct {
	port  string
	debug bool
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello World!"))
			return
		}

	case "/profile":
		switch r.Method {
		case http.MethodGet:
			myProfile := User{
				FirstName: "Idris",
				LastName:  "Akinsola",
				Email:     "akinsolaidris1999@mail.com",
				Bio:       "AI/ML enthusiast, backend developer in making.",
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			err := json.NewEncoder(w).Encode(myProfile)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			//w.Write(myProfile)
			return

		}
		return

	//fmt.Fprintln("Request: %s", r.URL.Path)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
func main() {
	s := &server{port: ":8080", debug: true}
	fmt.Println("Listening on port " + s.port)
	err := http.ListenAndServe(s.port, s)
	if err != nil && s.debug {
		panic(err)
	}
}
