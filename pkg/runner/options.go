package runner

import (
	"flag"
	"fmt"
	"github.com/yhy0/FuckFingerprint/pkg/config"
	"github.com/yhy0/FuckFingerprint/pkg/util"
	"os"
)

/**
  @author: yhy
  @since: 2022/9/17
  @desc: //TODO
**/

type Options struct {
	Target     string
	TargetFile string
	Proxy      string
	Output     string
	Global     bool
	Thread     int
}

func ParseArguments() Options {
	showBanner()
	option := Options{}
	flag.StringVar(&option.Target, "url", "", "Target. eg:http://127.0.0.1")
	flag.IntVar(&option.Thread, "t", 30, "Thread")
	flag.StringVar(&option.TargetFile, "f", "", "Targets. eg:t.txt")
	flag.StringVar(&option.Proxy, "p", "", "http(s) proxy. eg: http://127.0.0.1:8080")
	flag.StringVar(&option.Output, "o", "", "Output results to file.")
	flag.BoolVar(&option.Global, "g", true, "Use online fingerprint")
	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if option.Target == "" && option.TargetFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	if option.Proxy != "" {
		config.Proxy = option.Proxy
	}

	if util.PathExists(option.Output) == nil {
		_ = os.Remove(option.Output)
	}

	return option
}
