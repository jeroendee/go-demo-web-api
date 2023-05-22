# go-demo-web-api
Demo web-api written in Go

Effectively the *rest* example from the official chi project is used. 

## Mat Ryer

I tend to select for...
- Maintainability
- Glanceability (happy path line of sight)
- Code should be boring
- Code that is **clear to new Gophers**, or people new to the project
- Self-similar code

## Goal

Offer a simple API that can be run easily locally. This API can be a "System Under Test" for Deno related api testing.

## Architectural Decision(s)

The chosen http router is chi. This balances using an external dependency, that by itself is not dependant on other third-party modules/packages. 
