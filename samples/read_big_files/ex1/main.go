package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	//datas := []map[string]interface{}{}
	var file []interface{}
	file, _ := ioutil.ReadFile("../teste.json")
	err := json.Unmarshal(file, &file)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range f {
		z := v.(map[string]interface{})
		for k2, v2 := range z {
			fmt.Println("Key:", k2, "Value:", v2)
		}
	}
}

// func readJSON(fileName string, filter func(map[string]interface{}) bool) []map[string]interface{} {
// 	datas := []map[string]interface{}{}

// 	file, _ := ioutil.ReadFile(fileName)
// 	json.Unmarshal(file, &datas)
// 	datas := []map[string]interface{}{}
// 	filteredData := []map[string]interface{}{}

// 	for _, data := range datas {
// 		// Do some filtering
// 		if filter(data) {
// 			filteredData = append(filteredData, data)
// 		}
// 	}

// 	return filteredData
// }
