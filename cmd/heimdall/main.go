package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/alexflint/go-arg"
	"github.com/cryring/Heimdall/internal/config"
	"github.com/cryring/Heimdall/internal/log"
	"github.com/cryring/Heimdall/internal/server"
	"go.uber.org/zap"
)

func main() {
	// init cmdline args
	args := Config{}
	if err := arg.Parse(&args); err != nil {
		log.Errorf("parse args failed: %v", err)
		return
	}

	// init log
	if args.LogPath == "" {
		args.LogPath = "./log/bridge.log"
	}
	log.InitLogger("./log/bridge.log")

	// init config
	cfg := config.Config{
		TcpAddress:       args.TcpAddress,
		TlsAddress:       args.TlsAddress,
		WebsocketAddress: args.WebsocketAddress,
		ClusterAddress:   args.ClusterAddress,
	}
	err := config.Init(cfg)
	if err != nil {
		log.Errorf("parse config file failed: %v", err)
		return
	}

	// start server
	srv := server.New()
	srv.Run()

	s := waitForSignal()
	log.Info("signal received, broker closed.", zap.Any("signal", s))
}

func waitForSignal() os.Signal {
	ch := make(chan os.Signal, 1)
	defer close(ch)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	signal.Stop(ch)
	return s
}
