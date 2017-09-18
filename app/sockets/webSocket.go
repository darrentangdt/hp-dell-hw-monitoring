package sockets

import (
	"container/list"
	"time"
)

type Event struct {
	Host      string
	Component string
	Timestamp int
}

type Subscription struct {
	New <-chan Event
}

func (s Subscription) Cancel() {
	unsubscribe <- s.New
	drain(s.New)
}

func newEvent(host, components string) Event {
	return Event{host, components, int(time.Now().Unix())}
}

func Subscribe() Subscription {
	resp := make(chan Subscription)
	subscribe <- resp
	return <-resp
}

func Alert(host, components string) {
	publish <- newEvent(host, components)
}

var (
	subscribe   = make(chan (chan<- Subscription), 10)
	unsubscribe = make(chan (<-chan Event), 10)
	publish     = make(chan Event, 10)
)

func webSocket() {
	subscribers := list.New()

	for {
		select {
		case ch := <-subscribe:
			subscriber := make(chan Event, 10)
			subscribers.PushBack(subscriber)
			ch <- Subscription{subscriber}

		case event := <-publish:
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				ch.Value.(chan Event) <- event
			}

		case unsub := <-unsubscribe:
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				if ch.Value.(chan Event) == unsub {
					subscribers.Remove(ch)
					break
				}
			}
		}
	}
}

func InitWebSocket() {
	go webSocket()
}

func drain(ch <-chan Event) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			return
		}
	}
}
