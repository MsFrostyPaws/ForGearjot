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
	IsTrue []bool `json:"response"`
}

func IsPrimeNumber(c *gin.Context) {
	var reqBody Request
	if err := c.BindJSON(&reqBody); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "the given input is invalid"})
		return
	}

	results := make([]bool, len(reqBody.Numbers))
	for i, n := range reqBody.Numbers {
		results[i] = primes.IsPrime(n)
	}

	res := Response{IsTrue: results}
	c.JSON(http.StatusOK, res)
}
