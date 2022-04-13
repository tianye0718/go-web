package main

import (
	"net/http"

	"github.com/tianye0718/go-web/gee"
)

func main() {

	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=yetian
		c.String(http.StatusOK, "Hello %s, you are at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/yetian
		c.String(http.StatusOK, "Hello %s, you are at %s\n", c.Params["name"], c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Params["filepath"]})
	})

	r.Run(":9999")
}
