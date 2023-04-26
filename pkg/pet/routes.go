package pet

import (
	"github.com/drakondarquesse/pet-api-gateway/pkg/pet/routes"
	"github.com/go-chi/chi"
)

func MountRoutes(r chi.Router) {
	r.Get("/", routes.GetPets)
}
