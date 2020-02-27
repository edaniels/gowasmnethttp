package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"syscall/js"
)

func main() {
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	req.Header.Add("js.fetch:mode", "cors")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	rd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var res interface{}
	if err := json.Unmarshal(rd, &res); err != nil {
		log.Fatal(err)
	}
	js.Global().Set("result", res)
}
