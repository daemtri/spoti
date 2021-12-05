package logs

import "github.com/spf13/pflag"

type Options struct {
}

func NewLoggerOptions() *Options {
	return &Options{}
}

func (o Options) AddFlags(fs *pflag.FlagSet, c *Options) {

}

type Logger = zap.Logger

func NewLogger(opts *Options) *Logger {
	l, _ := zap.NewProduction()
	return l
}
