# goclean-boiler-plate
Boiler plate for go follow Clean architecture.

## Motivation
This is not Go example 101 about how to use Golang with _JWT_, _Rethinkdb_, _sendgrid_...
But instead about how to put all of those pieces in to a complete picture that use Clean architecture.

The target of this project are:
1. An example about architecture and how to integrate with other component.
2. A boiler plate to help creating a back-end REST api service quickly.

I myself use this service to make it possible to implement my idea in weekend.

## Component
1. Authentication: **JWT**
2. Database: **Rethinkdb**
3. Middleware
4. Mail sending service: **SendGrid**
5. Unit test & Integration test

## API
```
// Register with email
POST /auth/registerbyemail

// Login by email & password
POST /auth/login

// Request reseting pass
POST /auth/reqresetpass

// Reset pass with token
POST /auth/resetpass

// Access resource with token authentication
GET /users/{userId}
```

## Run Test
To test all unit test
```bash
go test goclean/... -short
```

To test full integration test with rethinkdb
```bash
rethinkdb --http-port 8000
go test goclean/... -short
```
