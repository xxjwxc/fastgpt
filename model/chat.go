// Package model 定义FastGPT API的请求和响应数据结构
//
// 该包包含了所有与FastGPT API交互所需的数据模型，包括：
// - 应用相关模型
// - 对话相关模型
// - 知识库相关模型
//
// 所有模型均使用JSON标签，用于序列化和反序列化API请求和响应。
package model

// ChatRequest 对话请求模型
//
// 用于向FastGPT发送对话请求，包含应用ID、消息列表和模型配置等。
type ChatRequest struct {
	ChatId             string                 `json:"chatId,omitempty"`             // 对话ID，可选，用于使用FastGPT提供的上下文功能
	Stream             bool                   `json:"stream,omitempty"`             // 是否使用流式响应，默认为false
	Detail             bool                   `json:"detail,omitempty"`             // 是否返回中间值，默认为false
	ResponseChatItemId string                 `json:"responseChatItemId,omitempty"` // 响应消息ID，可选，用于指定本次对话的响应消息ID
	Variables          map[string]interface{} `json:"variables,omitempty"`          // 模块变量，用于替换模块中的变量
	Messages           []Message              `json:"messages,omitempty"`           // 消息列表，包含历史对话记录
}

// Message 消息结构体
//
// 用于表示对话中的单条消息，包含角色和内容。
type Message struct {
	Role    string      `json:"role"`    // 消息角色，可选值：user, assistant, system
	Content interface{} `json:"content"` // 消息内容，支持字符串或结构化内容
}

// ContentItem 结构化内容项
//
// 用于表示消息中的结构化内容，如文本、图片URL或文件URL。
type ContentItem struct {
	Type     string    `json:"type"`                // 内容类型：text, image_url, file_url
	Text     string    `json:"text,omitempty"`      // 文本内容，当type为text时使用
	ImageURL *ImageURL `json:"image_url,omitempty"` // 图片URL，当type为image_url时使用
	FileURL  *FileURL  `json:"file_url,omitempty"`  // 文件URL，当type为file_url时使用
}

// ImageURL 图片URL结构体
//
// 用于表示消息中的图片URL。
type ImageURL struct {
	URL string `json:"url"` // 图片URL
}

// FileURL 文件URL结构体
//
// 用于表示消息中的文件URL。
type FileURL struct {
	Name string `json:"name"` // 文件名
	URL  string `json:"url"`  // 文件URL
}

// Delta 增量内容模型
//
// 用于表示流式响应中的增量内容，包含角色和内容。
type Delta struct {
	Role    string `json:"role,omitempty"`    // 角色，仅在第一条消息中出现
	Content string `json:"content,omitempty"` // 增量内容
}

// Choice 对话选择模型
//
// 用于表示对话响应中的选择项，包含增量内容、索引和结束原因。
type Choice struct {
	Delta        Delta  `json:"delta"`         // 增量内容
	Index        int    `json:"index"`         // 选择项索引
	FinishReason string `json:"finish_reason"` // 结束原因，如stop, length等
}

// AnswerEvent 回答事件模型
//
// 用于表示流式响应中的回答事件，包含生成的内容和相关元数据。
type AnswerEvent struct {
	ID      string   `json:"id"`      // 事件ID
	Object  string   `json:"object"`  // 对象类型，如chat.completion.chunk
	Created int64    `json:"created"` // 创建时间戳
	Model   string   `json:"model"`   // 使用的模型名称
	Choices []Choice `json:"choices"` // 生成的选择项列表
}

// FlowNodeStatusEvent 流程节点状态事件模型
//
// 用于表示流程节点的状态变化事件，包含节点状态和名称。
type FlowNodeStatusEvent struct {
	Status string `json:"status"` // 节点状态，如running, completed, failed等
	Name   string `json:"name"`   // 节点名称
}

// HistoryPreview 历史预览模型
//
// 用于表示历史对话的预览信息，包含对象类型和值。
type HistoryPreview struct {
	Obj   string `json:"obj"`   // 对象类型，如message
	Value string `json:"value"` // 对象值，如消息内容
}

// FlowResponse 流程响应模型
//
// 用于表示流程执行的响应结果，包含节点信息、模型使用情况和运行时间等。
type FlowResponse struct {
	NodeId          string           `json:"nodeId"`                 // 节点ID
	ModuleName      string           `json:"moduleName"`             // 模块名称
	ModuleType      string           `json:"moduleType"`             // 模块类型
	TotalPoints     float64          `json:"totalPoints"`            // 消耗的总点数
	Model           string           `json:"model"`                  // 使用的模型名称
	Tokens          int              `json:"tokens"`                 // 生成的token数
	Query           string           `json:"query"`                  // 查询内容
	MaxToken        int              `json:"maxToken"`               // 最大token数限制
	HistoryPreview  []HistoryPreview `json:"historyPreview"`         // 历史预览列表
	ContextTotalLen int              `json:"contextTotalLen"`        // 上下文总长度
	RunningTime     float64          `json:"runningTime"`            // 运行时间，单位秒
	PluginOutput    interface{}      `json:"pluginOutput,omitempty"` // 插件输出，可选
}

