package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type FancyPainterDBConfig struct {
	Dsn          string `yaml:"Dsn"`
	MaxIdleConns int    `yaml:"MaxIdleConns"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
	MaxLifetime  int    `yaml:"MaxLifetime"` // 单位秒
}

// SystemConfig 系统配置
type SystemConfig struct {
	AccessKeyID          string               `yaml:"AccessKeyID"`
	AccessKeySecret      string               `yaml:"AccessKeySecret"`
	ServerIP             string               `yaml:"ServerIP"`
	ServerPort           int32                `yaml:"ServerPort"`
	FancyPainterDBConfig FancyPainterDBConfig `yaml:"FancyPainterDBConfig"`
}

// gSystemConfig 系统配置
var gSystemConfig *SystemConfig

// InitSystemConfig 配置初始化
func InitSystemConfig() {
	log.Printf("InitSystemConfig begin1")
	file, err := ioutil.ReadFile("system_config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("InitSystemConfig begin2, file:%v", file)

	var systemConfig SystemConfig
	err = yaml.Unmarshal(file, &systemConfig)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("InitSystemConfig begin3")

	gSystemConfig = &systemConfig
	log.Printf("InitSystemConfig success, systemConfig:%v", systemConfig)
}

func GetAccessKeyID() string {
	if gSystemConfig != nil {
		return gSystemConfig.AccessKeyID
	}
	return ""
}

func GetAccessKeySecret() string {
	if gSystemConfig != nil {
		return gSystemConfig.AccessKeySecret
	}
	return ""
}

func GetServerIP() string {
	if gSystemConfig != nil {
		return gSystemConfig.ServerIP
	}
	return ""
}

func GetServerPort() int32 {
	if gSystemConfig != nil {
		return gSystemConfig.ServerPort
	}
	return 510010
}

func GetFancyPainterDBConfig() FancyPainterDBConfig {
	if gSystemConfig != nil {
		return gSystemConfig.FancyPainterDBConfig
	}
	return FancyPainterDBConfig{}
}
