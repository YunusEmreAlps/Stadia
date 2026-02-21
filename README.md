<!-- STADIA -->

![Stadia Logo](docs/assets/Stadia.png)

# Stadia

**Stadia** is a simulation project that allows users to experience the excitement of the Champions League group stage. Built with Go for the backend and Vue.js for the frontend, **Stadia** provides a realistic and interactive platform for simulating football matches, tracking league standings, and predicting championship outcomes.

## Meaning of "Stadia"

The name **"Stadia"** is derived from the Latin word **"stadium"** which historically referred to a large, open-air venue used for athletic competitions and public events. In the context of this project, **"Stadia"** symbolizes a virtual arena where football matches are simulated and fans can engage with the sport in an interactive way.

## Table of Contents

- [Stadia](#stadia)
  - [Meaning of "Stadia"](#meaning-of-stadia)
  - [Table of Contents](#table-of-contents)
  - [Built With](#built-with)
  - [Features](#features)
    - [Core Features](#core-features)
  - [Project Structure](#project-structure)
  - [Getting Started](#getting-started)
  - [API Documentation](#api-documentation)
  - [Testing](#testing)
  - [Deployment](#deployment)
  - [Development Workflow](#development-workflow)
  - [Contributing](#contributing)
  - [License](#license)
  - [Contact](#contact)

---

## Built With

- Backend:
  - [Go](https://golang.org/) (Go 1.21+)
  - [Gin](https://github.com/gin-gonic/gin) (HTTP web framework)
  - [Go Swagger](https://github.com/go-swagger/go-swagger) (API documentation)
  - [Go CORS](https://github.com/rs/cors) (CORS middleware)
  - [Go Testing](https://golang.org/pkg/testing/) (Unit testing)
  - [Viper](https://github.com/spf13/viper) (Configuration management)
  - [Go Modules](https://blog.golang.org/using-go-modules) (Dependency management)
  - Clean Architecture with service layer pattern
- Frontend:
  - [Vue.js](https://vuejs.org/) (Vue 3.4+ with Composition API)
  - [Pinia](https://pinia.vuejs.org/) (State management)
  - [Vite](https://vitejs.dev/) (Build tool)
  - [Axios](https://axios-http.com/) (HTTP client)
  - Custom CSS with CSS Variables for styling
  - [NPM](https://www.npmjs.com/) (Package management)

> EsLint and Prettier configurations are included for code quality and formatting consistency.

## Features

- **Realistic Match Simulation**: Considers team strength, home advantage, and statistical distributions
- **Intelligent Predictions**: Uses **Monte Carlo** simulation with 10,000 iterations for accurate probability calculations
- **Interactive UI**: Real-time updates, auto-play feature, and match result editing
- **Extensible Design**: Built with OOP principles and clean architecture
- **Comprehensive Testing**: Unit tests for critical components

### Core Features

- **League Initialization**
  - Select from 20 pre-configured teams with varying strengths
  - Customizable team selection (minimum 2, optimized for 4 teams)
  - Automatic fixture generation using round-robin algorithm
- **Match Simulation**
  - Week-by-week progression through the league
  - Auto-play feature to simulate all remaining matches
  - Realistic scoring based on team power and home advantage
  - Considers multiple factors: attack strength, defense capability, randomness
- **Live Standings**
  - Real-time league table updates
  - Premier League rules: 3 points for win, 1 for draw, 0 for loss
  - Tie-breaking: Points → Goal Difference → Goals For → Alphabetical
- **Championship Predictions** (from Week 4)
  - Monte Carlo simulation with 10,000 iterations
  - Considers current standings, remaining fixtures, and team strength
  - Real-time probability updates after each week
  - Visual probability bars and percentage displays
- **Match Result Editing**
  - Edit any played match result
  - Automatic recalculation of standings and predictions
  - Maintains data integrity throughout the league
- **Reset Functionality**
  - Reset league to initial state while keeping selected teams
  - Start fresh simulations with same or different teams

> For detailed implementation of the simulation and prediction algorithms, please refer to the [Simulation Guide](docs/SIMULATION.md) document.
> Early Stage - database not implemented yet, so all data is stored in memory. Future plans include integrating a database for persistent storage and enhanced performance.
---

## Project Structure

```bash
Stadia/
├── backend/
│   ├── models/          # Data models (Team, Match, League, Prediction)
│   ├── services/        # Business logic (Simulation, Fixture, League, Prediction)
│   ├── handlers/        # HTTP request handlers
│   ├── main.go          # Application entry point
│   ├── go.mod           # Go module definition
│   └── *_test.go        # Unit tests
│
├── frontend/
│   ├── src/
│   │   ├── components/  # Vue components
│   │   │   ├── InitializeLeague.vue
│   │   │   ├── LeagueControls.vue
│   │   │   ├── LeagueTable.vue
│   │   │   ├── Predictions.vue
│   │   │   └── MatchWeeks.vue
│   │   ├── stores/      # Pinia state management
│   │   ├── services/    # API service layer
│   │   ├── assets/      # CSS and static assets
│   │   ├── App.vue      # Root component
│   │   └── main.js      # Application entry point
│   ├── index.html
│   ├── vite.config.js
│   └── package.json
│
├── docs/
│   └── assets/
└── README.md
```

---

## Getting Started

For detailed setup instructions, please refer to the [Quick Start Guide](docs/QUICK_START.md)

## API Documentation

For detailed API endpoint information, please refer to the [API Documentation](docs/API.md)

## Testing

For detailed testing instructions, please refer to the [Testing Guide](docs/TESTING.md)

## Deployment

For detailed deployment instructions, please refer to the [Deployment Guide](docs/DEPLOYMENT.md)

## Development Workflow

In the **Stadia** project, we follow a structured development workflow to ensure efficient collaboration and code management. This workflow includes the following key components: branching strategy, versioning, and commit message conventions. By following these guidelines, we aim to maintain a clean and organized codebase that is easy to manage and contribute to. For more information, please refer to the [Development Workflow](docs/DEVELOPMENT_WORKFLOW.md) document.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Thanks to the following people who have contributed to this project:

- [Yunus Emre Alpu](https://www.linkedin.com/in/yunus-emre-alpu/) - Creator and Maintainer

---

<div align="center">
    Made with ❤️ and ⚽
</div>
