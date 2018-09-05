package randxString

import(
	"testing"
)



func TestRandXBitString(t *testing.T){
	x := RandXBitString(12)

	if len(x) != 12{
		t.Fail()
	}

	t.Log("success")
}


func BenchmarkRandXBitString(b *testing.B){
	for i:=0;i<b.N;i++{
		RandXBitString(i)
	}
}

func BenchmarkRandXBitStringOld(b *testing.B){
	for i:=0;i<b.N;i++{
		RandXBitStringOld(i)
	}
}

func BenchmarkRandXBitStringCopy(b *testing.B){
	for i:=0;i<b.N;i++{
		RandXBitStringCopy(i)
	}
}

