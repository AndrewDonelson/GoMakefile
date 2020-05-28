package main

import (
	"net/http"

	"github.com/AndrewDonelson/golog"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	golog.Log.HandlerLog(w, r)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	sent, err := w.Write([]byte(`{"message": "hello world"}`))
	if err != nil {
		golog.Log.Error(err)
	} else {
		golog.Log.Successf("Sent [%d] bytes", sent)
	}

}

func main() {
	golog.Log.SetModuleName("http-svc")
	golog.Log.SetEnvironment(golog.EnvDevelopment)

	s := &server{}
	http.Handle("/", s)
	golog.Log.Print("Listening at localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		golog.Log.Fatal(err.Error())
	}
}
