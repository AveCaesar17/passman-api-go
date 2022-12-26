package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testApiServer_HandleHello(test *testing.T) {
	server := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	server.handleHello().ServeHTTP(rec, req)
	assert.Equal(test, rec.Body.String(), "Hi")
}
