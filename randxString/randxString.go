// 生成n位随机字符串方法    仅数字以及大小写字母

package randxString

import(
	"math/rand"
	"time"
	//"os"
)

var src = rand.NewSource(time.Now().Unix())

const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
 const(
	letterIdxBits = 6 // 62 = 111110
	letterIdxMax = 63/letterIdxBits  // int63 可用分组
	letterIdxMask = 1<<letterIdxBits-1 // 截取6位的掩码 
 )





func RandXBitString(n int) string{
	if n<=0 {
		return ""
	}

	b := make([]byte,n)
	length := len(letterBytes)

	for i,cache,remain := n-1,rand.Int63(),letterIdxMax;i>=0;{
		if remain == 0{
			cache,remain = rand.Int63(),letterIdxMax
		}

		if idx := int(cache&letterIdxMask);idx < length{
			b[i] = letterBytes[idx]
			i--
		}

		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}


func RandXBitStringOld(n int)string{
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = letterBytes[b&61]
	}
	return string(bytes)
}


func RandXBitStringCopy(n int) string{
	if n<=0 {
		return ""
	}

	b := make([]byte,n)
	length := len(letterBytes)

	for i,cache,remain := n-1,src.Int63(),letterIdxMax;i>=0;{
		if remain == 0{
			cache,remain = src.Int63(),letterIdxMax
		}

		if idx := int(cache&letterIdxMask);idx < length{
			b[i] = letterBytes[idx]
			i--
		}

		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}