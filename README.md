# Bitasmbl-AI-Powered-Study-Notes-Generator-9de79a

## Description
Build a web application that allows students to input lectures, textbooks, or articles and automatically generate concise, structured study notes. The system uses AI summarization and keyword extraction to create easy-to-review content while maintaining readability and context.

## Tech Stack
- Go
- Nuxt.js
- Material-UI

## Requirements
- Build a web application that allows students to input lectures, textbooks, or articles
- Automatically generate concise, structured study notes
- Use AI summarization and keyword extraction
- Maintain readability and context

## Installation
bash
git clone https://github.com/MrBitasmblTester2/Bitasmbl-AI-Powered-Study-Notes-Generator-9de79a.git
cd Bitasmbl-AI-Powered-Study-Notes-Generator-9de79a
# Go backend setup
go mod tidy
# Nuxt.js frontend setup
npm install


## Usage
bash
# Run Go backend
go run ./...
# Run Nuxt.js frontend
npm run dev


## Implementation Steps
1. Initialize Go backend project structure.
2. Implement endpoints to receive text input and return summarized notes with keywords.
3. Initialize Nuxt.js project for the frontend.
4. Build UI using Material-UI components for text input and notes display.
5. Connect Nuxt.js frontend to Go backend endpoints.
6. Add basic error handling for invalid or empty inputs.
7. Test end-to-end flow from text input to generated notes.