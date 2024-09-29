package db

import (
	"context"
	"fmt"
	"os"
	beconfig "sec-kill/common/be_config"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var dbpool *pgxpool.Pool

func Init() error {
	logger := zap.L()
	if logger == nil {
		fmt.Println("logger is nil in db")
		os.Exit(-1)
	}
	// 获取配置项
	host, err := beconfig.GetConfig("dbms.postgres.host")
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("获取的数据库地址：" + host)

	port, err := beconfig.GetConfig("dbms.postgres.port")
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("获取的数据库端口：" + port)

	db, err := beconfig.GetConfig("dbms.postgres.db")
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("获取的数据库名称：" + db)

	user, err := beconfig.GetConfig("dbms.postgres.user")
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("获取的数据库用户：" + user)

	pwd, err := beconfig.GetConfig("dbms.postgres.password")
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("获取的数据库密码：" + pwd)

	// 连接数据库
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, pwd, host, port, db)
	// fmt.Println(connString)
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		logger.Error("数据库连接配置解析错误：" + err.Error())
		return err
	}
	// 最大连接数
	config.MaxConns = 16
	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		logger.Error("连接数据池失败:" + err.Error())
		return err
	}
	// 测试数据库连接池是否有效
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		logger.Error("数据库连接池无效：" + err.Error())
		return err
	}
	conn.Release()

	dbpool = pool
	logger.Info("数据库连接池初始化成功！")

	return nil
}

func GetDBPool() *pgxpool.Pool {
	logger := zap.L()
	if logger == nil {
		fmt.Println("logger is nil in db")
		os.Exit(-1)
	}
	if dbpool == nil {
		logger.Error("failed,数据库连接池还未初始化")
		return nil
	}
	// 测试数据库连接池是否有效
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		logger.Error("数据库连接池无效：" + err.Error())
		return nil
	}
	conn.Release()

	return dbpool
}
