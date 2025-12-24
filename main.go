package main

import (
	"fmt"
	"net/http"
)

func main() {

	app := &app{port: ":8080", debug: true}
	mux := http.NewServeMux()
	srv := &http.Server{Addr: app.port, Handler: mux}

	mux.HandleFunc("GET /users", app.GetUsersHandler)
	mux.HandleFunc("POST /users", app.CreateUsersHandler)
	mux.HandleFunc("GET /posts", app.GetPostsHandler)

	fmt.Println("Listening on port " + srv.Addr)
	//err := http.ListenAndServe(srv.Addr, srv)
	err := srv.ListenAndServe()
	if err != nil && app.debug {
		panic(err)
	}
}
