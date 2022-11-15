# Pintap Recruitment

This requirement from https://docs.google.com/document/d/18Yt-YBxnDJyuTjotIVaWJdBhpEeH3_89/edit?usp=sharing&ouid=110539423922657302860&rtpof=true&sd=true

## Summary
This app provides REST API for todo list stuff, This Rest API including simple CRUD (Create Read Update Delete) with unit testing, migration, and containerization using docker file.

The kit provides the following features right out of the box:

* RESTful endpoints in the widely accepted format
* Standard CRUD operations of a database table
* Environment dependent application configuration management
* Structured logging with contextual information
* Error handling with proper error response generation
* Database migration
* Data validation
* Full test coverage

The kit uses the following Go packages which can be easily replaced with your own favorite ones
since their usages are mostly localized and abstracted. 

* Routing: [httprouter](https://github.com/julienschmidt/httprouter)
* Database access: [postgres](gorm.io/driver/postgres)
* Database migration: [golang-migrate](https://github.com/golang-migrate/migrate)
* Data validation: [validator](https://github.com/go-playground/validator/)
* Environment: [env](https://github.com/joho/godotenv)
* Testing: [testify](https://github.com/stretchr/testify)
* Gorm [gorm](https://gorm.io/)

## REST API Career
- Find all list careers
- Get detail career by ID

## Installation
If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires **Go 1.13 or above**.

[Docker](https://www.docker.com/get-started) is also needed if you want to try the kit without setting up your
own database server. The kit requires **Docker 17.05 or higher** for the multi-stage build support.



# download the starter kit
git clone https://github.com/arsyadarmawan/manipulation-api.git

cd todo-list

# migrate a postgres and create your database first then run this command
Install migration for golang, 
```go install "postgres" https://github.com/golang-migrate/migrate/v4/cmd/migrate@latest```.

 Then after installed you just run this command 
```migrate -database "postgres://user:password@host:port/dbname?query" -path db/migrations up```. 

Example
```migrate -database "postgres://admin:admin123@localhost:5432/new-app?sslmode=disable" -path db/migrations up```

# Copy .env.example to .env
```cp .env.example .env```

# Fill in the blank .env credentials and this is for example
```
APPNAME=Cart
APP_URL=localhost:3002
DB_HOST=localhost
DB_PORT=2022
DB_NAME=todo
DB_USERNAME=postgres
DB_PASSWORD=admin123

```



## REST API Detail Documentation postman on api folder
* `POST /api/auth/register`: register users. Here the example

```
curl --location --request POST 'localhost:3000/api/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name" : "admin",
    "password" : "admin123",
    "username" : "admin"
}'
```

* `Login /api/auth/login`: Login users with regitered account.

```
curl --location --request POST 'localhost:3000/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username" : "admin",
    "password" : "admin123"
}'
```

* `Fetching Recruitment API /api/dans/career`: Fetching API. you can add location, full_time value true or false, description, and page with value of number.

```
curl --location --request GET 'localhost:3000/api/dans/career?full_time=false&location=Remote&description=Murphy&page=1' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwibmFtZSI6InRoYXJlZXEiLCJ1c2VybmFtZSI6InRoYXJlZXEiLCJleHAiOjE2Njg3Nzg5MTZ9.EyN-WmRvhofuptT2jHY-dx9VWD1xU09lgKIQNr95nPg'
```

* `Fetching Recruitment API /api/dans/career/:id`: Fetching data by ID.

```
curl --location --request GET 'localhost:3000/api/dans/career/7638eee4-4e75-4c06-a920-ea7619a311b5' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwibmFtZSI6InRoYXJlZXEiLCJ1c2VybmFtZSI6InRoYXJlZXEiLCJleHAiOjE2Njg3Nzg5MTZ9.EyN-WmRvhofuptT2jHY-dx9VWD1xU09lgKIQNr95nPg'
```


## Run the project
Download package
```console
make download
```

Run dependency manager

```console
make dep
```

Connect postgres with password admin123
```console
make postgres in docker
```

Create DB in docker
```console
make createdb
```

Migrate in docker
```console
make migrate_up
```

Running app
```console
make run
```


