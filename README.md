## Event Management Backend

A simple backend application built with Go, using the Gin framework and SQLite database. It supports user authentication with JWT and CRUD operations for events, along with event registration functionality.
Features

- User signup and login with JWT-based authentication
- Create, retrieve, update, and delete events
- Register users to existing events
- Authentication middleware for protected routes
- Built with Gin framework and SQLite database

### Prerequisites

- Go (version 1.16 or later)
- SQLite
- Git

### Install dependencies:

go mod tidy

### Run the application:

`go run main.go`

### API Endpoints

- Authentication

```
POST /signup - Register a new user
POST /login - Login and receive a JWT token
```

- Events

```
POST /events - Create a new event (requires authentication)
GET /events - Retrieve all events
GET /events/:id - Retrieve a specific event
PUT /events/:id - Update an event (requires authentication)
DELETE /events/:id - Delete an event (requires authentication)
POST /events/:id/register - Register a user to an event (requires authentication)
```

### Project Structure

```
├── main.go # Entry point of the application
├── handlers/ # Request handlers for API endpoints
├── models/ # Data models and database schema
├── middleware/ # Authentication middleware
├── db/ # Database setup and connection
├── utils/ # utils for hash password and jwt
├── go.mod # Go module dependencies
└── go.sum # Dependency checksums
```

### Dependencies

Gin - HTTP web framework
jwt-go - JWT authentication

### Usage

- Use a tool like Postman or cURL to interact with the API.
- Authenticate using /signup or /login to obtain a JWT token.
- Include the token in the Authorization header (Bearer <token>) for protected routes.

### Example

- Signup `curl -X POST http://localhost:8080/signup -d '{"username":"user","password":"pass"}'`

- Login `curl -X POST http://localhost:8080/login -d '{"username":"user","password":"pass"}'`

- Create Event

```
curl -X POST http://localhost:8080/events -H "Authorization: Bearer <token>" -d '{
  "name": "Tech Meetup 3",
  "description": "3 Discuss latest trends in AI and Golang",
  "location": "3 Bangalore Hub",
  "dateTime": "2025-09-05T10:30:00+05:30"
  }'
```
