package main

import (
	"GoGinRestApi/httpd/handler"
	"GoGinRestApi/platform/goginrestapi"

	"github.com/gin-gonic/gin"
)

func main() {
    restApi := goginrestapi.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/goginrestapi", handler.GoGinRestApiGet(restApi))
	r.POST("/goginrestapi", handler.GoGinRestApiPost(restApi))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// fmt.Println(restApi)
	// restApi.Add(goginrestapi.Item{"Hello", "How are you doing friend?"})
	// fmt.Println(restApi)
}