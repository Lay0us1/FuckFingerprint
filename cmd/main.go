package main

import "github.com/yhy0/FuckFingerprint/pkg/runner"

/**
  @author: yhy
  @since: 2022/9/19
  @desc: //TODO
**/

func main() {
	options := runner.ParseArguments()
	runner.Run(&options)
}
