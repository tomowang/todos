package main

import (
	"flag"
	"os"
	"time"
	"todos/config"
	"todos/db"

	"github.com/phuslu/log"
	"github.com/robfig/cron/v3"
)

func main() {
	executable, err := os.Executable()
	if err != nil {
		println("cannot get executable path")
		os.Exit(1)
	}

	var validate bool
	flag.BoolVar(&validate, "validate", false, "parse the config toml and exit")
	flag.Parse()
	err = config.Init()
	if err != nil {
		log.Fatal().Err(err).Str("filename", flag.Arg(0)).Msg("read config error")
	}

	if validate {
		os.Exit(0)
	}

	conf := config.GetConfig()
	if log.IsTerminal(os.Stderr.Fd()) {
		log.DefaultLogger = log.Logger{
			Level:      log.ParseLevel(conf.LogLevel),
			Caller:     1,
			TimeFormat: "15:04:05",
			Writer: &log.ConsoleWriter{
				ColorOutput: true,
			},
		}
	} else {
		log.DefaultLogger = log.Logger{
			Level: log.ParseLevel(conf.LogLevel),
			Writer: &log.FileWriter{
				Filename:   executable + ".log",
				MaxSize:    conf.LogMaxsize,
				MaxBackups: conf.LogBackups,
				LocalTime:  false,
			},
		}
	}
	runner := cron.New(cron.WithSeconds(), cron.WithLocation(time.UTC), cron.WithLogger(cron.PrintfLogger(&log.DefaultLogger)))
	// log rotating daily
	if !log.IsTerminal(os.Stderr.Fd()) {
		runner.AddFunc("0 0 0 * * *", func() { log.DefaultLogger.Writer.(*log.FileWriter).Rotate() })
	}
	db.Init()
	c := config.GetConfig()
	router := NewRouter()
	router.Run(c.Listen)
}
