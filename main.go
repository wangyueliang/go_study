package main

import (
	"fmt"
	m_test_0102 "go_study/20200112"
	"log"
	"net"
	"reflect"
	"sync"
)

// "go_test/m_channel"
//"employee"
func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
func map_test() {

	/*
		无须初始化，直接声明即可。
		sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，
		Store 表示存储，Load 表示获取，Delete 表示删除。

		使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中
		回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。
	*/
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	a, ok := scene.Load("london")
	if ok {
		log.Println("a  ", a)
	}

	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		log.Println("iterate:", k, v)
		return true
	})
}

type server struct {
}

func (s *server) Start() {
	listener, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Println("listen error: ", err)
		return
	}

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error: ", err)
			break
		}

		go s.handleConnect(conn)
	}
}
func (s *server) handleConnect(conn net.Conn) {

	defer conn.Close()

	for {
		// read from the connection
		// ... ...
		// write to the connection
		//... ...
	}
}

//接口测试

type VowelFinder interface {
	FindVowels() []rune //切片
	//rune 类型。rune 类型等价于 int32 类型。
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, r := range ms {
		if r == 'a' || r == 'w' {
			vowels = append(vowels, r)
		}
	}
	return vowels
}

type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describe(t Test) { //  传递的是 接口参数
	fmt.Printf("Interface type %T value %v\n", t, t)
}

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() { // 使用值接受者实现 wangxue
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { // 使用指针接受者实现  wangxue
	fmt.Printf("State %s Country %s", a.state, a.country)
}

/// 接口  2
// type SalaryCalculator interface {
// 	DisplaySalary()
// }

// type LeaveCalculator interface {
// 	CalculateLeavesLeft() int
// }

// type EmployeeOperations interface {
// 	SalaryCalculator
// 	LeaveCalculator
// }
//或是 下面

type EmployeeOperations interface {
	DisplaySalary()
	CalculateLeavesLeft() int
}
type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

//  上面  是  接口作为变量 赋值
//  下面是  接口 作为 函数的参数在传递
//  反射 reflect    包  2021 03 19

type order struct {
	ordId      int
	customerId int
	name       string
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func createQuery1(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		log.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			log.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))

			switch v.Field(i).Kind() {
			case reflect.Int:
				log.Println(v.Field(i).Int())
			case reflect.String:
				log.Println(v.Field(i).String())
			default:
				log.Println("Unsupported type")
				return
			}

		}

	}
}

func createQuery2(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	log.Println("Type ", t)
	log.Println("Value ", v)

	k := t.Kind()
	log.Println("Kind ", k)

}

func jiekou_test() {
	name := MyString("wangxue")
	var v VowelFinder //  接口 类型 变量
	v = name
	fmt.Println("vowels is ", v.FindVowels())
	fmt.Println("vowels is  string ", string(v.FindVowels()))
	var t Test
	f := MyFloat(89.7)
	t = f
	t.Tester()
	describe(t)
	tp := t.(Test) // 不是  string 类型   t.(T)
	log.Println("tp   is ", tp)

	// reflect.TypeOf(s.getter) 这个还不会

	// 接口 2
	log.Println("接口2 测试...........")
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}

	/* 如果下面一行取消注释会导致编译错误：
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)
	*/
	d2 = &a // 值接受和指针接受        的不同

	//d2 = &a // 这是合法的
	// 因为在第 22 行，Address 类型的指针实现了 Describer 接口
	d2.Describe()

	//  多个接口函数  或是 一个接口包含多个接口函数

	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmployeeOperations = e //  接口可以被 实现接口的结构体  赋值
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}

