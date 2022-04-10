package event

import (
	"context"
	"github.com/go-redis/redis"
)

const EventBusBufferCapacity = 512

type Event struct {
	Channel string
	Payload string
}

type PubSub interface {
	Publish(channel string, payload interface{}) error
	Subscribe(ctx context.Context, pattern string) (<-chan *Event, error)
	OnEvent(ctx context.Context, pattern string, onEvent func(payload string)) error
	Close()
}

type bus struct {
	client *redis.Client
}

func NewEventBus(address, password string) *bus {
	b := new(bus)
	r := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	b.client = r
	return b
}

func (b *bus) Publish(channel string, payload interface{}) error {
	b.client.Publish(channel, payload)
	return nil
}

func (b *bus) Subscribe(ctx context.Context, pattern string) (<-chan *Event, error) {
	ch := make(chan *Event, EventBusBufferCapacity)

	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			for message := range b.client.PSubscribe(pattern).Channel() {
				ch <- &Event{
					Channel: message.Channel,
					Payload: message.Payload,
				}
			}
		}
	}()

	return ch, nil
}

func (b *bus) OnEvent(ctx context.Context, pattern string, onEvent func(payload string)) error {
	eventCh, err := b.Subscribe(ctx, pattern)
	if err != nil {
		return err
	}

	go func() {
		for event := range eventCh {
			onEvent(event.Payload)
		}
	}()

	return nil
}

func (b *bus) Close() {
	b.client.Close()
}
