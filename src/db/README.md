# BalanceLife Database Architecture

This directory contains the database implementation for the BalanceLife application. The architecture follows a multi-tier approach with different storage options.

## Architecture Overview

The database layer is built with flexibility and performance in mind:

1. **Primary Storage**: MongoDB for persistent data storage
2. **Cache Layer**: Redis for fast access to frequently requested data
3. **Fallback Storage**: In-memory store for development or when databases are unavailable

## Components

### Store Interface (`store.go`)

The core interface that all storage implementations must satisfy. It defines methods for:
- User management
- Meal package retrieval
- Workout package retrieval
- Meal entry management
- Workout entry management

### MongoDB Integration (`mongodb.go`)

Provides persistent storage using MongoDB:
- Collection management for users, meal packages, workout packages, etc.
- BSON tagging for proper data mapping
- Index creation for performance optimization
- Sample data initialization

### Redis Caching (`redis.go`)

Implements caching for frequently accessed data:
- User data caching
- Package data caching (meals and workouts)
- User activity entries caching
- Automatic cache invalidation on writes

### CachedStore (`cached_store.go`)

Combines MongoDB and Redis:
- Uses the cache-aside pattern: check cache first, then database
- Graceful fallback if Redis is unavailable
- Cache updates on database reads for future requests
- Automatic cache invalidation for consistency

### Memory Store (`memory_store.go`)

For development and fallback only:
- In-memory implementation of the Store interface
- Contains sample data for development and testing
- Used when MongoDB is unavailable
- **IMPORTANT**: Data is not persisted between application restarts

## Usage

In normal operation, the application will:
1. Attempt to connect to MongoDB and Redis
2. If both succeed, use the `CachedStore` implementation
3. If Redis fails but MongoDB succeeds, use `CachedStore` with MongoDB only
4. If MongoDB fails, fall back to `MemoryStore`

## Configuration

Database connection settings are managed through environment variables or the configuration file:

- `MONGODB_URI`: MongoDB connection string
- `MONGODB_DATABASE`: Database name
- `MONGODB_CERT_PATH`: Optional TLS certificate path
- `REDIS_ADDR`: Redis server address
- `REDIS_PASSWORD`: Redis password
- `REDIS_DB`: Redis database number

See `src/config` for more details on configuration.

## Sample Data

Sample data initialization happens in two ways:
1. Automatically in memory store mode
2. Via `InitSampleData()` method when using MongoDB

Use the `INIT_SAMPLE_DATA=true` environment variable to load sample data on startup when using MongoDB. 