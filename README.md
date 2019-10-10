# Juno Inc., Test Task: In-memory cache

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

---

## In-memory web-cache

### Build for definite OS and architecture
`env GOOS=linux GOARCH=amd64 go build -o ./out/linux-amd64/web-cache ./main/web_server.go`

`env GOOS=windows GOARCH=amd64 go build -o ./out/windows-amd64/web-cache.exe ./main/web_server.go`

### Description of usage

#### Start cache using default port 8000
`go run ./main/web_server.go`

#### Start cache using user defined port
`go run ./main/web_server.go 8005`

#### Cache population with some key-value pairs
`curl -i -H "Accept: application/json" -H "Content-Type: application/json" -d '{"value": "Ivan", "ttl": 60000000000}' -X POST http://localhost:8000/items/name`

`curl -i -H "Accept: application/json" -H "Content-Type: application/json" -d '{"value": "27", "ttl": 60000000000}' -X POST http://localhost:8000/items/age`

`curl -i -H "Accept: application/json" -H "Content-Type: application/json" -d '{"value": "80.5", "ttl": 60000000000}' -X POST http://localhost:8000/items/weight`

#### Getting value by key from cache
`curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8000/items/name`

#### Getting all keys from cache
`curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8000/items/keys`

#### Deletion of some key from cache
`curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X DELETE http://localhost:8000/items/name`

#### Deletion of all keys from cache
`curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X DELETE http://localhost:8000/items/keys`
