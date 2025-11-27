// Package model 定义FastGPT API的请求和响应数据结构
//
// 该包包含了所有与FastGPT API交互所需的数据模型，包括：
// - 应用相关模型
// - 对话相关模型
// - 知识库相关模型
//
// 所有模型均使用JSON标签，用于序列化和反序列化API请求和响应。
package model

import "encoding/json"

// BaseResponse 基础响应模型
//
// 用于表示API响应的基础结构，包含状态码、状态文本、消息和数据。
type BaseResponse struct {
	Code       int             `json:"code"`       // 状态码，200表示成功
	StatusText string          `json:"statusText"` // 状态文本
	Message    string          `json:"message"`    // 消息
	Data       json.RawMessage `json:"data"`       // 响应数据，使用RawMessage避免二次序列化
}

// 训练订单相关模型

// DatasetTrainOrderRequest 训练订单创建请求模型
//
// 用于请求创建训练订单。
type DatasetTrainOrderRequest struct {
	DatasetId string `json:"datasetId"`      // 知识库ID
	Name      string `json:"name,omitempty"` // 可选，自定义订单名称
}

// 知识库相关模型

// DatasetCreateRequest 知识库创建请求模型
//
// 用于请求创建一个新的知识库。
type DatasetCreateRequest struct {
	ParentId    *string `json:"parentId,omitempty"`    // 父级ID，用于构建目录结构
	Type        string  `json:"type,omitempty"`        // dataset或者folder，代表普通知识库和文件夹
	Name        string  `json:"name"`                  // 知识库名（必填）
	Intro       string  `json:"intro,omitempty"`       // 介绍（可选）
	Avatar      string  `json:"avatar,omitempty"`      // 头像地址（可选）
	VectorModel string  `json:"vectorModel,omitempty"` // 向量模型（建议传空，用系统默认的）
	AgentModel  string  `json:"agentModel,omitempty"`  // 文本处理模型（建议传空，用系统默认的）
	VlmModel    string  `json:"vlmModel,omitempty"`    // 图片理解模型（建议传空，用系统默认的）
}

// VectorModel 向量模型信息
//
// 用于表示知识库的向量模型信息。
type VectorModel struct {
	Model            string  `json:"model"`            // 模型名称
	Name             string  `json:"name"`             // 模型显示名称
	CharsPointsPrice float64 `json:"charsPointsPrice"` // 字符点数价格
	DefaultToken     int     `json:"defaultToken"`     // 默认Token数
	MaxToken         int     `json:"maxToken"`         // 最大Token数
	Weight           int     `json:"weight"`           // 权重
}

// AgentModel 文本处理模型信息
//
// 用于表示知识库的文本处理模型信息。
type AgentModel struct {
	Model            string  `json:"model"`            // 模型名称
	Name             string  `json:"name"`             // 模型显示名称
	MaxContext       int     `json:"maxContext"`       // 最大上下文长度
	MaxResponse      int     `json:"maxResponse"`      // 最大响应长度
	CharsPointsPrice float64 `json:"charsPointsPrice"` // 字符点数价格
}

// DatasetInfo 知识库信息模型
//
// 用于表示知识库的详细信息。
type DatasetInfo struct {
	ID          string      `json:"_id"`                  // 知识库ID
	ParentId    *string     `json:"parentId"`             // 父级ID
	Avatar      string      `json:"avatar"`               // 头像地址
	Name        string      `json:"name"`                 // 知识库名称
	Intro       string      `json:"intro"`                // 知识库介绍
	Type        string      `json:"type"`                 // 知识库类型
	Permission  string      `json:"permission"`           // 权限
	CanWrite    bool        `json:"canWrite"`             // 是否可写
	IsOwner     bool        `json:"isOwner"`              // 是否是所有者
	VectorModel VectorModel `json:"vectorModel"`          // 向量模型信息
	AgentModel  *AgentModel `json:"agentModel,omitempty"` // 文本处理模型信息
	Status      string      `json:"status,omitempty"`     // 状态
	TeamId      string      `json:"teamId,omitempty"`     // 团队ID
	TmbId       string      `json:"tmbId,omitempty"`      // 成员ID
	UpdateTime  string      `json:"updateTime,omitempty"` // 更新时间
}

