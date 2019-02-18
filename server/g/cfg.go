package g

import (
	"encoding/json"

	"log"

	"sync"

	"github.com/toolkits/file"
)

type GlobalConfig struct {
	GRPC         *GRPCConfig `json:"grpc"`
	Http         *HttpConfig `json:"http"`
	ChannelCache int         `json:"channel_cache"`
	AliveClients []string    `json:"alive_clients"`
}

type GRPCConfig struct {
	Listen string `json:"listen"`
}

type HttpConfig struct {
	X_API_KEY string `json:"x-api-key"`
	Listen    string `json:"listen"`
}

type Message struct {
	ClientName   string
	Notification string
}

var NotifMessage chan Message

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

}
