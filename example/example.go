package main

import (
	"fmt"
	"log"

	"github.com/xxjwxc/fastgpt"
	"github.com/xxjwxc/fastgpt/model"
)

func main() {
	// 初始化FastGPT客户端
	fgpt := fastgpt.NewFastGPT("https://cloud.fastgpt.cn", "your-api-key")

	// 示例1：获取应用统计数据
	fmt.Println("=== 示例1：获取应用统计数据 ===")
	statsReq := &model.AppStatsRequest{
		StartTime: 1758585600000, // 2025-01-23 00:00:00
		EndTime:   1758672000000, // 2025-01-24 00:00:00
	}
	statsResp, err := fgpt.App.GetStats(statsReq)
	if err != nil {
		log.Printf("获取应用统计数据失败: %v\n", err)
	} else {
		fmt.Printf("应用统计数据获取成功，用户数据数量: %d\n", len(statsResp.Data.UserData))
	}

	// 示例2：发送对话请求
	fmt.Println("\n=== 示例2：发送对话请求 ===")
	chatReq := &model.ChatRequest{
		AppId: "your-app-id",
		Messages: []model.Message{
			{
				Role:    "user",
				Content: "你好",
			},
		},
		Stream: true,
	}

	err = fgpt.Chat.Chat(chatReq, func(eventType string, data interface{}) error {
		switch eventType {
		case "flowNodeStatus":
			statusEvent := data.(model.FlowNodeStatusEvent)
			fmt.Printf("流程节点状态: %s - %s\n", statusEvent.Name, statusEvent.Status)
		case "answer":
			if data == "[DONE]" {
				fmt.Println("对话结束")
				return nil
			}
			answerEvent := data.(model.AnswerEvent)
			for _, choice := range answerEvent.Choices {
				fmt.Print(choice.Delta.Content)
			}
		case "flowResponses":
			flowEvent := data.(model.FlowResponsesEvent)
			fmt.Printf("\n流程响应数量: %d\n", len(flowEvent.Responses))
		}
		return nil
	})
	if err != nil {
		log.Printf("对话请求失败: %v\n", err)
	}

	// 示例3：创建外部文件集合
	fmt.Println("\n=== 示例3：创建外部文件集合 ===")
	externalReq := &model.ExternalFileCollectionCreateRequest{
		ExternalFileUrl: "https://example.com/file.pdf",
		ExternalFileId:  "123456",
		Filename:        "示例文件.pdf",
		CreateTime:      "2025-01-01T00:00:00.000Z",
		DatasetId:       "your-dataset-id",
		TrainingType:    "chunk",
		ChunkSize:       1500,
		Tags:            []string{"tag1", "tag2"},
	}
	externalResp, err := fgpt.Dataset.CreateExternalFileCollection(externalReq)
	if err != nil {
		log.Printf("创建外部文件集合失败: %v\n", err)
	} else {
		fmt.Printf("外部文件集合创建成功，集合ID: %s\n", externalResp.CollectionId)
	}

	// 示例4：创建纯文本集合
	fmt.Println("\n=== 示例4：创建纯文本集合 ===")
	textReq := &model.TextCollectionCreateRequest{
		DatasetId:    "your-dataset-id",
		Text:         "这是一段测试纯文本内容，用于创建FastGPT知识库集合。",
		TrainingType: "chunk",
		ChunkSize:    1500,
		Tags:         []string{"text", "test"},
	}
	textResp, err := fgpt.Dataset.CreateTextCollection(textReq)
	if err != nil {
		log.Printf("创建纯文本集合失败: %v\n", err)
	} else {
		fmt.Printf("纯文本集合创建成功，集合ID: %s\n", textResp.CollectionId)
	}

	// 示例5：创建链接集合
	fmt.Println("\n=== 示例5：创建链接集合 ===")
	linkReq := &model.LinkCollectionCreateRequest{
		DatasetId:    "your-dataset-id",
		Link:         "https://example.com/article",
		TrainingType: "chunk",
		ChunkSize:    1500,
		Tags:         []string{"link", "article"},
	}
	linkResp, err := fgpt.Dataset.CreateLinkCollection(linkReq)
	if err != nil {
		log.Printf("创建链接集合失败: %v\n", err)
	} else {
		fmt.Printf("链接集合创建成功，集合ID: %s\n", linkResp.CollectionId)
	}
}
