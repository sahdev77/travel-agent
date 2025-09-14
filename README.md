# Travel Agent

A Go-based travel agent application built with Google's Genkit framework that helps users discover travel destinations using AI.

## Prerequisites

- Go 1.19 or later
- Node.js and npm (for Genkit CLI)
- Google AI API key

## Setup Instructions

### 1. Initialize the Go Module

```bash
go mod init travel-agent
```

### 2. Install Go Dependencies

```bash
go get github.com/firebase/genkit/go/genkit
go get github.com/firebase/genkit/go/plugins/googlegenai
```

### 3. Install Genkit CLI

```bash
npm install -g genkit-cli
```

### 4. Set Environment Variables

Set your Google AI API key:

```bash
export GOOGLE_AI_API_KEY="your-api-key-here"
```

**Note**: Replace `"your-api-key-here"` with your actual Google AI API key.

## Running the Application

### Start with Genkit UI

To start the application with the Genkit Developer UI:

```bash
genkit start -- go run .
```

This will start:
- The travel agent service on port 8080
- Genkit Developer UI on http://localhost:4000

You should see output similar to:
```
Genkit Developer UI: http://localhost:4000
2025/09/14 12:57:29 Starting Genkit server on port 8080
```

### Testing the API

Once the server is running, you can test the travel agent endpoint:

```bash
curl -X POST http://localhost:8080/travelAgent \
  -H "Content-Type: application/json" \
  -d '{"userQuery": "I want to fly somewhere cool. Where should I go?"}'
```

## Project Structure

```
travel-agent/
├── go.mod              # Go module file
├── go.sum              # Go dependencies checksum
├── main.go             # Main application entry point
├── prompts.go          # AI prompts and templates
├── tools.go            # Travel-related tools and functions
└── README.md           # This file
```

## Features

- AI-powered travel recommendations
- Integration with Google's Genkit framework
- RESTful API for travel queries
- Developer-friendly UI for testing and debugging

## Development

The application uses Google's Genkit framework to provide AI-powered travel recommendations. The Genkit Developer UI allows you to:

- Test different prompts and queries
- Monitor API calls and responses
- Debug the application flow
- Explore available tools and functions

## API Endpoints

### POST /travelAgent

**Request Body:**
```json
{
  "userQuery": "Your travel query here"
}
```

**Example Response:**
The AI will provide personalized travel recommendations based on your query.

## Troubleshooting

1. **Port Issues**: If port 8080 is already in use, the application will automatically try other ports.
2. **API Key**: Ensure your Google AI API key is properly set in the environment variables.
3. **Dependencies**: Run `go mod tidy` if you encounter dependency issues.