// DatasetListRequest 知识库列表请求模型
//
// 用于请求获取知识库列表。
type DatasetListRequest struct {
	ParentId string `json:"parentId"` // 父级ID，传空字符串或者null，代表获取根目录下的知识库
}

// DatasetDetailRequest 知识库详情请求模型
//
// 用于请求获取知识库详情。
type DatasetDetailRequest struct {
	Id string `json:"id"` // 知识库ID
}

// DatasetDeleteRequest 知识库删除请求模型
//
// 用于请求删除知识库。
type DatasetDeleteRequest struct {
	Id string `json:"id"` // 知识库ID
}

// 集合相关模型

// CollectionCreateRequest 集合创建请求模型
//
// 用于请求创建一个空的集合。
type CollectionCreateRequest struct {
	DatasetId string                 `json:"datasetId"`          // 知识库的ID(必填)
	ParentId  *string                `json:"parentId,omitempty"` // 父级ID，不填则默认为根目录
	Name      string                 `json:"name"`               // 集合名称（必填）
	Type      string                 `json:"type"`               // 集合类型：folder, virtual
	Metadata  map[string]interface{} `json:"metadata,omitempty"` // 元数据
}

// CollectionCreateTextRequest 纯文本集合创建请求模型
//
// 用于请求创建一个纯文本集合。
type CollectionCreateTextRequest struct {
	Text             string                 `json:"text"`                       // 原文本
	DatasetId        string                 `json:"datasetId"`                  // 知识库的ID(必填)
	ParentId         *string                `json:"parentId,omitempty"`         // 父级ID，不填则默认为根目录
	Name             string                 `json:"name"`                       // 集合名称（必填）
	TrainingType     string                 `json:"trainingType"`               // 数据处理方式：chunk, qa
	ChunkSettingMode string                 `json:"chunkSettingMode,omitempty"` // 分块参数模式：auto, custom
	ChunkSplitMode   string                 `json:"chunkSplitMode,omitempty"`   // 分块拆分模式：size, char
	ChunkSize        int                    `json:"chunkSize,omitempty"`        // 分块大小
	IndexSize        int                    `json:"indexSize,omitempty"`        // 索引大小
	ChunkSplitter    string                 `json:"chunkSplitter,omitempty"`    // 自定义最高优先分割符号
	QAPrompt         string                 `json:"qaPrompt,omitempty"`         // qa拆分提示词
	Metadata         map[string]interface{} `json:"metadata,omitempty"`         // 元数据
}

// CollectionCreateLinkRequest 链接集合创建请求模型
//
// 用于请求创建一个链接集合。
type CollectionCreateLinkRequest struct {
	Link             string                 `json:"link"`                       // 网络链接
	DatasetId        string                 `json:"datasetId"`                  // 知识库的ID(必填)
	ParentId         *string                `json:"parentId,omitempty"`         // 父级ID，不填则默认为根目录
	TrainingType     string                 `json:"trainingType"`               // 数据处理方式：chunk, qa
	ChunkSettingMode string                 `json:"chunkSettingMode,omitempty"` // 分块参数模式：auto, custom
	ChunkSplitMode   string                 `json:"chunkSplitMode,omitempty"`   // 分块拆分模式：size, char
	ChunkSize        int                    `json:"chunkSize,omitempty"`        // 分块大小
	IndexSize        int                    `json:"indexSize,omitempty"`        // 索引大小
	ChunkSplitter    string                 `json:"chunkSplitter,omitempty"`    // 自定义最高优先分割符号
	QAPrompt         string                 `json:"qaPrompt,omitempty"`         // qa拆分提示词
	Metadata         map[string]interface{} `json:"metadata,omitempty"`         // 元数据，包含webPageSelector等
}

