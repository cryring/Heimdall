package server

import (
	"bytes"
	"fmt"
	"net"

	"github.com/cryring/Heimdall/internal/config"
	"github.com/cryring/Heimdall/internal/log"
	"github.com/cryring/Heimdall/internal/router"

	"github.com/eclipse/paho.golang/packets"
	"go.uber.org/zap"
)

type Server struct {
	router *router.Router
}

func New() *Server {
	return &Server{
		router: router.New(),
	}
}

func (srv *Server) Run() error {
	if config.GetConfig().TcpAddress != "" {
		if err := srv.ListenTCP(config.GetConfig().TcpAddress); err != nil {
			return err
		}
	}
	if config.GetConfig().TlsAddress != "" {
		if err := srv.ListenTLS(config.GetConfig().TlsAddress); err != nil {
			return err
		}
	}
	if config.GetConfig().WebsocketAddress != "" {
		if err := srv.ListenWebsocket(config.GetConfig().WebsocketAddress); err != nil {
			return err
		}
	}
	return nil
}

func (srv *Server) ListenTCP(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	go srv.ServeTCP(l)

	log.Infof("tcp address: %v is running ...", address)
	return nil
}

func (srv *Server) ServeTCP(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok {
				log.Errorf("Client Accept Net Error(%v)", zap.Error(ne))
			} else {
				log.Errorf("Accept error", zap.Error(err))
			}
			continue
		}

		go func() {
			if err := srv.handleTcpConn(conn); err != nil {
				log.Infof("conn %v closed with error: %v", conn, err)
				conn.Close()
			}
		}()
	}
}

func (srv *Server) ListenTLS(address string) error {
	// TODO:
	return nil
}

func (srv *Server) ServeTLS(l net.Listener) {
	// TODO:
}

func (srv *Server) ListenWebsocket(address string) error {
	// TODO:
	return nil
}

func (srv *Server) ServeWebsocket(l net.Listener) {
	// TODO:
}

func (srv *Server) handleTcpConn(conn net.Conn) error {
	pkg, err := packets.ReadPacket(conn)
	if err != nil {
		return fmt.Errorf("read connect packet error: %v", err)
	}
	msg, ok := pkg.Content.(*packets.Connect)
	if !ok {
		return fmt.Errorf("received msg that was not Connect")
	}
	ack := packets.NewControlPacket(packets.CONNACK)
	content := ack.Content.(*packets.Connack)
	// TODO:
	// content.SessionPresent = msg.CleanStart
	// content.Properties.MaximumQOS = &mqtt.QoS0 // use config

	if err := srv.checkAuth(msg.ClientID, msg.Username, msg.Password); err != nil {
		content.ReasonCode = packets.ConnackNotAuthorized
		if _, err := content.WriteTo(conn); err != nil {
			return fmt.Errorf("send connack error: %v, clientID: %v, conn: %v", err, msg.ClientID, conn)
		}
		return fmt.Errorf("connect packet CheckConnectAuth failed with connack.ReasonCode: %v", content.ReasonCode)
	}

	if _, err := content.WriteTo(conn); err != nil {
		return fmt.Errorf("send connack error:%v,clientID:%v,conn:%v", err, msg.ClientID, conn)
	}

	c := newClient(srv.router)
	c.readLoop()
	return nil
}

func (srv *Server) checkAuth(clientID, username string, password []byte) error {
	// TODO: for test
	if clientID == "123" && username == "foo" && bytes.Equal(password, []byte("abc")) {
		return nil
	}
	return fmt.Errorf("client username password not match")
}
