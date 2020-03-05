package main

import (
    "fmt"
    "time"
    "github.com/kelvins/sunrisesunset"
)

func main() {
    // Set the parameters
    latitude  := -23.545570
    longitude := -46.704082
    utcOffset := -3.0
    date      := time.Date(2017, 3, 23, 0, 0, 0, 0, time.UTC)

    // Calculate the sunrise and sunset times
    sunrise, sunset, err := sunrisesunset.GetSunriseSunset(latitude, longitude, utcOffset, date)

    // If no error has occurred, print the results
    if err == nil {
        fmt.Println("Sunrise:", sunrise.Format("15:04:05")) // Sunrise: 06:11:44
        fmt.Println("Sunset:", sunset.Format("15:04:05")) // Sunset: 18:14:27
    } else {
        fmt.Println(err)
    }
}