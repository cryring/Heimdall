package router

import (
	"github.com/eclipse/paho.golang/packets"
)

type Router struct {
}

func New() *Router {
	return &Router{}
}

func (r *Router) Publish(srcClientID string, pkg *packets.Publish) error {
	topic := pkg.Topic
	r.parseTopic(topic)
	return nil
}

func (r *Router) PublishSystem(topic, payload string) error {
	return nil
}

func (r *Router) parseTopic(topic string) {
	// fields := strings.Split(topic, "/")
}
