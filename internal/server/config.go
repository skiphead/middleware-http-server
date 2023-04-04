package server

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {
	BindAddr string `json:"bind_addr"`
}

type Server struct {
	Params Config
	Router *http.ServeMux
}

func NewConfig() *Server {
	cfg := ParseConfig()
	if cfg.BindAddr == "" {
		cfg.BindAddr = "0.0.0.0:8080"
	}
	return &Server{
		Params: Config{
			BindAddr: cfg.BindAddr,
		},
		Router: http.NewServeMux(),
	}
}

var path string

func init() {
	flag.StringVar(&path, "config", "./config/config.json", "path to path config json")
	flag.Parse()
}

func ParseConfig() Config {
	var conf Config
	byteData := OpenConfig(path)
	err := json.Unmarshal(byteData, &conf)
	if err != nil {
		log.Println(err)
	}
	return conf
}

func OpenConfig(jsonPath string) []byte {
	conf, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Println(err)
	}
	return conf
}
