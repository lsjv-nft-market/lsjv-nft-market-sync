package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NFT 模型定义
type NFT struct {
	ID          uint   `gorm:"primaryKey"`
	TokenID     string `gorm:"uniqueIndex"`
	Name        string
	Description string
	ImageURL    string
}

// 初始化数据库
func setupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("nft.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移数据表
	err = db.AutoMigrate(&NFT{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// 连接以太坊
func connectToEthereum() (*ethclient.Client, error) {
	// 请替换为您的以太坊节点URL
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_KEY")
	if err != nil {
		return nil, err
	}

	return client, nil
}

// 设置API路由
func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 获取NFT列表
	r.GET("/nfts", func(c *gin.Context) {
		var nfts []NFT
		result := db.Find(&nfts)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, nfts)
	})

	// 获取单个NFT
	r.GET("/nfts/:id", func(c *gin.Context) {
		id := c.Param("id")
		var nft NFT
		result := db.First(&nft, "token_id = ?", id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "NFT not found"})
			return
		}

		c.JSON(http.StatusOK, nft)
	})

	// 添加NFT
	r.POST("/nfts", func(c *gin.Context) {
		var nft NFT
		if err := c.ShouldBindJSON(&nft); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := db.Create(&nft)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, nft)
	})

	return r
}

func main() {
	// 初始化数据库
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 连接以太坊
	ethClient, err := connectToEthereum()
	if err != nil {
		log.Fatalf("连接以太坊失败: %v", err)
	}

	// 获取以太坊区块号
	blockNumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Printf("获取区块号失败: %v", err)
	} else {
		log.Printf("当前以太坊区块号: %d", blockNumber)
	}

	// 设置路由
	r := setupRouter(db)

	// 启动服务
	log.Println("服务启动在 http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