// FlowResponsesEvent 流程响应事件模型
//
// 用于表示流程执行的响应事件，包含多个流程响应。
type FlowResponsesEvent struct {
	Responses []FlowResponse `json:"responses"` // 流程响应列表
}

// Usage 对话使用情况模型
//
// 用于表示对话的token使用情况。
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`     // 提示token数
	CompletionTokens int `json:"completion_tokens"` // 完成token数
	TotalTokens      int `json:"total_tokens"`      // 总token数
}

// ChatResponse 基本聊天响应模型
//
// 用于表示不带Detail的聊天响应。
type ChatResponse struct {
	ID      string   `json:"id"`      // 响应ID
	Model   string   `json:"model"`   // 模型名称
	Usage   Usage    `json:"usage"`   // 使用情况
	Choices []Choice `json:"choices"` // 选择项列表
}

// QuoteItem 引用列表项模型
//
// 用于表示对话响应中的引用内容。
type QuoteItem struct {
	DatasetID    string  `json:"dataset_id,omitempty"`   // 数据集ID
	ID           string  `json:"id,omitempty"`           // 引用ID
	Q            string  `json:"q,omitempty"`            // 引用问题
	A            string  `json:"a,omitempty"`            // 引用答案
	Source       string  `json:"source,omitempty"`       // 引用来源
	CollectionID string  `json:"collectionId,omitempty"` // 集合ID
	SourceName   string  `json:"sourceName,omitempty"`   // 来源名称
	SourceID     string  `json:"sourceId,omitempty"`     // 来源ID
	Score        float64 `json:"score,omitempty"`        // 相似度分数
}

// CompleteMessage 完整消息模型
//
// 用于表示对话响应中的完整消息。
type CompleteMessage struct {
	Obj   string `json:"obj"`   // 对象类型
	Value string `json:"value"` // 对象值
}

// ResponseDataItem 响应数据项模型
//
// 用于表示带Detail的响应中的数据项。
type ResponseDataItem struct {
	ModuleName       string            `json:"moduleName"`                 // 模块名称
	Price            float64           `json:"price,omitempty"`            // 价格
	Model            string            `json:"model,omitempty"`            // 模型名称
	Tokens           int               `json:"tokens,omitempty"`           // Token数
	Similarity       float64           `json:"similarity,omitempty"`       // 相似度
	Limit            int               `json:"limit,omitempty"`            // 限制
	Question         string            `json:"question,omitempty"`         // 问题
	Answer           string            `json:"answer,omitempty"`           // 答案
	MaxToken         int               `json:"maxToken,omitempty"`         // 最大Token数
	QuoteList        []QuoteItem       `json:"quoteList,omitempty"`        // 引用列表
	CompleteMessages []CompleteMessage `json:"completeMessages,omitempty"` // 完整消息列表
	NodeID           string            `json:"nodeId,omitempty"`           // 节点ID
	ModuleType       string            `json:"moduleType,omitempty"`       // 模块类型
	TotalPoints      float64           `json:"totalPoints,omitempty"`      // 总点数
	Query            string            `json:"query,omitempty"`            // 查询内容
	HistoryPreview   []HistoryPreview  `json:"historyPreview,omitempty"`   // 历史预览
	ContextTotalLen  int               `json:"contextTotalLen,omitempty"`  // 上下文总长度
	RunningTime      float64           `json:"runningTime,omitempty"`      // 运行时间
	PluginOutput     interface{}       `json:"pluginOutput,omitempty"`     // 插件输出
}

// ChatDetailResponse 带Detail的聊天响应模型
//
// 用于表示带Detail的聊天响应。
type ChatDetailResponse struct {
	ResponseData []ResponseDataItem     `json:"responseData"`           // 响应数据列表
	NewVariables map[string]interface{} `json:"newVariables,omitempty"` // 新变量
	ChatResponse                        // 嵌入基本聊天响应
}

// Interactive 交互节点响应模型
//
// 用于表示工作流中交互节点的响应。
type Interactive struct {
	Type   string      `json:"type"`   // 交互类型：userSelect, userInput
	Params interface{} `json:"params"` // 交互参数，根据type不同而不同
}

// UserSelectParams 用户选择参数模型
//
// 用于表示交互节点中的用户选择参数。
type UserSelectParams struct {
	Description       string             `json:"description"`       // 描述
	UserSelectOptions []UserSelectOption `json:"userSelectOptions"` // 用户选择选项
}

// UserSelectOption 用户选择选项模型
//
// 用于表示用户选择参数中的选项。
type UserSelectOption struct {
	Value string `json:"value"` // 选项值
	Key   string `json:"key"`   // 选项键
}

// InputFormItem 输入表单项模型
//
// 用于表示交互节点中的输入表单项。
type InputFormItem struct {
	Type         string       `json:"type"`         // 输入类型：input, numberInput
	Key          string       `json:"key"`          // 表单键
	Label        string       `json:"label"`        // 表单标签
	Description  string       `json:"description"`  // 表单描述
	Value        interface{}  `json:"value"`        // 表单值
	DefaultValue interface{}  `json:"defaultValue"` // 表单默认值
	ValueType    string       `json:"valueType"`    // 值类型：string, number
	Required     bool         `json:"required"`     // 是否必填
	List         []ListOption `json:"list"`         // 选项列表
}

// ListOption 列表选项模型
//
// 用于表示输入表单项中的选项列表。
type ListOption struct {
	Label string `json:"label"` // 选项标签
	Value string `json:"value"` // 选项值
}

// UserInputParams 用户输入参数模型
//
// 用于表示交互节点中的用户输入参数。
type UserInputParams struct {
	Description string          `json:"description"` // 描述
	InputForm   []InputFormItem `json:"inputForm"`   // 输入表单
}

// GetHistoriesRequest 获取历史记录请求模型
//
// 用于请求获取应用的历史对话记录。
type GetHistoriesRequest struct {
	AppId    string `json:"appId"`    // 应用ID
	Offset   int    `json:"offset"`   // 偏移量
	PageSize int    `json:"pageSize"` // 每页数量
	Source   string `json:"source"`   // 对话源，如api
}

// ChatHistory 聊天历史记录模型
//
// 用于表示应用的历史对话记录。
type ChatHistory struct {
	ChatId      string `json:"chatId"`      // 对话ID
	UpdateTime  string `json:"updateTime"`  // 更新时间
	AppId       string `json:"appId"`       // 应用ID
	CustomTitle string `json:"customTitle"` // 自定义标题
	Title       string `json:"title"`       // 标题
	Top         bool   `json:"top"`         // 是否置顶
}

// GetHistoriesResponse 获取历史记录响应模型
//
// 用于表示获取历史记录的响应。
type GetHistoriesResponse struct {
	List  []ChatHistory `json:"list"`  // 历史记录列表
	Total int           `json:"total"` // 总记录数
}

// UpdateHistoryRequest 更新历史记录请求模型
//
// 用于更新对话历史记录，如修改标题或置顶状态。
type UpdateHistoryRequest struct {
	AppId       string `json:"appId"`                 // 应用ID
	ChatId      string `json:"chatId"`                // 对话ID
	CustomTitle string `json:"customTitle,omitempty"` // 自定义标题
	Top         *bool  `json:"top,omitempty"`         // 是否置顶
}

// ChatInitResponse 获取对话初始化信息响应模型
//
// 用于表示获取单个对话初始化信息的响应。
type ChatInitResponse struct {
	ChatId    string                 `json:"chatId"`    // 对话ID
	AppId     string                 `json:"appId"`     // 应用ID
	Variables map[string]interface{} `json:"variables"` // 变量
	App       ChatAppInfo            `json:"app"`       // 应用信息
}

// ChatAppInfo 聊天应用信息模型
//
// 用于表示对话初始化信息中的应用信息。
type ChatAppInfo struct {
	ChatConfig   ChatConfig `json:"chatConfig"`   // 聊天配置
	ChatModels   []string   `json:"chatModels"`   // 聊天模型列表
	Name         string     `json:"name"`         // 应用名称
	Avatar       string     `json:"avatar"`       // 应用头像
	Intro        string     `json:"intro"`        // 应用介绍
	Type         string     `json:"type"`         // 应用类型
	PluginInputs []string   `json:"pluginInputs"` // 插件输入列表
}

// ChatConfig 聊天配置模型
//
// 用于表示聊天应用的配置信息。
type ChatConfig struct {
	QuestionGuide    bool             `json:"questionGuide"`    // 是否开启问题引导
	TTSConfig        TTSConfig        `json:"ttsConfig"`        // TTS配置
	WhisperConfig    WhisperConfig    `json:"whisperConfig"`    // Whisper配置
	ChatInputGuide   ChatInputGuide   `json:"chatInputGuide"`   // 聊天输入引导
	Instruction      string           `json:"instruction"`      // 指令
	Variables        []interface{}    `json:"variables"`        // 变量列表
	FileSelectConfig FileSelectConfig `json:"fileSelectConfig"` // 文件选择配置
	WelcomeText      string           `json:"welcomeText"`      // 欢迎文本
}

// TTSConfig TTS配置模型
//
// 用于表示聊天应用的TTS配置。
type TTSConfig struct {
	Type string `json:"type"` // TTS类型
}

// WhisperConfig Whisper配置模型
//
// 用于表示聊天应用的Whisper配置。
type WhisperConfig struct {
	Open            bool `json:"open"`            // 是否开启
	AutoSend        bool `json:"autoSend"`        // 是否自动发送
	AutoTTSResponse bool `json:"autoTTSResponse"` // 是否自动TTS响应
}

// ChatInputGuide 聊天输入引导模型
//
// 用于表示聊天应用的输入引导配置。
type ChatInputGuide struct {
	Open      bool     `json:"open"`      // 是否开启
	TextList  []string `json:"textList"`  // 文本列表
	CustomUrl string   `json:"customUrl"` // 自定义URL
}

// FileSelectConfig 文件选择配置模型
//
// 用于表示聊天应用的文件选择配置。
type FileSelectConfig struct {
	CanSelectFile bool `json:"canSelectFile"` // 是否可以选择文件
	CanSelectImg  bool `json:"canSelectImg"`  // 是否可以选择图片
	MaxFiles      int  `json:"maxFiles"`      // 最大文件数
}

// GetPaginationRecordsRequest 获取对话记录列表请求模型
//
// 用于请求获取对话记录列表。
type GetPaginationRecordsRequest struct {
	AppId               string `json:"appId"`               // 应用ID
	ChatId              string `json:"chatId"`              // 对话ID
	Offset              int    `json:"offset"`              // 偏移量
	PageSize            int    `json:"pageSize"`            // 每页数量
	LoadCustomFeedbacks bool   `json:"loadCustomFeedbacks"` // 是否加载自定义反馈
}

// ChatRecord 聊天记录模型
//
// 用于表示对话中的单条记录。
type ChatRecord struct {
	ID                   string        `json:"_id"`                            // 记录ID
	DataId               string        `json:"dataId"`                         // 数据ID
	Obj                  string        `json:"obj"`                            // 对象类型：Human, AI
	Value                interface{}   `json:"value"`                          // 记录值
	CustomFeedbacks      []interface{} `json:"customFeedbacks"`                // 自定义反馈
	LLMModuleAccount     int           `json:"llmModuleAccount,omitempty"`     // LLM模块账号
	TotalQuoteList       []interface{} `json:"totalQuoteList,omitempty"`       // 总引用列表
	TotalRunningTime     float64       `json:"totalRunningTime,omitempty"`     // 总运行时间
	HistoryPreviewLength int           `json:"historyPreviewLength,omitempty"` // 历史预览长度
}

// GetPaginationRecordsResponse 获取对话记录列表响应模型
//
// 用于表示获取对话记录列表的响应。
type GetPaginationRecordsResponse struct {
	List  []ChatRecord `json:"list"`  // 记录列表
	Total int          `json:"total"` // 总记录数
}

// UpdateUserFeedbackRequest 更新用户反馈请求模型
//
// 用于更新用户对对话记录的反馈，如点赞或点踩。
type UpdateUserFeedbackRequest struct {
	AppId            string `json:"appId"`                      // 应用ID
	ChatId           string `json:"chatId"`                     // 对话ID
	DataId           string `json:"dataId"`                     // 数据ID
	UserGoodFeedback string `json:"userGoodFeedback,omitempty"` // 用户点赞反馈
	UserBadFeedback  string `json:"userBadFeedback,omitempty"`  // 用户点踩反馈
}

// CreateQuestionGuideRequest 创建猜你想问请求模型
//
// 用于请求生成猜你想问的问题。
type CreateQuestionGuideRequest struct {
	AppId         string               `json:"appId"`                   // 应用ID
	ChatId        string               `json:"chatId"`                  // 对话ID
	QuestionGuide *QuestionGuideConfig `json:"questionGuide,omitempty"` // 问题引导配置
}

// QuestionGuideConfig 问题引导配置模型
//
// 用于配置问题引导的生成。
type QuestionGuideConfig struct {
	Open         bool   `json:"open"`                   // 是否开启
	Model        string `json:"model,omitempty"`        // 模型名称
	CustomPrompt string `json:"customPrompt,omitempty"` // 自定义提示
}

// CreateQuestionGuideResponse 创建猜你想问响应模型
//
// 用于表示生成猜你想问问题的响应。
type CreateQuestionGuideResponse struct {
	Questions []string `json:"questions"` // 生成的问题列表
}
