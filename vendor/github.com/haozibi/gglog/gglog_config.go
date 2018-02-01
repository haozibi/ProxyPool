// edit by haozibi
// 放置 gglog 新增方法

package gglog

import (
	"strconv"
	"strings"
)

var (
	prefix = ""
	// 输出形式default(默认),normal,simple,(复杂程度依次减少)
	outPutType     = outTypeDefault
	outTypeDefault = 0
	outTypeNormal  = 1
	outTypeSimple  = 2
)

var outPutTypes map[string]int = map[string]int{
	"default": outTypeDefault,
	"normal":  outTypeNormal,
	"simple":  outTypeSimple,
}

// 设置日志console 输出格式，DEFAULT,NORMAL,SIMPLE
func SetOutType(t string) {
	tmp := strings.ToLower(t)
	if _, ok := outPutTypes[tmp]; !ok {
		panic("Set Out type Error")
	}
	outPutType = outPutTypes[tmp]
}

//设置stderrThreshold值，只有大于等于 stderrThreshold 的级别才能正确输出，706行附近
// setOL => set output level
// 默认级别 ERROR
func SetOutLevel(value string) error {
	var threshold severity
	// Is it a known name?
	if v, ok := severityByName(value); ok {
		threshold = v
	} else {
		v, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		threshold = severity(v)
	}
	logging.stderrThreshold.set(threshold)
	return nil
}

// 设置日志输出目录，
func SetLogDir(dir string) {
	logDirs = append(logDirs, dir)
}

func SetPrefix(p string) {
	prefix = p
}
