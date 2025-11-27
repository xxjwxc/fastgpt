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
fgpt := fastgpt.NewFastGPT("https://cloud.fastgpt.cn", "your-api-key")

// 开启debug模式（可选）
fgpt.SetDebug(true)
```

## 应用接口

### 获取累积运行结果

```go
import "github.com/xxjwxc/fastgpt/model"

// 构建请求
totalDataReq := &model.AppTotalDataRequest{
    AppId: "your-app-id", // 应用ID
}

// 发送请求
totalDataResp, err := fgpt.App.GetTotalData(totalDataReq)
if err != nil {
    log.Printf("获取累积运行结果失败: %v\n", err)
} else {
    fmt.Printf("累积运行结果获取成功，总对话数: %d\n", totalDataResp.TotalChat)
}
```

### 获取应用日志看板

```go
// 构建请求
chartDataReq := &model.AppChartDataRequest{
    AppId:       "your-app-id", // 应用ID
    DateStart:   "2025-09-19T16:00:00.000Z", // 开始时间
    DateEnd:     "2025-09-27T15:59:59.999Z", // 结束时间
    UserTimespan: "day", // 用户统计时间粒度
    ChatTimespan: "day", // 对话统计时间粒度
    AppTimespan:  "day", // 应用统计时间粒度
}

// 发送请求
chartDataResp, err := fgpt.App.GetChartData(chartDataReq)
if err != nil {
    log.Printf("获取应用日志看板失败: %v\n", err)
} else {
    fmt.Printf("应用日志看板获取成功，用户数据数量: %d\n", len(chartDataResp.UserData))
}
```

## 对话接口

### 发送对话请求

```go
import "github.com/xxjwxc/fastgpt/model"

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

### 获取应用历史记录

```go
// 构建请求
historiesReq := &model.GetHistoriesRequest{
    AppId:    "your-app-id",
    Offset:   0,
    PageSize: 10,
    Source:   "api",
}

// 发送请求
historiesResp, err := fgpt.Chat.GetHistories(historiesReq)
if err != nil {
    log.Printf("获取应用历史记录失败: %v\n", err)
} else {
    fmt.Printf("获取应用历史记录成功，总数: %d\n", historiesResp.Total)
}
```

## 知识库接口

### 创建知识库

```go
// 构建创建知识库请求
datasetReq := &model.DatasetCreateRequest{
    Name:  "我的知识库",
    Type:  "dataset",
    Intro: "这是一个测试知识库",
}

// 发送请求
datasetId, err := fgpt.Dataset.CreateDataset(datasetReq)
if err != nil {
    log.Printf("创建知识库失败: %v\n", err)
} else {
    fmt.Printf("知识库创建成功，ID: %s\n", datasetId)
}
```

### 创建外部文件集合

```go
// 构建创建外部文件集合请求
externalReq := &model.CollectionCreateExternalFileRequest{
    ExternalFileUrl: "https://example.com/file.pdf", // 外部文件URL
    ExternalFileId:  "123456", // 外部文件ID
    Filename:        "示例文件.pdf", // 自定义文件名
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
pushReq := &model.DataPushRequest{
    CollectionId: "your-collection-id",
    TrainingType: "chunk",
    Data: []model.DatasetData{
        {
            Q: "主要数据",
            A: "辅助数据",
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
pushResp, err := fgpt.Dataset.PushData(pushReq)
if err != nil {
    log.Printf("批量添加数据失败: %v\n", err)
} else {
    fmt.Printf("批量添加数据成功，插入数量: %d\n", pushResp.InsertLen)
}
```

### 获取集合列表

```go
// 构建请求
listReq := &model.CollectionListRequest{
    DatasetId:  "your-dataset-id",
    Offset:     0,
    PageSize:   10,
}

// 发送请求
listResp, err := fgpt.Dataset.GetCollectionList(listReq)
if err != nil {
    log.Printf("获取集合列表失败: %v\n", err)
} else {
    fmt.Printf("获取集合列表成功，总数: %d\n", listResp.Total)
}
```

### 创建训练订单

```go
// 构建创建训练订单请求
trainReq := &model.DatasetTrainOrderRequest{
    DatasetId: "your-dataset-id",
    Name:      "文档训练-fastgpt.docx", // 可选，自定义订单名称
}

// 发送请求
trainOrderId, err := fgpt.Dataset.CreateTrainOrder(trainReq)
if err != nil {
    log.Printf("创建训练订单失败: %v\n", err)
} else {
    fmt.Printf("训练订单创建成功，ID: %s\n", trainOrderId)
}
```

## API文档

- [应用接口](https://doc.fastgpt.cn/docs/introduction/development/openapi/app)
- [对话接口](https://doc.fastgpt.cn/docs/introduction/development/openapi/chat)
- [知识库接口](https://doc.fastgpt.cn/docs/introduction/development/openapi/dataset)

## 许可证

MIT
