# Nearby Shops REST API

The Nearby Shops REST API is designed to help users find shops in their vicinity. It provides endpoints for retrieving information about shops, including their names, locations, and other relevant details. The API leverages geospatial data to provide accurate and efficient location-based queries.

## Tech Stack

- Go 1.22
- Chi (HTTP Router)
- SQLx (database library)
- Viper (configuration management)
- Air (hot reload for development)
- PostgreSQL (database)
- PostGIS (geospatial extension for PostgreSQL)

## Docker Compose

To run the project with Docker Compose, follow these steps:

1. Start the PostgreSQL Docker container by running:

```bash
docker-compose up database
```

2. Run database migrations and seed data by executing:

```bash
go run cmd/migrations/up/main.go && go run cmd/seeder/main.go
```

3. Finally, launch the application using Docker Compose:

```bash
docker-compose up app
```

This will start the application in a Docker container, allowing you to interact with the REST API. You can interact with the API by navigating to [http://localhost:8080](http://localhost:8080).

## Development Mode

To run the project in development mode, follow these steps:

1. Create a file named `app.env` and populate it with the contents from `app.env.example`. Alternatively, you can copy the contents using the following command:

```bash
cp app.env.example app.env
```

2. Make sure you have installed air for hot reload. Refer to [air docs](https://github.com/cosmtrek/air) for installation instructions.
3. Run the following command to start the application with air:

```bash
air
```

This will launch the application with hot reload enabled, allowing you to make changes to the code and see them reflected in real-time without restarting the server.

## Swagger Documentation

This project includes Swagger documentation to help you understand and interact with the API endpoints. You can access the Swagger UI by navigating to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html). The Swagger UI provides a user-friendly interface for exploring the API, viewing endpoint details, and making test requests.
