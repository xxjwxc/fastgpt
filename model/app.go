package model

// AppStatsRequest 应用统计请求
type AppStatsRequest struct {
    StartTime int64 `json:"startTime"`
    EndTime   int64 `json:"endTime"`
    AppId     string `json:"appId,omitempty"`
}

// SourceCountMap 来源统计映射
type SourceCountMap struct {
    Test            int `json:"test"`
    Online          int `json:"online"`
    Share           int `json:"share"`
    Api             int `json:"api"`
    CronJob         int `json:"cronJob"`
    Team            int `json:"team"`
    Feishu          int `json:"feishu"`
    OfficialAccount int `json:"official_account"`
    Wecom           int `json:"wecom"`
    Mcp             int `json:"mcp"`
}

// UserSummary 用户统计摘要
type UserSummary struct {
    UserCount       int             `json:"userCount"`
    NewUserCount    int             `json:"newUserCount"`
    RetentionUserCount int          `json:"retentionUserCount"`
    Points          float64         `json:"points"`
    SourceCountMap  SourceCountMap  `json:"sourceCountMap"`
}

// UserData 用户统计数据
type UserData struct {
    Timestamp int64        `json:"timestamp"`
    Summary   UserSummary  `json:"summary"`
}

// ChatSummary 对话统计摘要
type ChatSummary struct {
    ChatItemCount int     `json:"chatItemCount"`
    ChatCount     int     `json:"chatCount"`
    ErrorCount    int     `json:"errorCount"`
    Points        float64 `json:"points"`
}

// ChatData 对话统计数据
type ChatData struct {
    Timestamp int64       `json:"timestamp"`
    Summary   ChatSummary `json:"summary"`
}

// AppSummary 应用统计摘要
type AppSummary struct {
    GoodFeedBackCount  int     `json:"goodFeedBackCount"`
    BadFeedBackCount   int     `json:"badFeedBackCount"`
    ChatCount          int     `json:"chatCount"`
    TotalResponseTime  float64 `json:"totalResponseTime"`
}

// AppData 应用统计数据
type AppData struct {
    Timestamp int64       `json:"timestamp"`
    Summary   AppSummary  `json:"summary"`
}

// AppStatsResponse 应用统计响应
type AppStatsResponse struct {
    Code       int        `json:"code"`
    StatusText string     `json:"statusText"`
    Message    string     `json:"message"`
    Data       struct {
        UserData []UserData `json:"userData"`
        ChatData []ChatData `json:"chatData"`
        AppData  []AppData  `json:"appData"`
    } `json:"data"`
}
