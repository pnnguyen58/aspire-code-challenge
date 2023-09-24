# aspire-code-challenge

## Description
- The purpose of source to implement a mini loan application
- It allows authenticated users to go through a loan application
- After the loan is approved, the user must be able to submit the weekly loan repayments. 
It can be a simplified repay functionality, which won’t need to check if the dates are correct but will just set the weekly amount to be repaid.


## Prerequisites
1. Golang
2. https://docs.temporal.io/
3. Docker, docker compose
4. gRPC and http
5. Clean architecture

## How to run local
- Run services and workers: `docker compose up -d`
- Generate proto: `make`
- TODO: run services and worker separately for scaling

## Monitoring
- Workflow status: http://localhost:8080/namespaces/aspire-code-challenge/workflows

## References
- Project layout: https://github.com/golang-standards/project-layout
- Transcoding of HTTP/JSON to gRPC: https://adevait.com/go/transcoding-of-http-json-to-grpc-using-go
- Workflow framework: https://docs.temporal.io/
- Clean architecture: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

