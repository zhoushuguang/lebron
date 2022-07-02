package batcher

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

var ErrFull = errors.New("channel is full")

type Option interface {
	apply(*options)
}

type options struct {
	size     int
	buffer   int
	worker   int
	interval time.Duration
}

func (o options) check() {
	if o.size <= 0 {
		o.size = 100
	}
	if o.buffer <= 0 {
		o.buffer = 100
	}
	if o.worker <= 0 {
		o.worker = 5
	}
	if o.interval <= 0 {
		o.interval = time.Second
	}
}

type funcOption struct {
	f func(*options)
}

func (fo *funcOption) apply(o *options) {
	fo.f(o)
}

func newOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithSize(s int) Option {
	return newOption(func(o *options) {
		o.size = s
	})
}

func WithBuffer(b int) Option {
	return newOption(func(o *options) {
		o.buffer = b
	})
}

func WithWorker(w int) Option {
	return newOption(func(o *options) {
		o.worker = w
	})
}

func WithInterval(i time.Duration) Option {
	return newOption(func(o *options) {
		o.interval = i
	})
}

type msg struct {
	key string
	val interface{}
}

type Batcher struct {
	opts options

	Do       func(ctx context.Context, val map[string][]interface{})
	Sharding func(key string) int
	chans    []chan *msg
	wait     sync.WaitGroup
}

func New(opts ...Option) *Batcher {
	b := &Batcher{}
	for _, opt := range opts {
		opt.apply(&b.opts)
	}
	b.opts.check()

	b.chans = make([]chan *msg, b.opts.worker)
	for i := 0; i < b.opts.worker; i++ {
		b.chans[i] = make(chan *msg, b.opts.buffer)
	}
	return b
}

func (b *Batcher) Start() {
	if b.Do == nil {
		log.Fatal("Batcher: Do func is nil")
	}
	if b.Sharding == nil {
		log.Fatal("Batcher: Sharding func is nil")
	}
	b.wait.Add(len(b.chans))
	for i, ch := range b.chans {
		go b.merge(i, ch)
	}
}

func (b *Batcher) Add(key string, val interface{}) error {
	ch, msg := b.add(key, val)
	select {
	case ch <- msg:
	default:
		return ErrFull
	}
	return nil
}

func (b *Batcher) add(key string, val interface{}) (chan *msg, *msg) {
	sharding := b.Sharding(key) % b.opts.worker
	ch := b.chans[sharding]
	msg := &msg{key: key, val: val}
	return ch, msg
}

func (b *Batcher) merge(idx int, ch <-chan *msg) {
	defer b.wait.Done()

	var (
		msg        *msg
		count      int
		closed     bool
		lastTicker = true
		interval   = b.opts.interval
		vals       = make(map[string][]interface{}, b.opts.size)
	)
	if idx > 0 {
		interval = time.Duration(int64(idx) * (int64(b.opts.interval) / int64(b.opts.worker)))
	}
	ticker := time.NewTicker(interval)
	for {
		select {
		case msg = <-ch:
			if msg == nil {
				closed = true
				break
			}
			count++
			vals[msg.key] = append(vals[msg.key], msg.val)
			if count >= b.opts.size {
				break
			}
			continue
		case <-ticker.C:
			if lastTicker {
				ticker.Stop()
				ticker = time.NewTicker(b.opts.interval)
				lastTicker = false
			}
		}
		if len(vals) > 0 {
			ctx := context.Background()
			b.Do(ctx, vals)
			vals = make(map[string][]interface{}, b.opts.size)
			count = 0
		}
		if closed {
			ticker.Stop()
			return
		}
	}
}

func (b *Batcher) Close() {
	for _, ch := range b.chans {
		ch <- nil
	}
	b.wait.Wait()
}
