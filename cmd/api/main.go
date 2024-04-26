// Created by Petr Lozhkin

// api micro service

package main

import (
	"context"
	"github.com/golang/glog"
	"github.com/im7mortal/airports/pkg/airports/server"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

func main() {

	// Main context
	world, globalCancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		glog.Infof("Got signal %s; Shutdown all operations", sig.String())
		globalCancel()
	}()

	gracefulShutdownD := time.Second

	var exist bool
	var err error
	gracefulShutdown, exist := os.LookupEnv("GRACEFUL_SHUTDOWN")
	if exist {
		gracefulShutdownD, err = time.ParseDuration(gracefulShutdown)
		if err != nil {
			glog.Exitf("gracefulShutdownD: %s is not valid: %s", gracefulShutdown, err)
		}
	}

	port := ":8080"

	sdkEngine := server.New()

	router := sdkEngine.GetMainEngine()

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}
	go func() {
		if r := recover(); r != nil {
			glog.Errorf("panic %s \nstacktrace from panic: \n%s", r, string(debug.Stack()))
		}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Error(err)
			globalCancel()
		}
	}()
	select {
	case <-world.Done():

	}

	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownD)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		glog.Error("Server forced to shutdown:", err)
	}
	glog.Error("correct exit")
}
