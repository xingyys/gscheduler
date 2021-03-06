package main

import (
	"time"

	"github.com/lack-io/gscheduler"
)

func main() {
	// 构建一个符合正则表达式的任务
	gscheduler.JobBuilder().Name("cron-job").Spec("*/10 * * * * * *").Out()

	// 构建一次性延时任务
	gscheduler.JobBuilder().Name("delay-job").Delay(time.Now().Add(time.Hour * 3)).Out()

	// 构建间隔执行的任务
	gscheduler.JobBuilder().Name("duration-job").Duration(time.Second * 10).Out()

	// 构建多次执行的任务
	gscheduler.JobBuilder().Name("three-times-job").Duration(time.Second*5).Times(3).Out()
}
