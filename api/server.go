package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pxj/courseSystem/api/cache"
	"pxj/courseSystem/api/databases"
	"pxj/courseSystem/api/routes"
	"pxj/courseSystem/config"
	"syscall"
	"time"
)

func InitResource() (err error) {
	// init mysql
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConf.MysqlConf.User,
		config.AppConf.MysqlConf.Password,
		config.AppConf.MysqlConf.Host,
		config.AppConf.MysqlConf.Port,
		config.AppConf.MysqlConf.DB)
	err = databases.InitMysql(dsn)
	if err != nil {
		return err
	}
	// init redis
	err = cache.InitRedis(config.AppConf.RedisConf.Host, config.AppConf.RedisConf.Db, config.AppConf.RedisConf.Password)
	if err != nil {
		return err
	}
	return
}

func RunServer() {
	routes.InitRoutes()
	router := routes.Routes
	host := config.AppConf.ServerConf.Host
	port := config.AppConf.ServerConf.Port
	logrus.Infof("server run at: host = %v, port = %v", host, port)

	address := fmt.Sprintf("%s:%d", host, port)
	server := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    61 * time.Second,
		WriteTimeout:   61 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logrus.Infof("Runing server at", address)
	go func() {
		err := server.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("Server Listen Error: %s\n ", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
