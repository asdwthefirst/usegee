package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestRouter(t *testing.T) {
	t.Run("TestParsePattern", testParsePattern)
	t.Run("TestGetRoute", testGetRoute)
	t.Run("TestGetRoute2", testGetRoute2)
	t.Run("TestGetRoutes", testGetRoutes)
}

func testParsePattern(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"/p/:name", []string{"p", ":name"}},
		{"/p/*", []string{"p", "*"}},
		{"/p/*name/*", []string{"p", "*name"}},
	}

	for _, tc := range testCases {
		actual := parsePattern(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("parsePattern(%s) = %v; expected %v", tc.input, actual, tc.expected)
		}
	}
}

func testGetRoute(t *testing.T) {
	r := newTestRouter()
	testCases := []struct {
		method   string
		path     string
		expected string
	}{
		{"GET", "/", "/"},
		{"GET", "/hello/geektutu", "/hello/:name"},
		{"GET", "/hello/b/c", "/hello/b/c"},
		{"GET", "/hi/geektutu", "/hi/:name"},
		{"GET", "/assets/file1.txt", "/assets/*filepath"},
	}

	for _, tc := range testCases {
		n, ps := r.getRoute(tc.method, tc.path)
		if n == nil {
			t.Errorf("getRoute(%s, %s) = nil; expected %s", tc.method, tc.path, tc.expected)
			continue
		}
		if n.pattern != tc.expected {
			t.Errorf("getRoute(%s, %s) = %s; expected %s", tc.method, tc.path, n.pattern, tc.expected)
		}
		fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
	}
}

func testGetRoute2(t *testing.T) {
	r := newTestRouter()
	testCases := []struct {
		method   string
		path     string
		expected string
	}{
		{"GET", "/assets/file1.txt", "file1.txt"},
		{"GET", "/assets/css/test.css", "css/test.css"},
	}

	for _, tc := range testCases {
		n, ps := r.getRoute(tc.method, tc.path)
		if n == nil {
			t.Errorf("getRoute(%s, %s) = nil; expected pattern %s", tc.method, tc.path, "/assets/*filepath")
			continue
		}
		if n.pattern != "/assets/*filepath" {
			t.Errorf("getRoute(%s, %s) = %s; expected pattern %s", tc.method, tc.path, n.pattern, "/assets/*filepath")
		}
		if ps["filepath"] != tc.expected {
			t.Errorf("getRoute(%s, %s) = %s; expected filepath %s", tc.method, tc.path, ps["filepath"], tc.expected)
		}
	}
}

func testGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET")
	if len(nodes) != 5 {
		t.Errorf("getRoutes('GET') returned %d nodes; expected %d", len(nodes), 5)
	}
	for i, n := range nodes {
		fmt.Printf("%d. %v\n", i+1, n)
	}
}
