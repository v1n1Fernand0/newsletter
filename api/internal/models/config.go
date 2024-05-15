package models

type Config struct {
	Sites     []string `yaml:"sites"`
	Frequency string   `yaml:"frequency"`
	Email     struct {
		SMTPServer string `yaml:"smtp_server"`
		SMTPPort   int    `yaml:"smtp_port"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		From       string `yaml:"from"`
		To         string `yaml:"to"`
	} `yaml:"email"`
}
