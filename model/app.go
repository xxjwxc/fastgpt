// Package model 定义FastGPT API的请求和响应数据结构
//
// 该包包含了所有与FastGPT API交互所需的数据模型，包括：
// - 应用相关模型
// - 对话相关模型
// - 知识库相关模型
//
// 所有模型均使用JSON标签，用于序列化和反序列化API请求和响应。
package model

// AppTotalDataRequest 获取累积运行结果请求模型
//
// 用于请求获取应用的累积运行结果。
type AppTotalDataRequest struct {
	AppId string `json:"appId"` // 应用ID
}

// AppTotalDataResponse 获取累积运行结果响应模型
//
// 用于返回应用累积运行结果的响应。
type AppTotalDataResponse struct {
	Code       int    `json:"code"`       // 响应状态码，200表示成功
	StatusText string `json:"statusText"` // 状态文本
	Message    string `json:"message"`    // 响应消息
	Data       struct {
		TotalUsers  int `json:"totalUsers"`  // 累积使用用户数量
		TotalChats  int `json:"totalChats"`  // 累积对话数量
		TotalPoints int `json:"totalPoints"` // 累积积分消耗
	} `json:"data"` // 响应数据
}

// SourceCountMap 来源统计映射模型
//
// 用于统计不同来源的访问次数，包括测试、线上、分享等渠道。
type SourceCountMap struct {
	Test            int `json:"test"`             // 测试来源访问次数
	Online          int `json:"online"`           // 线上来源访问次数
	Share           int `json:"share"`            // 分享来源访问次数
	Api             int `json:"api"`              // API来源访问次数
	CronJob         int `json:"cronJob"`          // 定时任务来源访问次数
	Team            int `json:"team"`             // 团队来源访问次数
	Feishu          int `json:"feishu"`           // 飞书来源访问次数
	OfficialAccount int `json:"official_account"` // 公众号来源访问次数
	Wecom           int `json:"wecom"`            // 企业微信来源访问次数
	Mcp             int `json:"mcp"`              // MCP来源访问次数
}

// UserSummary 用户统计摘要模型
//
// 包含用户相关的统计数据，如用户数量、新增用户数、留存用户数等。
type UserSummary struct {
	UserCount          int            `json:"userCount"`          // 活跃用户数量
	NewUserCount       int            `json:"newUserCount"`       // 新用户数量
	RetentionUserCount int            `json:"retentionUserCount"` // 留存用户数量
	Points             float64        `json:"points"`             // 总积分消耗
	SourceCountMap     SourceCountMap `json:"sourceCountMap"`     // 各来源用户数量
}

// UserData 用户统计数据模型
//
// 包含指定时间点的用户统计摘要。
type UserData struct {
	Timestamp int64       `json:"timestamp"` // 时间戳，毫秒级
	Summary   UserSummary `json:"summary"`   // 用户统计摘要
}

// ChatSummary 对话统计摘要模型
//
// 包含对话相关的统计数据，如对话次数、对话项数量、错误次数等。
type ChatSummary struct {
	ChatItemCount int     `json:"chatItemCount"` // 对话次数
	ChatCount     int     `json:"chatCount"`     // 会话次数
	ErrorCount    int     `json:"errorCount"`    // 错误对话次数
	Points        float64 `json:"points"`        // 总积分消耗
}

// ChatData 对话统计数据模型
//
// 包含指定时间点的对话统计摘要。
type ChatData struct {
	Timestamp int64       `json:"timestamp"` // 时间戳，毫秒级
	Summary   ChatSummary `json:"summary"`   // 对话统计摘要
}

// AppSummary 应用统计摘要模型
//
// 包含应用相关的统计数据，如反馈数量、对话次数、响应时间等。
type AppSummary struct {
	GoodFeedBackCount int     `json:"goodFeedBackCount"` // 好评反馈数量
	BadFeedBackCount  int     `json:"badFeedBackCount"`  // 差评反馈数量
	ChatCount         int     `json:"chatCount"`         // 对话次数
	TotalResponseTime float64 `json:"totalResponseTime"` // 总响应时间
}

// AppData 应用统计数据模型
//
// 包含指定时间点的应用统计摘要。
type AppData struct {
	Timestamp int64      `json:"timestamp"` // 时间戳，毫秒级
	Summary   AppSummary `json:"summary"`   // 应用统计摘要
}

// AppChartDataRequest 获取应用日志看板请求模型
//
// 用于请求获取指定时间范围内的应用日志看板数据。
type AppChartDataRequest struct {
	AppId        string   `json:"appId"`        // 应用Id
	DateStart    string   `json:"dateStart"`    // 开始时间，ISO格式
	DateEnd      string   `json:"dateEnd"`      // 结束时间，ISO格式
	Offset       int      `json:"offset"`       // 用户留存偏移量
	Source       []string `json:"source"`       // 日志来源
	UserTimespan string   `json:"userTimespan"` // 用户数据时间跨度：day｜week｜month｜quarter
	ChatTimespan string   `json:"chatTimespan"` // 对话数据时间跨度：day｜week｜month｜quarter
	AppTimespan  string   `json:"appTimespan"`  // 应用数据时间跨度：day｜week｜month｜quarter
}

// AppChartDataResponse 获取应用日志看板响应模型
//
// 用于返回应用日志看板数据的响应，包含用户数据、对话数据和应用数据。
type AppChartDataResponse struct {
	Code       int    `json:"code"`       // 响应状态码，200表示成功
	StatusText string `json:"statusText"` // 状态文本
	Message    string `json:"message"`    // 响应消息
	Data       struct {
		UserData []UserData `json:"userData"` // 用户数据数组
		ChatData []ChatData `json:"chatData"` // 对话数据数组
		AppData  []AppData  `json:"appData"`  // 应用数据数组
	} `json:"data"` // 响应数据
}
