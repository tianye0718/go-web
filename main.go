package main

import (
	"net/http"

	"github.com/tianye0718/go-web/gee"
)

func main() {

	r := gee.New()
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Ye</h1>")
		})
		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=yetian
			c.String(http.StatusOK, "Hello %s, you are at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/yetian
			c.String(http.StatusOK, "Hello %s, you are at %s\n", c.Params["name"], c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":9999")
}
