package config

import (
	"fmt"
	"time"
)

type Config struct {
	PGConn           string          `json:"pg_conn"`
	WebhookPath      string          `json:"webhook_path"`
	WebhookURL       string          `json:"webhook_url"`
	Addr             string          `json:"addr"`
	Rpc              RpcConfig       `json:"rpc"`
	ShutdownTimeout  time.Duration   `json:"shutdown_timeout"`
	RootDir          string          `json:"root_dir"`
	MessagesURL      string          `json:"messages_url"`
	UsersURL         string          `json:"users_url"`
}

type RpcConfig struct {
	Host             string	         `json:"host"`
	Port             string          `json:"port"`
}

func (c RpcConfig) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}