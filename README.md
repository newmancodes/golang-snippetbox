# Let's Go by Alex Edwards

[![CI](https://github.com/newmancodes/golang-snippetbox/actions/workflows/main.yml/badge.svg)](https://github.com/newmancodes/golang-snippetbox/actions/workflows/main.yml)

Working through [Alex's book](https://lets-go.alexedwards.net/) to develop my knowledge and experience with Golang.

I've done a few experiments in the past but continue to be curious about this language.

## Getting Started

I'm using v1.25.4 of Go, so make sure you have that installed from [https://go.dev/doc/install](go.dev) then execute `go run ./cmd/web` in your terminal from the root of the repository.

## Notable Deviations

- I've decided to go with PostgreSQL rather than the suggested use of mysql and am using docker to host the database. There is a docker-compose.yml file in the root of this repository, use `docker compose -f ./docker-compose.yml up` to pull the image and kick start the database before running the application
  - Used [golang-migrate](https://github.com/golang-migrate/migrate) to manage the database schema. Check out the [freecodecamp tutorial](https://www.freecodecamp.org/news/database-migration-golang-migrate/) for more details.
