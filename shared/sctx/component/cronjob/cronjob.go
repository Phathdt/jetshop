package cronjob

import (
	"github.com/robfig/cron"
	"jetshop/shared/sctx"
)

type Cronjob interface {
	Enqueue(spec string, handler func())
}

type cronjob struct {
	id     string
	logger sctx.Logger
	cron   *cron.Cron
}

func NewCronjob(id string) *cronjob {
	return &cronjob{id: id, cron: cron.New()}
}

func (c *cronjob) ID() string {
	return c.id
}

func (c *cronjob) InitFlags() {}

func (c *cronjob) Activate(sc sctx.ServiceContext) error {
	c.logger = sc.Logger(c.id)

	go c.cron.Run()

	return nil
}

func (c *cronjob) Stop() error {
	c.cron.Stop()

	return nil
}

func (c *cronjob) Enqueue(spec string, handler func()) {
	_ = c.cron.AddFunc(spec, handler)
}
