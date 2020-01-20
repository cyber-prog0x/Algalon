package debian_base_chk

import (
	"fmt"
	"io/ioutil"
)

var passwd string = "/etc/passwd"

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func AbnormalUserCheck() {

	dat, err := ioutil.ReadFile(passwd)
	errorCheck(err)

	fmt.Println(dat)

}
