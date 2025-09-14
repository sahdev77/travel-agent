// main.go
// Package main implements a travel booking service using Firebase Genkit and Google AI.
// This service provides an HTTP API for booking travel by processing natural language queries
// and utilizing AI-powered tools for flight search and hotel suggestions.
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/firebase/genkit/go/plugins/server"
)

// TravelAgentInput defines the input structure for the travel agent flow.
// It contains the user's natural language query describing their travel needs.
type TravelAgentInput struct {
	// UserQuery is the natural language input from the user describing their travel requirements.
	// Examples: "Book a flight from NYC to LAX", "Find hotels in Paris for next week"
	UserQuery string `json:"userQuery"`
}

// main initializes and starts the travel booking service.
// It sets up the Genkit framework with Google AI plugin, defines the travel booking flow,
// and starts an HTTP server to handle travel booking requests.
//
// The service listens on the port specified by the PORT environment variable,
// defaulting to port 8080 if not set.
//
// Endpoints:
//   - POST /travelAgent: Processes travel booking requests with natural language input
//
// Environment Variables:
//   - PORT: The port number for the HTTP server (default: 8080)
func main() {
	ctx := context.Background()

	// Initialize Genkit with Google AI plugin and set default model to Gemini 2.5 Flash
	g := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
	)

	// Get the configured travel agent prompt with tools
	agentPrompt := travelAgentPrompt(g)

	// Define the main travel agent flow that processes user queries
	// and returns AI-generated responses with travel recommendations
	travelAgentFlow := genkit.DefineFlow(g,
		"travelAgent",
		func(ctx context.Context, input TravelAgentInput) (string, error) {
			var resp *ai.ModelResponse
			var err error

			// Execute the travel prompt with the user's query
			resp, err = agentPrompt.Execute(ctx, ai.WithInput(input))
			if err != nil {
				return "", err
			}

			// Return the AI-generated text response
			return resp.Text(), nil
		},
	)

	// Set up the HTTP endpoint for the travel agent flow with custom JSON parsing
	// This handler provides a REST API interface for travel booking requests.
	//
	// Endpoint: POST /travelAgent
	// Content-Type: application/json
	//
	// Request Body:
	//   {
	//     "userQuery": "Natural language travel request"
	//   }
	//
	// Response Body:
	//   {
	//     "result": "AI-generated travel recommendations and booking information"
	//   }
	//
	// Error Responses:
	//   - 405 Method Not Allowed: If request method is not POST
	//   - 400 Bad Request: If JSON parsing fails or required fields are missing
	//   - 500 Internal Server Error: If the AI flow execution fails
	http.HandleFunc("/travelAgent", func(w http.ResponseWriter, r *http.Request) {
		// Ensure only POST requests are accepted for travel agent queries
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		// Parse the JSON request body into TravelAgentInput struct
		// Expected format: {"userQuery": "user's travel request"}
		var input TravelAgentInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		
		// Execute the travel agent flow with the parsed input
		// This will process the user query using AI and available tools
		result, err := travelAgentFlow.Run(r.Context(), input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		// Return the AI-generated response as JSON
		// Format: {"result": "travel recommendations and booking details"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"result": result})
	})

	// Get the port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the HTTP server
	log.Println("Starting Genkit server on port " + port)
	log.Fatal(server.Start(ctx, "127.0.0.1:"+port, http.DefaultServeMux))
}
