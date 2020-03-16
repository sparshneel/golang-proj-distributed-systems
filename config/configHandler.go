package config

import(
	"github.com/spf13/viper"
	"sync"
)

var(
	cassandraConfig CassandraClusterConfiguration
	configurations *Configuration
	once sync.Once
)

type Configuration struct{
	CassandraConfig *CassandraClusterConfiguration
}

type CassandraClusterConfiguration struct{
	CassandraHost string
	CassandraPort string
}

func (config *Configuration)ParseConfig(configFilePath string) {
	viper.AddConfigPath(configFilePath)
	viper.SetConfigName("cassandra-cluster-config")
	viper.SetConfigType("yml")
	viper.Unmarshal(&cassandraConfig)
	config.CassandraConfig = &cassandraConfig
}

func GetInstanceConfiguration() {
	once.Do(func(){
		configurations = &Configuration{}
	})
}



