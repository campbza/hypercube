# hypercube
Golang project to learn about concurrency. We start by prompting the user for a power of two. Call 
this number `n` 
and create a hypercube with `n` nodes. We create `n` channels, and generate a `struct` called a 
`packet` which contains source node and destination node information. We randomly generate the 
sources and destinations, and concurrently forward these packets along the edges of the 
hypercube until they reach the destination node.

Read more about goroutines here: https://golangbot.com/goroutines/  

Usage:

`go build hypercube.go`

`go run hypercube.go` 
