// Created by Petr Lozhkin

package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/im7mortal/airports/pkg/airports/calc"
	"net/http"
	"time"
)

type sdkGin struct {
}

// Create sdk server
func New() SDKServer {
	return &sdkGin{}
}

func (sdk *sdkGin) GetMainEngine() http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(setCORS())

	router.POST("/calculate", sdk.calc())

	return router
}

// Set CORS policy
func setCORS() gin.HandlerFunc {
	d := cors.Config{
		AllowMethods:     []string{http.MethodOptions, "OPTIONS", "GET", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept-Language"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}
	return cors.New(d)
}

func (sdk *sdkGin) calc() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req [][]string

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, "invalid json, expected array of arrays (routes)")
			c.Abort()
			return
		}

		// we don't expect that every request will have 100 000 entries. Otherwise, I would prefer to do validation during calculation

		if err := validate(req); err != nil {
			c.JSON(http.StatusBadRequest, "invalid json, expected array of arrays (routes)")
			c.Abort()
			return
		}

		res, err := calc.ProcessFlights(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

var minAsciiValue = int('A')
var maxAsciiValue = int('Z')

func validateCode(airport string) error {
	if len(airport) != 3 {
		return fmt.Errorf("invalid airport code: %s (must be exactly 3 letters)", airport)
	}
	for _, letter := range airport {
		asciiValue := int(letter)
		if asciiValue < minAsciiValue || asciiValue > maxAsciiValue {
			return fmt.Errorf("invalid airport code: %s (all letters must be uppercase)", airport)
		}
	}
	return nil
}

func validate(ss [][]string) error {
	for i := range ss {
		if len(ss[i]) != 2 {
			return fmt.Errorf("every flight require 2 elements")
		}
		for j := range ss[i] {
			if err := validateCode(ss[i][j]); err != nil {
				return err
			}
		}
	}

	return nil
}
