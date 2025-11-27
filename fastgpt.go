package fastgpt

import (
	"github.com/xxjwxc/fastgpt/api/app"
	"github.com/xxjwxc/fastgpt/api/chat"
	"github.com/xxjwxc/fastgpt/api/dataset"
	"github.com/xxjwxc/fastgpt/client"
)

// FastGPT 主客户端
type FastGPT struct {
	Client  *client.Client
	App     *app.AppAPI
	Chat    *chat.ChatAPI
	Dataset *dataset.DatasetAPI
}

// NewFastGPT 创建FastGPT客户端实例
// baseURL: FastGPT服务地址，例如：https://api.fastgpt.cn
// apiKey: 你的API密钥
func NewFastGPT(baseURL, apiKey string) *FastGPT {
	c := client.NewClient(baseURL, apiKey)
	return &FastGPT{
		Client:  c,
		App:     app.NewAppAPI(c),
		Chat:    chat.NewChatAPI(c),
		Dataset: dataset.NewDatasetAPI(c),
	}
}
