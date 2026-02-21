# Stadia - Quick Start Guide

This guide will help you get Stadia up and running quickly.

## Prerequisites

- **Go**: Version 1.21 or higher ([Download](https://golang.org/dl/))
- **Node.js**: Version 18 or higher ([Download](https://nodejs.org/))
- **Git**: For cloning the repository

## Installation

```bash
# Move to your workspace
cd <your-workspace>

# Clone this project into your workspace
git clone <repository-url>

# Move to the project root directory
cd stadia

# Open the project in your favorite IDE
code . # For Visual Studio Code
```

### Using Make (Recommended)

If you have `make` installed:

```bash
# Install all dependencies
make install

# In terminal 1: Run backend
make run-backend

# In terminal 2: Run frontend
make run-frontend
```

### Manual Installation

#### Backend

```bash
# Navigate to backend directory
cd backend

# Download Go dependencies
go mod download

# Run the backend server
go run main.go
```

- The backend server will start on `http://localhost:8000`

#### Frontend

```bash
# Navigate to frontend directory
cd frontend

# Install Node dependencies
npm install

# Run the frontend development server
npm run dev
```

The frontend will start on `http://localhost:5173`

## Quick Test

1. Open `http://localhost:5173` in your browser
2. Select 4 teams from the list
3. Click "Start League"
4. Click "Play Next Week" to simulate matches
5. Watch the league table and predictions update

## Features to Try

- **Auto Play**: Automatically play all remaining weeks
- **Edit Results**: Click "Edit" on any played match to change the score
- **Predictions**: Available from Week 4 onwards
- **Reset**: Start over with the same or different teams

## Running Tests

```bash
# Backend tests
cd backend
go test ./...

# Backend tests with coverage
go test -v -cover ./...
```

## Building for Production

```bash
# Build backend
cd backend
go build -o stadia-backend main.go

# Build frontend
cd frontend
npm run build
```

## Docker Deployment

```bash
# Build and run with Docker Compose
docker-compose up --build
```

Access the application at http://localhost

## Troubleshooting

````bash
# make command not found 
`choco install make`

### Backend won't start
- Check if port 8000 is already in use
- Verify Go dependencies: `go mod tidy`

### Frontend won't start
- Check if port 5173 is already in use
- Delete node_modules and reinstall: `rm -rf node_modules && npm install`

### CORS errors
- Ensure backend is running on port 8000
- Check CORS configuration in backend/main.go

## Need Help?

Check the main [README.md](README.md) for detailed documentation.

## Project Structure

```
Stadia/
├── backend/       # Go backend
├── frontend/      # Vue.js frontend
├── docs/          # Documentation
├── Makefile       # Build commands
└── README.md      # Full documentation
```

Happy simulating! ⚽
