// Package app 提供FastGPT应用管理相关的API接口
//
// 该包封装了与应用管理相关的所有API，包括：
// - 累积运行结果查询
// - 应用日志看板获取
//
// 所有API均需要通过FastGPT客户端实例访问，使用前需先创建客户端。
package app

import (
	"fmt"

	"github.com/xxjwxc/fastgpt/client"
	"github.com/xxjwxc/fastgpt/model"
)

// AppAPI 应用接口结构体，封装了所有应用相关的API方法
//
// 该结构体通过组合HTTP客户端，提供了与FastGPT应用管理相关的所有功能。
type AppAPI struct {
	client *client.Client // HTTP客户端，用于发送API请求
}

// NewAppAPI 创建应用接口实例
//
// 参数：
//
//	c: HTTP客户端实例，由外部传入
//
// 返回值：
//
//	*AppAPI: 应用接口实例，用于访问应用相关API
//
// 使用示例：
//
//	c := client.NewClient("https://cloud.fastgpt.cn", "sk-xxx")
//	appAPI := app.NewAppAPI(c)
func NewAppAPI(c *client.Client) *AppAPI {
	return &AppAPI{client: c}
}

// GetTotalData 获取累积运行结果
//
// 该方法用于获取应用的累积运行结果，包括累积使用用户数量、累积对话数量和累积积分消耗。
//
// 参数：
//
//	req: 获取累积运行结果请求，包含应用ID
//
// 返回值：
//
//	*model.AppTotalDataResponse: 累积运行结果响应
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/app#%E8%8E%B7%E5%8F%96%E7%B4%AF%E7%A7%AF%E8%BF%90%E8%A1%8C%E7%BB%93%E6%9E%9C
//
// 使用示例：
//
//	req := &model.AppTotalDataRequest{
//	    AppId: "your-app-id",
//	}
//	resp, err := appAPI.GetTotalData(req)
func (api *AppAPI) GetTotalData(req *model.AppTotalDataRequest) (*model.AppTotalDataResponse, error) {
	// 发送HTTP请求到FastGPT服务器
	resp, err := api.client.DoRequest("GET", fmt.Sprintf("/api/proApi/core/app/logs/getTotalData?appId=%s", req.AppId), nil)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	// 解析响应数据
	var totalDataResp model.AppTotalDataResponse
	if err := api.client.ParseResponse(resp, &totalDataResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	return &totalDataResp, nil // 返回累积运行结果
}

// GetChartData 获取应用日志看板
//
// 该方法用于获取应用的日志看板数据，包括指定时间范围内的用户统计、对话统计和应用统计。
//
// 参数：
//
//	req: 获取应用日志看板请求，包含应用ID、开始时间、结束时间等
//
// 返回值：
//
//	*model.AppChartDataResponse: 应用日志看板响应，包含用户数据、对话数据和应用数据
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/app#%E8%8E%B7%E5%8F%96%E5%BA%94%E7%94%A8%E6%97%A5%E5%BF%97%E7%9C%8B%E6%9D%BF
//
// 使用示例：
//
//	req := &model.AppChartDataRequest{
//	    AppId:          "your-app-id",
//	    DateStart:      "2025-09-19T16:00:00.000Z",
//	    DateEnd:        "2025-09-27T15:59:59.999Z",
//	    Offset:         1,
//	    Source:         []string{"test", "online", "share", "api", "cronJob", "team", "feishu", "official_account", "wecom", "mcp"},
//	    UserTimespan:   "day",
//	    ChatTimespan:   "day",
//	    AppTimespan:    "day",
//	}
//	resp, err := appAPI.GetChartData(req)
func (api *AppAPI) GetChartData(req *model.AppChartDataRequest) (*model.AppChartDataResponse, error) {
	// 发送HTTP请求到FastGPT服务器
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/app/logs/getChartData", req)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	// 解析响应数据
	var chartDataResp model.AppChartDataResponse
	if err := api.client.ParseResponse(resp, &chartDataResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	return &chartDataResp, nil // 返回日志看板数据
}
