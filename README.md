# Authentication Server (REST API)

## <a href="https://web-cpv.ru">Speccy-Rom</a>

## Requirements
- docker & docker-compose

## Run Project

Set your signing key and hash salt in config file `pkg/config/config.yml` before running server

Use ```make run``` to build and run docker containers with application itself and mongodb instance

## Parse Token

You can import `github.com/Speccy-Rom/AuthenticationServer/pkg/parser` into your Go application and user `ParseToken()` to validate and parse token claims.

#### See example at `cmd/example/main.go -username=user -pass=pass`

You can run example using `go run main.go` but make sure you have golang installed on your machine

## Authentication Middleware Example Using Go/Gin

```golang
func Middleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err := parser.ParseToken(headerParts[1], SIGNING_KEY)
	if err != nil {
		status := http.StatusBadRequest
		if err == auth.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		c.AbortWithStatus(status)
		return
	}
}
```

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
