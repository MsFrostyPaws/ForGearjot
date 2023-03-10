package httpserver

import (
	"Gearjot/pkg/primes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	IsPrime []bool `json:"response"`
}

// Checking if number prime or not
func IsPrimeNumber(c *gin.Context) {

	var req []interface{}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "the given input is invalid."})
		return
	}

	results := make([]bool, len(req))

	for i, n := range req {
		switch w := n.(type) {
		case float64:
			results[i] = primes.IsPrime(int(w))
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("the given input is invalid, element on index %v is not a number", i+1)})
			return
		}
	}

	res := Response{IsPrime: results}
	c.JSON(http.StatusOK, res)
}
