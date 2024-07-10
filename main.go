package main

import (
    "log"
    "net/http"
    "os"

    "chatbot-backend/handlers"
    "chatbot-backend/utils"
)

func main() {
    utils.ConnectDatabase()
    defer utils.DB.Close()

    http.HandleFunc("/register", handlers.RegisterHandler)
    http.HandleFunc("/login", handlers.LoginHandler)
	// http.HandleFunc("/analyze", handlers.AnalyzeSentimentHandler)


    protected := http.NewServeMux()
    protected.HandleFunc("/ws", handlers.ChatbotHandler)
    protected.HandleFunc("/user/get", handlers.GetUserHandler)
    protected.HandleFunc("/analyze", handlers.AnalyzeSentimentHandler)

    http.Handle("/protected/", utils.AuthMiddleware(protected))

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server started on :%s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
