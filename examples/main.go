package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"net"
	"net/url"
	"time"

	relay "github.com/katallaxie/fiber-ws-relay/v3"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/katallaxie/pkg/server"
)

var addr = flag.String("addr", "localhost:4222", "server addr")

var _ server.Listener = (*svc)(nil)

type svc struct{}

func (s *svc) Start(ctx context.Context, _ server.ReadyFunc, _ server.RunFunc) func() error {
	return func() error {
		cfg := net.ListenConfig{
			KeepAlive: time.Minute,
		}

		l, err := cfg.Listen(ctx, "tcp4", *addr)
		if err != nil {
			return err
		}
		defer l.Close()

		c, err := l.Accept()
		if err != nil {
			return err
		}

		log.Printf("serving %s\n", c.RemoteAddr().String())

		for {
			buf := make([]byte, 65536)
			_, err := bufio.NewReader(c).Read(buf)
			if err != nil {
				return err
			}

			log.Println("recv:", string(buf))

			_, err = bufio.NewWriter(c).WriteString("got it")
			if err != nil {
				return err
			}
		}
	}
}

type ws struct{}

func (ws *ws) Start(_ context.Context, _ server.ReadyFunc, _ server.RunFunc) func() error {
	return func() error {
		app := fiber.New()

		app.Use(logger.New())
		app.Get("/ws", relay.New(relay.Config{}, "localhost:4222"))

		if err := app.Listen(":8080"); err != nil {
			return err
		}

		return nil
	}
}

type multiplex struct{}

func (m *multiplex) Start(ctx context.Context, _ server.ReadyFunc, _ server.RunFunc) func() error {
	return func() error {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		time.Sleep(time.Second * 5)

		u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
		log.Printf("connecting to %s", u.String())

		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			return err
		}
		defer c.Close()

		for {
			select {
			case t := <-ticker.C:
				err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
				if err != nil {
					return err
				}
			case <-ctx.Done():
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					return err
				}
			}
		}
	}
}

func main() {
	s, _ := server.WithContext(context.Background())

	svc := &svc{}
	s.Listen(svc, false)

	ws := &ws{}
	s.Listen(ws, false)

	m := &multiplex{}
	s.Listen(m, false)

	log.Fatal(s.Wait())
}
