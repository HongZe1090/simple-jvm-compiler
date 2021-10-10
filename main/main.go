package main

import (
	"fmt"
	"jvmgo/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	// 输入 -version命令，输出版本号cmd
	// 输入 -help命令，调用printUsage函数输出相关帮助信息cmd
	// 否则 正常启动，调用startJVM
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args: %v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)

	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Print("Could not find or load main class %s \n", cmd.class)
		return
	}

	fmt.Printf("class Date:%v\n", classData)
	// fmt.Printf("classpath:%s class:%s args:%v\n",
	// 	cmd.cpOption, cmd.class, cmd.args)
}
