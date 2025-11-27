package model

// DatasetBaseRequest 知识库基础请求
type DatasetBaseRequest struct {
	DatasetId string `json:"datasetId"`
	TeamId    string `json:"teamId,omitempty"`
}

// DatasetCreateRequest 创建知识库请求
type DatasetCreateRequest struct {
	Name        string   `json:"name"`                  // 知识库名称
	Description string   `json:"description,omitempty"` // 知识库描述
	Tags        []string `json:"tags,omitempty"`        // 知识库标签
	Type        string   `json:"type,omitempty"`        // 知识库类型，默认public
}

// DatasetInfo 知识库信息
type DatasetInfo struct {
	Id          string   `json:"id"`                    // 知识库ID
	Name        string   `json:"name"`                  // 知识库名称
	Description string   `json:"description,omitempty"` // 知识库描述
	Tags        []string `json:"tags,omitempty"`        // 知识库标签
	Type        string   `json:"type"`                  // 知识库类型
	CreateTime  string   `json:"createTime"`            // 创建时间
	UpdateTime  string   `json:"updateTime"`            // 更新时间
}

// DatasetListResponse 知识库列表响应
type DatasetListResponse struct {
	List     []DatasetInfo `json:"list"`               // 知识库列表
	Total    int           `json:"total"`              // 总数
	Page     int           `json:"page,omitempty"`     // 页码
	PageSize int           `json:"pageSize,omitempty"` // 每页大小
}

// CollectionCreateRequest 创建集合请求
type CollectionCreateRequest struct {
	DatasetId        string   `json:"datasetId"`                  // 知识库ID
	ParentId         string   `json:"parentId,omitempty"`         // 父级ID，不填则默认为根目录
	TrainingType     string   `json:"trainingType"`               // 数据处理方式：chunk-按文本长度分割; qa-问答对提取
	IndexPrefixTitle bool     `json:"indexPrefixTitle,omitempty"` // 是否自动生成标题索引
	CustomPdfParse   bool     `json:"customPdfParse,omitempty"`   // 是否开启PDF增强解析
	AutoIndexes      bool     `json:"autoIndexes,omitempty"`      // 是否自动生成索引(仅商业版支持)
	ImageIndex       bool     `json:"imageIndex,omitempty"`       // 是否自动生成图片索引(仅商业版支持)
	ChunkSettingMode string   `json:"chunkSettingMode,omitempty"` // 分块参数模式：auto-系统默认; custom-手动指定
	ChunkSplitMode   string   `json:"chunkSplitMode,omitempty"`   // 分块拆分模式：size-按长度; char-按字符
	ChunkSize        int      `json:"chunkSize,omitempty"`        // 分块大小，默认1500
	IndexSize        int      `json:"indexSize,omitempty"`        // 索引大小，默认512
	ChunkSplitter    string   `json:"chunkSplitter,omitempty"`    // 自定义最高优先分割符号
	QaPrompt         string   `json:"qaPrompt,omitempty"`         // qa拆分提示词
	Tags             []string `json:"tags,omitempty"`             // 集合标签
	CreateTime       string   `json:"createTime,omitempty"`       // 文件创建时间
}

// ExternalFileCollectionCreateRequest 创建外部文件集合请求
type ExternalFileCollectionCreateRequest struct {
	ExternalFileUrl string   `json:"externalFileUrl"`         // 文件访问链接
	ExternalFileId  string   `json:"externalFileId"`          // 外部文件ID
	Filename        string   `json:"filename,omitempty"`      // 自定义文件名，需要带后缀
	CreateTime      string   `json:"createTime,omitempty"`    // 文件创建时间
	DatasetId       string   `json:"datasetId"`               // 知识库ID
	ParentId        string   `json:"parentId,omitempty"`      // 父级ID
	Tags            []string `json:"tags,omitempty"`          // 集合标签
	TrainingType    string   `json:"trainingType"`            // 数据处理方式
	ChunkSize       int      `json:"chunkSize,omitempty"`     // 分块大小
	ChunkSplitter   string   `json:"chunkSplitter,omitempty"` // 自定义分割符号
	QaPrompt        string   `json:"qaPrompt,omitempty"`      // qa拆分提示词
}

// TextCollectionCreateRequest 创建纯文本集合请求
type TextCollectionCreateRequest struct {
	DatasetId        string   `json:"datasetId"`                  // 知识库ID
	ParentId         string   `json:"parentId,omitempty"`         // 父级ID，不填则默认为根目录
	Text             string   `json:"text"`                       // 纯文本内容
	TrainingType     string   `json:"trainingType"`               // 数据处理方式：chunk-按文本长度分割; qa-问答对提取
	IndexPrefixTitle bool     `json:"indexPrefixTitle,omitempty"` // 是否自动生成标题索引
	ChunkSettingMode string   `json:"chunkSettingMode,omitempty"` // 分块参数模式：auto-系统默认; custom-手动指定
	ChunkSplitMode   string   `json:"chunkSplitMode,omitempty"`   // 分块拆分模式：size-按长度; char-按字符
	ChunkSize        int      `json:"chunkSize,omitempty"`        // 分块大小，默认1500
	IndexSize        int      `json:"indexSize,omitempty"`        // 索引大小，默认512
	ChunkSplitter    string   `json:"chunkSplitter,omitempty"`    // 自定义最高优先分割符号
	QaPrompt         string   `json:"qaPrompt,omitempty"`         // qa拆分提示词
	Tags             []string `json:"tags,omitempty"`             // 集合标签
	CreateTime       string   `json:"createTime,omitempty"`       // 文件创建时间
}

