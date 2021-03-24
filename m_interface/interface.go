package m_interface

import (
	"fmt"
	"log"
	"runtime/debug"
	"time"
)

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

//  新增收益 流
type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}

func Test_interface_main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}

	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}

	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd} //切片
	calculateNetIncome(incomeStreams)                                          //  接口就是 结构体地址
}

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func hello(n int) {
	log.Println("n = ", n)
}
func Test_defer_main() {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)

	defer hello(0) //和ｃ++ 先构造后析构顺序一样
	defer hello(1)
	defer hello(2)
	defer hello(3)

}

//  自己定义错误
// type areaError struct {
//     err    string
//     radius float64
// }

// func (e *areaError) Error() string {
//     return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
// }

// err.(*areaError) go 断言
// var t interface{}
// t = functionOfSomeType()
// switch t := t.(type) {
// ｝ 断言类型

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("这里捕获异常 recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil { ///   last name 捕获异常
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func Test_panic_main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func Test_panic_main2() { // 不是同一个协程    不能捕获异常
	a()
	fmt.Println("normally returned from main")
}

////////////////////////////////

func r() {
	if r := recover(); r != nil {
		fmt.Println("捕获数组越界   .. Recovered", r)
		debug.PrintStack()
	}
}
func a3() {
	//defer r() //   捕获数组越界

	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}
func Test_panic_main3() {
	a3()
	fmt.Println("normally returned from main")
}

//闭包（Closure）是匿名函数的一个特例。当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包。
//闭包函数
func Test_bibao_main() {
	a := 100
	func() {
		log.Println("a = ", a)
	}()
}
