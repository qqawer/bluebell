package main

import (
	"WebApp/dao/mysql"
	"WebApp/dao/redis"
	"WebApp/global"
	"WebApp/logger"
	"WebApp/pkg/snowflake"
	"WebApp/router"
	"WebApp/settings"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	//1.加载配置
	if err:=settings.Init();err!=nil{
		fmt.Printf("init settings failed, err:%v\n",err)
		return
	}

	//2.初始化日志
	if err:=logger.Init(settings.AppConfig,settings.AppConfig.App.Mode);err!=nil{
		fmt.Printf("init logger failed, err:%v\n",err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//3.初始化Mysql连接
	if err:=mysql.Init(settings.AppConfig);err!=nil{
		fmt.Printf("Failed to initialize database, err:%v\n",err)
		return
	}
	
	//
	//4.初始化Redis连接
	if err:=redis.Init(settings.AppConfig);err!=nil{
		fmt.Printf("Failed to connect to Redis, err:%v",err)
		return
	}
	// defer global.RedisDB.Close()
	
	if err:=snowflake.Init(settings.AppConfig.App.StartTime,settings.AppConfig.App.MachineID);err!=nil{
		fmt.Printf("init snowflake failed, err:%v\n",err)
	}
	//5.注册路由
	r:=router.Setup(settings.AppConfig.App.Mode)

	 //6.启动服务(优雅关闭)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d",settings.AppConfig.App.Port),
		Handler: r,
	   }
	   
		go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	   }
	   }()
	   
	   quit := make(chan os.Signal, 1)
	   signal.Notify(quit, syscall.SIGINT,syscall.SIGTERM)
	   <-quit
	   zap.L().Info("Shutdown Server ...")
	   
	   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	   defer cancel()
	   if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown",zap.Error(err))
	   }
	   // 关闭数据库和 Redis 连接
	if global.Db != nil {
		sqlDB, err := global.Db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
	if global.RedisDB != nil {
		global.RedisDB.Close()
	}
	   zap.L().Info("Server exiting")


}
