package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func doEvery(d time.Duration, f func(time.Time)) {
	ticker := time.NewTicker(d)
	go func() {
		for x := range ticker.C {
			// post data
			f(x)
		}
	}()
	// stop
	time.Sleep(time.Second * 62)
	ticker.Stop()
}

func postData(t time.Time) {
	// random value
	value1 := rand.Intn(100)
	value2 := rand.Intn(100)

	data := map[string]interface{}{
		"water": value1,
		"wind":  value2,
	}

	client := http.Client{}

	// // konfersi ke json
	jsonByte, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// //buffer
	bufer := bytes.NewBuffer(jsonByte)

	// //request
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bufer)
	// add body
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	hasil := body[:30]
	hasil2 := body[42:44]

	log.Println(string(hasil))
	fmt.Println(string(hasil2))
	Status(value1, value2)

}

func Status(water int, wind int) {
	s_wind := ""
	s_water := ""

	switch {
	case water < 5:
		s_water = "aman"
	case water > 8:
		s_water = "bahaya"
	case water >= 6 || water <= 8:
		s_water = "siaga"
	}

	switch {
	case wind < 6:
		s_wind = "aman"
	case wind > 15:
		s_wind = "bahaya"
	case wind >= 7 || wind <= 15:
		s_wind = "siaga"
	}

	// print status

	fmt.Println("status water : ", s_water)
	fmt.Println("status wind :", s_wind)
}

func main() {
	doEvery(15*time.Second, postData)
}