// LinkCollectionCreateRequest 创建链接集合请求
type LinkCollectionCreateRequest struct {
	DatasetId        string   `json:"datasetId"`                  // 知识库ID
	ParentId         string   `json:"parentId,omitempty"`         // 父级ID，不填则默认为根目录
	Link             string   `json:"link"`                       // 网络链接
	TrainingType     string   `json:"trainingType"`               // 数据处理方式：chunk-按文本长度分割; qa-问答对提取
	IndexPrefixTitle bool     `json:"indexPrefixTitle,omitempty"` // 是否自动生成标题索引
	ChunkSettingMode string   `json:"chunkSettingMode,omitempty"` // 分块参数模式：auto-系统默认; custom-手动指定
	ChunkSplitMode   string   `json:"chunkSplitMode,omitempty"`   // 分块拆分模式：size-按长度; char-按字符
	ChunkSize        int      `json:"chunkSize,omitempty"`        // 分块大小，默认1500
	IndexSize        int      `json:"indexSize,omitempty"`        // 索引大小，默认512
	ChunkSplitter    string   `json:"chunkSplitter,omitempty"`    // 自定义最高优先分割符号
	QaPrompt         string   `json:"qaPrompt,omitempty"`         // qa拆分提示词
	Tags             []string `json:"tags,omitempty"`             // 集合标签
	CreateTime       string   `json:"createTime,omitempty"`       // 文件创建时间
}

// FileCollectionCreateRequest 创建文件集合请求
type FileCollectionCreateRequest struct {
	DatasetId        string   `json:"datasetId"`                  // 知识库ID
	ParentId         string   `json:"parentId,omitempty"`         // 父级ID，不填则默认为根目录
	FileId           string   `json:"fileId"`                     // 文件ID
	TrainingType     string   `json:"trainingType"`               // 数据处理方式：chunk-按文本长度分割; qa-问答对提取
	IndexPrefixTitle bool     `json:"indexPrefixTitle,omitempty"` // 是否自动生成标题索引
	CustomPdfParse   bool     `json:"customPdfParse,omitempty"`   // 是否开启PDF增强解析
	ChunkSettingMode string   `json:"chunkSettingMode,omitempty"` // 分块参数模式：auto-系统默认; custom-手动指定
	ChunkSplitMode   string   `json:"chunkSplitMode,omitempty"`   // 分块拆分模式：size-按长度; char-按字符
	ChunkSize        int      `json:"chunkSize,omitempty"`        // 分块大小，默认1500
	IndexSize        int      `json:"indexSize,omitempty"`        // 索引大小，默认512
	ChunkSplitter    string   `json:"chunkSplitter,omitempty"`    // 自定义最高优先分割符号
	QaPrompt         string   `json:"qaPrompt,omitempty"`         // qa拆分提示词
	Tags             []string `json:"tags,omitempty"`             // 集合标签
	CreateTime       string   `json:"createTime,omitempty"`       // 文件创建时间
}

// APICollectionCreateRequest 创建API集合请求
type APICollectionCreateRequest struct {
	DatasetId        string   `json:"datasetId"`                  // 知识库ID
	ParentId         string   `json:"parentId,omitempty"`         // 父级ID，不填则默认为根目录
	FileId           string   `json:"fileId"`                     // 文件ID
	TrainingType     string   `json:"trainingType"`               // 数据处理方式：chunk-按文本长度分割; qa-问答对提取
	IndexPrefixTitle bool     `json:"indexPrefixTitle,omitempty"` // 是否自动生成标题索引
	CustomPdfParse   bool     `json:"customPdfParse,omitempty"`   // 是否开启PDF增强解析
	ChunkSettingMode string   `json:"chunkSettingMode,omitempty"` // 分块参数模式：auto-系统默认; custom-手动指定
	ChunkSplitMode   string   `json:"chunkSplitMode,omitempty"`   // 分块拆分模式：size-按长度; char-按字符
	ChunkSize        int      `json:"chunkSize,omitempty"`        // 分块大小，默认1500
	IndexSize        int      `json:"indexSize,omitempty"`        // 索引大小，默认512
	ChunkSplitter    string   `json:"chunkSplitter,omitempty"`    // 自定义最高优先分割符号
	QaPrompt         string   `json:"qaPrompt,omitempty"`         // qa拆分提示词
	Tags             []string `json:"tags,omitempty"`             // 集合标签
	CreateTime       string   `json:"createTime,omitempty"`       // 文件创建时间
}

