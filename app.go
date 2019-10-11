package main

import infra "github.com/hukurou-s/user-auth-api-with-jwt/infrastructure"

func main() {
	infra.Echo.Logger.Fatal(infra.Echo.Start(":1322"))
}
