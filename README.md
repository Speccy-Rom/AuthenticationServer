# Authentication Server (REST API)

## <a href="https://web-cpv.ru">Speccy-Rom</a>

## Requirements
- docker & docker-compose

## Run Project

Set your signing key and hash salt in config file `pkg/config/config.yml` before running server

Use ```make run``` to build and run docker containers with application itself and mongodb instance


## API:

### POST /auth/sign-up

Registers new user

#### Example Input: 
```
{
	"username": "user",
	"password": "password"
}
```

#### Example Response: 
```
{
	"status": "ok",
	"message": "user created successfully"
}
```

### POST /auth/sign-in

Generates JWT Token

#### Example Input: 
```
{
	"username": "user",
	"password": "password"
}
```

#### Example Response: 
```
{
	"status": "ok",
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7InVzZXJuYW1lIjoidXNlciIsInBhc3N3b3JkIjoiNWJhYTYxZTRjOWI5M2YzZjA2ODIyNTBiNmNmODMzMWI3ZWU2OGZkOCJ9fQ.UvCbjhn7o17cvvYRK3rr6ih0Ro_VvZZpKWns1sOH-CE"
}
```
