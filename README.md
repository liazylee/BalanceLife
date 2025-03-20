# BalanceLife API

A calorie tracking application backend built with Go, Gin, MongoDB, and Redis.

## Overview

BalanceLife is a calorie tracking application that helps users track their meals and workouts using a package-based approach. The application provides pre-configured meal and workout packages that users can log with customizable portion sizes and intensity levels.

## Features

- User account management
- Pre-configured meal and workout packages
- Meal and workout tracking
- Calorie tracking with daily targets
- MongoDB for persistent storage
- Redis for caching and improved performance
- RESTful API
- Swagger/OpenAPI documentation

## Getting Started

### Prerequisites

- Go 1.16 or higher
- MongoDB (required)
- Redis (optional, improves performance if available)
- Docker and Docker Compose (for containerized setup)

### Installation

#### Option 1: Running Locally (Without Docker)

If you have MongoDB and Redis instances hosted online or elsewhere, you can run the application directly:

1. Clone the repository:
```bash
git clone https://github.com/zhenyili/BalanceLife.git
cd BalanceLife
```

2. Configure your database connections:
   - Edit `config/.env` file with your MongoDB and Redis connection details
   - Update `config/app.config.json` with the same information

3. Initialize the database:
```bash
./run-local.sh init
```

4. Start the application:
```bash
./run-local.sh run
```

The API will be available at http://localhost:8080.
The Swagger documentation will be available at http://localhost:8080/swagger/index.html.

#### Option 2: Quick Setup (Docker)

1. Clone the repository:
```bash
git clone https://github.com/zhenyili/BalanceLife.git
cd BalanceLife
```

2. Run the setup script:
```bash
./setup.sh
```

3. Start the API in development or production mode:
```bash
./run.sh dev   # For development
# OR
./run.sh prod  # For production
```

#### Option 3: Docker Setup (Detailed)

1. Clone the repository:
```bash
git clone https://github.com/zhenyili/BalanceLife.git
cd BalanceLife
```

2. Configure environment settings (optional):
   - Edit `config/.env` with your database credentials
   - Modify `config/app.config.json` if needed

3. Start the API with the environment of your choice:
```bash
./run.sh [environment]
```

Available environments:
- `dev` - Development environment with hot reloading
- `prod` - Production environment with optimized build

To stop all containers:
```bash
./run.sh stop
```

For more options:
```bash
./run.sh help
```

#### Option 4: Manual Setup

1. Clone the repository:
```bash
git clone https://github.com/zhenyili/BalanceLife.git
cd BalanceLife
```

2. Use the Makefile to set up your development environment:
```bash
make setup
```

3. Edit the configuration files in the `config` directory:
   - Update `config/.env` with your database credentials
   - Set appropriate values in `config/app.config.json`

4. Initialize MongoDB with sample data (optional):
```bash
make init-db
```

5. Run the application:
```bash
make run
```

The server will start on port 8080 by default. You can change this in your configuration files.

### Makefile Commands

The project includes a Makefile with useful commands:

- `make run` - Run the API server
- `make build` - Build the application
- `make clean` - Clean build artifacts
- `make test` - Run tests
- `make init-db` - Initialize sample data in MongoDB
- `make config` - Generate config files from templates
- `make deps` - Install dependencies
- `make setup` - Setup development environment
- `make help` - Show help message

## Docker Environments

### Development Environment

The development environment includes:
- Go application with hot reloading using Air
- MongoDB for data persistence 
- Redis for caching
- Source code mounted as a volume for live editing

```bash
./run.sh dev
```

### Production Environment

The production environment includes:
- Optimized Go application build
- MongoDB for data persistence
- Redis for caching
- Restart policies for reliability

```bash
./run.sh prod
```

## Configuration

The application supports configuration through:

1. Environment variables
2. Configuration file (`config/app.config.json`)
3. Command-line arguments

### Environment Variables

- `SERVER_PORT`: HTTP server port (default: 8080)
- `MONGODB_URI`: MongoDB connection string
- `MONGODB_DATABASE`: MongoDB database name (default: balancelife)
- `MONGODB_CERT_PATH`: Path to MongoDB certificate file
- `REDIS_ADDR`: Redis server address
- `REDIS_PASSWORD`: Redis password
- `INIT_SAMPLE_DATA`: Set to "true" to initialize sample data on startup

