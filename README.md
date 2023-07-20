| Clean Architecture 層 | 對應使用 Hexagonal Architecture 的 DDD 的層 |
| :-------------------- | :------------------------------------------ |
| Entity                | Domain                                      |
| Use Case              | Application Service                         |
| Adapter               | Infrastructure                              |
| Framework & Driver    | External Service (IO)                       |

# Golang DDD 的分層架構（六邊型架構）

├── cmd 存放 main.go 等
├── adapter
│   ├── grpc
│   └── http
│   └── facade  引用其他微服務
├── application 
│   ├── assembler   負責將內部領域模型轉化為可對外的DTO
│   └── cqe         Command、Query和Event --  輸入參數
│   └── dto         Application層的所有接口返回值為DTO -- 輸出參數
│   └── service     負責業務流程的編排，但本身不負責任何業務邏輯
├── domain
│   ├── aggregate   聚合
│   ├── entity      實體
│   ├── event       事件
│   │   ├── publish
│   │   └── subsctibe
│   ├── repo    接口
│   │   └── specification   统一封装查询
│   ├── service 领域服务
│   └── vo  值对象
└── infrastructure
│   ├── config  配置文件
│   ├── pkg     工具類封裝（DB,log,tool等）
│   └── repository
│   ├── converter domain 內部資料轉換 do {互轉}
│   └── do 資料庫轉換
└── types 封裝自定義的參數類型, 例如 phone 自校驗參數
