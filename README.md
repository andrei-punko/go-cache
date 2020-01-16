# Juno Inc., Test Task: In-memory cache
###### Task definition taken from [here](https://github.com/gojuno/test_tasks)

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

# Build instructions

##### Build for definite OS and architecture on Linux

    env GOOS=linux GOARCH=amd64 go build -o ./out/linux-amd64/web-cache ./web_server.go  
    env GOOS=windows GOARCH=amd64 go build -o ./out/windows-amd64/web-cache.exe ./web_server.go

##### Build for definite OS and architecture on Windows

    set GOOS=linux  
    set GOARCH=amd64  
    go build -o ./out/linux-amd64/web-cache ./web_server.go

    set GOOS=windows  
    set GOARCH=amd64  
    go build -o ./out/windows-amd64/web-cache.exe ./web_server.go

##### Build using Gradle

    ./gradlew goBuild

Result binaries will be placed into `./gogradle` folder

##### Put cache application into Docker image:

    docker build -t apunko/go-cache .

# Usage description

##### Start cache application using default port 8000:

    go run ./web_server.go

##### Start cache application using user defined port:

    go run ./web_server.go 8005

##### Start Docker image:

    docker run --rm -p 8000:8000 apunko/go-cache

#### Or using docker-compose:

    docker-compose up

##### Cache population with some key-value pairs:

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -d '{"value": "Ivan", "ttl": 60000000000}' -X POST http://localhost:8000/items/name

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -d '{"value": ["VISA", "Mastercard"], "ttl": 60000000000}' -X POST http://localhost:8000/items/cards

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -d '{"value": {"Math": "9", "English": "7"}, "ttl": 60000000000}' -X POST http://localhost:8000/items/marks

##### Getting value by key from cache:

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8000/items/name

##### Getting all keys from cache:

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8000/items/keys

##### Deletion of some key from cache:

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X DELETE http://localhost:8000/items/name

##### Deletion of all keys from cache:

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X DELETE http://localhost:8000/items/keys
