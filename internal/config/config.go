package config

import (
	"fmt"
	"errors"
	"encoding/json"
	"time"
)

type Config struct {
	PGConn           string          `json:"pg_conn"`
	WebhookPath      string          `json:"webhook_path"`
	WebhookURL       string          `json:"webhook_url"`
	Addr             string          `json:"addr"`
	Rpc              RpcConfig       `json:"rpc"`
	ShutdownTimeout  Duration        `json:"shutdown_timeout"`
	RootDir          string          `json:"root_dir"`
	MessagesURL      string          `json:"messages_url"`
	UsersURL         string          `json:"users_url"`
}

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) (err error) {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}

type RpcConfig struct {
	Host             string	         `json:"host"`
	Port             string          `json:"port"`
}

func (c RpcConfig) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}