package main

import "github.com/kataras/iris"

// Routes handles HTTP requests from iris.
type Routes struct {
	Shortner *Shortner
}

// Root returns the root.
func (r *Routes) Root(c *iris.Context) {
	c.Write("hello world")
}

// Get fetches a value for a given key.
func (r *Routes) Get(c *iris.Context) {
	key := c.Param("key")
	value := r.Shortner.Get([]byte(key))
	c.Write(string(value))
}

// Set writes a given value to the specified key.
func (r *Routes) Set(c *iris.Context) {
	key := []byte(c.Param("key"))
	value := []byte(c.Param("value"))
	err := r.Shortner.Set(key, value)

	if err == nil {
		c.Write("OK")
	} else {
		c.Log("ERROR: ", err.Error(), "\n")
		c.Write("Oops >_<")
	}
}

// Shorten attempts to shorten a URL to a UID
func (r *Routes) Shorten(c *iris.Context) {
	url := []byte(c.Param("url"))
	uid, err := r.Shortner.Shorten(url)

	if err == nil {
		c.Write("Shortened: %s", uid)
	} else {
		c.Log("ERROR: ", err.Error(), "\n")
		c.Write("Oops >_<")
	}
}

// Lookup attempts to fetch previously shortened URLs based on UID
func (r *Routes) Lookup(c *iris.Context) {
	uid := []byte(c.Param("uid"))
	url := r.Shortner.Lookup(uid)
	c.Write("URL: %s", url)
}
