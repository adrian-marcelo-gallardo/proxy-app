package middleware

import (
	"fmt"

	"github.com/adrian-marcelo-gallardo/proxy-app/api/models"
	"github.com/kataras/iris"
)

var reader = &models.DomainReader{}

// Creates domain priorities Map
var domainPriorities = reader.Read()

// Init Queue struct
var Queue = &models.DomainList{}

// ProxyMiddleware should queue our incoming requests
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}

	fmt.Println("FROM HEADER", domain)

	// Gets the domain priority object for current domain
	var domainPriority = domainPriorities[domain]

	// Invalid domain
	if domainPriority == nil {
		c.JSON(iris.Map{"status": 400, "result": "invalid-domain"})
		return
	}

	// Adds domain to priority Queue
	Queue.Add(domainPriority)

	fmt.Println("Domain Queue", Queue.Domains())

	c.Next()
}
