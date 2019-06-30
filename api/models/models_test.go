package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var reader = &DomainReader{}

// Creates domain priorities Map
var domainPriorities map[string]*DomainPriority

// Init Queue struct
var queue *DomainList

func setupTest() {
	domainPriorities = reader.Read()
	queue = &DomainList{}
}

func TestDomainListAdd(t *testing.T) {
	setupTest()

	cases := []struct {
		Domain string
		Output []string
	}{
		{Domain: "alpha", Output: []string{"alpha"}},
		{Domain: "beta", Output: []string{"beta", "alpha"}},
		{Domain: "delta", Output: []string{"delta", "beta", "alpha"}},
		{Domain: "omega", Output: []string{"delta", "beta", "omega", "alpha"}},
	}

	for _, singleCase := range cases {
		domainPriority := domainPriorities[singleCase.Domain]
		queue.Add(domainPriority)

		assert.Equal(t, singleCase.Output, queue.Domains())
	}
}

func TestDomainListShift(t *testing.T) {
	setupTest()

	queue.Add(domainPriorities["alpha"])
	queue.Add(domainPriorities["beta"])
	queue.Add(domainPriorities["delta"])
	queue.Add(domainPriorities["omega"])

	cases := []struct {
		Output string
	}{
		{Output: "delta"},
		{Output: "beta"},
		{Output: "omega"},
		{Output: "alpha"},
	}

	assert.Equal(t, len(cases), queue.Lenght)
	for _, singleCase := range cases {
		domainPriority, _ := queue.Shift()

		assert.Equal(t, singleCase.Output, domainPriority.Domain)
	}
	assert.Equal(t, 0, queue.Lenght)
}
