package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/EliasSantiago/api-cotacao/adapter/http/docs"
	"github.com/EliasSantiago/api-cotacao/adapter/postgres"
	"github.com/EliasSantiago/api-cotacao/di"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Api Cotação
// @version 1.0.0
// @contact.name Elias Fonseca
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001
// @BasePath /
func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()
	postgres.RunMigrations()
	quoteService := di.ConfigQuoteDI(conn)

	r := chi.NewRouter()

	r.Post("/quote", quoteService.Create)
	r.Get("/metrics", quoteService.Metrics)
	r.HandleFunc("/check", checkHandler)

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API ok!")
}
