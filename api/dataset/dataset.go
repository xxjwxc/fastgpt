package dataset

import (
	"github.com/xxjwxc/fastgpt/client"
	"github.com/xxjwxc/fastgpt/model"
)

// DatasetAPI 知识库接口
type DatasetAPI struct {
	client *client.Client
}

// NewDatasetAPI 创建知识库接口实例
func NewDatasetAPI(c *client.Client) *DatasetAPI {
	return &DatasetAPI{client: c}
}

// CreateDataset 创建知识库
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E7%9F%A5%E8%AF%86%E5%BA%93
func (api *DatasetAPI) CreateDataset(req *model.DatasetCreateRequest) (*model.DatasetInfo, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/create", req)
	if err != nil {
		return nil, err
	}

	var datasetInfo model.DatasetInfo
	if err := api.client.ParseResponse(resp, &datasetInfo); err != nil {
		return nil, err
	}

	return &datasetInfo, nil
}

// GetDatasetList 获取知识库列表
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E8%8E%B7%E5%8F%96%E7%9F%A5%E8%AF%86%E5%BA%93%E5%88%97%E8%A1%A8
func (api *DatasetAPI) GetDatasetList(page, pageSize int) (*model.DatasetListResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/list", map[string]interface{}{
		"page":     page,
		"pageSize": pageSize,
	})
	if err != nil {
		return nil, err
	}

	var listResp model.DatasetListResponse
	if err := api.client.ParseResponse(resp, &listResp); err != nil {
		return nil, err
	}

	return &listResp, nil
}

// GetDatasetDetail 获取知识库详情
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E8%8E%B7%E5%8F%96%E7%9F%A5%E8%AF%86%E5%BA%93%E8%AF%A6%E6%83%85
func (api *DatasetAPI) GetDatasetDetail(datasetId string) (*model.DatasetInfo, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/detail", map[string]string{
		"datasetId": datasetId,
	})
	if err != nil {
		return nil, err
	}

	var datasetInfo model.DatasetInfo
	if err := api.client.ParseResponse(resp, &datasetInfo); err != nil {
		return nil, err
	}

	return &datasetInfo, nil
}

// DeleteDataset 删除知识库
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%A0%E9%99%A4%E4%B8%80%E4%B8%AA%E7%9F%A5%E8%AF%86%E5%BA%93
func (api *DatasetAPI) DeleteDataset(datasetId string) error {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/delete", map[string]string{
		"datasetId": datasetId,
	})
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}

// CreateCollection 创建一个空的集合
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset
func (api *DatasetAPI) CreateCollection(req *model.CollectionCreateRequest) (*model.CollectionCreateResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/collection/create", req)
	if err != nil {
		return nil, err
	}

	var createResp model.CollectionCreateResponse
	if err := api.client.ParseResponse(resp, &createResp); err != nil {
		return nil, err
	}

	return &createResp, nil
}

// CreateExternalFileCollection 创建外部文件集合
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset
func (api *DatasetAPI) CreateExternalFileCollection(req *model.ExternalFileCollectionCreateRequest) (*model.CollectionCreateResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/collection/create/externalFileUrl", req)
	if err != nil {
		return nil, err
	}

	var createResp model.CollectionCreateResponse
	if err := api.client.ParseResponse(resp, &createResp); err != nil {
		return nil, err
	}

	return &createResp, nil
}

// CreateTextCollection 创建纯文本集合
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E7%BA%AF%E6%96%87%E6%9C%AC%E9%9B%86%E5%90%88
func (api *DatasetAPI) CreateTextCollection(req *model.TextCollectionCreateRequest) (*model.CollectionCreateResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/collection/create/text", req)
	if err != nil {
		return nil, err
	}

	var createResp model.CollectionCreateResponse
	if err := api.client.ParseResponse(resp, &createResp); err != nil {
		return nil, err
	}

	return &createResp, nil
}

// CreateLinkCollection 创建链接集合
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E9%93%BE%E6%8E%A5%E9%9B%86%E5%90%88
func (api *DatasetAPI) CreateLinkCollection(req *model.LinkCollectionCreateRequest) (*model.CollectionCreateResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/collection/create/link", req)
	if err != nil {
		return nil, err
	}

	var createResp model.CollectionCreateResponse
	if err := api.client.ParseResponse(resp, &createResp); err != nil {
		return nil, err
	}

	return &createResp, nil
}

