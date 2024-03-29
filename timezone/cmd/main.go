package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main()  {
	var timezone []string
	res, err := http.Get("http://worldtimeapi.org/api/timezone")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	
	if res.StatusCode != 200{
		log.Panicf("Got Status code: %d", res.StatusCode)
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if json.Unmarshal(body, &timezone) != nil{
		log.Panicf("Unable to parse json")
	}
}