## API Documentation

### Swagger/OpenAPI Documentation

BalanceLife provides interactive API documentation using Swagger/OpenAPI. You can access it at:

```
http://localhost:8080/swagger/index.html
```

This interactive documentation allows you to:
- View all available endpoints
- Read detailed descriptions of each API operation
- Explore request and response models
- Test the API directly from your browser

### Health Check

```
GET /health
```

Returns the status of the API server, including MongoDB and Redis connection status.

### User Management

#### Get All Users

```
GET /api/users
```

Returns all users in the system.

#### Get User by ID

```
GET /api/users/:id
```

Returns a specific user by ID.

#### Create User

```
POST /api/users
```

Creates a new user.

**Request Body:**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword",
  "gender": "MALE",
  "birthDate": "1990-01-01",
  "height": 180.0,
  "weight": 80.0,
  "activityLevel": "MODERATE",
  "goal": "LOSE"
}
```

### Meal Packages

#### Get All Meal Packages

```
GET /api/meals/packages
```

Returns all meal packages. Can be filtered by goal type using the `goalType` query parameter.

#### Get Meal Package by ID

```
GET /api/meals/packages/:id
```

Returns a specific meal package by ID.

### Meal Entries

#### Create Meal Entry

```
POST /api/meals/entries
```

Logs a meal entry for a user.

**Request Body:**

```json
{
  "userId": "usr1",
  "packageId": "meal1",
  "portionMultiplier": 1.0,
  "date": "2023-03-18"
}
```

#### Get Meal Entries

```
GET /api/meals/entries?userId=usr1&startDate=2023-03-01&endDate=2023-03-18
```

Returns meal entries for a user within a date range.

### Workout Packages

#### Get All Workout Packages

```
GET /api/workouts/packages
```

Returns all workout packages. Can be filtered by goal type using the `goalType` query parameter.

#### Get Workout Package by ID

```
GET /api/workouts/packages/:id
```

Returns a specific workout package by ID.

### Workout Entries

#### Create Workout Entry

```
POST /api/workouts/entries
```

Logs a workout entry for a user.

**Request Body:**

```json
{
  "userId": "usr1",
  "packageId": "workout1",
  "intensityMultiplier": 1.0,
  "durationMinutes": 30,
  "date": "2023-03-18"
}
```

#### Get Workout Entries

```
GET /api/workouts/entries?userId=usr1&startDate=2023-03-01&endDate=2023-03-18
```

Returns workout entries for a user within a date range.

## Data Storage Architecture

The application uses a multi-tier storage approach:

1. **Redis** - Fast in-memory cache for frequently accessed data (optional)
2. **MongoDB** - Persistent document storage for all application data (required)

### MongoDB Collections

- `users` - User profiles and account information
- `meal_packages` - Pre-configured meal package templates
- `workout_packages` - Pre-configured workout package templates
- `meal_entries` - User-logged meal records
- `workout_entries` - User-logged workout records

### Redis Cache Structure

- User profiles: `user:{userId}`
- Meal packages: `meal_package:{packageId}`
- Workout packages: `workout_package:{packageId}`
- Meal entries by date range: `meal_entries:{userId}:{startDate}:{endDate}`
- Workout entries by date range: `workout_entries:{userId}:{startDate}:{endDate}`

## Future Plans

- Authentication and authorization
- User profile management
- Expanded meal and workout package libraries
- Custom meal and workout entry support
- Analytics and progress tracking
- Goal adjustment capabilities

## Development

The application follows a standard Go project structure:

- `src/cmd/api`: Main application entry point
- `src/models`: Data models
- `src/handlers`: HTTP handlers for API routes
- `src/db`: Data storage implementations (MongoDB, Redis)
- `src/utils`: Utility functions
- `src/config`: Configuration management
- `scripts`: Utility scripts for database setup
- `docs`: Swagger/OpenAPI documentation

## Security Notes

- Configuration files containing sensitive data are excluded from version control
- MongoDB connection uses TLS with certificate authentication
- Redis connection is secured with password authentication
- Environment variables are used for sensitive configuration

## License

This project is licensed under the MIT License - see the LICENSE file for details.