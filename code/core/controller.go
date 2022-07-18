package core

import "github.com/gin-gonic/gin"

type Controller interface {
	Bean
	RegisterRoutes() map[string]gin.HandlerFunc
	// HandleRoutes(writer http.ResponseWriter, request *http.Request) (func(writer http.ResponseWriter, request *http.Request), bool)
}
