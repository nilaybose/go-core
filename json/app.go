package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var jsonBlob = []byte(`[
		{
		 "name": "Platypus",
		 "order": 1.9
		},
		{
		 "name": "Quoll",
		 "order": 2.8
		}
	   ]`)

	var dst bytes.Buffer
	json.Indent(&dst, jsonBlob, "", " ")

	os.Stdout.Write(dst.Bytes())

	type group struct {
		Name  string `json:"name"`
		Order float64
	}

	var groups []group

	err := json.Unmarshal(jsonBlob, &groups)

	fmt.Println(err)
	fmt.Printf("%+v", groups)

	// var result map[string]interface{}
	// json.Unmarshal(jsonBlob, &result)

	// fmt.Println(result)

}
