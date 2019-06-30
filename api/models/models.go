package models

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const PriorityHigh = 1
const PriorityMedium = 2
const PriorityLow = 3

// Stores the priority values for a given domain
type DomainPriority struct {
	Domain   string
	Weight   int
	Priority int
}

// Adds Type()
type PriorityType interface {
	Type() int
}

// Should return the priority based on Weight and Priority values
func (dp *DomainPriority) Type() int {
	if dp.Priority > 5 && dp.Weight > 5 {
		return PriorityHigh
	}
	if dp.Priority > 5 && dp.Weight <= 5 || dp.Weight > 5 && dp.Priority <= 5 {
		return PriorityMedium
	}
	return PriorityLow
}

// Interface for reading and parsing domain's priority values
type Repository interface {
	Read() map[string]*DomainPriority
}

// Will read domain's priority values
type DomainReader struct{}

// Reads domain's priority values from a configuration file
func (reader *DomainReader) Read() map[string]*DomainPriority {

	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))

	domainsFile, err := os.Open(apppath + "/models/domain.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer domainsFile.Close()
	scanner := bufio.NewScanner(domainsFile)

	priorityMap := map[string]*DomainPriority{}
	var dp *DomainPriority

	i := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			i = 0
			continue
		}

		switch i {
		case 0:
			dp = &DomainPriority{}
			dp.Domain = scanner.Text()
		case 1:
			weight, _ := strconv.Atoi(strings.Split(scanner.Text(), ":")[1])
			dp.Weight = weight
		case 2:
			priority, _ := strconv.Atoi(strings.Split(scanner.Text(), ":")[1])
			dp.Priority = priority

			priorityMap[dp.Domain] = dp
		}
		i++
	}
	return priorityMap
}

// Represent a node in a Linked list
type Node struct {
	Value *DomainPriority
	Next  *Node
}

// Linked list of DomainPriority
type DomainList struct {
	first  *Node
	Lenght int
}

// Sets the methods for the Linked list
type DomainListInterface interface {
	Add(priority *DomainPriority)
	Domains() []string
	Shift() *DomainPriority
}

// Should add a DomainPriority in the right position of linked list, sorted by priority
func (domainList *DomainList) Add(dp *DomainPriority) {
	node := &Node{}
	node.Value = dp

	if domainList.first == nil {
		domainList.first = &Node{}
	}

	curr := domainList.first
	if curr.Next == nil {
		curr.Next = node
	} else {
		for curr.Next != nil {

			if dp.Type() < curr.Next.Value.Type() {
				node.Next = curr.Next
				curr.Next = node
				break
			} else {
				curr = curr.Next
			}
		}
		curr.Next = node
	}
	domainList.Lenght++
}

// Should iterate the list and return an array of domains
func (domainList *DomainList) Domains() []string {
	domains := make([]string, domainList.Lenght)

	i := 0
	curr := domainList.first.Next

	for curr != nil {
		domains[i] = curr.Value.Domain
		curr = curr.Next
		i++
	}
	return domains
}

// Should remove and return the first element on the list
func (domainList *DomainList) Shift() (*DomainPriority, error) {
	if domainList.Lenght == 0 {
		return nil, errors.New("DomainList is empty")
	}
	node := domainList.first.Next
	domainList.first.Next = node.Next
	domainList.Lenght--

	return node.Value, nil
}
