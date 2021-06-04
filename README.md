# go-http-server

HTTP server for static files using golang

[![CodeQL](https://github.com/guionardo/go-http-server/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/guionardo/go-http-server/actions/workflows/codeql-analysis.yml)

[![Publish Docker image](https://github.com/guionardo/go-http-server/actions/workflows/go-http-server-docker.yml/badge.svg)](https://github.com/guionardo/go-http-server/actions/workflows/go-http-server-docker.yml)

## CLI usage

For serve files into ./html folder thru 8080 port, use:

```bash
go-http-server --port 8080 --folder ./html
```

## Environment variables

- PORT
- FOLDER

## Docker-compose sample

```yaml
version: "3"
services:
  http-server:
    build: .
    ports:
      - "8000:8000"
    environment:
      - PORT=8000
      - FOLDER=/html
      - GIN_MODE=release
    volumes:
      - ./html:/html
```
