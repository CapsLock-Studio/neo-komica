package config

import (
	"fmt"

	"github.com/CapsLock-Studio/neo-komica/util"
)

// ConnectInfo - the basic form for db config
type ConnectInfo struct {
	URI      string // e.g. redis:6379
	Username string
	Password string
	DBName   string
	Host     string // e.g. redis
}

// Gorm - the setting config for gorm
type Gorm struct {
	ConnectInfo
	Dialect string
	Log     bool
}

// Redis - the setting config for redis
type Redis struct {
	ConnectInfo
}

// NewGorm return Gorm default config with default setting
func NewGorm() *Gorm {
	g := &Gorm{}
	g.Dialect = util.Getenv("DB_DIALECT", "mysql")
	// g.URI = util.Getenv("DB_CONNECT_URI", "mysql:3306")
	g.Username = util.Getenv("DB_USERNAME", "root")
	g.Password = util.Getenv("DB_PASSWORD", "example")
	g.DBName = util.Getenv("DB_NAME", "kol_radar")
	g.Host = util.Getenv("DB_HOST", "localhost")
	g.Log = util.Getenv("GORM_LOG", "false") == "true"

	return g
}

// ConnectURI combines gorm setting to URI
func (m *Gorm) ConnectURI() string {
	switch m.Dialect {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.DBName)
	default:
		return m.URI
	}
}

// NewRedis return Redis config with default setting
func NewRedis() *Redis {
	r := &Redis{}
	r.URI = util.Getenv("REDIS_URI", "localhost:6379")
	r.Username = util.Getenv("REDIS_USERNAME", "")
	r.Password = util.Getenv("REDIS_PASS", "")

	return r
}
