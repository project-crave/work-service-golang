package main

import (
	"crave/hub/cmd/work/cmd/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.Use()
	go startLib(router)
}

func startLib(router *gin.Engine) {
	go lib.Start(router)
}
