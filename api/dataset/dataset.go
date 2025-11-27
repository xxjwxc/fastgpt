// Package dataset 提供FastGPT知识库相关的API接口
//
// 该包封装了与知识库管理相关的所有API，包括：
// - 知识库的创建、查询、修改和删除
// - 集合的创建、查询、修改和删除
// - 数据的批量添加、查询、修改和删除
// - 支持多种类型的集合创建（外部文件、纯文本、链接、文件、API）
//
// 所有API均需要通过FastGPT客户端实例访问，使用前需先创建客户端。
package dataset

import (
	"encoding/json"

	"github.com/xxjwxc/fastgpt/client"
	"github.com/xxjwxc/fastgpt/model"
)

// DatasetAPI 知识库接口结构体，封装了所有知识库相关的API方法
//
// 该结构体通过组合HTTP客户端，提供了与FastGPT知识库管理相关的所有功能，
// 包括知识库管理、集合管理和数据管理。
type DatasetAPI struct {
	client *client.Client // HTTP客户端，用于发送API请求
}

// NewDatasetAPI 创建知识库接口实例
//
// 参数：
//
//	c: HTTP客户端实例，由外部传入
//
// 返回值：
//
//	*DatasetAPI: 知识库接口实例，用于访问知识库相关API
//
// 使用示例：
//
//	c := client.NewClient("https://cloud.fastgpt.cn", "sk-xxx")
//	datasetAPI := dataset.NewDatasetAPI(c)
func NewDatasetAPI(c *client.Client) *DatasetAPI {
	return &DatasetAPI{client: c}
}

// CreateDataset 创建知识库
//
// 该方法用于创建一个新的知识库，可设置知识库名称、描述、标签和类型等信息。
//
// 参数：
//
//	req: 知识库创建请求，包含知识库名称、描述、标签和类型
//
// 返回值：
//
//	string: 知识库ID
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E7%9F%A5%E8%AF%86%E5%BA%93
//
// 使用示例：
//
//	req := &model.DatasetCreateRequest{
//	    Name:         "我的知识库",
//	    Type:         "dataset",
//	    Intro:        "这是一个测试知识库",
//	}
//	datasetId, err := datasetAPI.CreateDataset(req)
func (api *DatasetAPI) CreateDataset(req *model.DatasetCreateRequest) (string, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/dataset/create", req)
	if err != nil {
		return "", err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return "", err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为字符串类型的知识库ID
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return "", err // 转换失败，返回错误
	}

	var datasetId string
	if err := json.Unmarshal(dataBytes, &datasetId); err != nil {
		return "", err // 解析失败，返回错误
	}

	return datasetId, nil // 返回知识库ID
}

// GetDatasetList 获取知识库列表
//
// 该方法用于获取知识库列表，支持分页查询。
//
// 参数：
//
//	req: 知识库列表请求，包含父级ID
//
// 返回值：
//
//	[]model.DatasetInfo: 知识库列表
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E8%8E%B7%E5%8F%96%E7%9F%A5%E8%AF%86%E5%BA%93%E5%88%97%E8%A1%A8
//
// 使用示例：
//
//	req := &model.DatasetListRequest{
//	    ParentId: "",
//	}
//	datasetList, err := datasetAPI.GetDatasetList(req)
func (api *DatasetAPI) GetDatasetList(req *model.DatasetListRequest) ([]model.DatasetInfo, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/dataset/list", req)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为[]model.DatasetInfo类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var datasetList []model.DatasetInfo
	if err := json.Unmarshal(dataBytes, &datasetList); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return datasetList, nil // 返回知识库列表
}

