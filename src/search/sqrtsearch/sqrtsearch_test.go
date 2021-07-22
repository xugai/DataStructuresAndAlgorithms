package sqrtsearch

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestDecimal(t *testing.T) {
	a := decimal.NewFromFloat(2.345342)
	s := a.String()
	str := strings.Split(s, ".")
	fmt.Println(len(str[1]))
}

func TestSqrtsearch(t *testing.T) {
	var wg sync.WaitGroup
	target := 7.0
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(Sqrtsearch(decimal.NewFromFloat(target)))
		wg.Done()
	}()
	fmt.Println(math.Sqrt(target))
	wg.Wait()
}
