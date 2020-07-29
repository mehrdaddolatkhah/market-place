package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"

	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

// RouteHandler handle all routes and accessibility here
func RouteHandler(client *mongo.Client) http.Handler {

	// handle cors for go-chi
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	var tokenAuth *jwtauth.JWTAuth
	// define a new admin handler and repository
	adminRepo := repository.NewAdminRepo(client)
	adminHandler := business.NewAdminHandler(adminRepo)

	// define a new user handler and repository
	userRepo := repository.NewUserRepo(client)
	userHandler := business.NewUserHandler(userRepo)

	// define a new market handler and repository
	marketRepo := repository.NewMarketRepo(client)
	marketHandler := business.NewMarketHandler(marketRepo)

	router := chi.NewRouter()

	router.Use(cors.Handler)
	router.Use(middleware.Logger)

	// Protected routes
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))

		r.Use(jwtauth.Authenticator)

		AdminRestAPI(router, adminHandler)
		UserRestAPI(router, userHandler)
		MarketRestAPI(router, marketHandler)

	})

	// Public routes
	router.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome anonymous"))
		})
	})

	return router
}
