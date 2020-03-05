package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Elm struct {
	Registros []struct {
		ID            string `json:"id"`
		Data          string `json:"data"`
		CodigoSensor  string `json:"codigo_sensor"`
		Latitude      string `json:"latitude"`
		Longitude     string `json:"longitude"`
		IDEquipamento string `json:"id_equipamento"`
	} `json:"Registros"`
}

func (e *Elm) Unmarshal(b []byte) error {
	return json.Unmarshal(b, e)
}

func main() {
	// Open our jsonFile
	//jsonFile, err := os.Open("../teste.json")
	//jsonFile, err := os.Open("/media/andrecamilo/Dados/dbs/locacao/bases_locacao/locacao_20200131.json")
	jsonFile, err := os.Open("/home/andrecamilo/Projects/locacao_20200131.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var e Elm
	json.Unmarshal(byteValue, &e)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	fmt.Println(e)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
}

// func main() {

// 	var basket Elm
// 	var file []interface{}
// 	file, _ := ioutil.ReadFile("../teste.json")
// 	// err := json.Unmarshal(file, &basket)
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// }
// 	fmt.Println(basket.Registros)
// }
