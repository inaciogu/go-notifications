package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/inaciogu/go-notifications/internal/application/usecase"
	"github.com/inaciogu/go-notifications/internal/infra/messaging/consumer"
	"github.com/inaciogu/go-notifications/internal/infra/repository/postgres"
	"github.com/inaciogu/go-notifications/internal/infra/web/handlers"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	db, err := sql.Open("postgres", "postgres://root:password@localhost:5432/notifications?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repo := postgres.NewPostgresNotificationRepository(db)

	createNotification := usecase.NewCreateNotification(repo)
	listRecipientNotifications := usecase.NewListRecipientNotifications(repo)

	notificationHandler := handlers.NewNotificationHandler(createNotification, listRecipientNotifications)

	notificationConsumer := consumer.NewNotificationConsumer(createNotification)

	go notificationConsumer.Start()

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Post("/notifications", notificationHandler.Create)
	router.Get("/notifications/from/{id}", notificationHandler.ListManyByRecipient)

	fmt.Println("Server running on port 3001")

	http.ListenAndServe(":3001", router)
}
