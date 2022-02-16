package solution

import (
	"testing"
)
const okResult = "Hello 🗺️ !!!"
func TestOk(in *testing.T){
	res:=GetMessage()
if res!=okResult{
	in.Errorf("test finished with error")
	return
}
}

