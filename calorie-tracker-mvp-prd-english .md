# BalanceLife Calorie Tracker - MVP Product Requirements Document (PRD)

## Document Information

| Property | Details |
|---------|------|
| Document Name | BalanceLife Calorie Tracker MVP Product Requirements Document |
| Version | V1.0.0 |
| Status | Draft |
| Creation Date | 2025-03-17 |
| Last Updated | 2025-03-17 |
| Author | Product Team |
| Reviewers | TBD |

## 1. Introduction

### 1.1 Purpose

This document details the product requirements for the Minimum Viable Product (MVP) version of "BalanceLife," a calorie tracking application. It serves as a guide for the design and development teams.

### 1.2 Product Overview

"BalanceLife" is a mobile application designed to help users easily track their calorie intake and exercise expenditure, providing a clear view of their calorie balance to assist in achieving specific fitness goals (weight loss or weight gain).

### 1.3 MVP Scope

The MVP version focuses on providing a core set of features with a unique meal and workout package approach to validate key product hypotheses and gather initial user feedback for future iterations.

### 1.4 Assumptions & Constraints

- Target users have basic smartphone usage capabilities
- Users are willing to spend at least 5 minutes daily recording food and exercise
- MVP supports iOS and Android platforms only
- Initial focus on English-speaking markets
- MVP will prioritize pre-configured meal and workout packages over extensive food databases

## 2. User Scenarios & Use Cases

### 2.1 Target User Personas

**Persona 1: Weight Loss Seeker**
- 25-45 years old
- Has 10-30 pounds to lose
- Busy professional with limited time for meal planning
- Wants simple, structured approach to calorie deficit

**Persona 2: Muscle Builder**
- 18-35 years old
- Wants to gain lean muscle mass
- Regular gym-goer looking for nutrition structure
- Needs consistent calorie surplus with adequate protein

### 2.2 Core Use Cases

#### UC-01: Select Goal and Create Profile
**Main Flow:**
1. User downloads and opens the app
2. Selects primary goal (weight loss or muscle gain)
3. Enters basic profile information (age, gender, height, weight, activity level)
4. System calculates recommended calorie target
5. User confirms or adjusts target
6. System generates recommended meal and workout packages

#### UC-02: Log Pre-configured Meal Package
**Main Flow:**
1. User taps "Add Meal" button
2. Views list of meal packages suitable for their goal
3. Selects a meal package (e.g., "Chicken, Broccoli & Rice Pack")
4. Confirms portion size (0.5x, 1x, 1.5x, 2x)
5. System records corresponding calories and macros
6. User receives confirmation of logged meal

#### UC-03: Log Pre-configured Workout Package
**Main Flow:**
1. User taps "Add Workout" button
2. Views workout packages aligned with their goal
3. Selects a workout package (e.g., "30-min HIIT Session")
4. Confirms completion or modifies duration/intensity
5. System calculates calories burned
6. User receives confirmation of logged workout

#### UC-04: View Calorie Balance Dashboard
**Main Flow:**
1. User opens app home screen
2. Views daily calorie summary (target, consumed, burned, remaining)
3. Checks weekly trend chart
4. Taps on specific date to view detailed logs

## 3. Feature Specifications

### 3.1 User Registration & Setup

#### 3.1.1 User Registration/Login
- Support email registration/login
- Support third-party login (Apple ID, Google, Facebook)
- Basic information collection: age, gender, height, weight, activity level

#### 3.1.2 Goal Setting
- Two primary goal paths: Weight Loss and Muscle Gain
- Daily calorie target automatically calculated
- Macro distribution recommendations based on goal
- Option for manual target adjustment

### 3.2 Meal Package System

#### 3.2.1 Meal Package Library
- 20 pre-configured meal packages for Weight Loss goal
- 20 pre-configured meal packages for Muscle Gain goal
- Each package contains complete nutritional information
- Packages categorized by meal type (breakfast, lunch, dinner, snack)

#### 3.2.2 Meal Package Details
- Name and description
- High-quality image
- Calorie content
- Macro breakdown (protein, carbs, fat)
- Portion size options (0.5x, 1x, 1.5x, 2x)
- Preparation instructions (optional link/popup)

#### 3.2.3 Custom Meal Entry
- Basic food search functionality (limited database)
- Manual entry of calories and macros
- Option to save as custom meal package

### 3.3 Workout Package System

#### 3.3.1 Workout Package Library
- 15 pre-configured workout packages for Weight Loss goal
- 15 pre-configured workout packages for Muscle Gain goal
- Each package details estimated calorie burn
- Time-based packages (15, 30, 45, 60 minutes)

#### 3.3.2 Workout Package Details
- Name and description
- Illustrated guide or simple animation
- Estimated calorie burn based on user's weight
- Intensity level selection (light, moderate, intense)
- Duration modification option
- Basic instructions or tips

