package middleware

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// Queue
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

// Que declaration
var Que []string

// Repository should implement common methods
type Repository interface {
	Read() []*Queue
}

func (q *Queue) Read() []*Queue {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/middleware/domain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	queue := []*Queue{}
	var queueItem *Queue
	i := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			fmt.Println("OUT")
			i = 0
			continue
		}
		fmt.Println("IN", scanner.Text())

		switch i {
		case 0:
			queueItem = &Queue{}
			queueItem.Domain = scanner.Text()
		case 1:
			weight, _ := strconv.Atoi(strings.Split(scanner.Text(), ":")[1])
			queueItem.Weight = weight
		case 2:
			priority, _ := strconv.Atoi(strings.Split(scanner.Text(), ":")[1])
			queueItem.Priority = priority

			queue = append(queue, queueItem)
		}
		i++
	}
	return queue
}

// MockQueue should mock an Array of Queues
func MockQueue() []*Queue {
	return []*Queue{
		{
			Domain:   "alpha",
			Weight:   5,
			Priority: 5,
		},
		{
			Domain:   "omega",
			Weight:   1,
			Priority: 5,
		},
		{
			Domain:   "beta",
			Weight:   5,
			Priority: 1,
		},
	}
}

// ProxyMiddleware should queue our incoming requests
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	var repo Repository
	repo = &Queue{}
	fmt.Println("FROM HEADER", domain)
	for _, row := range repo.Read() {
		fmt.Println("FROM SOURCE", row.Domain, row.Weight, row.Priority)

		//  ALGORITHM HERE...
		//  USE QUE

	}
	Que = append(Que, domain)

	c.Next()
}
