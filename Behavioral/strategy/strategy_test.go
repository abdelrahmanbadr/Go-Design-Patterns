package strategy

import (
	"fmt"
)

func ExampleStrategy() {
	ctx := new(Context)
	ctx.SetStrategy(new(Add))
	fmt.Println(ctx.Calc(3, 4))
	ctx.SetStrategy(new(Subtract))
	fmt.Println(ctx.Calc(8, 4))
	ctx.SetStrategy(new(Multiply))
	fmt.Println(ctx.Calc(2, 5))

	// Output:
	// 7
	// 4
	// 10

}
