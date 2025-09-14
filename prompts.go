// Package main provides travel booking functionality using Firebase Genkit
package main

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

// travelAgentPrompt creates and configures a travel agent assistant prompt
// that can help users with flight searches and hotel recommendations.
//
// Parameters:
//   - g: A Genkit instance used to define the prompt and register tools
//
// Returns:
//   - ai.Prompt: A configured prompt that acts as a travel agent assistant
//
// The prompt includes:
//   - System instructions for professional travel agent behavior
//   - Access to searchFlights and suggestHotel tools
//   - Input handling for BookTravelInput requests
func travelAgentPrompt(g *genkit.Genkit) ai.Prompt {
	// Define system prompt with clear instructions for the travel agent assistant
	// This prompt establishes the AI's role and available capabilities
	sysPrompt := `You are a professional and courteous travel agent assistant.
			You have access to the following tools:
			- searchFlights: to find flights between two cities.
			- suggestHotel: to recommend a hotel in a specific city.

			Your goal is to fulfill the user's travel request by intelligently using the tools at your disposal.

			- If the user asks for a flight, use the searchFlights tool.
			- If the user asks for a hotel, use the suggestHotel tool.
			- If the user asks for both, you should call both tools sequentially or in parallel as needed.
			- If you are missing any information (e.g., a city or destination), you MUST ask the user for it.
			- If the user's request is not related to travel, respond politely that you can only help with travel-related queries.`

	// Create and configure the travel prompt with system instructions,
	// input processing function, and available tools
	prompt := genkit.DefinePrompt(g, "travelAgentPrompt",
		// Set the system prompt that defines the AI's behavior and capabilities
		ai.WithSystem(sysPrompt),

		// Define how user input is processed and formatted for the AI
		// Takes TravelAgentInput and extracts the UserQuery field
		ai.WithPromptFn(func(ctx context.Context, input any) (string, error) {
			return fmt.Sprintf("The user's request is: %s", input.(TravelAgentInput).UserQuery), nil
		}),

		// Register the available tools that the AI can use to fulfill requests
		// searchFlightsTool: for finding flights between cities
		// suggestHotelTool: for recommending hotels in specific locations
		ai.WithTools(searchFlightsTool(g), suggestHotelTool(g)),
	)
	return prompt
}