#### 3.3.3 Custom Workout Entry
- Common exercise selection from limited database
- Duration and intensity input
- Calculated calorie burn
- Option to save as custom workout package

### 3.4 Data Analysis & Display

#### 3.4.1 Main Dashboard
- Daily calorie target
- Daily calorie consumption total
- Daily calorie expenditure total
- Remaining calorie balance
- Goal progress ring chart

#### 3.4.2 Trend Analysis
- Seven-day calorie balance trend chart
- Weekly average calorie surplus/deficit
- Basic macro distribution chart (protein/carbs/fat)

#### 3.4.3 History Records
- Date-based view of historical records
- Support for editing/deleting historical records
- Calendar view for quick navigation

### 3.5 Basic Achievement System

#### 3.5.1 Achievements
- Streak badges (consecutive logging days)
- Milestone badges (first week completed, first month completed)
- Goal-specific achievements (first deficit week, first surplus week)

### 3.6 Settings & Configuration

#### 3.6.1 System Settings
- Push notification controls
- Measurement unit selection (metric/imperial)
- Data export functionality
- Privacy settings

## 4. Non-Functional Requirements

### 4.1 Performance Requirements
- App launch time under 3 seconds
- Package list loading under 1 second
- Local data operations response time under 0.5 seconds

### 4.2 Usability Requirements
- App compatible with iOS 14+ and Android 9.0+
- Compliance with WCAG 2.1 AA accessibility standards
- Support for light/dark mode

### 4.3 Security Requirements
- Encrypted user data storage
- HTTPS protocol for network transmissions
- Compliance with relevant data protection regulations

### 4.4 Reliability Requirements
- App crash rate below 0.5%
- Data synchronization success rate above 99%
- Daily automatic backup functionality

## 5. User Interface

### 5.1 UI Component Library

| Component Name | Description | Priority |
|----------------|-------------|----------|
| Meal Package Card | Displays meal package image, name, calories | High |
| Workout Package Card | Displays workout type, duration, calorie burn | High |
| Calorie Progress Ring | Visualizes daily calorie goal progress | High |
| Portion Size Selector | For adjusting meal package quantities | High |
| Trend Chart Component | For displaying trends and distribution data | Medium |
| Achievement Badge | Displays user achievements | Low |

### 5.2 Core Pages

#### 5.2.1 Home/Dashboard
- Calorie summary card (consumed/burned/remaining)
- Goal progress ring
- Today's logged meals summary
- Today's logged workouts summary
- Quick add buttons (meal/workout)

#### 5.2.2 Meal Package Selection Page
- Goal-filtered package tabs
- Meal type filter (breakfast/lunch/dinner/snack)
- Meal package cards with images
- Portion size selector
- Nutritional preview

#### 5.2.3 Workout Package Selection Page
- Goal-filtered package tabs
- Duration/type filter
- Workout package cards
- Intensity selector
- Calorie burn calculator

#### 5.2.4 History Page
- Date selector/calendar
- Daily logs by meal and workout
- Daily summary card
- Edit/delete options

#### 5.2.5 Profile Settings Page
- Personal information form
- Goal settings options
- System settings entry
- Achievement badge display

### 5.3 Wireframes
[Note: This section would include wireframes for key pages, simplified versions for MVP documentation]

## 6. Data Model

### 6.1 Core Data Structures

#### User Information
```
User {
  userId: String
  name: String
  gender: String
  birthDate: Date
  height: Number
  weight: Number
  activityLevel: Enum
  goal: GoalInfo
  createdAt: Date
  lastLoginAt: Date
}

GoalInfo {
  type: Enum (LOSE/GAIN)
  targetCalories: Number
  targetProtein: Number
  targetCarbs: Number
  targetFat: Number
  startDate: Date
  startWeight: Number (optional)
  targetWeight: Number (optional)
}
```

#### Meal Package Entry
```
MealEntry {
  entryId: String
  userId: String
  packageId: String
  portionMultiplier: Number
  calories: Number
  protein: Number
  carbs: Number
  fat: Number
  mealType: Enum
  date: Date
  createdAt: Date
}
```

#### Workout Package Entry
```
WorkoutEntry {
  entryId: String
  userId: String
  packageId: String
  intensityMultiplier: Number
  durationMinutes: Number
  caloriesBurned: Number
  date: Date
  createdAt: Date
}
```

#### Meal Package Library
```
MealPackage {
  packageId: String
  name: String
  description: String
  goalType: Enum (LOSE/GAIN/BOTH)
  mealType: Enum
  baseCalories: Number
  baseProtein: Number
  baseCarbs: Number
  baseFat: Number
  imageUrl: String
  preparationSteps: Array<String> (optional)
  ingredients: Array<String> (optional)
}
```

#### Workout Package Library
```
WorkoutPackage {
  packageId: String
  name: String
  description: String
  goalType: Enum (LOSE/GAIN/BOTH)
  workoutType: String
  baseDurationMinutes: Number
  baseCaloriesBurn: Number
  caloriesBurnFormula: String
  imageUrl: String
  instructions: Array<String> (optional)
}
```

