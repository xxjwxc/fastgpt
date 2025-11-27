package app

import (
	"github.com/xxjwxc/fastgpt/client"
	"github.com/xxjwxc/fastgpt/model"
)

// AppAPI 应用接口
type AppAPI struct {
	client *client.Client
}

// NewAppAPI 创建应用接口实例
func NewAppAPI(c *client.Client) *AppAPI {
	return &AppAPI{client: c}
}

// GetStats 获取应用统计数据
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/app
func (api *AppAPI) GetStats(req *model.AppStatsRequest) (*model.AppStatsResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/app/stats", req)
	if err != nil {
		return nil, err
	}

	var statsResp model.AppStatsResponse
	if err := api.client.ParseResponse(resp, &statsResp); err != nil {
		return nil, err
	}

	return &statsResp, nil
}
