package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValueMetric(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/value/counter/example", nil)
	w := httptest.NewRecorder()
	ValueMetric(w, req)
	res := w.Result()
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Error")
	}
}
