package model

type Subscriber interface {
	Update()
}

type Publisher interface {
	Attach(Subscriber)
	Detach(Subscriber)
	NotifySubscribers()
}