// CollectionCreateAPRequest API集合创建请求模型
//
// 用于请求创建一个API集合。
type CollectionCreateAPRequest struct {
	Name          string  `json:"name"`                    // 集合名，建议就用文件名，必填
	ApiFileId     string  `json:"apiFileId"`               // 文件的ID，必填
	DatasetId     string  `json:"datasetId"`               // 知识库的ID(必填)
	ParentId      *string `json:"parentId,omitempty"`      // 父级ID，不填则默认为根目录
	TrainingType  string  `json:"trainingType"`            // 训练模式（必填）
	ChunkSize     int     `json:"chunkSize,omitempty"`     // 每个chunk的长度
	ChunkSplitter string  `json:"chunkSplitter,omitempty"` // 自定义最高优先分割符号
	QAPrompt      string  `json:"qaPrompt,omitempty"`      // qa拆分自定义提示词
}

// CollectionCreateExternalFileRequest 外部文件集合创建请求模型
//
// 用于请求创建一个外部文件库集合（商业版）。
type CollectionCreateExternalFileRequest struct {
	ExternalFileUrl string   `json:"externalFileUrl"`          // 文件访问链接（可以是临时链接）
	ExternalFileId  string   `json:"externalFileId,omitempty"` // 外部文件ID
	Filename        string   `json:"filename,omitempty"`       // 自定义文件名，需要带后缀
	CreateTime      string   `json:"createTime,omitempty"`     // 文件创建时间
	DatasetId       string   `json:"datasetId"`                // 知识库的ID(必填)
	ParentId        *string  `json:"parentId,omitempty"`       // 父级ID，不填则默认为根目录
	Tags            []string `json:"tags,omitempty"`           // 集合标签
	TrainingType    string   `json:"trainingType"`             // 数据处理方式：chunk, qa
	ChunkSize       int      `json:"chunkSize,omitempty"`      // 分块大小
	ChunkSplitter   string   `json:"chunkSplitter,omitempty"`  // 自定义最高优先分割符号
	QAPrompt        string   `json:"qaPrompt,omitempty"`       // qa拆分提示词
}

// CollectionCreateResult 集合创建结果模型
//
// 用于表示集合创建的结果。
type CollectionCreateResult struct {
	InsertLen int      `json:"insertLen"` // 插入的块数量
	OverToken []string `json:"overToken"` // 超出token的项
	Repeat    []string `json:"repeat"`    // 重复的项
	Error     []string `json:"error"`     // 错误的项
}

// CollectionCreateResponse 集合创建响应模型
//
// 用于表示集合创建的响应。
type CollectionCreateResponse struct {
	CollectionId string                 `json:"collectionId"` // 新建的集合ID
	Results      CollectionCreateResult `json:"results"`      // 创建结果
}

// CollectionPermission 集合权限模型
//
// 用于表示集合的权限信息。
type CollectionPermission struct {
	Value        int  `json:"value"`        // 权限值
	IsOwner      bool `json:"isOwner"`      // 是否是所有者
	HasManagePer bool `json:"hasManagePer"` // 是否有管理权限
	HasWritePer  bool `json:"hasWritePer"`  // 是否有写权限
	HasReadPer   bool `json:"hasReadPer"`   // 是否有读权限
}

// CollectionInfo 集合信息模型
//
// 用于表示集合的详细信息。
type CollectionInfo struct {
	ID             string               `json:"_id"`                      // 集合ID
	ParentId       *string              `json:"parentId"`                 // 父级ID
	TmbId          string               `json:"tmbId"`                    // 成员ID
	Type           string               `json:"type"`                     // 集合类型
	Name           string               `json:"name"`                     // 集合名称
	UpdateTime     string               `json:"updateTime"`               // 更新时间
	DataAmount     int                  `json:"dataAmount"`               // 数据量
	TrainingAmount int                  `json:"trainingAmount"`           // 训练量
	ExternalFileId string               `json:"externalFileId,omitempty"` // 外部文件ID
	Tags           []string             `json:"tags,omitempty"`           // 标签
	Forbid         bool                 `json:"forbid"`                   // 是否禁用
	TrainingType   string               `json:"trainingType"`             // 训练类型
	Permission     CollectionPermission `json:"permission"`               // 权限信息
	RawLink        string               `json:"rawLink,omitempty"`        // 原始链接
	DatasetId      interface{}          `json:"datasetId,omitempty"`      // 知识库ID
	TeamId         string               `json:"teamId,omitempty"`         // 团队ID
	RawTextLength  int                  `json:"rawTextLength,omitempty"`  // 原始文本长度
	HashRawText    string               `json:"hashRawText,omitempty"`    // 原始文本哈希
	CreateTime     string               `json:"createTime,omitempty"`     // 创建时间
	CanWrite       bool                 `json:"canWrite,omitempty"`       // 是否可写
	SourceName     string               `json:"sourceName,omitempty"`     // 来源名称
	ChunkSize      int                  `json:"chunkSize,omitempty"`      // 分块大小
	ChunkSplitter  string               `json:"chunkSplitter,omitempty"`  // 分块分割符
	QAPrompt       string               `json:"qaPrompt,omitempty"`       // QA提示词
}

