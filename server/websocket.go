package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Data struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type ErrorData struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,

		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	priceUpdateChan = make(chan float64, 2)
)

// common jsonResponse method

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	jR, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the response with the specified status code
	w.WriteHeader(status)
	w.Write(jR)
}

// Middleware function for authentication

func authMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorization := r.Header.Get("Authorization")

		fmt.Println("Auth======", authorization)

		if authorization == "" {
			jsonResponse(w, http.StatusUnauthorized, ErrorData{"Unauthorized Request!", http.StatusUnauthorized})
			// http.Error(w, "Unauthorized Request!", http.StatusUnauthorized)
			return
		}

		fmt.Println("executed middleware")

		// call the next handler
		next.ServeHTTP(w, r)
	})

}

func handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Continuously read and handle messages from the client

	// Read message from the client
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error reading message:", err)
	}

	log.Printf("Received message from client: %s", msg)

	if string(msg) == "Ping" {
		conn.WriteJSON("Pong")
	}

	// Continuously send price updates to the client
	for price := range priceUpdateChan {
		err := conn.WriteJSON(price)
		if err != nil {
			log.Println("Error writing JSON:", err)
			return
		}
	}
}

func sendPriceUpdates() {
	for {
		// Simulate price updates
		price := 100 + rand.Float64()*10 - 1 // Example price update logic
		fmt.Println(price)
		priceUpdateChan <- price

		time.Sleep(2 * time.Second) // Simulate a delay between updates
	}
}

func handleHello(w http.ResponseWriter, r *http.Request) {

	// decode the json data from request body

	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	// Send the same data in the response
	jsonResponse(w, http.StatusOK, data)

}

func main() {
	fmt.Println("Websocket programming in Go")

	// Start a goroutine to send price updates
	go sendPriceUpdates()

	// define router using mux
	router := mux.NewRouter()

	// Apply the authentication middleware to all routes
	router.Use(authMiddleware)

	// Handle WebSocket connections
	router.HandleFunc("/ws", handleWebSocketConnection)

	router.HandleFunc("/api/hello", handleHello).Methods("POST")

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/", fs)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}

//
