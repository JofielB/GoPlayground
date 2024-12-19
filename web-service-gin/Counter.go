package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Counter struct {
	count int
}

func (c *Counter) GetCount(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, c.count)
}

func (c *Counter) Increment(g *gin.Context) {
	c.count++

	c.GetCount(g)
}

func (c *Counter) Decrement(g *gin.Context) {
	c.count--

	c.GetCount(g)
}