func fanshe_test() {
	// 反射 学习
	o := order{
		ordId:      1234,
		customerId: 567,
		name:       "wangxue",
	}
	log.Println(createQuery(o))
	createQuery(o)
	createQuery1(o)
	createQuery2(o)
}

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) { //
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func increment_chanel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func select_test() {
	var w sync.WaitGroup
	var m sync.Mutex
	//ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m) // 共享变量 建议用锁
		//go increment_chanel(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
func hello_channel(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func sendData(sendch chan<- int) { //   双向  信道 可以转化为单向信道，反之不行
	sendch <- 10
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

//mv $WORK\b001\exe\a.out.exe C:\Users\Administrator\go\bin\go_test.exe

// go 中文网
// Go 系列教程 —— 23. 缓冲信道和工作池（Buffered Channels and Worker Pools）
// 接口 管道  缓冲管道工作池 锁  结构体继承多态  panic&recover  defer() refelect
// bytes 包   sync.Map   tcp server listen  connectfd(goroutinue)
//
func main() {
	log.Println(" test  main() ...")
	//map_test()

	var query string = "wangxue"
	query = fmt.Sprintf("%s, %d", query, 100) //格式化 输出
	log.Println(query)
	/*
		// s := &server{}
		// s.Start()
		//jiekou_test()
		//fanshe_test()
		select_test()

		// e := m_employee.Employee{ //模块测试中
		// 	FirstName:   "Sam",
		// 	LastName:    "Adolf",
		// 	TotalLeaves: 30,
		// 	LeavesTaken: 20,
		// }
		// e.LeavesRemaining()

		//  var e employee.Employee
		e := m_employee.New("Sam", "Adolf", 30, 20)
		e.LeavesRemaining()

		m_employee.Test_jicheng_main() //    模块.函数名()

		m_interface.Test_interface_main()

		m_interface.Test_defer_main() //defer 测试
		//  断言　　自定义　error
		// 	// recover 复原        panic 恐慌
		// 	在本教程里，我们还会接着讨论，当程序发生 panic 时，使用 recover 可以重新获得对该程序的控制。
		// 可以认为 panic 和 recover 与其他语言中的 try-catch-finally 语句类似，只不过一般我们很少使用
		// panic 和 recover。而当我们使用了 panic 和 recover 时，也会比 try-catch-finally 更加优雅，代码更加整洁。

		m_interface.Test_panic_main() // 捕获 异常

		//m_interface.Test_panic_main2() // 不是同一个协程    不能捕获异常

		//m_interface.Test_panic_main3() //数组 越界 触发  panic

		m_interface.Test_bibao_main()

		//bytes.buffer
		var s string = "name:wangxue."
		w := bytes.NewBufferString(s)
		fmt.Fprintf(w, "hello.m3u8")
		log.Println("bytes string is ", w.String())

		// 管道
		done := make(chan bool)
		go hello_channel(done)
		flag := <-done //   阻塞等待 goroutinue 结束
		fmt.Println("main function wait goroutinue  ...   ", flag)

		number := 5672
		sum := 0
		for number != 0 {
			digit := number % 10
			sum += digit //* digit/ * digit
			number /= 10
			log.Println("digit is ", digit, " sum: ", sum)
		}

		// ch := make(chan int)
		// ch <- 5 // 死锁

		cha1 := make(chan int)
		go sendData(cha1)
		fmt.Println("int chan  recv  data: ", <-cha1)

		ch := make(chan int)
		go producer(ch)
		for {
			v, ok := <-ch
			if ok == false { //   从关闭的信道读取数据 返回false
				log.Println("Received ", ok, v)
				break
			}
			log.Println("Received ", ok, v)
		}

		log.Println("for range  接收数据直到 信道关闭  ....")
		ch1 := make(chan int)
		go producer(ch1)
		for v := range ch1 {
			log.Println("for range Received ", v)
		}
	*/

	ch := make(chan string, 3)
	ch <- "hello"
	ch <- "world"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))

	//管道是协程安全的   一个生产者 多个消费者  就是 消费者池
	//m_channel.Test_channelPool_main()

	m_test_0102.Test_jichu_main()
}
