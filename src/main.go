package main

import (
	"github.com/rs/cors"
	"net-http/myapp/route"
	"net/http"
)

func main() {
	router := route.Router{Mutex: http.NewServeMux()}
	router.GetRouter()
	router.GetAuthRouter()
	router.GetCourseRouter()
	// corsについて https://maku77.github.io/p/goruwy4/
	corsOrigin := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://maku77.github.io"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	////全てを許可する Access-Control-Allow-Origin: *
	//corsOrigin := cors.Default()
	handler := corsOrigin.Handler(router.Mutex)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}
