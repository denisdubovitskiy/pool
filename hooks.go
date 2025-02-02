package pool

import (
	"time"
)

type OnCloseHook interface {
	OnClose(age time.Duration, reason CloseReason, err error)
}

type OnDialHook interface {
	OnDial(duration time.Duration, err error)
}

type OnGetTimeoutHook interface {
	OnGetTimeout(timeout time.Duration)
}

func newDefaultHooks() *hooks {
	return &hooks{
		onClose:      []func(age time.Duration, reason CloseReason, err error){},
		onDial:       []func(duration time.Duration, err error){},
		onGetTimeout: []func(duration time.Duration){},
	}
}

type hooks struct {
	onClose      []func(age time.Duration, reason CloseReason, err error)
	onDial       []func(duration time.Duration, err error)
	onGetTimeout []func(duration time.Duration)
}

func (c *hooks) addOnGetTimeoutHook(hook OnGetTimeoutHook) {
	c.onGetTimeout = append(c.onGetTimeout, hook.OnGetTimeout)
}

func (c *hooks) addOnDialHook(hook OnDialHook) {
	c.onDial = append(c.onDial, hook.OnDial)
}

func (c *hooks) addOnCloseHook(hook OnCloseHook) {
	c.onClose = append(c.onClose, hook.OnClose)
}

func (c *hooks) OnClose(age time.Duration, reason CloseReason, err error) {
	for _, h := range c.onClose {
		h(age, reason, err)
	}
}

func (c *hooks) OnDial(duration time.Duration, err error) {
	for _, h := range c.onDial {
		h(duration, err)
	}
}

func (c *hooks) OnGetTimeout(duration time.Duration) {
	for _, h := range c.onGetTimeout {
		h(duration)
	}
}
