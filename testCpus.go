package main 

import(
	"fmt"
	"runtime"
	"time"
)

func calculate(from,to int,num []int,ch chan int){
	temp := 0
	
	for i:=from;i<to;i++{
		temp +=num[i]
	}
	
	ch<-temp
}

func cal(num []int) int{
	len := len(num)
	fmt.Println(len)
	
	
	cpus := runtime.NumCPU()
	fmt.Println(cpus)
	
	chs := make([]chan int,cpus)
	t1 :=time.Now()
	fmt.Println(t1)
	
	for i:=0;i<cpus;i++{
		chs[i] = make(chan int)
		go calculate(i*100/cpus,(i+1)*100/cpus,num[0:],chs[i])
	
	}
	result := 0
	
	for i:=0;i<cpus;i++{
		temp:=<-chs[i]
		fmt.Println(temp)
		result += temp
	}
	
	t2 :=time.Now()
	fmt.Println(t2)
	
	return result
}


func main(){
	var num [300]int
	
	for i:=0;i<100;i++{
		num[i] = i
	}
	
	fmt.Println(cal(num[0:]))
}