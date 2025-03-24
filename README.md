这是一个用于同步NFT市场数据的Go语言服务。

## 已安装的包

### 1. 以太坊客户端库 (github.com/ethereum/go-ethereum)

用于连接以太坊区块链，监听事件，查询智能合约等。

示例用法:
```go
import (
    "github.com/ethereum/go-ethereum/ethclient"
)

func connectToEthereum() {
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_KEY")
    if err != nil {
        log.Fatal(err)
    }
    
    // 使用client与以太坊交互
}
```

### 2. Gin Web框架 (github.com/gin-gonic/gin)

用于构建RESTful API和Web服务。

示例用法:
```go
import (
    "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    
    r.GET("/nfts", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "NFT列表",
        })
    })
    
    return r
}

func main() {
    r := setupRouter()
    r.Run(":8080") // 监听并在0.0.0.0:8080上启动服务
}
```

### 3. GORM ORM框架 (gorm.io/gorm)

用于数据库操作，本项目配置了SQLite驱动。

示例用法:
```go
import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type NFT struct {
    ID          uint   `gorm:"primaryKey"`
    TokenID     string `gorm:"uniqueIndex"`
    Name        string
    Description string
    ImageURL    string
}

func setupDatabase() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("nft.db"), &gorm.Config{})
    if err != nil {
        panic("无法连接数据库")
    }
    
    // 自动迁移
    db.AutoMigrate(&NFT{})
    
    return db
}
```

## 项目结构

待开发...

## 如何运行

待开发...