// GetDatasetDetail 获取知识库详情
//
// 该方法用于获取指定知识库的详细信息。
//
// 参数：
//
//	req: 知识库详情请求，包含知识库ID
//
// 返回值：
//
//	*model.DatasetInfo: 知识库详细信息
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E8%8E%B7%E5%8F%96%E7%9F%A5%E8%AF%86%E5%BA%93%E8%AF%A6%E6%83%85
//
// 使用示例：
//
//	req := &model.DatasetDetailRequest{
//	    Id: "your-dataset-id",
//	}
//	datasetDetail, err := datasetAPI.GetDatasetDetail(req)
func (api *DatasetAPI) GetDatasetDetail(req *model.DatasetDetailRequest) (*model.DatasetInfo, error) {
	resp, err := api.client.DoRequest("GET", "/api/core/dataset/detail?id="+req.Id, nil)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为DatasetInfo类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var datasetInfo model.DatasetInfo
	if err := json.Unmarshal(dataBytes, &datasetInfo); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return &datasetInfo, nil // 返回知识库详情
}

// DeleteDataset 删除知识库
//
// 该方法用于删除指定的知识库。
//
// 参数：
//
//	req: 知识库删除请求，包含知识库ID
//
// 返回值：
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%A0%E9%99%A4%E4%B8%80%E4%B8%AA%E7%9F%A5%E8%AF%86%E5%BA%93
//
// 使用示例：
//
//	req := &model.DatasetDeleteRequest{
//	    Id: "your-dataset-id",
//	}
//	err := datasetAPI.DeleteDataset(req)
func (api *DatasetAPI) DeleteDataset(req *model.DatasetDeleteRequest) error {
	resp, err := api.client.DoRequest("DELETE", "/api/core/dataset/delete?id="+req.Id, nil)
	if err != nil {
		return err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return err // 响应解析失败，返回错误
	}

	return nil // 删除成功
}

// CreateCollection 创建一个空的集合
//
// 该方法用于在指定知识库中创建一个空的集合。
//
// 参数：
//
//	req: 集合创建请求，包含知识库ID、父级ID、数据处理方式等
//
// 返回值：
//
//	string: 集合ID
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E7%A9%BA%E7%9A%84%E9%9B%86%E5%90%88
//
// 使用示例：
//
//	req := &model.CollectionCreateRequest{
//	    DatasetId:    "your-dataset-id",
//	    Name:         "测试集合",
//	    Type:         "virtual",
//	    TrainingType: "chunk",
//	}
//	collectionId, err := datasetAPI.CreateCollection(req)
func (api *DatasetAPI) CreateCollection(req *model.CollectionCreateRequest) (string, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/dataset/collection/create", req)
	if err != nil {
		return "", err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return "", err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为字符串类型的集合ID
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return "", err // 转换失败，返回错误
	}

	var collectionId string
	if err := json.Unmarshal(dataBytes, &collectionId); err != nil {
		return "", err // 解析失败，返回错误
	}

	return collectionId, nil // 返回集合ID
}

// CreateTextCollection 创建一个纯文本集合
//
// 该方法用于通过纯文本内容创建集合，系统会自动处理文本并生成索引。
//
// 参数：
//
//	req: 纯文本集合创建请求，包含文本内容、知识库ID等
//
// 返回值：
//
//	*model.CollectionCreateResponse: 集合创建响应，包含创建的集合ID和处理结果
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E7%BA%AF%E6%96%87%E6%9C%AC%E9%9B%86%E5%90%88
//
// 使用示例：
//
//	req := &model.CollectionCreateTextRequest{
//	    Text:         "这是一段测试文本",
//	    DatasetId:    "your-dataset-id",
//	    Name:         "测试文本集合",
//	    TrainingType: "chunk",
//	}
//	createResp, err := datasetAPI.CreateTextCollection(req)
func (api *DatasetAPI) CreateTextCollection(req *model.CollectionCreateTextRequest) (*model.CollectionCreateResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/dataset/collection/create/text", req)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为CollectionCreateResponse类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var createResp model.CollectionCreateResponse
	if err := json.Unmarshal(dataBytes, &createResp); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return &createResp, nil // 返回集合创建响应
}

