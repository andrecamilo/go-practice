package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Message struct {
	DATA []struct {
		ID              string `json:"id"`
		Data            string `json:"data"`
		CodigoSensor    string `json:"codigo_sensor"`
		Latitude        string `json:"latitude"`
		Longitude       string `json:"longitude"`
		IDPlanoTrabalho string `json:"id_plano_trabalho"`
		IDEquipamento   string `json:"id_equipamento"`
	} `json:"DATA"`
}

func main() {
	fileName := "/media/andrecamilo/Dados/dbs/locacao/bases_locacao/locacao_JAN_01_10.JSON"

	dec := json.NewDecoder(strings.NewReader(fileName))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.DATA)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

}
