package main

import (
	"flag"
	"net/http"

	"github.com/go-workshop/kata1/internal/dictionary"
	"github.com/go-workshop/kata1/internal/server"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	bind := flag.String("b", ":8080", "HTTP server bind")
	flag.Parse()

	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			logrus.Infof("%s [%s] %s", request.RemoteAddr, request.Method, request.URL.String())
			next.ServeHTTP(writer, request)
		})
	})

	router.HandleFunc("/healthz", server.NewHealthzHandler().Handle)

	reader := dictionary.NewEntriesReader()
	router.HandleFunc("/word/{word}", dictionary.NewHandler(reader).Handle)

	logrus.Infof("Server starting at %s", *bind)
	s := server.NewServer(router, *bind)
	logrus.Fatal(s.ListenAndServe())
}