// CreateLinkCollection 创建一个链接集合
//
// 该方法用于通过网络链接创建集合，系统会自动爬取链接内容并处理。
//
// 参数：
//
//	req: 链接集合创建请求，包含网络链接、知识库ID等
//
// 返回值：
//
//	*model.CollectionCreateResponse: 集合创建响应，包含创建的集合ID和处理结果
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E9%93%BE%E6%8E%A5%E9%9B%86%E5%90%88
//
// 使用示例：
//
//	req := &model.CollectionCreateLinkRequest{
//	    Link:         "https://example.com/article",
//	    DatasetId:    "your-dataset-id",
//	    TrainingType: "chunk",
//	}
//	createResp, err := datasetAPI.CreateLinkCollection(req)
func (api *DatasetAPI) CreateLinkCollection(req *model.CollectionCreateLinkRequest) (*model.CollectionCreateResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/dataset/collection/create/link", req)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为CollectionCreateResponse类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var createResp model.CollectionCreateResponse
	if err := json.Unmarshal(dataBytes, &createResp); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return &createResp, nil // 返回集合创建响应
}

// CreateAPICollection 创建一个API集合
//
// 该方法用于通过API文件创建集合，系统会自动处理API文件并生成索引。
//
// 参数：
//
//	req: API集合创建请求，包含文件ID、知识库ID等
//
// 返回值：
//
//	*model.CollectionCreateResponse: 集合创建响应，包含创建的集合ID和处理结果
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AAapi%E9%9B%86%E5%90%88
//
// 使用示例：
//
//	req := &model.CollectionCreateAPRequest{
//	    Name:         "API集合",
//	    ApiFileId:    "your-api-file-id",
//	    DatasetId:    "your-dataset-id",
//	    TrainingType: "chunk",
//	}
//	createResp, err := datasetAPI.CreateAPICollection(req)
func (api *DatasetAPI) CreateAPICollection(req *model.CollectionCreateAPRequest) (*model.CollectionCreateResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/dataset/collection/create/apiCollection", req)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为CollectionCreateResponse类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var createResp model.CollectionCreateResponse
	if err := json.Unmarshal(dataBytes, &createResp); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return &createResp, nil // 返回集合创建响应
}

// CreateExternalFileCollection 创建一个外部文件库集合（商业版）
//
// 该方法用于通过外部文件URL创建集合，系统会自动下载并处理外部文件。
//
// 参数：
//
//	req: 外部文件集合创建请求，包含外部文件URL、知识库ID等
//
// 返回值：
//
//	*model.CollectionCreateResponse: 集合创建响应，包含创建的集合ID和处理结果
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E5%A4%96%E9%83%A8%E6%96%87%E4%BB%B6%E5%BA%93%E9%9B%86%E5%90%88
//
// 使用示例：
//
//	req := &model.CollectionCreateExternalFileRequest{
//	    ExternalFileUrl: "https://example.com/file.pdf",
//	    ExternalFileId:  "123456",
//	    Filename:        "示例文件.pdf",
//	    DatasetId:       "your-dataset-id",
//	    TrainingType:    "chunk",
//	}
//	createResp, err := datasetAPI.CreateExternalFileCollection(req)
func (api *DatasetAPI) CreateExternalFileCollection(req *model.CollectionCreateExternalFileRequest) (*model.CollectionCreateResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/collection/create/externalFileUrl", req)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为CollectionCreateResponse类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var createResp model.CollectionCreateResponse
	if err := json.Unmarshal(dataBytes, &createResp); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return &createResp, nil // 返回集合创建响应
}

