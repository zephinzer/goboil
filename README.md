# Golang Tooling Boilerplate
An opinionated tooling boilerplate for Golang using Docker.

# Scope
## Technical
- [x] Start application
- [x] Start application with live-reloading
- [x] Environment configuration
- [x] Test application with coverage
- [x] Test application with live-reloading
- [x] Compile binary
- [x] Create production Docker image bundle

# Development
Run `make start` to start the application in development.

# Configuration
Use the `./.env` file to configure the environment

# Testing
Run `make test` to run the tests once.

Run `make testc` to run the tests once and output the coverage.

Run `make testw` to run the tests in automated live-reload mode.

# Building
Run `make build` to create the binary.

# Deployment
Either use the generated binary from `make build`, or create a production image by running `make dkbuild.prd` to create the production image.
