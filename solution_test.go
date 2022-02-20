package solution

import (
	"testing"
)

func TestOk(in *testing.T){
	expected := string([]byte{72, 101, 108, 108, 111, 32, 240, 159, 151, 186, 239, 184, 143, 32, 33})
	res:=GetMessage()
	if res!=expected{
	in.Errorf("test finished with error \n %q \n %q ",res,expected)
	return
	in.Log("Ok")
}
}

