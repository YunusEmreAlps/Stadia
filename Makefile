# Stadia Makefile

.PHONY: help install-backend install-frontend install run-backend run-frontend run test-backend clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: install-backend install-frontend ## Install all dependencies

install-backend: ## Install backend dependencies
	@echo "Installing Go dependencies..."
	cd backend && go mod download
	@echo "Backend dependencies installed!"

install-frontend: ## Install frontend dependencies
	@echo "Installing Node.js dependencies..."
	cd frontend && npm install
	@echo "Frontend dependencies installed!"

run-backend: ## Run backend server
	@echo "Starting backend server on :8000..."
	cd backend && go run main.go

run-frontend: ## Run frontend development server
	@echo "Starting frontend server on :5173..."
	cd frontend && npm run dev

run: ## Run both backend and frontend (in separate terminals)
	@echo "Please run 'make run-backend' and 'make run-frontend' in separate terminals"

dev-backend: ## Run backend with auto-reload (requires air)
	@echo "Starting backend with live reload..."
	cd backend && air

test-backend: ## Run backend tests
	@echo "Running backend tests..."
	cd backend && go test -v ./...

test-backend-coverage: ## Run backend tests with coverage
	@echo "Running backend tests with coverage..."
	cd backend && go test -v -cover ./...

build-backend: ## Build backend binary
	@echo "Building backend binary..."
	cd backend && go build -o ../bin/stadia-backend main.go
	@echo "Backend binary created at bin/stadia-backend"

build-frontend: ## Build frontend for production
	@echo "Building frontend for production..."
	cd frontend && npm run build
	@echo "Frontend build created at frontend/dist"

build: build-backend build-frontend ## Build both backend and frontend

clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	rm -rf backend/bin
	rm -rf frontend/dist
	rm -rf frontend/node_modules
	@echo "Clean complete!"

docker-build: ## Build Docker images
	@echo "Building Docker images..."
	docker build -t stadia-backend:latest ./backend
	@echo "Docker images built!"

docker-run: ## Run with Docker Compose
	@echo "Starting services with Docker Compose..."
	docker-compose up

lint-backend: ## Lint backend code
	@echo "Linting backend code..."
	cd backend && go fmt ./...
	cd backend && go vet ./...

lint-frontend: ## Lint frontend code
	@echo "Linting frontend code..."
	cd frontend && npm run lint
