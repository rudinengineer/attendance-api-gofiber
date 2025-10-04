package config

type Config struct {
	Server   Server
	Database Database
	JWT      JWT
}

type Server struct {
	Host     string
	Port     string
	Timezone string
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type JWT struct {
	SecretKey   string
	ExpiredTime string
}
