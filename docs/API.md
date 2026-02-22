# API Documentation

This document provides an overview of the API endpoints available in the Stadia Championship Predictor application. The API allows clients to interact with the backend services for managing teams, fixtures, simulations, and league standings.

## Base URL

```bash
http://localhost:8000/api
```

## Endpoints

### Initialize League

```http
POST /api/league/initialize
```

**Request Body:**

```json
{
    "teams": [
        {
            "name": "Manchester City",
            "power": 92,
            "logo": "GB-ENG"
        },
        {
            "name": "Bayern Munich",
            "power": 90,
            "logo": "DE"
        },
        {
            "name": "Barcelona",
            "power": 86,
            "logo": "ES"
        },
        {
            "name": "Liverpool",
            "power": 87,
            "logo": "GB-ENG"
        },
        {
            "name": "Real Madrid",
            "power": 91,
            "logo": "ES"
        },
        {
            "name": "Paris Saint-Germain",
            "power": 88,
            "logo": "FR"
        }
    ]
}
```

**Response:** League object with initialized state

---

### Get League State

```http
GET /api/league
```

**Response:** Complete league state including teams, fixtures, and current week

---

### Get Standings

```http
GET /api/league/standings
```

**Response:**

```json
{
  "standings": [
    {
      "id": "uuid",
      "name": "Manchester City",
      "power": 92,
      "played": 3,
      "won": 2,
      "drawn": 1,
      "lost": 0,
      "goalsFor": 7,
      "goalsAgainst": 2,
      "points": 7
    }
  ]
}
```

---

### Play Next Week

```http
POST /api/league/play-next-week
```

**Response:** Updated league state after simulating next week's matches

---

### Play All Weeks

```http
POST /api/league/play-all-weeks
```

**Response:** Final league state after simulating all remaining weeks

---

### Update Match Result

```http
PUT /api/league/match/:id
```

**Request Body:**

```json
{
  "homeScore": 3,
  "awayScore": 1
}
```

**Response:** Updated league state with recalculated standings

---

### Get Predictions

```http
GET /api/league/predictions
```

**Response:**

```json
{
  "week": 4,
  "predictions": [
    {
      "teamId": "uuid",
      "teamName": "Manchester City",
      "probability": 67.5
    }
  ]
}
```

---

### Reset League

```http
POST /api/league/reset
```

**Response:** League reset to initial state (no matches played)

---

### Health Check

```http
GET /health
```

**Response:**

```json
{
  "status": "healthy",
  "service": "stadia-backend"
}
```

## Swagger UI

- Access Swagger UI: <http://localhost:8000/swagger/index.html>

![Swagger UI Screenshot](/docs/assets/swagger.jpeg)
