package main

import (
	"fmt"
	"github.com/codingsince1985/geo-golang/yandex"
	"strings"
)

func runGetYandex(adr string) (string, string, error) {
	fullAdr := strings.Split(adr, " ")
	//Москва,+Тверская+улица,+дом+7
	//https://geocode-maps.yandex.ru/1.x/?apikey=ваш API-ключ&geocode=Москва, улица Новый Арбат, дом 24
	//url := "https://geocode-maps.yandex.ru/1.x/?apikey=" + YandexAPIKey + "&geocode="
	url := ""
	for i, value := range fullAdr {
		if i == len(fullAdr)-1 {
			url = url + value
		} else {
			url = url + value + "+"
		}

	}
	location := yandex.Geocoder(YandexAPIKey)

	loc, _ := location.Geocode(url)
	fmt.Println(loc)
	return fmt.Sprintf("%f", loc.Lng), fmt.Sprintf("%f", loc.Lat), nil
}
