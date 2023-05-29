package gee

import (
	"reflect"
	"testing"
)

func newTestEngine() *Engine {
	r := New()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestEngine(t *testing.T) {
	t.Run("TestGroup", TestRouterGroup_Group)

}

func TestRouterGroup_Group(t *testing.T) {
	e := newTestEngine()
	group := e.Group("/hello")
	group.GET("/b/a", nil)
	route, _ := e.router.getRoute("GET", "/hello/b/a")
	if !reflect.DeepEqual(route.pattern, "/hello/b/a") {
		t.Errorf("pattern:%v!= pattern:%v", route.pattern, "/hello/b/a")
	}

}
