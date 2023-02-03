package test

import (
	"context"
	"github.com/AnirudhAgnihotri2902/architecture-setup-go/ent"
	"github.com/AnirudhAgnihotri2902/architecture-setup-go/handlers"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Densityasmt password=Anirudh@123 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	http.HandleFunc("/signup", handlers.Signup)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/refresh", handlers.Refresh)
	http.HandleFunc("/logout", handlers.Logout)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
