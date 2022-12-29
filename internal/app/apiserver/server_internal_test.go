package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUserCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedcode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"username": "caesar_Test",
				"pubkey":   "pubkey_Test",
			},
			expectedcode: http.StatusCreated,
		},
		{
			name:         "invalid",
			payload:      "invalid",
			expectedcode: http.StatusBadRequest,
		},
		{
			name: "invalid_params",
			payload: map[string]string{
				"username": "usertecset",
			},
			expectedcode: http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/create_user", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedcode, rec.Code)

		})
	}
	// rec := httptest.NewRecorder()
	// req, _ := http.NewRequest(http.MethodPost, "/test", nil)
	// s := newServer(teststore.New())
	// s.ServeHTTP(rec, req)
	// assert.Equal(t, rec.Code, http.StatusOK)
}
