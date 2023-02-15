package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func AddHandler(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	param := addParam{}
	json.Unmarshal(b, &param)
	ret := param.X + param.Y
	respBytes, _ := json.Marshal(addResult{Code: 200, Data: ret})
	w.Write(respBytes)
}

func main() {
	http.HandleFunc("/add", AddHandler)
	http.ListenAndServe(":8001", nil)
}
