package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"
)

type Elm struct {
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

// type Elm struct {
// 	Registros []struct {
// 		ID            string `json:"id"`
// 		Data          string `json:"data"`
// 		CodigoSensor  string `json:"codigo_sensor"`
// 		Latitude      string `json:"latitude"`
// 		Longitude     string `json:"longitude"`
// 		IDEquipamento string `json:"id_equipamento"`
// 	} `json:"Registros"`
// }

func (e *Elm) Unmarshal(b []byte) error {
	return json.Unmarshal(b, e)
}

func main() {
	start := time.Now()

	//fileName := os.Getenv("file_base")
	//fileName := "/media/andrecamilo/Dados/dbs/locacao/bases_locacao/locacao_JAN_01_10.JSON"
	fileName := "/media/andrecamilo/Dados/dbs/locacao/bases_locacao/locacao_20200131.json"
	//fileName := "../teste.json"
	//fileName := "/home/andrecamilo/Projects/locacao_20200131.json"
	f, err := os.Open(fileName)

	byteValue, _ := ioutil.ReadAll(f)
	e := &Elm{}
	json.Unmarshal(byteValue, &e)

	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatalf("Could not obtain stat, handle error: %v", err.Error())
	}

	r := bufio.NewReader(f)
	d := json.NewDecoder(r)
	i := 0

	d.Token()
	for d.More() {
		elm := &Elm{}
		d.Decode(elm)
		//err := elm.Unmarshal(d);
		fmt.Printf("%v \n", elm)
		i++
	}
	d.Token()
	elapsed := time.Since(start)

	fmt.Printf("Total of [%v] object created.\n", i)
	fmt.Printf("The [%s] is %s long\n", fileName, FileSize(fi.Size()))
	fmt.Printf("To parse the file took [%v]\n", elapsed)
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+"%s", val, suffix)
}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1024, sizes)
}
