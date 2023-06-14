package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://blog.csdn.net/xinRCNN/article/details/129048119")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()            //defer关键字修饰的语句会推迟到执行return语句或函数执行完毕以及发生错误之后才会执行,defer关键字修饰的匿名内部函数可以读取其外层函数函数的返回值
	body, err := io.ReadAll(resp.Body) //读取流信息
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	re := regexp.MustCompile("<title>(.*)</title>") //正则
	title := re.FindStringSubmatch(string(body))[0] //
	fmt.Println("Title:", title)
}
