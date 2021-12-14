# Go REST API - Bucket list - built with go-chi, Docker, and PostgreSQL

### Requirements
* Docker and Go
* [golang-migrate/migrate](https://github.com/golang-migrate/migrate) 

### Usage
Clone the repository with:
```bash
git clone github.com/zurcacielos/bucketeer
```

Copy the `env.example` file to a `.env` file.
```bash
$ cp .env.example .env
```
Update the postgres variables declared in the new `.env` to match your preference. 
There's a handy guide on the [Postgres' DockerHub](https://hub.docker.com/_/postgres).

Build and start the services with:
```bash
$ docker-compose up --build
```

## Test
###POST
```bash
curl -X POST http://localhost:8080/items -H "Content-type: application/json" -d '{ "name": "swim across the River Benue", "description": "ho ho ho"}'
```

```bash
curl -X POST http://localhost:8080/items -H "Content-type: application/json" -d '{ "name": "read all the books in the world", "description": "get an ebook reader"}'
```

###GET
```bash
curl http://localhost:8080/items/1
```

The database migration files are in `db/migrations` and will be run by composer up. 
Beside that feel free to simply source them directly. Alternatively, you can apply them using `migrate` by running:
```bash
$ export POSTGRESQL_URL="postgres://$PG_USER:$PG_PASS@localhost:5432/$PG_DB?sslmode=disable"
$ migrate -database ${POSTGRESQL_URL} -path db/migrations up
```
_**NOTE:** Remember to replace the `$PG*` variables with their actual values_
### Development
After making your changes, you can rebuild the `server` service by running the commands below
```bash
$ docker-compose stop server
$ docker-compose build server
$ docker-compose up --no-start server
$ docker-compose start server
```
