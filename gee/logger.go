package gee

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		//log.Printf("Goroutine:%v\n", GoID())
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
