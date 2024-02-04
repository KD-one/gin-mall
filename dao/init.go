package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var DB *gorm.DB

// Database 初始化数据库
func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connWrite, // 相当于定义主数据库
		DefaultStringSize:         256,       // 字符串默认长度
		DisableDatetimePrecision:  true,      // 禁用时间戳精度
		DontSupportRenameIndex:    true,      // 禁用重命名索引
		DontSupportRenameColumn:   true,      // 禁用重命名列
		SkipInitializeWithVersion: true,      // 跳过使用版本初始化
	}), &gorm.Config{
		Logger: ormLogger, // 设置日志级别
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名命名规则不要复数（不加s）
		},
	})

	if err != nil {
		panic("failed to connect database")
	}

	// 获取数据库连接
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}

	// 设置数据库连接池
	sqlDB.SetMaxIdleConns(10)
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(20)
	// 设置数据库连接最大存活时间
	sqlDB.SetConnMaxLifetime(60 * time.Second)

	DB = db

	// 主从配置
	_ = DB.Use(dbresolver.Register(dbresolver.Config{
		// 主库（写）
		Sources: []gorm.Dialector{mysql.Open(connWrite)},
		// 从库（读）
		Replicas: []gorm.Dialector{mysql.Open(connRead)},
		// 策略
		Policy: dbresolver.RandomPolicy{},
	}))
	// 数据库迁移
	migration()
}
