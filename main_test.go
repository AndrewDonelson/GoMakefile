package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndrewDonelson/golog"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	golog.Log.HandlerLog(w, r)
	golog.Log.HandlerLogf(w, r, "")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message": "hello world"}`))
}
func TestHandlers(t *testing.T) {
	golog.Log.SetModuleName("test-svc")
	golog.Log.SetEnvironment(golog.EnvDevelopment)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeHTTP)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"message": "hello world"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

/*********************** BENCHMARKS *****************************/
func BenchmarkLoggerNew(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		log := golog.NewLogger(nil)
		if log == nil {
			panic(fmt.Errorf("BenchmarkLoggerNew failed to create NewLogger"))
		}
	}
}

func BenchmarkLoggerNewLogger(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		log := golog.NewLogger(nil)
		if log == nil {
			panic(fmt.Errorf("BenchmarkLoggerNew failed to create NewLogger"))
		}
		log.Options.Module = "BenchNewLogger"
		log.SetEnvironment(0)
	}
}
