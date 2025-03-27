# Home Health Visit Transcript Analyzer

This application consists of two parts:
1. A GraphQL service in Go that analyzes home health care visit transcripts
2. A React frontend that provides a user interface for entering transcripts and viewing analysis results

## Features

- Extract vital signs from visit transcripts
- Parse OASIS-related elements
- Generate visit summaries
- Modern, responsive UI
- Real-time analysis

## Prerequisites

- Go 1.16 or later
- Node.js 14 or later
- npm 6 or later

## Setup

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Generate GraphQL code:
   ```bash
   go run github.com/99designs/gqlgen generate
   ```

4. Start the server:
   ```bash
   go run server.go
   ```

The GraphQL server will be available at `http://localhost:8080/query`

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm start
   ```

The React app will be available at `http://localhost:3000`

## Usage

1. Open your browser and navigate to `http://localhost:3000`
2. Enter a home health visit transcript in the text area
3. Click "Analyze Transcript" to process the text
4. View the extracted vital signs, OASIS elements, and visit summary

## Example Transcript

Here's an example of a transcript that can be analyzed:

```
Patient presented for follow-up visit. Vital signs: BP 120/80, HR: 72, Temp: 98.6, RR: 16, O2 Sat: 98.
Patient lives alone and is independent with ADLs. No new complaints or concerns.
Visit duration: 60 minutes.
```

## Development

### Backend Development

The backend uses the following structure:
- `schema.graphql`: GraphQL schema definition
- `graph/`: Contains resolvers and generated code
- `server.go`: Main server entry point

### Frontend Development

The frontend uses:
- React with TypeScript
- Material-UI for components
- Apollo Client for GraphQL integration

## Testing

### Backend Tests

Run the backend tests:
```bash
cd backend
go test ./...
```

### Frontend Tests

Run the frontend tests:
```bash
cd frontend
npm test
```
