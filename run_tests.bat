rem Run usual tests
go test ./... -v

rem Run benchmark tests
go test -bench=. ./... > benchmark.out
