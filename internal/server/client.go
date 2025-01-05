package server

import (
	"net"

	"github.com/cryring/Heimdall/internal/log"
	"github.com/cryring/Heimdall/internal/router"
	"github.com/eclipse/paho.golang/packets"
	"go.uber.org/zap"
)

type client struct {
	conn     net.Conn
	clientID string
	dying    chan struct{}

	router *router.Router
}

func newClient(router *router.Router) *client {
	return &client{
		dying:  make(chan struct{}),
		router: router,
	}
}

func (c *client) readLoop() {
	for {
		select {
		case <-c.dying:
			return
		default:
		}

		pkg, err := packets.ReadPacket(c.conn)
		if err != nil {
			log.Error("read packet error: ", zap.Error(err), zap.String("ClientID", c.clientID))
			// TODO:
			return
		}
		c.process(pkg)
	}
}

func (c *client) process(pkg *packets.ControlPacket) {
	switch pkg.Content.(type) {
	case *packets.Connect:
	case *packets.Connack:
	case *packets.Publish:
		pkt := pkg.Content.(*packets.Publish)
		c.processPublish(pkt)
	case *packets.Puback:
	case *packets.Pubrec:
	case *packets.Pubrel:
	case *packets.Pubcomp:
	case *packets.Subscribe:
		pkt := pkg.Content.(*packets.Subscribe)
		c.processSubscribe(pkt)
	case *packets.Suback:
	case *packets.Unsubscribe:
	case *packets.Unsuback:
	case *packets.Pingreq:
	case *packets.Pingresp:
	case *packets.Disconnect:
	}
}

func (c *client) processPublish(pkg *packets.Publish) {
	c.router.Publish(c.clientID, pkg)
}

func (c *client) processSubscribe(pkg *packets.Subscribe) {
	// TODO: check subscribe info
}
