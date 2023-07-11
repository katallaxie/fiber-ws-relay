# :ticket: Websocket-TCP Relay

[![Test & Build](https://github.com/katallaxie/fiber-ws-relay/actions/workflows/main.yml/badge.svg)](https://github.com/katallaxie/fiber-ws-relay/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/katallaxie/fiber-ws-relay)](https://goreportcard.com/report/github.com/katallaxie/fiber-ws-relay)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)

Websocket-TCP Relay is a simple relay that allows to connect to a TCP server via a websocket connection. It is built with [Fiber](https://gofiber.io/).

## Usage

This creates a middleware relay that connects to a TCP server on `localhost:4222` and relays the data between the websocket and the TCP server.

```go
app := fiber.New()

app.Use(logger.New())
app.Get("/ws", relay.New(relay.Config{}, "localhost:4222"))

if err := app.Listen(":8080"); err != nil {
	return err
}
```

## License

[MIT](/LICENSE)