### 6.2 Data Relationship Diagram
[Note: This section would include an entity relationship diagram, simplified for MVP documentation]

## 7. API Specifications

### 7.1 User API

#### Register User
- **Endpoint**: `/api/users`
- **Method**: POST
- **Request Body**:
  ```json
  {
    "name": "string",
    "email": "string",
    "password": "string",
    "gender": "string",
    "birthDate": "date",
    "height": "number",
    "weight": "number",
    "activityLevel": "string",
    "goal": "string"
  }
  ```
- **Response**: 201 Created

#### Get User Profile
- **Endpoint**: `/api/users/{userId}`
- **Method**: GET
- **Response**: 200 OK

### 7.2 Meal Package API

#### Get Meal Packages
- **Endpoint**: `/api/meal-packages`
- **Method**: GET
- **Parameters**: `goalType`, `mealType`, `limit`, `offset`
- **Response**: 200 OK

#### Log Meal Package
- **Endpoint**: `/api/meal-entries`
- **Method**: POST
- **Request Body**:
  ```json
  {
    "userId": "string",
    "packageId": "string",
    "portionMultiplier": "number",
    "mealType": "string",
    "date": "date"
  }
  ```
- **Response**: 201 Created

#### Get Meal Entries
- **Endpoint**: `/api/meal-entries`
- **Method**: GET
- **Parameters**: `userId`, `startDate`, `endDate`
- **Response**: 200 OK

### 7.3 Workout Package API

#### Get Workout Packages
- **Endpoint**: `/api/workout-packages`
- **Method**: GET
- **Parameters**: `goalType`, `workoutType`, `limit`, `offset`
- **Response**: 200 OK

#### Log Workout Package
- **Endpoint**: `/api/workout-entries`
- **Method**: POST
- **Request Body**:
  ```json
  {
    "userId": "string",
    "packageId": "string",
    "intensityMultiplier": "number",
    "durationMinutes": "number",
    "date": "date"
  }
  ```
- **Response**: 201 Created

#### Get Workout Entries
- **Endpoint**: `/api/workout-entries`
- **Method**: GET
- **Parameters**: `userId`, `startDate`, `endDate`
- **Response**: 200 OK

## 8. Implementation Plan

### 8.1 Development Phases

#### Phase 1: Core Framework (Weeks 1-3)
- User registration and profile setup
- Basic data models and storage
- Simple dashboard UI

#### Phase 2: Package System Implementation (Weeks 4-6)
- Meal package library development
- Workout package library development
- Package logging functionality

#### Phase 3: Analytics & Polish (Weeks 7-8)
- Trend analysis implementation
- Achievement system
- UI refinement and performance optimization

### 8.2 MVP Success Metrics

| Metric | Target |
|--------|--------|
| Daily Active Users / Monthly Active Users | > 0.4 |
| 7-Day Retention Rate | > 30% |
| 30-Day Retention Rate | > 20% |
| Average Daily Logging Sessions | > 2 |
| Users Setting Custom Goals | > 15% |
| Crash-free Users | > 99% |

## 9. Appendices

### 9.1 Sample Meal Packages

#### Weight Loss Meal Packages
1. **Protein Breakfast Bowl** (300 calories)
   - Greek yogurt, berries, low-sugar granola
   - 25g protein, 30g carbs, 7g fat

2. **Lean Chicken Lunch** (400 calories)
   - Grilled chicken breast, mixed veggies, quinoa
   - 35g protein, 40g carbs, 8g fat

3. **Salmon Dinner** (450 calories)
   - Baked salmon fillet, steamed broccoli, sweet potato
   - 30g protein, 35g carbs, 15g fat

#### Muscle Gain Meal Packages
1. **Power Breakfast** (600 calories)
   - Egg whites, oatmeal, banana, protein shake
   - 40g protein, 70g carbs, 10g fat

2. **Muscle Lunch** (700 calories)
   - Lean beef, brown rice, mixed vegetables
   - 45g protein, 80g carbs, 15g fat

3. **Recovery Dinner** (650 calories)
   - Chicken thighs, whole grain pasta, leafy greens
   - 40g protein, 65g carbs, 18g fat

### 9.2 Sample Workout Packages

#### Weight Loss Workout Packages
1. **Fat-Burning HIIT** (30 minutes)
   - High-intensity interval training
   - 8 rounds of 20 seconds max effort, 10 seconds rest
   - Estimated calorie burn: 300-400 calories
   - Intensity levels: Beginner, Intermediate, Advanced

2. **Cardio Blast** (45 minutes)
   - Steady-state cardio with interval bursts
   - Walking/jogging alternating with sprints
   - Estimated calorie burn: 400-500 calories
   - Suitable for all fitness levels

   