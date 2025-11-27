package chat

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xxjwxc/fastgpt/client"
	"github.com/xxjwxc/fastgpt/model"
)

// ChatAPI 对话接口
type ChatAPI struct {
	client *client.Client
}

// NewChatAPI 创建对话接口实例
func NewChatAPI(c *client.Client) *ChatAPI {
	return &ChatAPI{client: c}
}

// ChatEventHandler SSE事件处理函数类型
type ChatEventHandler func(eventType string, data interface{}) error

// Chat 发送对话请求
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat
func (api *ChatAPI) Chat(req *model.ChatRequest, handler ChatEventHandler) error {
	// 发送对话请求
	resp, err := api.client.DoRequest("POST", "/api/proApi/chat", req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取SSE流
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// 解析SSE事件
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			continue
		}

		eventType := parts[0]
		dataStr := parts[1]

		// 根据事件类型处理数据
		switch eventType {
		case "event":
			// 处理事件类型
			eventName := dataStr
			// 下一行应该是data字段
			if !scanner.Scan() {
				break
			}
			dataLine := scanner.Text()
			if !strings.HasPrefix(dataLine, "data: ") {
				break
			}
			dataContent := strings.TrimPrefix(dataLine, "data: ")

			// 根据事件名称解析数据
			switch eventName {
			case "flowNodeStatus":
				var statusEvent model.FlowNodeStatusEvent
				if err := json.Unmarshal([]byte(dataContent), &statusEvent); err != nil {
					return err
				}
				if err := handler(eventName, statusEvent); err != nil {
					return err
				}

			case "answer":
				// 检查是否是[DONE]结束标志
				if dataContent == "[DONE]" {
					if err := handler(eventName, "[DONE]"); err != nil {
						return err
					}
					continue
				}
				var answerEvent model.AnswerEvent
				if err := json.Unmarshal([]byte(dataContent), &answerEvent); err != nil {
					return err
				}
				if err := handler(eventName, answerEvent); err != nil {
					return err
				}

			case "flowResponses":
				var flowEvent model.FlowResponsesEvent
				if err := json.Unmarshal([]byte(dataContent), &flowEvent); err != nil {
					return err
				}
				if err := handler(eventName, flowEvent); err != nil {
					return err
				}

			default:
				// 未知事件类型，跳过
				if err := handler(eventName, dataContent); err != nil {
					return err
				}
			}

		default:
			// 其他字段，跳过
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取SSE流失败: %w", err)
	}

	return nil
}
