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

## Technologies and Frameworks
- Backend language: Go (1.20)
- SDK: temporal.io
- Clean architecture
- Network: gRPC, HTTP
- Docker, docker compose
- Database: postgres

## Directory structure
    📁 aspire-code-challenge
    |__ 📁 api // adapter layer, write endpoints
    |__ 📁 cmd
        |__ 📁 api
        |__ 📁 worker
    |__ 📁 config
        |__ .env
        |__ app.go // app config
        |__ db.go // config for database
        |__ temporal.go config for temporal client
    |__ 📁 internal 
        |__ 📁 activities // temporal activities
        |__ 📁 app // application layer
        |__ 📁 models // entitiy layer
        |__ 📁 repositories // database repositories
        |__ 📁 workflows // temporal workflows
        |__ 📁 defined // types and constants
    |__ 📁 pkg
        |__ 📁 clients // temporal client, and others
        |__ 📁 logger
        |__ 📁 persistence // database
        |__ 📁 proto // define grpc proto and http transcoding
    |__ 📁 scripts // scripts for init database, others
    |__ 📁 infra // temporal workers
        |__ Dockerfile-api // for run api
        |__ Dockerfile-worker // for run worker
    |__ docker-compose.yml
    |__ go.mod
    |__ go.sum
    |__ README.md

## How to run local
- Run services and workers: `docker compose up --build -d`. 
If `aspire-api` and `aspire-worker` not running, maybe initialize not done yet, please 
re-run the docker compose command again.
- Generate proto: `make`
- TODO: run services and worker separately for scaling

## How to test
- HTTP: Install postman.
  - Import `aspire-code-challenge.json`
  - Host: localhost
  - Port: 9001
- gRPC: Install postman or BloomRPC
  - Host: localhost
  - Port: 8001

## Monitoring
- Workflow status: http://localhost:8080/namespaces/aspire-code-challenge/workflows

## Conceptual model
![img.png](docs/aspire-code-challenge-conceptual-model.png)

## TODO tasks
- Add a policy check to make sure that the customers can view them own loan only.
- Implement more unit tests
- Implement exceed amount repayment
- Fix docker network and wait dependencies healthy
- Implement error abstract


## References
- Project layout: https://github.com/golang-standards/project-layout
- Transcoding of HTTP/JSON to gRPC: https://adevait.com/go/transcoding-of-http-json-to-grpc-using-go
- Workflow framework: https://docs.temporal.io/
- Clean architecture: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