// CollectionListRequest 集合列表请求模型
//
// 用于请求获取集合列表。
type CollectionListRequest struct {
	Offset     int     `json:"offset"`               // 偏移量
	PageSize   int     `json:"pageSize"`             // 每页数量，最大30
	DatasetId  string  `json:"datasetId"`            // 知识库的ID(必填)
	ParentId   *string `json:"parentId,omitempty"`   // 父级Id
	SearchText string  `json:"searchText,omitempty"` // 模糊搜索文本
}

// CollectionListResponse 集合列表响应模型
//
// 用于表示集合列表的响应。
type CollectionListResponse struct {
	List  []CollectionInfo `json:"list"`  // 集合列表
	Total int              `json:"total"` // 总记录数
}

// CollectionUpdateRequest 集合更新请求模型
//
// 用于请求更新集合信息。
type CollectionUpdateRequest struct {
	ID             string   `json:"id,omitempty"`             // 集合的ID
	DatasetId      string   `json:"datasetId,omitempty"`      // 知识库ID
	ExternalFileId string   `json:"externalFileId,omitempty"` // 外部文件ID
	ParentId       *string  `json:"parentId,omitempty"`       // 修改父级ID
	Name           string   `json:"name,omitempty"`           // 修改集合名称
	Tags           []string `json:"tags,omitempty"`           // 修改集合标签
	Forbid         bool     `json:"forbid,omitempty"`         // 修改集合禁用状态
	CreateTime     string   `json:"createTime,omitempty"`     // 修改集合创建时间
}

// CollectionDeleteRequest 集合删除请求模型
//
// 用于请求删除集合。
type CollectionDeleteRequest struct {
	CollectionIds []string `json:"collectionIds"` // 集合的ID列表
}

// 数据相关模型

// Index 索引模型
//
// 用于表示数据的向量索引。
type Index struct {
	Type   string `json:"type,omitempty"`   // 索引类型：default, custom, summary, question, image
	DataId string `json:"dataId,omitempty"` // 关联的向量ID
	Text   string `json:"text"`             // 文本内容
	ID     string `json:"_id,omitempty"`    // 索引ID
}

// DatasetData 数据集数据模型
//
// 用于表示知识库中的数据。
type DatasetData struct {
	ID            string  `json:"_id,omitempty"`           // 数据ID
	TeamId        string  `json:"teamId,omitempty"`        // 团队ID
	TmbId         string  `json:"tmbId,omitempty"`         // 成员ID
	DatasetId     string  `json:"datasetId,omitempty"`     // 知识库ID
	CollectionId  string  `json:"collectionId,omitempty"`  // 集合ID
	Q             string  `json:"q"`                       // 主要数据
	A             string  `json:"a,omitempty"`             // 辅助数据
	FullTextToken string  `json:"fullTextToken,omitempty"` // 分词
	Indexes       []Index `json:"indexes"`                 // 向量索引
	UpdateTime    string  `json:"updateTime,omitempty"`    // 更新时间
	ChunkIndex    int     `json:"chunkIndex,omitempty"`    // 分块索引
	SourceName    string  `json:"sourceName,omitempty"`    // 来源名称
	SourceId      string  `json:"sourceId,omitempty"`      // 来源ID
	IsOwner       bool    `json:"isOwner,omitempty"`       // 是否是所有者
	CanWrite      bool    `json:"canWrite,omitempty"`      // 是否可写
}

