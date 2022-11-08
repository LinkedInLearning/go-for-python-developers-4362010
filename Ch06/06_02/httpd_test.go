package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	healthHandler(w, r)

	res := w.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
}
