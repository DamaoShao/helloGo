package main

import "fmt"

/*
go的价值:
+ 多核硬件架构
+ 超大规模分布式计算集群
+ web模式导致的开发规模和速度
*/

//入口：main package 的 main函数
//main函数不支持返回值 也不支持传入参数
//通os.Exit()来返回
//通过os.args来接受参数

func main() {
	fmt.Println("hello world!")
}