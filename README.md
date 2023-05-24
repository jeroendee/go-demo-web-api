# Go demo web-api
Demo web-api written in Go

Effectively the *rest* example from the official chi project is used and adapted according to need. 

## Goal

Offer a simple API that can be run easily locally. This API can be treated as a "System Under Test" for api testing.

## Architectural Decision(s)

*Principle*: don't rely on modules that by themselves rely on external depencendies. 

### http pipeline
The chosen http router is [chi](https://go-chi.io/#/). This package provides convenience over having to write the http pipeline boilerplate using the Go std lib. This module is using only the Go std lib, [and thus has no external dependencies](https://deps.dev/go/github.com%2Fgo-chi%2Fchi%2Fv5/v5.0.8/dependencies).


### fixture data
[gofakeit](https://github.com/brianvoe/gofakeit) is used to seed fixture struct data. This module is using only the Go std lib, [and thus has no external dependencies](https://deps.dev/go/github.com%2Fbrianvoe%2Fgofakeit%2Fv6/v6.21.0/dependencies).

### Slices functions
Since the introduction of generics in Go 1.18 an experimental x/exp/slices package that offers convenience functions when using slices. This package has been [accepted to be included in the official Go std lib](https://github.com/golang/go/issues/57433#issuecomment-1423528134)

## Usage
1. Build the binary in /cmd/server 
2. Run the `server` binary in your terminal of choice => this will display: "Server starting on localhost:3000"

Endpoints to investigate:

GET:
- /
- /ping
- /clients
- /clients/{id}

POST: 
- /clients {payload}

PUT:
- /clients/{id} {payload}

DELETE:
- /clients/{id}

**Payload example**

```json
{
    "id": string,
    "name": string,
    "surname": string,
    "birthdate": string,
    "email": string,
    "phone": string,
    "age": number
}
```