package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pxj/courseSystem/config"
	"syscall"
	"time"
)

func InitResource() {

}

func RunServer() {
	router := gin.Default()
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
