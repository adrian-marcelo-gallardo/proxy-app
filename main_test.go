package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	handlers "github.com/edenriquez/proxy-app/api/handlers"
	server "github.com/edenriquez/proxy-app/api/server"
	utils "github.com/edenriquez/proxy-app/api/utils"
	"github.com/stretchr/testify/assert"
)

func init() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		utils.LoadEnv()
		app := server.SetUp()
		handlers.HandlerRedirection(app)
		wg.Done()
		server.RunServer(app)
	}(wg)
	wg.Wait()
	fmt.Println("init")
}

type Response struct {
	Status       int            `json:"status,omitempty"`
	Response     string         `json:"result,omitempty"`
	ResponseText []ResponseText `json:"res,omitempty"`
}

type ResponseText struct {
}

func TestAlgorithmn(t *testing.T) {
	fmt.Println("test")

	cases := []struct {
		// Attr
		Domain string
		Output string
	}{
		// Values
		{Domain: "alpha", Output: `["alpha"]`},
		{Domain: "", Output: "error"},
	}
	fmt.Println(cases)

	client := http.Client{}

	for _, singleCase := range cases {
		req, err1 := http.NewRequest("GET", "http://localhost:8080/ping", nil)
		req.Header.Add("domain", singleCase.Domain)
		assert.Nil(t, err1)

		response, _ := client.Do(req)
		bytes, err2 := ioutil.ReadAll(response.Body)
		assert.Nil(t, err2)

		valuesToCompare := &Response{}

		err3 := json.Unmarshal(bytes, valuesToCompare)
		assert.Nil(t, err3)

		fmt.Println("Response:", valuesToCompare.ResponseText)
		fmt.Println("Status:", valuesToCompare.Status)

		assert.Equal(t, singleCase.Output, valuesToCompare.Response)
	}
}
