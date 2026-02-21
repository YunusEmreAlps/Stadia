# Simulation Algorithm

The simulation algorithm is designed to provide a realistic and engaging experience. It takes into account various factors such as team strength, home advantage, and randomness to generate match outcomes that reflect real-world football dynamics. The prediction algorithm uses a Monte Carlo simulation approach, running 10,000 iterations to calculate the probabilities of each team winning the championship based on the current league standings and remaining fixtures.

## Match Simulation Approach

The match simulation uses a sophisticated algorithm that considers multiple factors:

### 1. **Team Power Integration**

```go
// Each team has a power rating (1-100)
homePower := float64(homeTeam.Power) * (1 + homeAdvantage/100)
awayPower := float64(awayTeam.Power)
```

### 2. **Home Advantage**

- Home teams receive a 10% power boost
- Reflects real-world statistics showing home teams win ~46% of matches

### 3. **Expected Goals Calculation**

```go
powerRatio := attackPower / defensePower
expectedGoals := baseGoals * math.Pow(powerRatio, 0.4)
```

- Uses power ratio with dampening factor (0.4) to prevent unrealistic scores
- Adds randomness factor (0.8-1.2) for unpredictability

### 4. **Goal Generation**

- Uses Poisson distribution approximation
- Strong teams can occasionally lose (realistic unpredictability)
- Weak teams can occasionally pull upsets

### 5. **Factors Considered**

- Attack power vs defense power
- Home/away status
- Random variance (the unpredictability of football)
- Statistical probability distributions

### Example Scenarios

- **Strong vs Weak (Power 90 vs 40)**
  - Expected outcome: Strong team wins ~85% of times
  - Typical scoreline: 2-0, 3-1, or similar
  - Upset possible but rare (~5% chance)

- **Equal Teams (Power 75 vs 75)**
  - Expected outcome: Home team slight advantage (~55%)
  - More draws likely
  - Either team can win

---

## Prediction Algorithm

### Monte Carlo Simulation

The prediction system uses Monte Carlo simulation for accurate probability calculation:

#### Algorithm Steps

1. **Snapshot Current State**
   - Capture current points, goal difference, and remaining fixtures

2. **Run Simulations** (10,000 iterations)

   ```go
   for i := 0; i < 10000; i++ {
       winner := simulateRemainingMatches()
       wins[winner.ID]++
   }
   ```

3. **Calculate Probabilities**
```go
probability := float64(winCount) / 10000.0
```

4. **Consider Multiple Factors**
   - Current points (60% weight)
   - Goal difference (20% weight)
   - Team power (20% weight)
   - Remaining fixtures difficulty

5. **Real-time Updates**
   - Predictions update after each week from Week 3 onwards
   - More accurate as season progresses

### Prediction Accuracy

- **Early Season (Week 3-4)**: Moderate accuracy, many possibilities
- **Mid Season (Week 4-5)**: Higher accuracy, clearer leaders
- **Late Season (Week 5-6)**: Very high accuracy, winners often clear

### Example Prediction Scenarios

**Scenario 1: Clear Leader**

```plain
Week 5 Standings:
1. Team A - 12 points
2. Team B - 6 points
3. Team C - 4 points
4. Team D - 1 point

Predictions:
Team A: 95% (almost certain)
Team B: 4% (needs miracle)
Team C: 1%
Team D: 0%
```

**Scenario 2: Close Race**

```plain
Week 4 Standings:
1. Team A - 7 points
2. Team B - 7 points
3. Team C - 6 points
4. Team D - 3 points

Predictions:
Team A: 42% (slight edge)
Team B: 38% (strong contender)
Team C: 19% (outsider chance)
Team D: 1%
```
