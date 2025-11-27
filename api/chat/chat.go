// Package chat 提供FastGPT对话相关的API接口
//
// 该包封装了与对话交互相关的所有API，主要功能是发送对话请求并处理SSE（Server-Sent Events）响应。
// 支持实时对话流处理，包括：
// - 发送对话请求
// - 处理流式响应
// - 解析各种事件类型
// - 调用用户提供的事件处理函数
//
// 所有API均需要通过FastGPT客户端实例访问，使用前需先创建客户端。
package chat

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xxjwxc/fastgpt/client"
	"github.com/xxjwxc/fastgpt/model"
)

// ChatAPI 对话接口结构体，封装了所有对话相关的API方法
//
// 该结构体通过组合HTTP客户端，提供了与FastGPT对话交互相关的所有功能，
// 主要用于发送对话请求并处理SSE流式响应。
type ChatAPI struct {
	client *client.Client // HTTP客户端，用于发送API请求
}

// NewChatAPI 创建对话接口实例
//
// 参数：
//
//	c: HTTP客户端实例，由外部传入
//
// 返回值：
//
//	*ChatAPI: 对话接口实例，用于访问对话相关API
//
// 使用示例：
//
//	c := client.NewClient("https://cloud.fastgpt.cn", "sk-xxx")
//	chatAPI := chat.NewChatAPI(c)
func NewChatAPI(c *client.Client) *ChatAPI {
	return &ChatAPI{client: c}
}

// ChatEventHandler SSE事件处理函数类型
//
// 该类型定义了处理SSE事件的回调函数签名，当收到SSE事件时，会调用该函数进行处理。
//
// 参数：
//
//	eventType: 事件类型，如"flowNodeStatus"、"answer"、"flowResponses"等
//	data: 事件数据，根据事件类型不同，数据类型也不同
//
// 返回值：
//
//	error: 如果处理失败，返回错误信息，将终止整个对话流程
type ChatEventHandler func(eventType string, data interface{}) error

