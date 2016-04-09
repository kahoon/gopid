// A simple go utility to run a service pid to /var/run on linux systems
package gopid

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	PATH = "/var/run"
	EXT  = "pid"
)

var (
	ErrNameEmpty = errors.New("gopid: empty service name")
)

func Run(svcname string) error {
	if len(svcname) == 0 {
		return ErrNameEmpty
	}
	return ioutil.WriteFile(fmt.Sprintf("%s/%s.%s", PATH, svcname, EXT), []byte(strconv.Itoa(os.Getpid())), 644)
}
