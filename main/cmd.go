package main

import (
	"flag"
	"fmt"
	"os"
)

// 结构体表示命令行参数和选项
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

//  对命令的执行
func parseCmd() *Cmd {
	cmd := &Cmd{} //初始化Cmd结构体

	// flag实现了命令行参数的解析
	// 第二个参数即是 -后面的值，后面依次是arg[]数组中的值
	// flag.Prase 来解析命令行参数写入注册的flag里。
	flag.Usage = printUsage //在解析失败时调用printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()

	args := flag.Args() //获取参数数组
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}
