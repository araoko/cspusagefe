package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	BEServerName string   `json:"back_end_server_name"`
	BEServerPort int      `json:"back_end_server_port"`
	ServerName   string   `json:"server_name"`
	ServerPort   int      `json:"server_port"`
	AppRoot      string   `json:"server_approot"`
	CertFilePath string   `json:"ssl_cert_path"`
	KeyFilePath  string   `json:"ssl_key_path"`
	CaCerts      []string `json:"ssl_ca_pool"`
	LogFile      string   `json:"logFile"`
}

func LoadConfigFromFile(path string) (*Config, error) {
	c := &Config{
		BEServerName: restServerName,
		BEServerPort: restServerPort,
		ServerName:   serverName,
		ServerPort:   serverPort,
		AppRoot:      appRoot,
		CertFilePath: sslCert,
		KeyFilePath:  sslKey,
		CaCerts:      []string{ca},
	}

	confb, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(confb, c)
	if err != nil {
		return c, err
	}
	return c, nil

}

const (
	serverName     = "localhost"
	serverPort     = 8080
	appRoot        = ""
	restServerName = "localhost"
	restServerPort = 9090
	sslCert        = "cert.pem"
	sslKey         = "cert.key"
	ca             = "cacert.crt"
)
