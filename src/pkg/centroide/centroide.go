package centroide

import (
	"github.com/gofiber/fiber/v2"
)

type CentroideRequest struct {
	Cluster []Cluster `json:"cluster"`
	//Longitude float64 `json:"lng"`
	//Latitude  float64 `json:"lat"`
}

func GetCentroide(c *fiber.Ctx) error {
	var (
		centroide             CentroideRequest
		lat                   float64
		lon                   float64
		counter               float64
		centroideResponse     CentroideResponse
		listCentroideResponse []CentroideResponse
	)

	counter = 0
	lat = 0
	lon = 0
	err := c.BodyParser(&centroide)
	if err != nil {
		c.Status(503)
		return err
	}

	for _, clusters := range centroide.Cluster {
		for _, centro := range clusters.Directions {
			lati := centro.Latitude
			long := centro.Longitude
			lat = lat + lati
			lon = lon + long
			counter++
		}
		centroideResponse.SetCLuster(clusters.IdCluster)
		centroideResponse.SetLat(lat / counter)
		centroideResponse.SetLon(lon / counter)
		listCentroideResponse = append(listCentroideResponse, centroideResponse)
		counter = 0
		lat = 0
		lon = 0
	}

	return c.JSON(listCentroideResponse)
}
