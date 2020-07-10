// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package config

// Config 系统配置
type Config struct {
	Server *Server
	DB     *DB
}

// DB 数据库配置
type DB struct {
	Type   string
	Dir    string
	User   string
	Pwd    string
	Host   string
	Port   int
	Dbname string
}

// Server 服务运行配置
type Server struct {
	Host string
	Port int
}
