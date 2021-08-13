# Library Management System - User Service
This responsibility of this service to provide endpoints to perform CRUD operations on user models.  It only accepts requests from the Node.js authServer application.

## Tools Used For This Project
---
* Golang
* SQL
* JSON for requests and responses
* Gorilla Mux

## Project Structure
---
* A standard layered structure.
  * The handler functions are in the main file.
  * The handler functions interact with the service layer and the service layer interacts with the data access layer.