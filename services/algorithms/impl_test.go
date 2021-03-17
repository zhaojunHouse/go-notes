package algorithms

import (
	"testing"
)

func Test_MD5(t *testing.T) {
	t.Log(HashToNumber("5erw32523efw"))
	t.Log(HashToNumber("5erw32523efe"))
	t.Log(HashToNumber("5erw32523efr"))
}
