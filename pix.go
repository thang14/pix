package pix

import (
  "github.com/thang14/pix/router"
  "github.com/thang14/pix/middleware
)

type Config struct {
  logging bool
}

// Server contains instance details for the server
type Server struct {
  cfg *Config
  routers []Router
  middlewares []middleware.Middleware
  servers []*HTTPServer
}
