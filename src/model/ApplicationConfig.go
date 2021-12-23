package model

type ConnectionConfig struct {
	ConnectionName string `json:"connectionName"`
	HostName       string `json:"hostName"`
	HostPort       string `json:"hostPort"`
	UserName       string `json:"userName"`
	Password       string `json:"password"`
	Topic          string `json:"topic"`
}
