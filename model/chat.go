package model

// ChatRequest 对话请求
type ChatRequest struct {
    AppId     string    `json:"appId"`
    Messages  []Message `json:"messages"`
    Stream    bool      `json:"stream,omitempty"`
    Temperature float64  `json:"temperature,omitempty"`
    TopP      float64   `json:"topP,omitempty"`
    MaxTokens int       `json:"maxTokens,omitempty"`
}

// Message 消息结构体
type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

// Delta 增量内容
type Delta struct {
    Role    string `json:"role,omitempty"`
    Content string `json:"content,omitempty"`
}

// Choice 对话选择
type Choice struct {
    Delta        Delta  `json:"delta"`
    Index        int    `json:"index"`
    FinishReason string `json:"finish_reason"`
}

// AnswerEvent 回答事件
type AnswerEvent struct {
    ID      string   `json:"id"`
    Object  string   `json:"object"`
    Created int64    `json:"created"`
    Model   string   `json:"model"`
    Choices []Choice `json:"choices"`
}

// FlowNodeStatusEvent 流程节点状态事件
type FlowNodeStatusEvent struct {
    Status string `json:"status"`
    Name   string `json:"name"`
}

// HistoryPreview 历史预览
type HistoryPreview struct {
    Obj   string `json:"obj"`
    Value string `json:"value"`
}

// FlowResponse 流程响应
type FlowResponse struct {
    NodeId         string           `json:"nodeId"`
    ModuleName     string           `json:"moduleName"`
    ModuleType     string           `json:"moduleType"`
    TotalPoints    float64          `json:"totalPoints"`
    Model          string           `json:"model"`
    Tokens         int              `json:"tokens"`
    Query          string           `json:"query"`
    MaxToken       int              `json:"maxToken"`
    HistoryPreview []HistoryPreview `json:"historyPreview"`
    ContextTotalLen int             `json:"contextTotalLen"`
    RunningTime    float64          `json:"runningTime"`
    PluginOutput   interface{}      `json:"pluginOutput,omitempty"`
}

// FlowResponsesEvent 流程响应事件
type FlowResponsesEvent struct {
    Responses []FlowResponse `json:"responses"`
}