// GetCollectionList 获取集合列表
//
// 该方法用于获取指定知识库中的集合列表，支持分页查询。
//
// 参数：
//
//	req: 集合列表请求，包含知识库ID、父级ID、页码和每页大小
//
// 返回值：
//
//	*model.CollectionListResponse: 集合列表响应，包含集合列表和总数
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E8%8E%B7%E5%8F%96%E9%9B%86%E5%90%88%E5%88%97%E8%A1%A8
//
// 使用示例：
//
//	req := &model.CollectionListRequest{
//	    DatasetId:  "your-dataset-id",
//	    Offset:     0,
//	    PageSize:   10,
//	}
//	collectionList, err := datasetAPI.GetCollectionList(req)
func (api *DatasetAPI) GetCollectionList(req *model.CollectionListRequest) (*model.CollectionListResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/dataset/collection/listV2", req)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为CollectionListResponse类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var listResp model.CollectionListResponse
	if err := json.Unmarshal(dataBytes, &listResp); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return &listResp, nil // 返回集合列表
}

// GetCollectionDetail 获取集合详情
//
// 该方法用于获取指定集合的详细信息。
//
// 参数：
//
//	collectionId: 集合ID
//
// 返回值：
//
//	*model.CollectionInfo: 集合详细信息
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E8%8E%B7%E5%8F%96%E9%9B%86%E5%90%88%E8%AF%A6%E6%83%85
//
// 使用示例：
//
//	collectionInfo, err := datasetAPI.GetCollectionDetail("your-collection-id")
func (api *DatasetAPI) GetCollectionDetail(collectionId string) (*model.CollectionInfo, error) {
	resp, err := api.client.DoRequest("GET", "/api/core/dataset/collection/detail?id="+collectionId, nil)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为CollectionInfo类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var collectionInfo model.CollectionInfo
	if err := json.Unmarshal(dataBytes, &collectionInfo); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return &collectionInfo, nil // 返回集合详情
}

// UpdateCollection 修改集合信息
//
// 该方法用于修改指定集合的信息。
//
// 参数：
//
//	req: 集合更新请求，包含集合ID、知识库ID、外部文件ID等
//
// 返回值：
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E4%BF%AE%E6%94%B9%E9%9B%86%E5%90%88%E4%BF%A1%E6%81%AF
//
// 使用示例：
//
//	req := &model.CollectionUpdateRequest{
//	    Id:        "your-collection-id",
//	    Name:      "更新后的集合名称",
//	    Tags:      []string{"tag1", "tag2"},
//	}
//	err := datasetAPI.UpdateCollection(req)
func (api *DatasetAPI) UpdateCollection(req *model.CollectionUpdateRequest) error {
	resp, err := api.client.DoRequest("PUT", "/api/core/dataset/collection/update", req)
	if err != nil {
		return err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return err // 响应解析失败，返回错误
	}

	return nil // 更新成功
}

// DeleteCollection 删除一个集合
//
// 该方法用于删除指定知识库中的集合。
//
// 参数：
//
//	req: 集合删除请求，包含集合ID列表
//
// 返回值：
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%A0%E9%99%A4%E4%B8%80%E4%B8%AA%E9%9B%86%E5%90%88
//
// 使用示例：
//
//	req := &model.CollectionDeleteRequest{
//	    CollectionIds: []string{"your-collection-id"},
//	}
//	err := datasetAPI.DeleteCollection(req)
func (api *DatasetAPI) DeleteCollection(req *model.CollectionDeleteRequest) error {
	resp, err := api.client.DoRequest("POST", "/api/core/dataset/collection/delete", req)
	if err != nil {
		return err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return err // 响应解析失败，返回错误
	}

	return nil // 删除成功
}

