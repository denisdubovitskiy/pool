package pool

import "time"

const (
	defaultMaxActiveConn = 10
	defaultMaxIdleConn   = 10
	defaultMaxConnAge    = 0
	defaultIdleTimeout   = 0
	defaultDialTimeout   = 0
	defaultDialRetries   = 0
)

func newPoolOptions(opts ...Option) *poolOptions {
	opt := &poolOptions{
		maxActiveConn: defaultMaxIdleConn,
		maxIdleConn:   defaultMaxIdleConn,
		maxConnAge:    defaultMaxConnAge,
		idleTimeout:   defaultIdleTimeout,
		dialTimeout:   defaultDialTimeout,
		dialRetries:   defaultDialRetries,
		dialer:        newDefaultDialer(),
		hooks:         newDefaultHooks(),
	}

	for _, o := range opts {
		o(opt)
	}

	if opt.maxActiveConn < opt.maxIdleConn {
		opt.maxIdleConn = opt.maxActiveConn
	}

	return opt
}

type poolOptions struct {
	maxActiveConn int
	maxIdleConn   int
	dialRetries   int
	maxConnAge    time.Duration
	idleTimeout   time.Duration
	dialTimeout   time.Duration
	dialer        Dialer
	hooks         *hooks
}

type Option func(options *poolOptions)

func WithMaxActiveConn(maxConn int) Option {
	return func(p *poolOptions) {
		p.maxActiveConn = maxConn
	}
}

func WithMaxIdleConn(maxIdleConn int) Option {
	return func(p *poolOptions) {
		p.maxIdleConn = maxIdleConn
	}
}

func WithMaxConnAge(maxConnAge time.Duration) Option {
	return func(p *poolOptions) {
		p.maxConnAge = maxConnAge
	}
}

func WithIdleTimeout(timeout time.Duration) Option {
	return func(p *poolOptions) {
		if timeout > 0 {
			p.idleTimeout = timeout
		}
	}
}

func WithDialTimeout(timeout time.Duration) Option {
	return func(p *poolOptions) {
		if timeout > 0 {
			p.dialTimeout = timeout
		}
	}
}

func WithDialer(d Dialer) Option {
	return func(p *poolOptions) {
		if d != nil {
			p.dialer = d
		}
	}
}

func WithDialRetries(retries int) Option {
	return func(p *poolOptions) {
		p.dialRetries = retries
	}
}

func WithOnCloseHook(hook OnCloseHook) Option {
	return func(p *poolOptions) {
		p.hooks.addOnCloseHook(hook)
	}
}

func WithOnDialHook(hook OnDialHook) Option {
	return func(p *poolOptions) {
		p.hooks.addOnDialHook(hook)
	}
}

func WithOnGetTimeoutHook(hook OnGetTimeoutHook) Option {
	return func(p *poolOptions) {
		p.hooks.addOnGetTimeoutHook(hook)
	}
}
