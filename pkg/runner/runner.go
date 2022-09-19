package runner

import (
	"fmt"
	"github.com/yhy0/FuckFingerprint/fingerPrints"
	"github.com/yhy0/FuckFingerprint/pkg/afrog"
	"github.com/yhy0/FuckFingerprint/pkg/logging"
	"github.com/yhy0/FuckFingerprint/pkg/util"
	"os"
	"strings"
	"sync"
)

/**
  @author: yhy
  @since: 2022/9/19
  @desc: //TODO
**/

func Run(options *Options) {
	var targets []string
	if options.TargetFile != "" {
		targets = util.LoadFile(options.TargetFile)
	}
	if options.Target != "" {
		targets = append(targets, options.Target)
	}

	var f *os.File
	if options.Output != "" {
		var err error
		f, err = os.Create(options.Output)
		if err != nil {
			logging.Logger.Fatal("Could not create output file '%s': %s\n", options.Output, err)
		}
		defer f.Close() //nolint
	}

	if options.Global {
		logging.Logger.Infoln("Please wait while online fingerprint is obtained...")
		err := fingerPrints.GetEHoleFingerDataOnline()
		if err != nil {
			logging.Logger.Warnln("Use default data(EHoleFinger)")
			fingerPrints.NewEHoleFinger()
		}
		err = fingerPrints.GetLocalFingerDataOnline()
		if err != nil {
			logging.Logger.Warnln("Use default data(LocalFinger)")
			fingerPrints.NewLocalFinger()
		}
		err = fingerPrints.GetAfrogFingerDataOnline()
		if err != nil {
			logging.Logger.Warnln("Use default data(AfrogFinger)")
			fingerPrints.NewAfrogFinger()
		}
	} else {
		fingerPrints.NewEHoleFinger()
		fingerPrints.NewLocalFinger()
		fingerPrints.NewAfrogFinger()
	}

	wg := sync.WaitGroup{}
	limiter := make(chan bool, options.Thread)

	for _, target := range targets {
		wg.Add(1)
		limiter <- true

		go func(url string) {
			var t string
			if !strings.Contains(url, "http://") && !strings.Contains(url, "https://") {
				t = "http://" + url
			}

			res, title, err := afrog.Run(t)

			if len(res) == 0 && title == "" {
				t = "https://" + url
				res, title, err = afrog.Run(t)
				if f != nil && err == nil {
					f.WriteString(fmt.Sprintf("%s %s %v\n", target, title, res))
				}
			} else {
				if f != nil && err == nil {
					f.WriteString(fmt.Sprintf("%s %s %v\n", target, title, res))
				}
			}

			<-limiter
			wg.Done()
		}(target)

	}

	wg.Wait()

}
