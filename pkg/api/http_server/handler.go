package httpserver

import (
	"Gearjot/pkg/primes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Numbers []int `json:"numbers"`
}

type Response struct {
	IsPrime []bool `json:"response"`
}

func IsPrimeNumber(c *gin.Context) {
	var req Request
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "the given input is invalid."})
		return
	}

	results := make([]bool, len(req.Numbers))
	for i, n := range req.Numbers {
		results[i] = primes.IsPrime(n)
	}

	res := Response{IsPrime: results}
	c.JSON(http.StatusOK, res)
}
