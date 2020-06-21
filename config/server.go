package config

import (
	"strings"

	"github.com/CapsLock-Studio/neo-komica/util"
)

// Server - the main config for web-server.
type Server struct {
	AllowOrigins []string
	SecretKey    string
	AESKey       string
	AppURL       string
}

// NewServer - get server config with default value.
func NewServer() *Server {
	s := &Server{}

	s.AllowOrigins = strings.Split(util.Getenv("ALLOW_ORIGINS", "http://localhost:3000"), ",")
	s.SecretKey = util.Getenv("SECRET_KEY", "password")
	s.AESKey = util.Getenv("AES_KEY", "3s6v9y$B&E)H@McQ")
	s.AppURL = util.Getenv("APP_URL", "http://localhost:3000")

	return s
}
