package main

import "fmt"
import "log"


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
func init(){


	log.SetFlags(log.LstdFlags | log.Lshortfile |log.LUTC)
}


func test_2(){
	log.Println("test_2")
	log.Println()
	
	
}



func test_1(){
	log.Println("test_1")
	
	
	
}




func  main(){

   fmt.Println("main")
   test_1()
   test_2()
   


}

