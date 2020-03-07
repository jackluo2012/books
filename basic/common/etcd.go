package common

// Etcd 配置
type Etcd struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	User    string    `json:"user"`
	Pass    string    `json:"pass"`
}
