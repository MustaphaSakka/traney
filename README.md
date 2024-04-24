TRANEY
------

TRANEY (Transfert-Money) is a web application that handles money transfert from client to client through simple transactions.

This project is based on HEXAGONAL architecture to isolate the core application from the outside dependencies through ports and adapters.

### Project HEXAGONAL architecture

- `app` : routes, handlers and starting server
- `domain` : business side, core of application 
- `dto` : objects that transfer data between the different layers of application
- `resources` : docker configuration and initialization of the database
- `service` : services that interact with domain

### USEFUL COMMANDS

- Start project: `go run main.go`
- Generate mocks before tests: `./generate-mocks.sh`
- Run unit tests: `./run-tests.sh`