// DataPushRequest 数据推送请求模型
//
// 用于请求为集合批量添加数据。
type DataPushRequest struct {
	CollectionId string        `json:"collectionId"`     // 集合ID（必填）
	TrainingType string        `json:"trainingType"`     // 训练模式（必填）
	Prompt       string        `json:"prompt,omitempty"` // 可选。qa拆分引导词，chunk模式下忽略
	BillId       string        `json:"billId,omitempty"` // 可选。如果有这个值，本次的数据会被聚合到一个订单中
	Data         []DatasetData `json:"data"`             // 具体数据
}

// DataPushResponse 数据推送响应模型
//
// 用于表示数据推送的响应。
type DataPushResponse struct {
	InsertLen int      `json:"insertLen"` // 最终插入成功的数量
	OverToken []string `json:"overToken"` // 超出token的
	Repeat    []string `json:"repeat"`    // 重复的数量
	Error     []string `json:"error"`     // 其他错误
}

// DataListRequest 数据列表请求模型
//
// 用于请求获取集合的数据列表。
type DataListRequest struct {
	Offset       int    `json:"offset"`               // 偏移量
	PageSize     int    `json:"pageSize"`             // 每页数量，最大30
	CollectionId string `json:"collectionId"`         // 集合的ID（必填）
	SearchText   string `json:"searchText,omitempty"` // 模糊搜索词
}

// DataListResponse 数据列表响应模型
//
// 用于表示数据列表的响应。
type DataListResponse struct {
	List  []DatasetData `json:"list"`  // 数据列表
	Total int           `json:"total"` // 总记录数
}

// DataDetailRequest 数据详情请求模型
//
// 用于请求获取单条数据详情。
type DataDetailRequest struct {
	Id string `json:"id"` // 数据的id
}

// DataUpdateRequest 数据更新请求模型
//
// 用于请求修改单条数据。
type DataUpdateRequest struct {
	DataId  string  `json:"dataId"`            // 数据的id
	Q       string  `json:"q,omitempty"`       // 主要数据
	A       string  `json:"a,omitempty"`       // 辅助数据
	Indexes []Index `json:"indexes,omitempty"` // 自定义索引
}

// DataDeleteRequest 数据删除请求模型
//
// 用于请求删除单条数据。
type DataDeleteRequest struct {
	Id string `json:"id"` // 数据的id
}

// 搜索测试相关模型

// DatasetSearchTestRequest 搜索测试请求模型
//
// 用于请求测试知识库搜索功能。
type DatasetSearchTestRequest struct {
	DatasetId                        string  `json:"datasetId"`                                  // 知识库ID
	Text                             string  `json:"text"`                                       // 需要测试的文本
	Limit                            int     `json:"limit"`                                      // 最大tokens数量
	Similarity                       float64 `json:"similarity,omitempty"`                       // 最低相关度（0~1，可选）
	SearchMode                       string  `json:"searchMode"`                                 // 搜索模式：embedding | fullTextRecall | mixedRecall
	UsingReRank                      bool    `json:"usingReRank"`                                // 使用重排
	DatasetSearchUsingExtensionQuery bool    `json:"datasetSearchUsingExtensionQuery,omitempty"` // 使用问题优化
	DatasetSearchExtensionModel      string  `json:"datasetSearchExtensionModel,omitempty"`      // 问题优化模型
	DatasetSearchExtensionBg         string  `json:"datasetSearchExtensionBg,omitempty"`         // 问题优化背景描述
}

// DatasetSearchTestResult 搜索测试结果模型
//
// 用于表示搜索测试的结果。
type DatasetSearchTestResult struct {
	ID           string  `json:"id"`           // 数据ID
	Q            string  `json:"q"`            // 主要数据
	A            string  `json:"a"`            // 辅助数据
	DatasetId    string  `json:"datasetId"`    // 知识库ID
	CollectionId string  `json:"collectionId"` // 集合ID
	SourceName   string  `json:"sourceName"`   // 来源名称
	SourceId     string  `json:"sourceId"`     // 来源ID
	Score        float64 `json:"score"`        // 相似度分数
}
