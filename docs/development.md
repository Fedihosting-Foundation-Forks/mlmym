# Development Guide

## Prerequisites

- Go 1.23.8 or later
- Basic understanding of Go web applications

## Project Structure

- `main.go` - HTTP server, middleware, templates
- `routes.go` - Handlers, Lemmy API integration  
- `state.go` - User sessions, preferences
- `templates/` - HTML templates
- `public/` - CSS, JS, assets
- `routes_test.go` - Link rewriting tests

## Dependencies

- `httprouter` - HTTP routing
- `go-lemmy` - Lemmy API client
- `goldmark` - Markdown processing
- `html/template` - Template rendering

## Development Commands

```bash
# Build the binary
make mlmym
# or
go build -v -o mlmym

# Run tests
go test
go test -v

# Development mode with hot reload
go run . --addr 0.0.0.0:8008 -w

# Make commands
make dev     # Start development environment
make serve   # Run server with file watching (port 8008)
make reload  # Watch for template/asset changes (port 8009)
```

## Key Features

- Templates cached in production, reloaded with `-w` flag
- Multi-instance vs single-instance modes via LEMMY_DOMAIN
- Link rewriting for Lemmy URLs
- Image handling for pictrs/imgur
- Session management via cookies

## Configuration

For configuration options including deployment modes, environment variables, and usage examples, see the [Configuration section in README.md](../README.md#️-configuration).
