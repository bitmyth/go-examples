// https://yourbasic.org/golang/switch-statement/
// Execution order
// First the switch expression is evaluated once.
// Then case expressions are evaluated left-to-right and top-to-bottom:
// the first one that equals the switch expression triggers execution of the statements of the associated case,
// the other cases are skipped.
package main
import(
	"fmt"
)

// Foo prints and returns n.
func Foo(n int) int {
    fmt.Println(n)
    return n
}

func main() {
    switch Foo(2) {
    case Foo(1), Foo(2), Foo(3):
        fmt.Println("First case")
        fallthrough
    case Foo(4):
        fmt.Println("Second case")
    }
}
