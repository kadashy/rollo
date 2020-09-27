package v1

import (
	"rollo/pkg/direction"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {

	//TBD
	app.Post("/api/v1/costTravel", direction.GetDirections)
	app.Post("/api/v1/costTravels", direction.GetDirectionMultiple)

}
