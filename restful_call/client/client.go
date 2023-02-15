package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type addResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		url := "http://127.0.0.1:8001/add"
		param := addParam{10, 20}
		paramBytes, _ := json.Marshal(param)
		resp, err := http.Post(url, "application/json", bytes.NewReader(paramBytes))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		respBytes, _ := ioutil.ReadAll(resp.Body)
		var respData addResult
		json.Unmarshal(respBytes, &respData)
	}
	d := time.Since(start)
	fmt.Println(d)

}