// CollectionDetailRequest 获取集合详情请求
type CollectionDetailRequest struct {
	DatasetId    string `json:"datasetId"`    // 知识库ID
	CollectionId string `json:"collectionId"` // 集合ID
}

// CollectionUpdateRequest 修改集合信息请求
type CollectionUpdateRequest struct {
	DatasetId    string   `json:"datasetId"`             // 知识库ID
	CollectionId string   `json:"collectionId"`          // 集合ID
	Name         string   `json:"name,omitempty"`        // 集合名称
	Description  string   `json:"description,omitempty"` // 集合描述
	Tags         []string `json:"tags,omitempty"`        // 集合标签
}

// CollectionCreateResponse 创建集合响应
type CollectionCreateResponse struct {
	CollectionId string `json:"collectionId"` // 新建的集合ID
	Results      struct {
		InsertLen int      `json:"insertLen"` // 插入的块数量
		OverToken []string `json:"overToken"` // 超出token限制的内容
		Repeat    []string `json:"repeat"`    // 重复的内容
		Error     []string `json:"error"`     // 错误的内容
	} `json:"results"`
}

// Index 向量索引
type Index struct {
	Type   string `json:"type,omitempty"`   // 索引类型：default-默认; custom-自定义; summary-总结; question-问题; image-图片
	DataId string `json:"dataId,omitempty"` // 关联的向量ID
	Text   string `json:"text"`             // 文本内容
}

// DatasetData 知识库数据
type DatasetData struct {
	TeamId        string  `json:"teamId"`                  // 团队ID
	TmbId         string  `json:"tmbId"`                   // 成员ID
	DatasetId     string  `json:"datasetId"`               // 知识库ID
	CollectionId  string  `json:"collectionId"`            // 集合ID
	Q             string  `json:"q"`                       // 主要数据
	A             string  `json:"a,omitempty"`             // 辅助数据
	FullTextToken string  `json:"fullTextToken,omitempty"` // 分词
	Indexes       []Index `json:"indexes"`                 // 向量索引
	UpdateTime    string  `json:"updateTime,omitempty"`    // 更新时间
	ChunkIndex    int     `json:"chunkIndex,omitempty"`    // 分块下标
}

// DatasetDataBatchRequest 批量添加数据请求
type DatasetDataBatchRequest struct {
	DatasetId    string        `json:"datasetId"`    // 知识库ID
	CollectionId string        `json:"collectionId"` // 集合ID
	Data         []DatasetData `json:"data"`         // 数据列表，每次最多200条
}

// DataDetailRequest 获取单条数据详情请求
type DataDetailRequest struct {
	DatasetId    string `json:"datasetId"`    // 知识库ID
	CollectionId string `json:"collectionId"` // 集合ID
	DataId       string `json:"dataId"`       // 数据ID
}

// DataUpdateRequest 修改单条数据请求
type DataUpdateRequest struct {
	DatasetId    string  `json:"datasetId"`         // 知识库ID
	CollectionId string  `json:"collectionId"`      // 集合ID
	DataId       string  `json:"dataId"`            // 数据ID
	Q            string  `json:"q,omitempty"`       // 主要数据
	A            string  `json:"a,omitempty"`       // 辅助数据
	Indexes      []Index `json:"indexes,omitempty"` // 向量索引
}

// DataDeleteRequest 删除单条数据请求
type DataDeleteRequest struct {
	DatasetId    string `json:"datasetId"`    // 知识库ID
	CollectionId string `json:"collectionId"` // 集合ID
	DataId       string `json:"dataId"`       // 数据ID
}

// CollectionListRequest 获取集合列表请求
type CollectionListRequest struct {
	DatasetId string `json:"datasetId"`          // 知识库ID
	ParentId  string `json:"parentId,omitempty"` // 父级ID
	Page      int    `json:"page,omitempty"`     // 页码
	PageSize  int    `json:"pageSize,omitempty"` // 每页大小
}

// CollectionInfo 集合信息
type CollectionInfo struct {
	Id           string   `json:"id"`           // 集合ID
	Name         string   `json:"name"`         // 集合名称
	DatasetId    string   `json:"datasetId"`    // 知识库ID
	ParentId     string   `json:"parentId"`     // 父级ID
	TrainingType string   `json:"trainingType"` // 数据处理方式
	Tags         []string `json:"tags"`         // 集合标签
	CreateTime   string   `json:"createTime"`   // 创建时间
	UpdateTime   string   `json:"updateTime"`   // 更新时间
}

// CollectionListResponse 获取集合列表响应
type CollectionListResponse struct {
	List     []CollectionInfo `json:"list"`     // 集合列表
	Total    int              `json:"total"`    // 总数
	Page     int              `json:"page"`     // 页码
	PageSize int              `json:"pageSize"` // 每页大小
}
