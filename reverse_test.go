package reverse_test

import (
	"testing"

	"github.com/youyo/reverse"
)

func TestBool(t *testing.T) {
	boolT := reverse.Bool(false)
	boolF := reverse.Bool(true)
	if boolT && !boolF {
		t.Log("ok")
	} else {
		t.Error("failed")
	}
}
