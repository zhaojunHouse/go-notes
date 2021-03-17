package config

type Config struct {
	DB    DB
	Mongo Mongo
}

type DB struct {
	Name     string
	User     string `default:"root"`
	Password string `required:"true" env:"DBPassword"`
	Port     uint   `default:"3306"`
}

type Mongo struct {
	Host string
	Name string
	Url  string
}
