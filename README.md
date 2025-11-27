# go-fastgpt

FastGPT API的Go客户端库，支持应用接口、对话接口和知识库接口。

## 安装

```bash
go get github.com/xxjwxc/fastgpt
```

## 快速开始

### 初始化客户端

```go
import "github.com/xxjwxc/fastgpt"

// 初始化FastGPT客户端
fgpt := fastgpt.NewFastGPT("https://api.fastgpt.cn", "your-api-key")
```

## 应用接口

### 获取应用统计数据

```go
import "github.com/xxjwxc/fastgpt/model"

// 构建请求
statsReq := &model.AppStatsRequest{
    StartTime: 1758585600000, // 开始时间戳
    EndTime:   1758672000000, // 结束时间戳
}

// 发送请求
statsResp, err := fgpt.App.GetStats(statsReq)
if err != nil {
    log.Printf("获取应用统计数据失败: %v\n", err)
} else {
    fmt.Printf("应用统计数据获取成功，用户数据数量: %d\n", len(statsResp.Data.UserData))
}
```

## 对话接口

### 发送对话请求

```go
// 构建对话请求
chatReq := &model.ChatRequest{
    AppId: "your-app-id",
    Messages: []model.Message{
        {
            Role:    "user",
            Content: "你好",
        },
    },
    Stream: true, // 开启流式响应
}

// 发送对话请求并处理SSE事件
err = fgpt.Chat.Chat(chatReq, func(eventType string, data interface{}) error {
    switch eventType {
    case "flowNodeStatus":
        statusEvent := data.(model.FlowNodeStatusEvent)
        fmt.Printf("流程节点状态: %s - %s\n", statusEvent.Name, statusEvent.Status)
    case "answer":
        if data == "[DONE]" {
            fmt.Println("对话结束")
            return nil
        }
        answerEvent := data.(model.AnswerEvent)
        for _, choice := range answerEvent.Choices {
            fmt.Print(choice.Delta.Content) // 打印流式响应内容
        }
    case "flowResponses":
        flowEvent := data.(model.FlowResponsesEvent)
        fmt.Printf("\n流程响应数量: %d\n", len(flowEvent.Responses))
    }
    return nil
})
if err != nil {
    log.Printf("对话请求失败: %v\n", err)
}
```

## 知识库接口

### 创建外部文件集合

```go
// 构建创建外部文件集合请求
externalReq := &model.ExternalFileCollectionCreateRequest{
    ExternalFileUrl: "https://example.com/file.pdf", // 外部文件URL
    ExternalFileId:  "123456", // 外部文件ID
    Filename:        "示例文件.pdf", // 自定义文件名
    CreateTime:      "2025-01-01T00:00:00.000Z", // 创建时间
    DatasetId:       "your-dataset-id", // 知识库ID
    TrainingType:    "chunk", // 数据处理方式：chunk-按文本长度分割; qa-问答对提取
    ChunkSize:       1500, // 分块大小
    Tags:            []string{"tag1", "tag2"}, // 集合标签
}

// 发送请求
externalResp, err := fgpt.Dataset.CreateExternalFileCollection(externalReq)
if err != nil {
    log.Printf("创建外部文件集合失败: %v\n", err)
} else {
    fmt.Printf("外部文件集合创建成功，集合ID: %s\n", externalResp.CollectionId)
}
```

### 为集合批量添加数据

```go
// 构建批量添加数据请求
batchReq := &model.DatasetDataBatchRequest{
    DatasetId:    "your-dataset-id",
    CollectionId: "your-collection-id",
    Data: []model.DatasetData{
        {
            TeamId:       "your-team-id",
            TmbId:        "your-tmb-id",
            Q:            "主要数据",
            A:            "辅助数据",
            Indexes: []model.Index{
                {
                    Type: "custom",
                    Text: "自定义索引文本",
                },
            },
        },
        // 可以添加更多数据，每次最多200条
    },
}

// 发送请求
err := fgpt.Dataset.BatchAddData(batchReq)
if err != nil {
    log.Printf("批量添加数据失败: %v\n", err)
} else {
    fmt.Println("批量添加数据成功")
}
```

### 获取集合列表

```go
// 构建请求
listReq := &model.CollectionListRequest{
    DatasetId: "your-dataset-id",
    Page:      1,
    PageSize:  10,
}

// 发送请求
listResp, err := fgpt.Dataset.GetCollectionList(listReq)
if err != nil {
    log.Printf("获取集合列表失败: %v\n", err)
} else {
    fmt.Printf("获取集合列表成功，总数: %d\n", listResp.Total)
}
```

## API文档

- [应用接口](https://doc.fastgpt.cn/docs/introduction/development/openapi/app)
- [对话接口](https://doc.fastgpt.cn/docs/introduction/development/openapi/chat)
- [知识库接口](https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset)

## 许可证

MIT
