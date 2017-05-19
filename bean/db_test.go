package bean

import "testing"
import "fmt"

func TestDB(t *testing.T) {
	en, er := GetDBConenct()
	if er != nil {
		fmt.Println(er)
		t.Fail()
	} else {
		fmt.Println(en.Ping())
	}

}
