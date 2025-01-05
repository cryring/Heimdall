package router

type Topic struct {
}

func NewTopic() *Topic {
	return &Topic{}
}

func (t *Topic) Match(topic string) bool {
	return false
}
