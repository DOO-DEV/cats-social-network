package event

import "meower/schema"

type Store interface {
	Close()
	PublishMeowCreated(meow schema.Meow) error
	SubscribeMeowCreated() (<-chan MeowCreatedMessage, error)
	OnMeowCreated(f func(MeowCreatedMessage)) error
}

var implement Store

func SetEventStore(es Store) {
	implement = es
}

func Close() {
	implement.Close()
}

func PublishMeowCreated(meow schema.Meow) error {
	return implement.PublishMeowCreated(meow)
}

func SubscribeMeowCreated() (<-chan MeowCreatedMessage, error) {
	return implement.SubscribeMeowCreated()
}

func OnMeowCreated(f func(MeowCreatedMessage)) error {
	return implement.OnMeowCreated(f)
}
