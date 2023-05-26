package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Req        *http.Request
	Writer     http.ResponseWriter
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string
}

func newContext(req *http.Request, writer http.ResponseWriter) *Context {
	return &Context{
		Req:    req,
		Writer: writer,
		Path:   req.URL.Path,
		Method: req.Method,
		Params: make(map[string]string),
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(statusCode int, data interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(statusCode)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(data); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) Write(data []byte) {
	c.Writer.Write(data)
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
