package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.USE(gee.Logger())
	engine.USE(gee.Recovery())
	engine.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	})

	engine.GET("/hi", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	engine.GET("/header", func(c *gee.Context) {
		for k, v := range c.Req.Header {
			//fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			c.Write([]byte(fmt.Sprintf("Header[%q] = %q\n", k, v)))
		}
		c.Status(http.StatusOK)
	})

	engine.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	engine.GET("/panic", func(c *gee.Context) {
		names := []string{"gee"}
		c.String(http.StatusOK, names[100])
	})

	engine.Run(":9999")
}
