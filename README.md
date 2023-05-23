# go-demo-web-api
Demo web-api written in Go

Effectively the *rest* example from the official chi project is used. 

## Goal

Offer a simple API that can be run easily locally. This API can be a "System Under Test" for Deno related api testing.

## Architectural Decision(s)

The chosen http router is [chi](https://go-chi.io/#/). This balances using an external dependency, that by itself is not dependant on other third-party modules/packages. 

[gofakeit](https://github.com/brianvoe/gofakeit) is used to seed struct data. This module has no external dependencies by itself. 