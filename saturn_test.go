package saturn_test

import (
	"testing"

	"github.com/saturnengine/saturn-engine"
)

func TestHello(t *testing.T) {
	if saturn.Hello() != "It is not ghost fighter!" {
		t.Errorf("Unexpected greeting from Saturn engine")
	}
}
