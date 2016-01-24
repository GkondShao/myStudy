package main

import (
    "fmt"
    "time"
    "math/rand"
    "crypto/md5"
    "strconv"
	"encoding/hex"

)

func main() {
	sessionId := make(chan string)
    go sessionIdGet(sessionId);
	fmt.Println(<-sessionId)
}

func sessionIdGet(sessionId chan string){
	
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
	//func (r *Rand) Int63() int64 { return r.src.Int63() }
	sessionNum := rand.Int63()//生成一个正的伪随机int63作为int64的数
	//int转string
	//1、int32位，strconv.Itoa
	//2、大于32位，strconv.FormatInt()
	//纳秒数的时间以及以其为种子的随机数作为sessionId的明文。
	boot :=strconv.FormatInt(sessionNum,10)+strconv.FormatInt(nano,10)
	hashMd5 := md5.New()
	hashMd5.Write([]byte(boot))
	result := hex.EncodeToString(hashMd5.Sum([]byte("")))
	
	sessionId <- result	
}