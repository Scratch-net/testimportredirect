package b

import "gopkg.in/testimportredirect.v0/a"

func B() {
	a.Foo()
}
