package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	//command := exec.Command("git","--version")
	command := exec.Command("ls", "-al")

	closer, e := command.StdoutPipe()
	if e != nil {
		log.Fatal(e)
	}
	//及时关闭
	defer closer.Close()
	//运行
	start := command.Run()
	if start != nil {
		log.Fatal(start)
	}
	//读取输出结果
	readAll, e := ioutil.ReadAll(closer)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("读取到命令输出： ", string(readAll))

}
