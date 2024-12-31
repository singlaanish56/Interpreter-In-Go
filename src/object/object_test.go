package object

import (
	"testing"
)

func TestHashKey(t *testing.T){
	hello1 := &String{Value: "Hello world"}
	hello2 := &String{Value: "Hello world"}
	dif1 := &String{Value: "my name is anthony"}
	
	if hello1.HashKey() != hello2.HashKey(){
		t.Errorf("same strings should have the same hashkey, str1=%s, st2=%s", hello1.Value, hello2.Value)
	}

	if hello1.HashKey() == dif1.HashKey(){
		t.Errorf("different strings cant have the same hashkey , str1=%s, str2=%s", hello1.Value, dif1.Value)
	}
}