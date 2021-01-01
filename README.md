# Locale Assignment
## Dependencies
 - Redis - Redis must be installed and running as the task queue depends on redis.
 - PostgreSQL - PostgreSQL is the persistance layer of choice to store data.

## Directory Structure

- The `cmd/` directory contains the main binary for the server.
- The `pkg/` directory contains all the libray code.
- `writeup.pdf` contains the solution writeup.

## Code Structure
The code structre is based upon very popular [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

- Entities are the blueprint of the data (model).
- Service is the layer that abstracts away business logic.
- Handler handles connections, in this case we have only HTTP handlers but the functionality can be extended to support other protocols easily.
- Repository is the layer that abstracts away the persistance layer. We can easily switch our persistance layer from PostgreSQL to something like MongoDB without breaking other layers.

## Starting The Server

- Clone this repository and `cd` into project root.
- Use `go run cmd/api/main.go`

**Note:** Make sure that all the above dependencies are installed and running.

**Note:** Make sure create the `.env` file in the project root and popluate it with the environment variables. The sample file is present in `env`.
