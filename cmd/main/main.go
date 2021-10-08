package main

import (
	"context"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"monitoring-service/internal/app/core"
	"monitoring-service/internal/app/database/mongodb"
	"monitoring-service/internal/app/provider/coincap"
	"monitoring-service/internal/app/schedule"
	"monitoring-service/internal/app/security"
	"monitoring-service/internal/app/server"
	"monitoring-service/internal/app/service/impl"
	"time"
)

func main() {
	log.Println("Starting monitoring service")
	config := core.NewConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DatabaseURL()))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Ping database on address: ", config.DatabaseURL())
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	db := mongodb.NewMongoDB(client, config.Database())
	coincapClient := coincap.NewRestCoinCapAPI()
	jwt := security.NewJWTSecurity(config.SecretKey())
	services := impl.NewServiceImpl(db, coincapClient, config.Username(), config.Password(), jwt)
	sch := schedule.NewSchedule(services)
	log.Println("Starting task schedule")
	go sch.TaskUpdateCryptocurrenciesRates()
	srv := server.NewServer(services, jwt)
	srv.ConfigureRoutes()
	log.Println("Starting HTTP server on port", config.Port())
	if err := fasthttp.ListenAndServe(config.Port(), srv.Router().Handler); err != nil {
		log.Fatal(err)
	}
}
