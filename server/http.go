package server

import (
	"context"
	"fmt"

	"github.com/erry-az/go-graceful"
	"github.com/labstack/echo/v4"
)

// HttpServerRoute defining http route to register
type HttpServerRoute func(e *echo.Echo)

// HttpConfig config for http server
type HttpConfig struct {
	// ListenAddress address setting for http server 0.0.0.0
	ListenAddress string `mapstructure:"listen_address"`
	// Port http server port
	Port int `mapstructure:"port"`
}

// AddressPort build address port from http config
func (hc *HttpConfig) AddressPort() string {
	return fmt.Sprintf("%s:%d", hc.ListenAddress, hc.Port)
}

// HttpServer struct that hold http config, server and watcher
type HttpServer struct {
	*echo.Echo
	cfg     *HttpConfig
	watcher *graceful.Graceful
}

// NewHttp init http server by defined config
func NewHttp(cfg *HttpConfig) *HttpServer {
	e := echo.New()
	e.Static("/assets", "assets")

	return &HttpServer{
		cfg:     cfg,
		Echo:    e,
		watcher: graceful.New(),
	}
}

// AddRoute register route to http server
func (hs *HttpServer) AddRoute(route HttpServerRoute) {
	route(hs.Echo)
}

// Start running http server
func (hs *HttpServer) Start() error {
	hs.watcher.RegisterProcess(func() error {
		return hs.Echo.Start(hs.cfg.AddressPort())
	})

	hs.watcher.RegisterShutdownProcessWithTag(func(ctx context.Context) error {
		return hs.Echo.Shutdown(ctx)
	}, "shutdown-server-echo")

	return hs.watcher.Wait()
}
