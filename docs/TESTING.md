# Testing

## Backend Tests

### Run All Tests

```bash
cd backend
go test ./...
```

### Run Tests with Coverage

```bash
go test -v -cover ./...
```

### Run Specific Test Package

```bash
go test -v ./models
go test -v ./services
```

### Test Coverage

The project includes comprehensive unit tests for:

- **Team Model**: Stats updates, goal difference, reset functionality
- **Simulation Service**: Match outcomes, goal generation, win probabilities
- **Fixture Service**: Round-robin generation, optimized scheduling
- **League Service**: Week progression, standings calculation

### Example Test Output

```bash
=== RUN   TestSimulateMatch
--- PASS: TestSimulateMatch (1.23s)
    Strong wins: 872, Draws: 89, Weak wins: 39
    Strong team win rate: 87.20%
    
PASS
coverage: 85.6% of statements
ok      stadia-backend/services    1.234s
```
