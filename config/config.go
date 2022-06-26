package config

var Conf Config

type Config struct {
	LastFM struct {
		Url    string
		ApiKey string
	}
}
