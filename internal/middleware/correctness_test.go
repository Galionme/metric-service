package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCorrectnessType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/value/counter/example", nil)
	w := httptest.NewRecorder()

	fakeHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {})

	correctedHandler := CorrectnessType(fakeHandler)

	correctedHandler.ServeHTTP(w, req)

	if w.Code == http.StatusNotFound {
		t.Errorf("Expected status less than %d; got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCorrectnessName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/value/counter/example", nil)
	w := httptest.NewRecorder()

	fakeHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {})

	correctedHandler := CorrectnessType(fakeHandler)

	correctedHandler.ServeHTTP(w, req)

	if w.Code == http.StatusNotFound {
		t.Errorf("Expected status less than %d; got %d", http.StatusBadRequest, w.Code)
	}
}
