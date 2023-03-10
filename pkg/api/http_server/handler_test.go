package httpserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCheckPrimesHandler(t *testing.T) {
	router := gin.New()
	router.POST("/", IsPrimeNumber)

	testCases := []struct {
		name             string
		input            Request
		expectedResponse Response
	}{
		{
			name: "empty",
			input: Request{
				Numbers: []int{},
			},
			expectedResponse: Response{
				IsTrue: []bool{},
			},
		},
		{
			name: "primes and no primes",
			input: Request{
				Numbers: []int{2, 3, 4, 5, 6},
			},
			expectedResponse: Response{
				IsTrue: []bool{true, true, false, true, false},
			},
		},
		{
			name: "primes",
			input: Request{
				Numbers: []int{7, 11, 13, 17, 19},
			},
			expectedResponse: Response{
				IsTrue: []bool{true, true, true, true, true},
			},
		},
		{
			name: "no primes",
			input: Request{
				Numbers: []int{8, 10, 12, 14, 15},
			},
			expectedResponse: Response{
				IsTrue: []bool{false, false, false, false, false},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			inputBytes, _ := json.Marshal(tt.input)
			req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(inputBytes))

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			var res Response
			if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
				t.Errorf("error decoding: %v", err)
			}

			if len(res.IsTrue) != len(tt.expectedResponse.IsTrue) {
				t.Errorf("expected %d results, but %d", len(tt.expectedResponse.IsTrue), len(res.IsTrue))
			}

		})
	}
}
