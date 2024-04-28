package embedding

import "fmt"

type CommonConfigInterface interface {
	GetNum() int
}

type CommonConfig struct {
	Num int
}

func (c *CommonConfig) GetNum() int {
	return c.Num
}

type ClientConfig struct {
	CommonConfigInterface
}

type ServerConfig struct {
	CommonConfigInterface
}

func takeCommonConfig(configInstance CommonConfigInterface) {
	fmt.Println(configInstance.GetNum())
}

func executeEmbedding() {
	clientConfig := &ClientConfig{
		CommonConfigInterface: &CommonConfig{
			Num: 10,
		},
	}

	fmt.Println(clientConfig.GetNum())
	fmt.Println(clientConfig.CommonConfigInterface.GetNum())
	takeCommonConfig(clientConfig)
}
