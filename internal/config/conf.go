package config

import (
	"database/sql"
	"log"
	"path/filepath"
	"time"

	"github.com/daheige/inject-demo/internal/infras/db"
	"github.com/spf13/viper"
)

// AppConfig app config
type AppConfig struct {
	Port      int
	DBDsn     string
	AppEnv    string
	AppDebug  bool
	GraceWait time.Duration
}

var (
	// CfgFile 配置文件路径
	CfgFile string
)

// LoadConfig read config from file.
func LoadConfig() (*viper.Viper, error) {
	viperEntry := viper.New()
	if CfgFile != "" {
		// Use config file from the flag.
		viperEntry.SetConfigFile(CfgFile)
	} else {
		configDir, err := filepath.Abs("./")
		if err != nil {
			return nil, err
		}

		log.Println("config_dir: ", configDir)
		// Search config in ./ directory with name "app.yaml"
		viperEntry.AddConfigPath(configDir)
		viperEntry.SetConfigType("yaml")
		viperEntry.SetConfigName("app")
	}

	viperEntry.AutomaticEnv()
	if err := viperEntry.ReadInConfig(); err != nil {
		return nil, err
	}

	log.Println("using config file:", viperEntry.ConfigFileUsed())

	return viperEntry, nil
}

// InitAppConf 初始化配置
func InitAppConf(viperEntry *viper.Viper) (*AppConfig, error) {
	conf := &AppConfig{}
	if err := viperEntry.UnmarshalKey("AppConfig", conf); err != nil {
		return nil, err
	}

	return conf, nil
}

// InitDBConf init db
func InitDBConf(viperEntry *viper.Viper) (*sql.DB, error) {
	conf := &db.DBConf{}
	if err := viperEntry.UnmarshalKey("DBConfig", conf); err != nil {
		return nil, err
	}

	return conf.ConnectDB()
}
