package database

import "github.com/bela333/Vigne/errors"

type Welcomer struct {
	d *Database
}

func (d *Database) Welcomer() (*Welcomer, error) {
	exists := d.Redis.Exists(d.Decorate("welcomer:main")).Val()
	if exists == 0 {
		return nil, errors.NoWelcomer
	}
	exists = d.Redis.Exists(d.Decorate("welcomer:text:after")).Val()
	if exists == 0 {
		return nil, errors.NoWelcomer
	}
	exists = d.Redis.Exists(d.Decorate("welcomer:text:before")).Val()
	if exists == 0 {
		return nil, errors.NoWelcomer
	}
	return &Welcomer{d:d}, nil
}

func (w *Welcomer) GetMain() string {
	return w.d.Redis.Get(w.d.Decorate("welcomer:main")).Val()
}

func (w *Welcomer) GetSecret() string {
	return w.d.Redis.Get(w.d.Decorate("welcomer:secret")).Val()
}

func (w *Welcomer) GetTextAfter() string {
	return w.d.Redis.Get(w.d.Decorate("welcomer:text:after")).Val()
}

func (w *Welcomer) GetTextBefore() string {
	return w.d.Redis.Get(w.d.Decorate("welcomer:text:before")).Val()
}