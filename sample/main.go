package main

import (
	_ "go/sample/matchers"
	"go/sample/search"
	"log"
	"os"
)

//init 在main函数之前调用
func init() {
	log.SetOutput(os.Stdout)
}

//main是整个程序的入口
func main() {
	//使用特定的项搜索
	search.Run("president")
}
