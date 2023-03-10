package httpserver

import (
	"bytes"
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCheckPrimesHandler(t *testing.T) {
	t.Parallel()

	router := gin.New()
	router.POST("/", IsPrimeNumber)

	testCases := [...]struct {
		name             string
		input            Request
		expectedStatus   int
		expectedResponse Response
	}{
		{
			name: "empty",
			input: Request{
				Numbers: []int{},
			},
			expectedStatus: http.StatusOK,
			expectedResponse: Response{
				IsPrime: []bool{},
			},
		},
		{
			name: "primes and no primes",
			input: Request{
				Numbers: []int{2, 3, 4, 5, 6},
			},
			expectedStatus: http.StatusOK,
			expectedResponse: Response{
				IsPrime: []bool{true, true, false, true, false},
			},
		},
		{
			name: "primes",
			input: Request{
				Numbers: []int{7, 11, 13, 17, 19},
			},
			expectedStatus: http.StatusOK,
			expectedResponse: Response{
				IsPrime: []bool{true, true, true, true, true},
			},
		},
		{
			name: "no primes",
			input: Request{
				Numbers: []int{math.MaxInt64, 10, 12, 14, 15},
			},
			expectedStatus: http.StatusOK,
			expectedResponse: Response{
				IsPrime: []bool{false, false, false, false, false},
			},
		},
		{
			name: "zero and negative",
			input: Request{
				Numbers: []int{0, 1, -5, -100, -74},
			},
			expectedStatus: http.StatusOK,
			expectedResponse: Response{
				IsPrime: []bool{false, false, false, false, false},
			},
		},
		{
			name: "zero and negative",
			input: Request{
				Numbers: []int{-11, math.MaxInt64, 8, 41, 37},
			},
			expectedStatus: http.StatusOK,
			expectedResponse: Response{
				IsPrime: []bool{false, false, false, true, true},
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			inputBytes, err := json.Marshal(tt.input)
			if err != nil {
				t.Errorf("an error happened: %v", err)
			}
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(inputBytes))

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			var res Response
			if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
				t.Errorf("error decoding: %v", err)
			}

			if len(res.IsPrime) != len(tt.expectedResponse.IsPrime) {
				t.Errorf("expected %d results, but %d", len(tt.expectedResponse.IsPrime), len(res.IsPrime))
			}

		})
	}
}
