package direction

import (
	"context"
	"log"
	"os"
	"rollo/pkg/response"
	"time"

	"github.com/gofiber/fiber"
	"googlemaps.github.io/maps"
)

type MapDirection struct {
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	Mode          maps.Mode `json:"mode"`
	DepartureTime string    `json:"departureTime"`
	PointId       string    `json:"pointID"`
}

func GetDirections(c *fiber.Ctx) {
	var mapRespon response.DirectionResponse

	apikey := os.Getenv("APYKEY")

	mapRequest := new(MapDirection)
	if err := c.BodyParser(mapRequest); err != nil {
		c.Status(503).Send(err)
		return
	}

	start := time.Now()
	log.Println("Start time: ", start)
	ma, err := maps.NewClient(maps.WithAPIKey(apikey))
	if err != nil {
		log.Fatalf("fatal error 1: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:        mapRequest.Origin,
		Destination:   mapRequest.Destination,
		Mode:          mapRequest.Mode,
		DepartureTime: mapRequest.DepartureTime,
	}
	route, _, err := ma.Directions(context.Background(), r)
	if err != nil {
		c.Status(409).Send("fatal error 1: %s", err)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Println("Execution time: ", elapsed)

	mapRespon.SetPoint(mapRequest.PointId)
	for _, rout := range route {
		for _, leg := range rout.Legs {
			mapRespon.SetDistance(leg.Distance.Meters)
			mapRespon.SetDuration(leg.Duration)

		}
	}

	c.JSON(mapRespon)
}

func GetDirectionMultiple(c *fiber.Ctx) {
	var mapRequest []MapDirection
	var mapResponse []response.DirectionResponse
	var mapRespon response.DirectionResponse

	err := c.BodyParser(&mapRequest)
	if err != nil {
		c.Status(503).Send(err)
		return
	}

	for _, marReq := range mapRequest {
		mapRespon = GetMapsDirection(marReq)
		mapResponse = append(mapResponse, mapRespon)
	}

	c.JSON(mapResponse)
}

func GetMapsDirection(mapRequest MapDirection) response.DirectionResponse {
	var response response.DirectionResponse
	apikey := os.Getenv("APYKEY")

	start := time.Now()
	log.Println("Start time: ", start)
	ma, err := maps.NewClient(maps.WithAPIKey(apikey))
	if err != nil {
		log.Fatalf("fatal error 1: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:        mapRequest.Origin,
		Destination:   mapRequest.Destination,
		Mode:          mapRequest.Mode,
		DepartureTime: mapRequest.DepartureTime,
	}
	route, _, err := ma.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error 1: %s", err)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Println("Execution time: ", elapsed)
	response.SetPoint(mapRequest.PointId)
	for _, rout := range route {
		for _, leg := range rout.Legs {
			response.SetDistance(leg.Distance.Meters)
			response.SetDuration(leg.Duration)
			return response

		}
	}

	return response
}
