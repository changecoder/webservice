package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// Ws WS
type Ws struct {
	JwtSecret string
	PageSize  int

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
}

// Server Server
type Server struct {
	RunMode      string
	SitePort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Database DB
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

// Redis Usage
type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

// WsSetting WS Setting
var WsSetting = &Ws{}

// ServerSetting ServerSetting
var ServerSetting = &Server{}

// DatabaseSetting DatabaseSetting
var DatabaseSetting = &Database{}

// RedisSetting Redis
var RedisSetting = &Redis{}

var cfg *ini.File

// Setup StartUp
func Setup() {
	var err error

	cfg, err = ini.Load("conf/config.ini")

	if err != nil {
		log.Fatalf("Setting Setup Fail")
	}

	mapTo("ws", WsSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	WsSetting.ImageMaxSize = WsSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
