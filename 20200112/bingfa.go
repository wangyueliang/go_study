package m_test_0102

import (
	"fmt"
	"log"
	"time"
)

//程序 小任务  goroutine channel 通信
//struct  interface 封装  继承 多态

//封装

type Foo struct {
	name string
}

func (f *Foo) echo() {
	log.Println(f.name)
}

//继承
type Bar struct {
	Foo
}

//多态  用到interface
type Ifoo interface {
	qux()
}

type Ibar struct {
	name string
}
type Ibaz struct {
	name string
}

func (b Ibar) qux() { //实现 了 接口函数
	log.Println("接口 Ibar 类型 一样 实现不一样 ")
}
func (b Ibaz) qux() {
	log.Println("接口 Ibaz Ibaz 类型 一样 实现不一样 2")
}

//墓库网 go并发编程案例分析
func jiekou() {
	log.Println("bing fa ..................封装...")

	f := Foo{name: "hello 封装对象测试"}
	f.echo()

	log.Println("bing fa ......................继承...")
	b := Bar{Foo{name: "hello 继承对象测试"}}
	b.echo()

	log.Println("bing fa ........................................多态...")
	var f1 Ifoo
	f1 = Ibar{} // 类型 一样 实现不一样
	log.Println(f1)
	f1.qux()
	f1 = Ibaz{} //类型 一样 实现不一样
	log.Println(f1)
	f1.qux()

}

func channel_select_test() {

	c := make(chan string) //创建 没有缓存的 channel

	c1 := make(chan string) //创建 没有缓存的 channel

	go func() {

		time.Sleep(1 * time.Second)
		c <- "message from test 1 second"
	}()

	go func() {

		time.Sleep(2 * time.Second)
		c1 <- "message from test 2 second"
	}()

	time.Sleep(3 * time.Second)

	select {
	case msg := <-c: // 阻塞接收数据
		log.Println(msg)

	case msg1 := <-c1: // 阻塞接收数据
		log.Println(msg1)
	default: //可选
		log.Println("no read ...")

	}
	/*
		switch msgType{
			case "":
			break
			case "":
			break
			default:
			break
			}

			go 微信 测试工具 像是模拟微信客户端 和 公众号 测试
			//http://github.com/weixinhost  侯斯特——微信公众号的应用商店
	*/

}

func printHello(i int, ch chan string) {
	//for {
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("hello world from "+"goroutine %d", i)
	//}
}
func channel_routine_test() { //  缓冲池  生产者 消费者  一个消费了5个

	log.Println("channel_routine_test .......... ")
	ch := make(chan string, 1)

	for i := 0; i < 5; i++ {
		//重点 记录
		//重点 记录
		//重点 记录
		go printHello(i, ch) //这个也是 routine 否则不返回啊  重点 记录
	}

	for {
		msg := <-ch
		log.Println(msg)
	} //主程序 结束这些 routine  都会被杀死

}
