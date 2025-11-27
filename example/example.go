// Package main 提供FastGPT SDK的使用示例
//
// 该示例文件展示了FastGPT SDK的主要功能，包括：
// 1. 获取应用累积运行结果
// 2. 获取应用日志看板
// 3. 发送对话请求
// 4. 创建知识库
// 5. 创建纯文本集合
// 6. 创建链接集合
// 7. 为集合批量添加数据
// 8. 搜索测试
//
// 使用前需将示例中的API密钥、应用ID和知识库ID替换为实际值。
package main

import (
	"fmt"
	"log"

	"github.com/xxjwxc/fastgpt"       // FastGPT SDK主包
	"github.com/xxjwxc/fastgpt/model" // 数据模型定义
)

// main 函数是示例程序的入口点
//
// 该函数演示了FastGPT SDK的各种使用场景，包括应用管理、对话交互和知识库管理。
func main() {
	// 初始化FastGPT客户端
	// 参数1: FastGPT服务地址，此处使用云服务地址
	// 参数2: API密钥，需替换为实际的API密钥
	fgpt := fastgpt.NewFastGPT("https://cloud.fastgpt.cn", "your-api-key")
	// 方法1：通过SetDebug方法开启debug模式
	fgpt.SetDebug(true)

	// 示例1：获取应用累积运行结果
	// 演示如何获取应用的累积运行结果
	fmt.Println("=== 示例1：获取应用累积运行结果 ===")
	totalDataReq := &model.AppTotalDataRequest{
		AppId: "your-app-id", // 应用ID，需替换为实际的应用ID
	}
	// 调用App.GetTotalData方法获取应用累积运行结果
	totalDataResp, err := fgpt.App.GetTotalData(totalDataReq)
	if err != nil {
		log.Printf("获取应用累积运行结果失败: %v\n", err)
	} else {
		fmt.Printf("应用累积运行结果获取成功:\n")
		fmt.Printf("  累积使用用户数量: %d\n", totalDataResp.Data.TotalUsers)
		fmt.Printf("  累积对话数量: %d\n", totalDataResp.Data.TotalChats)
		fmt.Printf("  累积积分消耗: %d\n", totalDataResp.Data.TotalPoints)
	}

	// 示例2：获取应用日志看板
	// 演示如何获取应用的日志看板数据
	fmt.Println("\n=== 示例2：获取应用日志看板 ===")
	chartDataReq := &model.AppChartDataRequest{
		AppId:     "your-app-id",              // 应用ID，需替换为实际的应用ID
		DateStart: "2025-09-19T16:00:00.000Z", // 开始时间
		DateEnd:   "2025-09-27T15:59:59.999Z", // 结束时间
		Offset:    1,                          // 用户留存偏移量
		Source: []string{
			"test",
			"online",
			"share",
			"api",
			"cronJob",
			"team",
			"feishu",
			"official_account",
			"wecom",
			"mcp",
		}, // 日志来源
		UserTimespan: "day", // 用户数据时间跨度
		ChatTimespan: "day", // 对话数据时间跨度
		AppTimespan:  "day", // 应用数据时间跨度
	}
	// 调用App.GetChartData方法获取应用日志看板数据
	chartDataResp, err := fgpt.App.GetChartData(chartDataReq)
	if err != nil {
		log.Printf("获取应用日志看板失败: %v\n", err)
	} else {
		fmt.Printf("应用日志看板获取成功:\n")
		fmt.Printf("  用户数据数量: %d\n", len(chartDataResp.Data.UserData))
		fmt.Printf("  对话数据数量: %d\n", len(chartDataResp.Data.ChatData))
		fmt.Printf("  应用数据数量: %d\n", len(chartDataResp.Data.AppData))
	}

	// 示例3：发送对话请求
	// 演示如何发送对话请求并处理流式响应
	fmt.Println("\n=== 示例3：发送对话请求 ===")
	chatReq := &model.ChatRequest{
		ChatId: "my_chatId", // 对话ID，可选
		Stream: true,        // 启用流式响应
		Detail: false,       // 是否返回中间值
		Messages: []model.Message{
			{
				Role:    "user", // 消息角色，此处为用户
				Content: "你好",   // 消息内容
			},
		},
	}

	// 调用Chat.Chat方法发送对话请求，使用回调函数处理流式事件
	err = fgpt.Chat.Chat(chatReq, func(eventType string, data interface{}) error {
		// 根据事件类型处理不同的事件
		switch eventType {
		case "flowNodeStatus":
			// 处理流程节点状态事件
			statusEvent := data.(model.FlowNodeStatusEvent)
			fmt.Printf("流程节点状态: %s - %s\n", statusEvent.Name, statusEvent.Status)
		case "answer":
			// 处理回答事件
			if data == "[DONE]" {
				// 对话结束标志
				fmt.Println("对话结束")
				return nil
			}
			// 解析回答内容
			answerEvent := data.(model.AnswerEvent)
			for _, choice := range answerEvent.Choices {
				// 输出增量内容
				fmt.Print(choice.Delta.Content)
			}
		case "flowResponses":
			// 处理流程响应事件
			flowEvent := data.(model.FlowResponsesEvent)
			fmt.Printf("\n流程响应数量: %d\n", len(flowEvent.Responses))
		}
		return nil
	})
	if err != nil {
		log.Printf("对话请求失败: %v\n", err)
	}

	// 示例4：创建知识库
	// 演示如何创建一个新的知识库
	fmt.Println("\n=== 示例4：创建知识库 ===")
	datasetCreateReq := &model.DatasetCreateRequest{
		Name:  "测试知识库",     // 知识库名称
		Type:  "dataset",   // 知识库类型
		Intro: "这是一个测试知识库", // 知识库介绍
	}
	// 调用Dataset.CreateDataset方法创建知识库
	datasetId, err := fgpt.Dataset.CreateDataset(datasetCreateReq)
	if err != nil {
		log.Printf("创建知识库失败: %v\n", err)
	} else {
		fmt.Printf("知识库创建成功，知识库ID: %s\n", datasetId)
		// 使用新创建的知识库ID进行后续操作
		// 注意：在实际使用中，您需要将此处的datasetId替换为实际创建的知识库ID
		// 或者使用您已有的知识库ID
		datasetId = "your-dataset-id" // 替换为实际的知识库ID
	}

	// 示例5：创建纯文本集合
	// 演示如何通过纯文本内容创建知识库集合
	fmt.Println("\n=== 示例5：创建纯文本集合 ===")
	textReq := &model.CollectionCreateTextRequest{
		Text:             "这是一段测试纯文本内容，用于创建FastGPT知识库集合。", // 纯文本内容
		DatasetId:        datasetId,                       // 知识库ID
		Name:             "测试文本集合",                        // 集合名称
		TrainingType:     "chunk",                         // 数据处理方式：chunk-按文本长度分割
		ChunkSettingMode: "auto",                          // 分块参数模式：auto-系统默认参数
	}
	// 调用Dataset.CreateTextCollection方法创建纯文本集合
	textResp, err := fgpt.Dataset.CreateTextCollection(textReq)
	if err != nil {
		log.Printf("创建纯文本集合失败: %v\n", err)
	} else {
		fmt.Printf("纯文本集合创建成功，集合ID: %s\n", textResp.CollectionId)
		fmt.Printf("插入块数量: %d\n", textResp.Results.InsertLen)
	}

	// 示例6：创建链接集合
	// 演示如何通过网络链接创建知识库集合
	fmt.Println("\n=== 示例6：创建链接集合 ===")
	linkReq := &model.CollectionCreateLinkRequest{
		Link:             "https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset", // 网络链接
		DatasetId:        datasetId,                                                              // 知识库ID
		TrainingType:     "chunk",                                                                // 数据处理方式：chunk-按文本长度分割
		ChunkSettingMode: "auto",                                                                 // 分块参数模式：auto-系统默认参数
	}
	// 调用Dataset.CreateLinkCollection方法创建链接集合
	linkResp, err := fgpt.Dataset.CreateLinkCollection(linkReq)
	if err != nil {
		log.Printf("创建链接集合失败: %v\n", err)
	} else {
		fmt.Printf("链接集合创建成功，集合ID: %s\n", linkResp.CollectionId)
		fmt.Printf("插入块数量: %d\n", linkResp.Results.InsertLen)
	}

	// 示例7：为集合批量添加数据
	// 演示如何为指定集合批量添加数据
	fmt.Println("\n=== 示例7：为集合批量添加数据 ===")
	pushDataReq := &model.DataPushRequest{
		CollectionId: "your-collection-id", // 集合ID，需替换为实际的集合ID
		TrainingType: "chunk",              // 训练模式：chunk-按文本长度分割
		Data: []model.DatasetData{
			{
				Q: "FastGPT是什么？",
				A: "FastGPT是一个基于大语言模型的知识库管理和对话系统。",
				Indexes: []model.Index{
					{
						Text: "FastGPT 知识库管理 对话系统",
					},
				},
			},
			{
				Q: "FastGPT支持哪些功能？",
				A: "FastGPT支持知识库管理、对话交互、文件导入、搜索测试等功能。",
				Indexes: []model.Index{
					{
						Text: "FastGPT 功能 知识库管理 对话交互 文件导入 搜索测试",
					},
				},
			},
		},
	}
	// 调用Dataset.PushData方法为集合批量添加数据
	pushDataResp, err := fgpt.Dataset.PushData(pushDataReq)
	if err != nil {
		log.Printf("为集合批量添加数据失败: %v\n", err)
	} else {
		fmt.Printf("为集合批量添加数据成功:\n")
		fmt.Printf("  插入数量: %d\n", pushDataResp.InsertLen)
		fmt.Printf("  超出token数量: %d\n", len(pushDataResp.OverToken))
		fmt.Printf("  重复数量: %d\n", len(pushDataResp.Repeat))
		fmt.Printf("  错误数量: %d\n", len(pushDataResp.Error))
	}

	// 示例8：搜索测试
	// 演示如何测试知识库搜索功能
	fmt.Println("\n=== 示例8：搜索测试 ===")
	searchTestReq := &model.DatasetSearchTestRequest{
		DatasetId:   datasetId,     // 知识库ID
		Text:        "FastGPT是什么？", // 需要测试的文本
		Limit:       5000,          // 最大tokens数量
		SearchMode:  "embedding",   // 搜索模式：embedding-向量搜索
		UsingReRank: false,         // 是否使用重排
	}
	// 调用Dataset.SearchTest方法进行搜索测试
	searchTestResp, err := fgpt.Dataset.SearchTest(searchTestReq)
	if err != nil {
		log.Printf("搜索测试失败: %v\n", err)
	} else {
		fmt.Printf("搜索测试成功，结果数量: %d\n", len(searchTestResp))
		for i, result := range searchTestResp {
			fmt.Printf("  结果 %d: 相似度 %.2f, 问题: %s\n", i+1, result.Score, result.Q)
		}
	}
}