// PushData 为集合批量添加数据
//
// 该方法用于为指定集合批量添加数据，每次最多支持200条。
//
// 参数：
//
//	req: 批量添加数据请求，包含集合ID、训练类型和数据列表
//
// 返回值：
//
//	*model.DataPushResponse: 批量添加数据响应，包含插入数量、超出token数量、重复数量和错误数量
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E4%B8%BA%E9%9B%86%E5%90%88%E6%89%B9%E9%87%8F%E6%B7%BB%E5%8A%A0%E6%B7%BB%E5%8A%A0%E6%95%B0%E6%8D%AE
//
// 使用示例：
//
//	req := &model.DataPushRequest{
//	    CollectionId: "your-collection-id",
//	    TrainingType: "chunk",
//	    Data: []model.DatasetData{
//	        {
//	            TeamId:       "your-team-id",
//	            TmbId:        "your-tmb-id",
//	            DatasetId:    "your-dataset-id",
//	            CollectionId: "your-collection-id",
//	            Q:            "问题1",
//	            A:            "答案1",
//	            UpdateTime:   "2024-01-01T00:00:00.000Z",
//	            Indexes: []model.Index{
//	                {
//	                    Text: "默认索引",
//	                },
//	            },
//	        },
//	    },
//	}
//	pushResp, err := datasetAPI.PushData(req)
func (api *DatasetAPI) PushData(req *model.DataPushRequest) (*model.DataPushResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/dataset/data/pushData", req)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为DataPushResponse类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var pushResp model.DataPushResponse
	if err := json.Unmarshal(dataBytes, &pushResp); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return &pushResp, nil // 返回批量添加数据响应
}

// GetDataList 获取集合的数据列表
//
// 该方法用于获取指定集合中的数据列表，支持分页查询。
//
// 参数：
//
//	req: 数据列表请求，包含集合ID、偏移量和每页大小
//
// 返回值：
//
//	*model.DataListResponse: 数据列表响应，包含数据列表和总数
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E8%8E%B7%E5%8F%96%E9%9B%86%E5%90%88%E7%9A%84%E6%95%B0%E6%8D%AE%E5%88%97%E8%A1%A8
//
// 使用示例：
//
//	req := &model.DataListRequest{
//	    CollectionId: "your-collection-id",
//	    Offset:      0,
//	    PageSize:    10,
//	}
//	dataList, err := datasetAPI.GetDataList(req)
func (api *DatasetAPI) GetDataList(req *model.DataListRequest) (*model.DataListResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/dataset/data/v2/list", req)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为DataListResponse类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var dataList model.DataListResponse
	if err := json.Unmarshal(dataBytes, &dataList); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return &dataList, nil // 返回数据列表
}

// GetDataDetail 获取单条数据详情
//
// 该方法用于获取指定集合中的单条数据详情。
//
// 参数：
//
//	req: 数据详情请求，包含数据ID
//
// 返回值：
//
//	*model.DatasetData: 数据详情
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E8%8E%B7%E5%8F%96%E5%8D%95%E6%9D%A1%E6%95%B0%E6%8D%AE%E8%AF%A6%E6%83%85
//
// 使用示例：
//
//	req := &model.DataDetailRequest{
//	    Id: "your-data-id",
//	}
//	dataDetail, err := datasetAPI.GetDataDetail(req)
func (api *DatasetAPI) GetDataDetail(req *model.DataDetailRequest) (*model.DatasetData, error) {
	resp, err := api.client.DoRequest("GET", "/api/core/dataset/data/detail?id="+req.Id, nil)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为DatasetData类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var dataDetail model.DatasetData
	if err := json.Unmarshal(dataBytes, &dataDetail); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return &dataDetail, nil // 返回数据详情
}

// UpdateData 修改单条数据
//
// 该方法用于修改指定集合中的单条数据。
//
// 参数：
//
//	req: 数据更新请求，包含数据ID、主要数据、辅助数据和索引
//
// 返回值：
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E4%BF%AE%E6%94%B9%E5%8D%95%E6%9D%A1%E6%95%B0%E6%8D%AE
//
// 使用示例：
//
//	req := &model.DataUpdateRequest{
//	    DataId: "your-data-id",
//	    Q:      "更新后的问题",
//	    A:      "更新后的答案",
//	}
//	err := datasetAPI.UpdateData(req)
func (api *DatasetAPI) UpdateData(req *model.DataUpdateRequest) error {
	resp, err := api.client.DoRequest("PUT", "/api/core/dataset/data/update", req)
	if err != nil {
		return err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return err // 响应解析失败，返回错误
	}

	return nil // 更新成功
}

