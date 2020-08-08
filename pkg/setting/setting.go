package setting

import (
	"path/filepath"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(cfg string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	if cfg != "" {
		p, f := filepath.Split(cfg)
		vp.AddConfigPath(p)
		if f != "" {
			vp.SetConfigName(f)
		}
	}

	viper.AutomaticEnv()
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
