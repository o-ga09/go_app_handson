package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/taiti09/go_app_handson/config"
)


func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v",err)
		os.Exit(1)
	}
} 

func helloHandle(w http.ResponseWriter,r *http.Request) {
	//こまんドララインから実験するため
	time.Sleep(5*time.Second)
	fmt.Fprintf(w,"Hello %s!",r.URL.Path[1:])
}

//テスト用コード
func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx,os.Interrupt,syscall.SIGTERM)
	defer stop()
	cfg, err := config.New()
	if err != nil {
		return err
	}
	listener, err := net.Listen("tcp",fmt.Sprintf(":%d",cfg.Port))
	if err != nil {
		log.Fatalf("failed to terminate server: %v",err)
	}
	url := fmt.Sprintf("http://%s",listener.Addr().String())
	log.Printf("start with: %v",url)

	mux, cleanup, err := NewMux(ctx,cfg)
	if err != nil {
		return err
	}
	defer cleanup()
	s := NewServer(listener,mux)
	return s.Run(ctx)
}