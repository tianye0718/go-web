package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/tianye0718/go-web/gee"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate processing time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {

	// r := gee.New()
	// r.Use(gee.Logger())
	r := gee.Default()
	// prepare HTML render and static handler
	r.SetFuncMap(template.FuncMap{"FormatAsDate": FormatAsDate})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	// prepare test data
	stu1 := &student{Name: "Ye", Age: 20}
	stu2 := &student{Name: "Vivian", Age: 20}

	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2022, 4, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	// test panic -- index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"Ye"}
		c.String(http.StatusOK, names[2])
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "css.tmpl", nil)
		})
		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=yetian
			c.String(http.StatusOK, "Hello %s, you are at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
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
