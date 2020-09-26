package service

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber"
)

func GetOne(c *fiber.Ctx) {
	//TBD

	//url := os.Getenv("SSXX")
	url := "https://maps.googleapis.com/maps/api/directions/json"
	origin := "-33.8690094,151.2092614"
	destination := "-33.8691196,151.210407"
	apikey := "AIzaSyBMhy5POh_I-joxZnOSTUHU_43D8bjhCeY"
	//Start time
	start := time.Now()
	log.Println("Start time: ", start)

	timeout := time.Duration(50 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", url, nil)
	q := request.URL.Query()
	q.Add("origin", origin)
	q.Add("destination", destination)
	q.Add("key", apikey)
	request.URL.RawQuery = q.Encode()

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	//Response
	c.JSON(resp.Body)

	t := time.Now()
	elapsed := t.Sub(start)
	log.Println("Execution time: ", elapsed)
}

func GetAll(c *fiber.Ctx) {

	//Start time
	start := time.Now()
	log.Println("Start time: ", start)
	counter := "Cod: "
	for i := 0; i <= 5; i++ {
		go func() {
			execHttpRequest()
		}()

	}
	//Response
	c.Send("OK", counter)

	t := time.Now()
	elapsed := t.Sub(start)
	log.Println("Execution time: ", elapsed)

}

func execHttpRequest() *http.Response {
	url := "https://maps.googleapis.com/maps/api/directions/json?origin=Toledo&destination=Madrid&region=es&key=AIzaSyBMhy5POh_I-joxZnOSTUHU_43D8bjhCeY"

	timeout := time.Duration(50 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	return resp

}
