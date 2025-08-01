package log

import (
	"sync"
	"sync/atomic"
)

type Complex struct {
	transports []EntryTransporter
	locker     sync.RWMutex
	maxLevel   Level
}

func NewComplex(transports ...EntryTransporter) *Complex {
	c := &Complex{
		transports: nil,
		locker:     sync.RWMutex{},
		//transportMap: make(map[string]EntryTransporter),
		//buildLocker:  sync.Mutex{},
		maxLevel: PanicLevel,
	}
	_ = c.Reset(transports...)
	return c
}

func (c *Complex) Reset(transporters ...EntryTransporter) error {

	ts := make([]EntryTransporter, 0, len(transporters))
	maxLevel := PanicLevel
	for _, transporter := range transporters {

		if transporter.Level() > maxLevel {
			maxLevel = transporter.Level()
		}
		ts = append(ts, transporter)
	}

	c.setLevel(maxLevel)

	c.locker.Lock()
	c.transports = ts
	c.locker.Unlock()
	return nil
}

func (c *Complex) Transport(entry *Entry) error {

	c.locker.RLock()
	ts := c.transports
	c.locker.RUnlock()

	for _, t := range ts {
		if t.Level() >= entry.Level {
			_ = t.Transport(entry)
		}
	}
	return nil
}

// SetLevel sets the logger level.
func (c *Complex) setLevel(level Level) {
	atomic.StoreUint32((*uint32)(&c.maxLevel), uint32(level))
}
func (c *Complex) Level() Level {
	return Level(atomic.LoadUint32((*uint32)(&c.maxLevel)))
}

func (c *Complex) Close() error {

	c.locker.Lock()
	ts := c.transports
	c.transports = nil
	c.locker.Unlock()

	for _, t := range ts {
		_ = t.Close()
	}
	return nil
}