// Chat 发送对话请求并处理SSE流式响应
//
// 该方法用于发送对话请求，并通过SSE（Server-Sent Events）协议接收实时响应。
// 支持处理多种事件类型，包括节点状态、回答内容和流程响应等。
//
// 参数：
//
//	req: 对话请求，包含应用ID、消息列表、模型配置等
//	handler: SSE事件处理函数，用于处理接收到的各种事件
//
// 返回值：
//
//	error: 如果请求失败或事件处理失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat
//
// 使用示例：
//
//	req := &model.ChatRequest{
//	    AppId: "your-app-id",
//	    Messages: []model.Message{
//	        {Role: "user", Content: "你好"},
//	    },
//	}
//
//	err := chatAPI.Chat(req, func(eventType string, data interface{}) error {
//	    switch eventType {
//	    case "answer":
//	        if data == "[DONE]" {
//	            fmt.Println("对话结束")
//	            return nil
//	        }
//	        answerEvent := data.(model.AnswerEvent)
//	        fmt.Print(answerEvent.Delta.Content)
//	    case "flowNodeStatus":
//	        statusEvent := data.(model.FlowNodeStatusEvent)
//	        fmt.Printf("节点状态: %s\n", statusEvent.Status)
//	    case "flowResponses":
//	        flowEvent := data.(model.FlowResponsesEvent)
//	        fmt.Printf("流程响应: %+v\n", flowEvent)
//	    }
//	    return nil
//	})
func (api *ChatAPI) Chat(req *model.ChatRequest, handler ChatEventHandler) error {
	// 发送对话请求到FastGPT服务器
	resp, err := api.client.DoRequest("POST", "/api/v1/chat/completions", req)
	if err != nil {
		return err // 请求发送失败，返回错误
	}
	defer resp.Body.Close() // 确保响应体被关闭

	// 创建扫描器，用于逐行读取SSE流
	scanner := bufio.NewScanner(resp.Body)

	// 循环读取SSE流中的每一行，处理SSE事件
	var currentEvent string // 当前事件名称，默认为"message"
	var currentData []string // 当前事件的数据行

	for scanner.Scan() {
		line := scanner.Text()

		// 空行表示当前事件结束，处理累积的事件数据
		if line == "" {
			// 如果有累积的数据，处理当前事件
			if len(currentData) > 0 {
				// 合并多行data
				dataContent := strings.Join(currentData, "")
				
				// 根据事件名称解析数据
				switch currentEvent {
				case "flowNodeStatus":
					// 处理节点状态事件
					var statusEvent model.FlowNodeStatusEvent
					if err := json.Unmarshal([]byte(dataContent), &statusEvent); err != nil {
						return err // JSON解析失败，返回错误
					}
					// 调用事件处理函数
					if err := handler(currentEvent, statusEvent); err != nil {
						return err // 事件处理失败，返回错误
					}

				case "answer", "fastAnswer":
					// 处理回答事件和快速回答事件
					// 检查是否是对话结束标志
					if dataContent == "[DONE]" {
						if err := handler(currentEvent, "[DONE]"); err != nil {
							return err // 事件处理失败，返回错误
						}
						// 重置当前事件状态
						currentEvent = "message"
						currentData = []string{}
						continue // 继续处理下一个事件
					}

					// 解析回答事件数据
					var answerEvent model.AnswerEvent
					if err := json.Unmarshal([]byte(dataContent), &answerEvent); err != nil {
						return err // JSON解析失败，返回错误
					}
					// 调用事件处理函数
					if err := handler(currentEvent, answerEvent); err != nil {
						return err // 事件处理失败，返回错误
					}

				case "flowResponses":
					// 处理流程响应事件
					var flowEvent model.FlowResponsesEvent
					if err := json.Unmarshal([]byte(dataContent), &flowEvent); err != nil {
						return err // JSON解析失败，返回错误
					}
					// 调用事件处理函数
					if err := handler(currentEvent, flowEvent); err != nil {
						return err // 事件处理失败，返回错误
					}

				case "toolCall", "toolParams", "toolResponse", "updateVariables", "error":
					// 处理工具调用、工具参数、工具响应、更新变量和错误事件
					// 这些事件直接传递原始数据，由调用者自行解析
					if err := handler(currentEvent, dataContent); err != nil {
						return err // 事件处理失败，返回错误
					}

				case "interactive":
					// 处理交互节点事件
					var interactiveEvent model.Interactive
					if err := json.Unmarshal([]byte(dataContent), &interactiveEvent); err != nil {
						return err // JSON解析失败，返回错误
					}
					// 调用事件处理函数
					if err := handler(currentEvent, interactiveEvent); err != nil {
						return err // 事件处理失败，返回错误
					}

				default:
					// 处理未知事件类型，直接传递原始数据
					if err := handler(currentEvent, dataContent); err != nil {
						return err // 事件处理失败，返回错误
					}
				}

				// 重置当前事件状态
				currentEvent = "message"
				currentData = []string{}
			}
			continue
		}

		// 解析SSE事件行
		var field, value string
		if strings.HasPrefix(line, "data:") {
			// data字段，可能是多行
			field = "data"
			value = strings.TrimPrefix(line, "data:")
			// 如果是"data:"开头但没有值，使用空字符串
			if value == "" || strings.HasPrefix(value, " ") {
				value = strings.TrimPrefix(value, " ")
			}
		} else if strings.HasPrefix(line, "event:") {
			// event字段，设置当前事件名称
			field = "event"
			value = strings.TrimPrefix(line, "event:")
			if value != "" {
				value = strings.TrimPrefix(value, " ")
				currentEvent = value
			}
		} else {
			// 其他字段，如id、retry等，暂时忽略
			continue
		}

		// 处理当前字段
		switch field {
		case "data":
			// 添加到当前数据行
			currentData = append(currentData, value)
		case "event":
			// 已经处理过，设置了currentEvent
		default:
			// 忽略其他字段
		}
	}

	// 检查扫描过程中是否发生错误
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取SSE流失败: %w", err) // 包装错误信息
	}

	return nil // 对话处理成功
}

