package common

import (
	"context"
	"go-test-basic/model"
	"log"
	"path/filepath"
	"testing"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/testcontainers/testcontainers-go"
	cmysql "github.com/testcontainers/testcontainers-go/modules/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gloablDB *gorm.DB

func GetDB() *gorm.DB {
	return gloablDB
}

func NewDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/hello?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.Item{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}

func InitDB() {
	gloablDB = NewDB()
}

func InitTestDB(t *testing.T) {
	db := memory.NewDatabase("hello")
	pro := memory.NewDBProvider(db)
	engine := sqle.NewDefault(pro)

	config := server.Config{
		Protocol: "tcp",
		Address:  "localhost:3306",
	}
	s, err := server.NewServer(config, engine, memory.NewSessionBuilder(pro), nil)
	if err != nil {
		panic(err)
	}
	go func() {
		if err = s.Start(); err != nil {
			panic(err)
		}
	}()

	gloablDB = NewDB()

	t.Cleanup(func() {
		gloablDB = nil
		if err = s.Close(); err != nil {
			panic(err)
		}
	})
}

func InitTestDBWithContainer(t *testing.T) {
	ctx := context.Background()

	mysqlContainer, err := cmysql.Run(ctx,
		"mysql:8.0.36",
		cmysql.WithConfigFile(filepath.Join("testdata", "my_8.cnf")),
		cmysql.WithDatabase("foo"),
		cmysql.WithUsername("root"),
		cmysql.WithPassword("password"),
		cmysql.WithScripts(filepath.Join("testdata", "schema.sql")),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(mysqlContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

}
