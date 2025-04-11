package mev

import (
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"
	"time"
)

type Config struct {
	Enabled      bool
	ServerAddr   string
	PollInterval time.Duration
}

const (
	flagEnabled      = "mev.enabled"
	flagServerAddr   = "mev.server_addr"
	flagPollInterval = "mev.poll_interval"
)

var DefaultConfig = Config{
	Enabled:      false,
	ServerAddr:   "",
	PollInterval: 200 * time.Millisecond,
}

func ReadConfig(opts servertypes.AppOptions) (Config, error) {
	cfg := DefaultConfig // copy
	var err error
	if v := opts.Get(flagEnabled); v != nil {
		if cfg.Enabled, err = cast.ToBoolE(v); err != nil {
			return cfg, err
		}
	}
	if !cfg.Enabled {
		return cfg, nil
	}
	if v := opts.Get(flagServerAddr); v != nil {
		if cfg.ServerAddr, err = cast.ToStringE(v); err != nil {
			return cfg, err
		}
	}
	if v := opts.Get(flagPollInterval); v != nil {
		if cfg.PollInterval, err = cast.ToDurationE(v); err != nil {
			return cfg, err
		}
	}
	return cfg, nil
}