// GetHistories 获取应用历史记录
//
// 该方法用于获取应用的历史对话记录，支持分页查询。
//
// 参数：
//
//	req: 获取历史记录请求，包含应用ID、偏移量、每页数量和对话源
//
// 返回值：
//
//	*model.GetHistoriesResponse: 历史记录响应，包含历史记录列表和总记录数
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat#%E8%8E%B7%E5%8F%96%E6%9F%90%E4%B8%AA%E5%BA%94%E7%94%A8%E5%8E%86%E5%8F%B2%E8%AE%B0%E5%BD%95
func (api *ChatAPI) GetHistories(req *model.GetHistoriesRequest) (*model.GetHistoriesResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/chat/getHistories", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result model.GetHistoriesResponse
	if err := api.client.ParseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateHistory 更新历史记录
//
// 该方法用于更新对话历史记录，如修改标题或置顶状态。
//
// 参数：
//
//	req: 更新历史记录请求，包含应用ID、对话ID、自定义标题或置顶状态
//
// 返回值：
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat#%E4%BF%AE%E6%94%B9%E6%9F%90%E4%B8%AA%E5%AF%B9%E8%AF%9D%E7%9A%84%E6%A0%87%E9%A2%98
func (api *ChatAPI) UpdateHistory(req *model.UpdateHistoryRequest) error {
	resp, err := api.client.DoRequest("POST", "/api/core/chat/updateHistory", req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := api.client.ParseResponse(resp, nil); err != nil {
		return err
	}

	return nil
}

// DeleteHistory 删除单个历史记录
//
// 该方法用于删除单个对话历史记录。
//
// 参数：
//
//	appId: 应用ID
//
//	chatId: 对话ID
//
// 返回值：
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat#%E5%88%A0%E9%99%A4%E6%9F%90%E4%B8%AA%E5%8E%86%E5%8F%B2%E8%AE%B0%E5%BD%95
func (api *ChatAPI) DeleteHistory(appId, chatId string) error {
	resp, err := api.client.DoRequest("DELETE", fmt.Sprintf("/api/core/chat/delHistory?chatId=%s&appId=%s", chatId, appId), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := api.client.ParseResponse(resp, nil); err != nil {
		return err
	}

	return nil
}

// ClearHistories 清空所有历史记录
//
// 该方法用于清空应用的所有历史对话记录。
//
// 参数：
//
//	appId: 应用ID
//
// 返回值：
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat#%E6%B8%85%E7%A9%BA%E6%89%80%E6%9C%89%E5%8E%86%E5%8F%B2%E8%AE%B0%E5%BD%95
func (api *ChatAPI) ClearHistories(appId string) error {
	resp, err := api.client.DoRequest("DELETE", fmt.Sprintf("/api/core/chat/clearHistories?appId=%s", appId), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := api.client.ParseResponse(resp, nil); err != nil {
		return err
	}

	return nil
}

// GetInit 获取单个对话初始化信息
//
// 该方法用于获取单个对话的初始化信息。
//
// 参数：
//
//	appId: 应用ID
//
//	chatId: 对话ID
//
// 返回值：
//
//	*model.ChatInitResponse: 对话初始化信息响应
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat#%E8%8E%B7%E5%8F%96%E5%8D%95%E4%B8%AA%E5%AF%B9%E8%AF%9D%E5%88%9D%E5%A7%8B%E5%8C%96%E4%BF%A1%E6%81%AF
func (api *ChatAPI) GetInit(appId, chatId string) (*model.ChatInitResponse, error) {
	resp, err := api.client.DoRequest("GET", fmt.Sprintf("/api/core/chat/init?appId=%s&chatId=%s", appId, chatId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result model.ChatInitResponse
	if err := api.client.ParseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetPaginationRecords 获取对话记录列表
//
// 该方法用于获取对话记录列表，支持分页查询。
//
// 参数：
//
//	req: 获取对话记录列表请求，包含应用ID、对话ID、偏移量、每页数量和是否加载自定义反馈
//
// 返回值：
//
//	*model.GetPaginationRecordsResponse: 对话记录列表响应，包含记录列表和总记录数
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat#%E8%8E%B7%E5%8F%96%E5%AF%B9%E8%AF%9D%E8%AE%B0%E5%BD%95%E5%88%97%E8%A1%A8
func (api *ChatAPI) GetPaginationRecords(req *model.GetPaginationRecordsRequest) (*model.GetPaginationRecordsResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/chat/getPaginationRecords", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result model.GetPaginationRecordsResponse
	if err := api.client.ParseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetResData 获取单个对话记录运行详情
//
// 该方法用于获取单个对话记录的运行详情。
//
// 参数：
//
//	appId: 应用ID
//
//	chatId: 对话ID
//
//	dataId: 对话记录ID
//
// 返回值：
//
//	[]model.ResponseDataItem: 对话记录运行详情列表
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat#%E8%8E%B7%E5%8F%96%E5%8D%95%E4%B8%AA%E5%AF%B9%E8%AF%9D%E8%AE%B0%E5%BD%95%E8%BF%90%E8%A1%8C%E8%AF%A6%E6%83%85
func (api *ChatAPI) GetResData(appId, chatId, dataId string) ([]model.ResponseDataItem, error) {
	resp, err := api.client.DoRequest("GET", fmt.Sprintf("/api/core/chat/getResData?appId=%s&chatId=%s&dataId=%s", appId, chatId, dataId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []model.ResponseDataItem
	if err := api.client.ParseResponse(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteItem 删除对话记录
//
// 该方法用于删除单个对话记录。
//
// 参数：
//
//	appId: 应用ID
//
//	chatId: 对话ID
//
//	contentId: 对话记录ID
//
// 返回值：
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat#%E5%88%A0%E9%99%A4%E5%AF%B9%E8%AF%9D%E8%AE%B0%E5%BD%95
func (api *ChatAPI) DeleteItem(appId, chatId, contentId string) error {
	resp, err := api.client.DoRequest("DELETE", fmt.Sprintf("/api/core/chat/item/delete?contentId=%s&chatId=%s&appId=%s", contentId, chatId, appId), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := api.client.ParseResponse(resp, nil); err != nil {
		return err
	}

	return nil
}

// UpdateUserFeedback 更新用户反馈
//
// 该方法用于处理用户对对话记录的反馈，如点赞或点踩。
//
// 参数：
//
//	req: 更新用户反馈请求，包含应用ID、对话ID、数据ID和反馈信息
//
// 返回值：
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat#%E7%82%B9%E8%B5%9E--%E5%8F%96%E6%B6%88%E7%82%B9%E8%B5%9E
func (api *ChatAPI) UpdateUserFeedback(req *model.UpdateUserFeedbackRequest) error {
	resp, err := api.client.DoRequest("POST", "/api/core/chat/feedback/updateUserFeedback", req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := api.client.ParseResponse(resp, nil); err != nil {
		return err
	}

	return nil
}

// CreateQuestionGuide 创建猜你想问
//
// 该方法用于生成猜你想问的问题，支持自定义配置。
//
// 参数：
//
//	req: 创建猜你想问请求，包含应用ID、对话ID和问题引导配置
//
// 返回值：
//
//	*model.CreateQuestionGuideResponse: 猜你想问响应，包含生成的问题列表
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/chat#%E7%8C%9C%E4%BD%A0%E6%83%B3%E9%97%AE
func (api *ChatAPI) CreateQuestionGuide(req *model.CreateQuestionGuideRequest) (*model.CreateQuestionGuideResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/ai/agent/v2/createQuestionGuide", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result model.CreateQuestionGuideResponse
	if err := api.client.ParseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
