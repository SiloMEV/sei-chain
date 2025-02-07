package mev

import (
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"
)

type Config struct {
	ListenAddr string
}

const (
	flagListenAddr = "mev.listen_addr"
)

var DefaultConfig = Config{
	ListenAddr: ":22137",
}

func ReadConfig(opts servertypes.AppOptions) (Config, error) {
	cfg := DefaultConfig // copy
	var err error
	if v := opts.Get(flagListenAddr); v != nil {
		if cfg.ListenAddr, err = cast.ToStringE(v); err != nil {
			return cfg, err
		}
	}
	return cfg, nil
}
