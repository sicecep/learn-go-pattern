package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/sicecep/learn-go-pattern/routes"
	env "github.com/spf13/viper"
)

func init() {
	env.SetConfigName("env")
	env.AddConfigPath(".")

	err := env.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file : %s \n", err))
	}
}

func main() {
	fmt.Println(env.GetString("app.env"))
	port := env.GetInt("app.port")

	r := new(routes.Route)
	route := r.Init()

	srv := &http.Server{
		Handler:      route,
		Addr:         ":" + strconv.Itoa(port),
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
