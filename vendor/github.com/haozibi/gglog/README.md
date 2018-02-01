## gglog

[![Build Status](https://travis-ci.org/haozibi/gglog.svg?branch=master)](https://travis-ci.org/haozibi/gglog) ![](https://img.shields.io/badge/language-go-red.svg)

基于 **[glog](https://github.com/golang/glog)** 进行改进的轻量级日志系统。

### 改进内容

* 添加 DEGUG 级别，五种日志等级 DEBUG < INFO < WARING < ERROR < FATAL
* 增加 SetOutLevel 方法，通过 SetOutLevel 设置 stderrThreshold 值，级别大于等于 stderrThreshold 则会在控制台输出信息，默认级别 ERROR
* 增加 SetLogDir 方法，通过 SetLogDir 设置日志文本输出路径
* 增加 SetOutType 方法，通过 SetOutType 设置 console 输出格式，默认为完整输出(DEFAULT)，DEFAULT > NORMAL > SIMPLE
* 增加 SetPrefix 方法，通过 SetPrefix 设置日志前缀
* 优化日志输出名字，把“级别”标签放在靠前的位置
* **其他操作与 [glog](https://github.com/golang/glog) 完全一样**

### 示例
```
package main

import (
    "flag"

    "github.com/haozibi/gglog"
)

func main() {
    flag.Parse()
    defer gglog.Flush()

    // 设置控制台输出级别，比此级别大的都会在控制台输出
    // DEBUG < INFO < WARING < ERROR < FATAL , 默认ERROR级别
    gglog.SetOutLevel("DEBUG")

    // 设置前缀
    gglog.SetPrefix("[gglog] ")

    // 设置log输出路径，路径必须存在
    gglog.SetLogDir("log")

    // 设置输出格式，DEFAULT > NORMAL > SIMPLE, 默认 DEFAULT 格式
    gglog.SetOutType("DEFAULT")
    // gglog.SetOutType("NORMAL")
    // gglog.SetOutType("SIMPLE")

    gglog.Info("Hello gglog")
    gglog.Debug("This is a Debug log")
    gglog.Warning("This is a Warning log")
    gglog.Error("This is a Error log")

    gglog.Infof("info %d", 1)
    gglog.Warningf("warning %d", 2)
    gglog.Errorf("error %d", 3)

    gglog.V(1).Infoln("level 1")
    gglog.V(2).Infoln("level 2")
}

```

### 截图

![gglog](https://i.loli.net/2018/01/22/5a6556a05ff34.jpg)
