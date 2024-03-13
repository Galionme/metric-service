package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateMetric(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/update/counter/example/2", nil)
	w := httptest.NewRecorder()
	UpdateMetric(w, req)
	res := w.Result()
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Error")
	}
}
