package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeMetrics(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ValueMetric(w, req)
	res := w.Result()
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Error")
	}
}
