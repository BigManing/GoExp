package algorithm

import (
	"testing"
)

func TestIter(t *testing.T) {
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	th.Iter()
	//v:=<-th.Iter()
	//fmt.Sprintf("%s%v","ch",v)

}