// CreateFileCollection 创建文件集合
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA%E6%96%87%E4%BB%B6%E9%9B%86%E5%90%88
func (api *DatasetAPI) CreateFileCollection(req *model.FileCollectionCreateRequest) (*model.CollectionCreateResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/collection/create/file", req)
	if err != nil {
		return nil, err
	}

	var createResp model.CollectionCreateResponse
	if err := api.client.ParseResponse(resp, &createResp); err != nil {
		return nil, err
	}

	return &createResp, nil
}

// CreateAPICollection 创建API集合
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AAapi%E9%9B%86%E5%90%88
func (api *DatasetAPI) CreateAPICollection(req *model.APICollectionCreateRequest) (*model.CollectionCreateResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/collection/create/api", req)
	if err != nil {
		return nil, err
	}

	var createResp model.CollectionCreateResponse
	if err := api.client.ParseResponse(resp, &createResp); err != nil {
		return nil, err
	}

	return &createResp, nil
}

// GetCollectionDetail 获取集合详情
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E8%8E%B7%E5%8F%96%E9%9B%86%E5%90%88%E8%AF%A6%E6%83%85
func (api *DatasetAPI) GetCollectionDetail(req *model.CollectionDetailRequest) (*model.CollectionInfo, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/collection/detail", req)
	if err != nil {
		return nil, err
	}

	var collectionInfo model.CollectionInfo
	if err := api.client.ParseResponse(resp, &collectionInfo); err != nil {
		return nil, err
	}

	return &collectionInfo, nil
}

// UpdateCollection 修改集合信息
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E4%BF%AE%E6%94%B9%E9%9B%86%E5%90%88%E4%BF%A1%E6%81%AF
func (api *DatasetAPI) UpdateCollection(req *model.CollectionUpdateRequest) error {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/collection/update", req)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}

// BatchAddData 为集合批量添加数据
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset
// 注意：每次最多推送200组数据
func (api *DatasetAPI) BatchAddData(req *model.DatasetDataBatchRequest) error {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/data/batch", req)
	if err != nil {
		return err
	}

	// 只需要检查状态码，不需要解析响应体
	resp.Body.Close()
	return nil
}

// GetCollectionList 获取集合列表
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset
func (api *DatasetAPI) GetCollectionList(req *model.CollectionListRequest) (*model.CollectionListResponse, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/collection/list", req)
	if err != nil {
		return nil, err
	}

	var listResp model.CollectionListResponse
	if err := api.client.ParseResponse(resp, &listResp); err != nil {
		return nil, err
	}

	return &listResp, nil
}

// DeleteCollection 删除一个集合
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset
func (api *DatasetAPI) DeleteCollection(datasetId, collectionId string) error {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/collection/delete", map[string]string{
		"datasetId":    datasetId,
		"collectionId": collectionId,
	})
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}

// GetCollectionDataList 获取集合的数据列表
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E8%8E%B7%E5%8F%96%E9%9B%86%E5%90%88%E7%9A%84%E6%95%B0%E6%8D%AE%E5%88%97%E8%A1%A8
func (api *DatasetAPI) GetCollectionDataList(datasetId, collectionId string, page, pageSize int) ([]model.DatasetData, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/data/list", map[string]interface{}{
		"datasetId":    datasetId,
		"collectionId": collectionId,
		"page":         page,
		"pageSize":     pageSize,
	})
	if err != nil {
		return nil, err
	}

	var dataList []model.DatasetData
	if err := api.client.ParseResponse(resp, &dataList); err != nil {
		return nil, err
	}

	return dataList, nil
}

// GetDataDetail 获取单条数据详情
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E8%8E%B7%E5%8F%96%E5%8D%95%E6%9D%A1%E6%95%B0%E6%8D%AE%E8%AF%A6%E6%83%85
func (api *DatasetAPI) GetDataDetail(req *model.DataDetailRequest) (*model.DatasetData, error) {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/data/detail", req)
	if err != nil {
		return nil, err
	}

	var dataDetail model.DatasetData
	if err := api.client.ParseResponse(resp, &dataDetail); err != nil {
		return nil, err
	}

	return &dataDetail, nil
}

// UpdateData 修改单条数据
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E4%BF%AE%E6%94%B9%E5%8D%95%E6%9D%A1%E6%95%B0%E6%8D%AE
func (api *DatasetAPI) UpdateData(req *model.DataUpdateRequest) error {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/data/update", req)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}

// DeleteData 删除单条数据
// 接口文档：https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset#%E5%88%A0%E9%99%A4%E5%8D%95%E6%9D%A1%E6%95%B0%E6%8D%AE
func (api *DatasetAPI) DeleteData(req *model.DataDeleteRequest) error {
	resp, err := api.client.DoRequest("POST", "/api/proApi/core/dataset/data/delete", req)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}
