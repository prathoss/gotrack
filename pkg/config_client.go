package pkg

func NewClientConfigFromViper() (ClientConfig, error) {
	cc := ClientConfig{}

	if err := loadCfg(&cc); err != nil {
		return ClientConfig{}, err
	}

	return cc, nil
}

func GenerateClientConfig() (string, error) {
	return generateCfgTemplate(ClientConfig{})
}

type ClientConfig struct {
	Url string `mapstructure:"url" env:"GOTRACK_URL"`
}
