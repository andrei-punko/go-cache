# Test Task (Juno Inc.): In-memory cache
[![Go CI with Gradle](https://github.com/andrei-punko/go-cache/actions/workflows/go.yml/badge.svg)](https://github.com/andrei-punko/go-cache/actions/workflows/go.yml)

Task definition taken from [here](https://github.com/gojuno/test_tasks)

## Description

Simple implementation of Redis-like in-memory cache

Desired features:
- Key-value storage with string, lists, dict support
- Per-key TTL
- Operations:
  - Get
  - Set
  - Remove
- Golang API client
- Provide some tests, API spec, some deployment manual, some examples of client usages would be nice.

Optional features:
- Telnet-like/HTTP-like API protocol
- performance tests
- Operations:
  - Keys

## Build using Gradle
```bash
./gradlew goBuild
```

Result binaries will be placed into `./gogradle` folder

## Put application into Docker image:
```bash
docker build -t apunko/go-cache .
```

## Start application

### Start application on default port 8000:
On Linux OS:
```bash
./.gogradle/linux_amd64_go-cache
```

On Win OS:
```bash
./.gogradle/windows_amd64_go-cache.exe
```

### Start application on user defined port (8005 for example):
On Linux OS:
```bash
./.gogradle/linux_amd64_go-cache 8005
```

On Win OS:
```bash
./.gogradle/windows_amd64_go-cache.exe 8005
```

### Start Docker container using `docker` command:
```bash
docker run --rm -p 8000:8000 apunko/go-cache
```

### Or start Docker container using `docker-compose` command:
```bash
docker-compose up
```

## Interact with application

### Cache population with some key-value pairs:
```bash
curl -i -X POST -H "Accept: application/json" -H "Content-Type: application/json" -d '{"value": "Ivan", "ttl": 60000000000}' http://localhost:8000/items/name
```

```bash
curl -i -X POST -H "Accept: application/json" -H "Content-Type: application/json" -d '{"value": ["VISA", "Mastercard"], "ttl": 60000000000}' http://localhost:8000/items/cards
```

```bash
curl -i -X POST -H "Accept: application/json" -H "Content-Type: application/json" -d '{"value": {"Math": "9", "English": "7"}, "ttl": 60000000000}' http://localhost:8000/items/marks
```

### Getting value by key=name from cache:
```bash
curl -i http://localhost:8000/items/name
```

### Getting all keys from cache:
```bash
curl -i http://localhost:8000/items/keys
```

### Deletion of some key-value pair from cache for key=name:
```bash
curl -i -X DELETE http://localhost:8000/items/name
```

### Cache cleanup:
```bash
curl -i -X DELETE http://localhost:8000/items/keys
```
