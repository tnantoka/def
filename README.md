# Def

## Development

```
$ docker compose up
$ docker compose exec app bash

# Format
$ go fmt ./...

# Test
$ go test ./... -coverprofile=cover.out && \
  go tool cover -html=cover.out -o coverage.html

# REPL
$ go run main.go
```

## Acknowledgements

- https://interpreterbook.com/
- https://compilerbook.com/
