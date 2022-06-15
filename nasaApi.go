package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func runGetNASA(lon string, lat string) (string, error) {
	fmt.Printf("lloll")
	url := "https://api.nasa.gov/planetary/earth/assets?" + "lon=" + lon + "&" + "lat=" + lat + "&date=2020-02-01&api_key=" + NasaAPIKey
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend", err)
		return "error happend", err
	}
	defer resp.Body.Close() // важный пункт!
	respBody, err := ioutil.ReadAll(resp.Body)

	//fmt.Printf("http.Get body %#v\n\n\n", string(respBody))
	urlMap := string(respBody)
	urlMap = urlMap[strings.LastIndex(urlMap, "url\":\"")+6 : strings.LastIndex(urlMap, "\"}")]
	fmt.Printf(urlMap)
	return urlMap, nil
}
