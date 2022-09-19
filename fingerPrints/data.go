package fingerPrints

import (
	_ "embed"
	"encoding/json"
	"errors"
	"github.com/yhy0/FuckFingerprint/pkg/config"
	"github.com/yhy0/FuckFingerprint/pkg/logging"
	"github.com/yhy0/FuckFingerprint/pkg/util"
)

/**
  @author: yhy
  @since: 2022/9/19
  @desc: //TODO
**/

//go:embed web_fingerprint_v3.json
var AfrogFinger []byte

//go:embed eHoleFinger.json
var EHoleFinger string

//go:embed localFinger.json
var LocalFinger string

var EholeFinpx *Packjson
var LocalFinpx *Packjson

type Packjson struct {
	Fingerprint []VscanFingerprint
}

type VscanFingerprint struct {
	Cms      string
	Method   string
	Location string
	Keyword  []string
}

var AfrogFingerPrintMap map[string][]AfrogFingerPrint

type AfrogFingerPrint struct {
	Name           string            `json:"name"`
	Path           string            `json:"path"`
	RequestMethod  string            `json:"request_method"`
	RequestHeaders map[string]string `json:"request_headers"`
	RequestData    string            `json:"request_data"`
	StatusCode     int               `json:"status_code"`
	Headers        map[string]string `json:"headers"`
	Keyword        []string          `json:"keyword"`
	FaviconHash    []string          `json:"favicon_hash"`
	Priority       int               `json:"priority"`
}

// NewEHoleFinger 只加载一次，聚合 path, 减少请求(为了集成到扫描器中考虑)
func NewEHoleFinger() {
	var config Packjson
	err := json.Unmarshal([]byte(EHoleFinger), &config)
	if err != nil {
		logging.Logger.Error("vscan EHoleFinger err ", err)
	}

	EholeFinpx = &config
}

// NewLocalFinger 只加载一次，聚合 path, 减少请求(为了集成到扫描器中考虑)
func NewLocalFinger() {
	var config Packjson
	err := json.Unmarshal([]byte(LocalFinger), &config)
	if err != nil {
		logging.Logger.Error("vscan LocalFinger err ", err)
	}

	LocalFinpx = &config
}

// NewAfrogFinger 只加载一次，聚合 path, 减少请求(为了集成到扫描器中考虑)
func NewAfrogFinger() {
	var fpSlice []AfrogFingerPrint
	if err := json.Unmarshal(AfrogFinger, &fpSlice); err != nil {
		logging.Logger.Error("afrog init err ", err)
	}

	AfrogFingerPrintMap = make(map[string][]AfrogFingerPrint)

	for _, v := range fpSlice { // 聚合 path
		AfrogFingerPrintMap[v.Path] = append(AfrogFingerPrintMap[v.Path], v)
	}
}

func GetEHoleFingerDataOnline() error {
	requset, err := util.HttpRequset(config.EHoleFingerDataOnline, "get", "", false, nil)
	if err != nil {
		logging.Logger.Error("Online data(EHoleFinger) acquisition failed,", err)
		return err
	}
	if requset.Body != "" {
		var config Packjson
		err := json.Unmarshal([]byte(EHoleFinger), &config)
		if err != nil {
			logging.Logger.Error("vscan EHoleFinger err ", err)
			return err
		}
		EholeFinpx = &config
	} else {
		return errors.New("no data")
	}

	return nil
}

func GetLocalFingerDataOnline() error {
	requset, err := util.HttpRequset(config.LocalFingerDataOnline, "get", "", false, nil)
	if err != nil {
		logging.Logger.Error("Online data(LocalFinger) acquisition failed,", err)
		return err
	}

	if requset.Body != "" {
		var config Packjson
		err := json.Unmarshal([]byte(LocalFinger), &config)
		if err != nil {
			logging.Logger.Error("vscan LocalFinger err ", err)
			return err
		}
		LocalFinpx = &config
	} else {
		return errors.New("no data")
	}
	return nil
}

func GetAfrogFingerDataOnline() error {
	requset, err := util.HttpRequset(config.AfrogFingerDataOnline, "get", "", false, nil)
	if err != nil {
		logging.Logger.Error("Online data(AfrogFinger) acquisition failed,", err)
		return err
	}

	if requset.Body != "" {
		var fpSlice []AfrogFingerPrint
		if err := json.Unmarshal(AfrogFinger, &fpSlice); err != nil {
			logging.Logger.Error("afrog init1 err ", err)
			return err
		} else {
			AfrogFingerPrintMap = make(map[string][]AfrogFingerPrint)
			for _, v := range fpSlice { // 聚合 path
				AfrogFingerPrintMap[v.Path] = append(AfrogFingerPrintMap[v.Path], v)
			}
		}
	} else {
		return errors.New("no data")
	}

	return nil
}
