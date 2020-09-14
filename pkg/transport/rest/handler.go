package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"

	"github.com/mehrdaddolatkhah/cafekala_server/pkg/business"
	"github.com/mehrdaddolatkhah/cafekala_server/pkg/repository"
	routes "github.com/mehrdaddolatkhah/cafekala_server/pkg/transport/rest/routes"
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
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	// define a new admin handler and repository
	adminRepo := repository.NewAdminRepo(client)
	adminHandler := business.NewAdminHandler(adminRepo)

	// define a new user handler and repository
	userRepo := repository.NewUserRepo(client)
	userHandler := business.NewUserHandler(userRepo)

	// define a new market handler and repository
	marketRepo := repository.NewMarketRepo(client)
	marketHandler := business.NewMarketHandler(marketRepo)

	// define a new market handler and repository
	authRepo := repository.NewAuthRepo(client)
	authHandler := business.NewAuthHandler(authRepo)

	router := chi.NewRouter()

	router.Use(cors.Handler)
	router.Use(middleware.Logger)

	// Protected routes
	router.Group(func(r chi.Router) {

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		// Mount the admin sub-router
		r.Mount("/api/v1/admin", routes.AdminRouter(router, adminHandler))

		// Mount the market sub-router
		r.Mount("/api/v1/market", routes.MarketRouter(router, marketHandler))

		// Mount the user sub-router
		r.Mount("/api/v1/user", routes.UserRouter(router, userHandler))

	})

	// Public routes
	router.Group(func(r chi.Router) {
		r.Mount("/api/v1", routes.PublicRouter(router, authHandler))
	})

	return router
}
