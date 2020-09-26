package v1

import (
	"rollo/pkg/direction"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {

	//TBD
	app.Get("/api/v1/costTravel", direction.GetDirections)
	app.Get("/api/v1/costTravels", direction.GetDirectionMultiple)

}
