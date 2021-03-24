package m_test_0102

import (
	"fmt"
	"log"
	"time"
)

/*
const (
    Ldate         = 1 << iota     //日期示例： 2009/01/23
    Ltime                         //时间示例: 01:23:23
    Lmicroseconds                 //毫秒示例: 01:23:23.123123.
    Llongfile                     //绝对路径和行号: /a/b/c/d.go:23
    Lshortfile                    //文件和行号: d.go:23.
    LUTC                          //日期时间转为0时区的
    LstdFlags     = Ldate | Ltime //Go提供的标准抬头信息
)

*/
func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

func test_2() {
	log.Println("test_2")
	log.Println()

}

func test_1() {
	log.Println("http://tour.studygolang.com/welcome/1 go 中国")
}

func test_3() {

}

//  go run  test_0112.go bingfa.go  编译
func Test_jichu_main() {

	fmt.Println("main")
	test_1()
	test_2()
	test_3()
	//bingfa()//go 并发编程实例解析 日志监控系统
	//go channel_select_test()

	go channel_routine_test()

	time.Sleep(4 * time.Second) // add done wait 等待routine 运行完成

	// reader writer
	//  go tool tour //http://docscn.studygolang.com/doc

	//go 并发编程实战 例子
	//goc2p ds pds  现实文件目录的文件
	//项目地址 http://github.com/hyper-carrot/goc2p

	// devops  自动化部署到阿里云
	// http://github.com/avenssi/newweb

	// flag.parse 记得调用 新需求尽量用 json //有良知的知识分子  老罗

}
