package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(issuer, audience string) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)
	r.Handle("/info/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewPart))).Methods(http.MethodGet)

	r.Handle("/info/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchParts))).Methods(http.MethodGet)

	r.Handle("/info", mw.Handler(http.HandlerFunc(CreatePart))).Methods(http.MethodPost)
	r.Handle("/info/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdatePart))).Methods(http.MethodPut)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
