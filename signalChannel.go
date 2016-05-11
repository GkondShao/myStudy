package main

import(
	"fmt"
	"os"
	"io"
	"errors"
	"bufio"
	"strings"
)
/**
	写着发现handle没用版本
*/
type Person struct{
	Name string
	Age string
	Address Addr
}

type Addr struct{
	city string
	
}

type PersonHandle interface{
	Batch(prigs <- chan Person) <- chan Person
	Handle(<- chan Person)
}


type PersonHandleImpl struct{}

func ( p PersonHandleImpl) Batch(origs <- chan Person) <- chan Person{
	descs := make(chan Person,100)
	
	go func(){
		fmt.Println("Person 处理开始")
		for{
			
			p,ok := <- origs
			if !ok{
				break
			}
			
			p.Age = "33"
			
			descs <- p
		}
			fmt.Println("Person 处理结束")
			close(descs)
	}()
	
	return descs
	
}

func (p PersonHandleImpl) Handle(<- chan Person){

}

func getPersonHandler() PersonHandle{
	return new(PersonHandleImpl)
}


func fetchPerson(origs chan <- Person){
	exist := IsExist("persons.txt")
	
	if !exist{
		checkErr(errors.New("待处理文本不存在"))
	}
	
	
	
	go func(){
		file ,err :=  os.Open("persons.txt")
		checkErr(err)
		defer file.Close()
		r := bufio.NewReader(file)
		fmt.Println("Person 录入开始")
		for{
		
			b,_,err := r.ReadLine()
			if err != nil{
				if err == io.EOF{
					break
				}
				panic(err)
			}
			
			str := string(b)
			strs := strings.Split(str,",")
			
			p := new(Person)
			p.Name = strs[0]
			p.Age = strs[1]
			p.Address.city = strs[2]
			origs <- *p
			
		}
		fmt.Println("Person 录入结束")
		close(origs)
		
	}()
	
	
	
	
}

func savePerson(descs <- chan Person) chan byte{
	sign := make(chan byte,2)
	file ,err :=  os.Create("persons2.txt")
	checkErr(err)
	defer file.Close()
	var strs string
	fmt.Println("Person 存储开始")
	for{
		
		p,ok := <- descs
		if !ok{
			break
		}
		str := p.Name+","+p.Age+","+p.Address.city+"\r\n"
		strs +=str
	}
	
	_,err = file.WriteString(strs)
	checkErr(err)
	
	fmt.Println("Person 存储结束")
	sign <- 1
	return sign

}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}


func IsExist(file string) bool{
	_,err := os.Stat(file)
	return err == nil ||os.IsExist(err)
}


func main(){
	handler := getPersonHandler()
	origs := make(chan Person,100)
	
	
	fetchPerson(origs)
	descs := handler.Batch(origs)
	
	sign := savePerson(descs)
	<- sign
	
	
}