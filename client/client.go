// Package client 提供HTTP客户端功能，负责处理FastGPT API的请求发送和响应处理
//
// 该包封装了HTTP请求的创建、发送和响应解析逻辑，为上层API模块提供基础支持。
// 主要功能包括：
// - 创建配置化的HTTP客户端
// - 发送JSON格式的API请求
// - 处理API响应的解析
// - 管理请求头和身份验证
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/xxjwxc/fastgpt/model"
)

// Client FastGPT HTTP客户端结构体，负责处理所有API请求和响应
//
// 该结构体封装了与FastGPT服务器通信的所有细节，包括请求构建、身份验证、
// 超时设置和响应处理。
type Client struct {
	BaseURL    string       // FastGPT服务基础URL，例如：https://api.fastgpt.cn
	APIKey     string       // API密钥，用于身份验证
	HTTPClient *http.Client // 底层HTTP客户端，用于发送请求
	Debug      bool         // 是否开启debug模式，开启后会打印HTTP请求和响应
}

// NewClient 创建新的FastGPT HTTP客户端实例
//
// 参数：
//
//	baseURL: FastGPT服务基础URL，例如：https://cloud.fastgpt.cn
//	apiKey: API密钥，用于身份验证
//
// 返回值：
//
//	*Client: 配置好的HTTP客户端实例
//
// 使用示例：
//
//	c := client.NewClient("https://cloud.fastgpt.cn", "sk-xxx")
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second, // 设置30秒超时
		},
		Debug: false, // 默认关闭debug模式
	}
}

// DoRequest 发送HTTP请求到FastGPT服务器
//
// 参数：
//
//	method: HTTP方法，如"GET"、"POST"等
//	path: API路径，如"/api/proApi/app/stats"
//	body: 请求体数据，将被序列化为JSON格式
//
// 返回值：
//
//	*http.Response: HTTP响应对象，需要调用者处理响应体
//	error: 如果请求发送失败，返回错误信息
//
// 处理流程：
// 1. 如果请求体不为空，将其序列化为JSON格式
// 2. 创建HTTP请求，设置URL、方法和请求体
// 3. 添加请求头，包括Authorization、Content-Type和User-Agent
// 4. 发送请求并返回响应
func (c *Client) DoRequest(method, path string, body interface{}) (*http.Response, error) {
	var reqBody io.Reader

	// 如果请求体不为空，将其序列化为JSON
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err // 序列化失败，返回错误
		}
		reqBody = bytes.NewBuffer(jsonBody) // 创建字节缓冲区
	}

	// 创建HTTP请求
	req, err := http.NewRequest(method, c.BaseURL+path, reqBody)
	if err != nil {
		return nil, err // 请求创建失败，返回错误
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+c.APIKey) // 添加身份验证头
	req.Header.Set("Content-Type", "application/json")  // 设置内容类型为JSON
	req.Header.Set("User-Agent", "go-fastgpt-client")   // 设置用户代理

	// 发送请求并返回响应
	return c.HTTPClient.Do(req)
}

// ParseResponse 解析HTTP响应体为指定的结构体
//
// 参数：
//
//	resp: HTTP响应对象，由DoRequest方法返回
//	v: 用于存储解析结果的结构体指针
//
// 返回值：
//
//	error: 如果解析失败，返回错误信息
//
// 注意事项：
// - 该方法会自动关闭响应体
// - 响应体必须是JSON格式
// - v必须是结构体指针
// - 该方法会检查BaseResponse的Code字段，200表示成功，其他状态码返回错误
//
// 优化说明：
// 1. 对于标准BaseResponse格式：
//   - 第一次解析：检查状态码
//   - 第二次解析：将Data字段解析为目标结构体（使用json.RawMessage避免二次序列化）
//
// 2. 对于非标准格式：
//   - 只解析一次，直接解析为目标结构体
//
// 3. 内存优化：使用io.ReadAll读取响应体，便于debug模式打印完整响应
func (c *Client) ParseResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close() // 确保响应体被关闭

	// 读取响应体内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// 如果开启了debug模式，打印HTTP返回结果
	if c.Debug {
		fmt.Printf("HTTP Response: %s\n", string(body))
	}

	// 首先解析为BaseResponse，检查状态码
	var baseResp model.BaseResponse
	if err := json.Unmarshal(body, &baseResp); err != nil {
		// 如果不是BaseResponse格式，直接解析为目标结构体
		return json.Unmarshal(body, v)
	}

	// 检查状态码，200表示成功，其他状态码返回错误
	if baseResp.Code != 200 {
		return fmt.Errorf("API error: %s (code: %d)", baseResp.Message, baseResp.Code)
	}

	// 如果状态码是200，直接将Data字段解析为目标结构体
	// 由于Data字段是json.RawMessage类型，这里避免了二次序列化
	return json.Unmarshal(baseResp.Data, v)
}
