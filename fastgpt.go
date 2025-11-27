// Package fastgpt 提供FastGPT SDK的主入口和核心功能
//
// FastGPT是一个强大的AI应用开发平台，提供对话、知识库、应用管理等功能。
// 该SDK封装了FastGPT的REST API，方便开发者在Go项目中快速集成FastGPT服务。
//
// 主要功能模块：
// - App: 应用管理和统计
// - Chat: 对话交互
// - Dataset: 知识库管理
//
// 使用示例：
//
//	import "github.com/xxjwxc/fastgpt"
//
//	// 创建客户端实例
//	fgpt := fastgpt.NewFastGPT("https://cloud.fastgpt.cn", "your-api-key")
//
//	// 使用各功能模块
//	// fgpt.App.GetStats(...)
//	// fgpt.Chat.Chat(...)
//	// fgpt.Dataset.CreateCollection(...)
package fastgpt

import (
	"github.com/xxjwxc/fastgpt/api/app"
	"github.com/xxjwxc/fastgpt/api/chat"
	"github.com/xxjwxc/fastgpt/api/dataset"
	"github.com/xxjwxc/fastgpt/client"
)

// FastGPT 主客户端结构体，封装了所有FastGPT SDK功能
// 
// 通过该结构体可以访问FastGPT的所有API模块，包括应用管理、对话交互和知识库管理。
type FastGPT struct {
	Client  *client.Client      // HTTP客户端，负责处理所有API请求
	App     *app.AppAPI         // 应用API，用于应用管理和统计
	Chat    *chat.ChatAPI       // 对话API，用于与AI模型进行交互
	Dataset *dataset.DatasetAPI // 知识库API，用于管理和操作知识库
}

// SetDebug 设置debug模式
// 
// 参数：
//   debug: 是否开启debug模式，开启后会打印HTTP请求和响应
// 
// 使用示例：
//
//     fgpt := fastgpt.NewFastGPT("https://api.fastgpt.cn", "sk-xxx")
//     fgpt.SetDebug(true) // 开启debug模式
//
func (f *FastGPT) SetDebug(debug bool) {
	f.Client.Debug = debug
}

// NewFastGPT 创建FastGPT客户端实例
//
// 参数：
//
//	baseURL: FastGPT服务地址，例如：https://cloud.fastgpt.cn
//	apiKey: 你的API密钥，用于身份验证
//
// 返回值：
//
//	*FastGPT: FastGPT客户端实例，用于访问所有API功能
//
// 使用示例：
//
//	fgpt := fastgpt.NewFastGPT("https://cloud.fastgpt.cn", "sk-xxx")
func NewFastGPT(baseURL, apiKey string) *FastGPT {
	// 创建HTTP客户端
	c := client.NewClient(baseURL, apiKey)

	// 初始化各API模块
	return &FastGPT{
		Client:  c,                        // HTTP客户端实例
		App:     app.NewAppAPI(c),         // 应用API实例
		Chat:    chat.NewChatAPI(c),       // 对话API实例
		Dataset: dataset.NewDatasetAPI(c), // 知识库API实例
	}
}
