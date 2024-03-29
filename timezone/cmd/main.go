package main

import (
	// "fmt"
	"log"
	"net/http"
)

func main()  {
	res, err := http.Get("http://worldtimeapi.org/api/timezone")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	
	if res.StatusCode != 200{
		log.Panicf("Got Status code: %d", res.StatusCode)
	}

}
