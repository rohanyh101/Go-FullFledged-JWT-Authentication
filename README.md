# GoLang JWT Authentication with Gin and MongoDB

This project is a full-fledged JWT authentication system built using GoLang with the Gin framework for HTTP routing and MongoDB for data storage. It provides secure user authentication and authorization functionalities.

## Features

- User Registration
- User login
- JWT token generation and validation
- Password hashing for enhanced security
- MongoDB integration for user data storage
- Secure middleware for authentication endpoints
- Error handling for authentication processes

## Pre Requisites

- GoLang installed on your machine
- MongoDB installed and running
- Basic understanding of GoLang, Gin, and MongoDB

## Installation

1. Clone the repository

2. Navigate to the project directory

4. Install dependencies:
`go mod tidy`

5. Set up your MongoDB connection in the configuration file.

6. Run the application:
`go run main.go`

## Registered Routes
- POST `/users/signin` endpoint to add a user to the database.
- POST `/users/login` endpoint to log in a user and regenerate the user credentials
- GET `/users` endpoint to retrieve all users (accessible by admin).
- GET `/users/:user_id` endpoint to retrieve a specific user by ID (accessible by the user or admin).

## Usage
1. Register a new user using the `users/signup` endpoint.
2. Login using the `users/login` endpoint with the registered user credentials.
4. Upon successful login, you will receive a JWT token and a Refresh Token.
5. Use this token to access protected endpoints by including it in the request headers as follows.

## Configuration

You can configure the application settings in the `.env` file. You can specify the MongoDB connection URI, server port, and other relevant parameters here.

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

This project is licensed under the [MIT License](LICENSE).


