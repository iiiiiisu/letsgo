package conf

type Database struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Redis struct {
	Network string `json:"network"`
	Addr    string `json:"addr"`
	Pwd     string `json:"password"`
}

type ServeMux struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Config struct {
	AppName   string     `json:"app_name"`
	Version   string     `json:"version"`
	Addr      string     `json:"addr"`
	Databases []Database `json:"databases"`
	Redis     []Redis    `json:"redis"`
}
