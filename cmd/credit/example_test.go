package credit_test

import (
	"fmt"
	_ "fmt"
	"github.com/NataliaOleinik/CreditAndExamples/pkg/credit"
)

func ExampleCalculate() {
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20))
	fmt.Println(credit.Calculate(10_000_00, 36, 20))
	// Output:
	// 3718400 33862300 133862300
	// 37184 338623 1338623
}
