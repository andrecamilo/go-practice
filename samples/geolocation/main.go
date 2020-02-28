package main

import (
	"fmt"

	"github.com/kelvins/geocoder"
)

func main() {
	// Voce precisa definir a chave da API como especificado no seguinte link:
	// https://developers.google.com/maps/documentation/geocoding/get-api-key
	// geocoder.ApiKey = "YOUR_API_KEY"

	// Veja todos os campos da estrutura Address na documentacao
	address := geocoder.Address{
		Street:  "Central Park West",
		Number:  115,
		City:    "New York",
		State:   "New York",
		Country: "United States",
	}

	// Converte o endereco para a localizacao (latitude e longitude)
	locacaoon, err := geocoder.Geocoding(address)

	if err != nil {
		fmt.Println("Nao foi possivel obter a localizacao: ", err)
	} else {
		fmt.Println("Latitude: ", locacaoon.Latitude)
		fmt.Println("Longitude: ", locacaoon.Longitude)
	}

	// Define a latitude e longitude
	locacaoon = geocoder.locacaoon{
		Latitude:  40.775807,
		Longitude: -73.97632,
	}

	// Converte a localizacao (latitude e longitude) para um slice the enderecos
	addresses, err := geocoder.GeocodingReverse(locacaoon)

	if err != nil {
		fmt.Println("Nao foi possivel obter os enderecos: ", err)
	} else {
		// Geralmente o primeiro endereco retornado pela API
		// eh mais detalhado, entao vamos trabalhar com ele
		address = addresses[0]

		// Mostra o endereco formatado pelo pacote geocoder
		fmt.Println(geocoder.FormatAddress(address))
		// Mostra o endereco direto da API
		fmt.Println(address.FormattedAddress)
		// Mostra o tipo do endereco
		fmt.Println(address.Types)
	}
}
