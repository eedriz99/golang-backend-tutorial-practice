package main

import (
	"fmt"
	"net/http"
)

//type server struct {
//	port  string
//	debug bool
//}

//	func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//		switch r.URL.Path {
//		case "/":
//			switch r.Method {
//			case http.MethodGet:
//				w.Write([]byte("Hello World!"))
//				return
//			}
//
//		case "/profile":
//			switch r.Method {
//			case http.MethodGet:
//				myProfile := User{
//					FirstName: "Idris",
//					LastName:  "Akinsola",
//					Email:     "akinsolaidris1999@mail.com",
//					Bio:       "AI/ML enthusiast, backend developer in making.",
//				}
//				w.Header().Set("Content-Type", "application/json; charset=utf-8")
//				err := json.NewEncoder(w).Encode(myProfile)
//				if err != nil {
//					http.Error(w, err.Error(), http.StatusInternalServerError)
//				}
//				w.WriteHeader(http.StatusOK)
//				//w.Write(myProfile)
//				return
//
//			}
//			return
//
//		//fmt.Fprintln("Request: %s", r.URL.Path)
//		default:
//			w.WriteHeader(http.StatusNotFound)
//		}
//	}

func main() {

	app := &app{port: ":8080", debug: true}
	mux := http.NewServeMux()
	srv := &http.Server{Addr: app.port, Handler: mux}

	mux.HandleFunc("GET /users", app.GetUsersHandler)
	mux.HandleFunc("POST /users", app.CreateUsersHandler)

	fmt.Println("Listening on port " + srv.Addr)
	//err := http.ListenAndServe(srv.Addr, srv)
	err := srv.ListenAndServe()
	if err != nil && app.debug {
		panic(err)
	}
}
