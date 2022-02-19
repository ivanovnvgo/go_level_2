package doc

import (
	"fmt"
)

func ExampleHelloWorld() string {
	var name string = enterName()
	return fmt.Sprintf("Hello, %s!!!", name)
	//Output: Hello, name!!!
}
