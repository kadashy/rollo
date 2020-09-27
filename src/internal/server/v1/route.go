package v1

import (
	"rollo/pkg/centroide"
	"rollo/pkg/direction"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	//TBD
	app.Post("/api/v1/costTravel", direction.GetDirections)

	app.Post("/api/v1/costTravels", direction.GetDirectionMultiple)

	app.Post("/api/v1/centroide", centroide.GetCentroide)

}
