// Package main provides travel-related tools for flight search and hotel suggestions.
// Note: These are mock implementations for demonstration purposes.
// In production, these would integrate with real APIs like Amadeus, Expedia, etc.
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

// SearchFlightsInput defines the input for the flight search tool.
type SearchFlightsInput struct {
	Departure string `json:"departure"` // Departure city
	Arrival   string `json:"arrival"`   // Arrival city
}

// SuggestHotelInput defines the input for the hotel suggestion tool.
type SuggestHotelInput struct {
	Destination string `json:"destination"` // Destination city
}

// searchFlightsTool creates a mock flight search tool.
// In production, this would integrate with real flight booking APIs.
func searchFlightsTool(g *genkit.Genkit) ai.Tool {

	flightSearchTool := genkit.DefineTool(g,
		"searchFlights",
		"Searches for flights between a departure and arrival city.",
		func(ctx *ai.ToolContext, input SearchFlightsInput) (string, error) {
			log.Printf("Tool called: searchFlights from %s to %s", input.Departure, input.Arrival)
			return fmt.Sprintf("Found flights from %s to %s. A non-stop flight is available for $550.", input.Departure, input.Arrival), nil
		},
	)

	return flightSearchTool
}

// suggestHotelTool creates a mock hotel suggestion tool.
// In production, this would integrate with real hotel booking APIs.
func suggestHotelTool(g *genkit.Genkit) ai.Tool {
	hotelSuggestionTool := genkit.DefineTool(g,
		"suggestHotel",
		"Suggests a popular and well-rated hotel in a given destination.",
		func(ctx *ai.ToolContext, input SuggestHotelInput) (string, error) {
			log.Printf("Tool called: suggestHotel in %s", input.Destination)
			switch strings.ToLower(input.Destination) {
			case "london":
				return "The Savoy is a highly-rated luxury hotel with excellent reviews.", nil
			case "tokyo":
				return "The Park Hyatt is a great choice with stunning city views.", nil
			default:
				return fmt.Sprintf("I'm sorry, I don't have a specific hotel recommendation for %s.", input.Destination), nil
			}
		},
	)
	return hotelSuggestionTool
}
