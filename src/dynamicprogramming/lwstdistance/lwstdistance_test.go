package lwstdistance

import (
	"fmt"
	"testing"
)

func TestLwstDistance(t *testing.T) {
	a, b := "mtacnu", "mitcmu"
	fmt.Println(lwstdistance(a, b, len(a), len(b)))
}