// DeleteData 删除单条数据
//
// 该方法用于删除指定集合中的单条数据。
//
// 参数：
//
//	req: 数据删除请求，包含数据ID
//
// 返回值：
//
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%A0%E9%99%A4%E5%8D%95%E6%9D%A1%E6%95%B0%E6%8D%AE
//
// 使用示例：
//
//	req := &model.DataDeleteRequest{
//	    Id: "your-data-id",
//	}
//	err := datasetAPI.DeleteData(req)
func (api *DatasetAPI) DeleteData(req *model.DataDeleteRequest) error {
	resp, err := api.client.DoRequest("DELETE", "/api/core/dataset/data/delete?id="+req.Id, nil)
	if err != nil {
		return err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return err // 响应解析失败，返回错误
	}

	return nil // 删除成功
}

// SearchTest 搜索测试
//
// 该方法用于测试知识库搜索功能，返回相关度最高的结果。
//
// 参数：
//
//	req: 搜索测试请求，包含知识库ID、测试文本、搜索模式等
//
// 返回值：
//
//	[]model.DatasetSearchTestResult: 搜索测试结果列表
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E6%90%9C%E7%B4%A2%E6%B5%8B%E8%AF%95
//
// 使用示例：
//
//	req := &model.DatasetSearchTestRequest{
//	    DatasetId:    "your-dataset-id",
//	    Text:         "测试搜索文本",
//	    Limit:        5000,
//	    SearchMode:   "embedding",
//	    UsingReRank:  false,
//	}
//	searchResults, err := datasetAPI.SearchTest(req)
func (api *DatasetAPI) SearchTest(req *model.DatasetSearchTestRequest) ([]model.DatasetSearchTestResult, error) {
	resp, err := api.client.DoRequest("POST", "/api/core/dataset/searchTest", req)
	if err != nil {
		return nil, err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return nil, err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为[]model.DatasetSearchTestResult类型
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, err // 转换失败，返回错误
	}

	var searchResults []model.DatasetSearchTestResult
	if err := json.Unmarshal(dataBytes, &searchResults); err != nil {
		return nil, err // 解析失败，返回错误
	}

	return searchResults, nil // 返回搜索测试结果
}

// CreateTrainOrder 创建训练订单
//
// 该方法用于创建训练订单，用于记录训练使用情况。
//
// 参数：
//
//	req: 训练订单创建请求，包含知识库ID和可选的自定义订单名称
//
// 返回值：
//
//	string: 训练订单ID
//	error: 如果请求失败，返回错误信息
//
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E8%AE%AD%E7%BB%83%E8%AE%A2%E5%8D%95
//
// 使用示例：
//
//	req := &model.DatasetTrainOrderRequest{
//	    DatasetId: "your-dataset-id",
//	    Name:      "文档训练-fastgpt.docx", // 可选
//	}
//	trainOrderId, err := datasetAPI.CreateTrainOrder(req)
func (api *DatasetAPI) CreateTrainOrder(req *model.DatasetTrainOrderRequest) (string, error) {
	resp, err := api.client.DoRequest("POST", "/api/support/wallet/usage/createTrainingUsage", req)
	if err != nil {
		return "", err // 请求发送失败，返回错误
	}

	var baseResp model.BaseResponse
	if err := api.client.ParseResponse(resp, &baseResp); err != nil {
		return "", err // 响应解析失败，返回错误
	}

	// 将baseResp.Data转换为字符串类型的订单ID
	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return "", err // 转换失败，返回错误
	}

	var orderId string
	if err := json.Unmarshal(dataBytes, &orderId); err != nil {
		return "", err // 解析失败，返回错误
	}

	return orderId, nil // 返回训练订单ID
}
