package main

import "fmt"

func main() {
	adams := 42
	fmt.Printf("42 as binary: %b\n", adams)
	fmt.Printf("42 as hexadecimal: %x\n", adams)

	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	fmt.Printf("a as binary: %b\n", a)
	fmt.Printf("b as binary: %b\n", b)
	fmt.Printf("c as binary: %b\n", c)
	fmt.Printf("d as binary: %b\n", d)
	fmt.Printf("e as binary: %b\n", e)
	fmt.Printf("f as binary: %b\n", f)

	fmt.Printf("a as hexadecimal: %x\n", a)
	fmt.Printf("b as hexadecimal: %x\n", b)
	fmt.Printf("c as hexadecimal: %x\n", c)
	fmt.Printf("d as hexadecimal: %x\n", d)
	fmt.Printf("e as hexadecimal: %x\n", e)
	fmt.Printf("f as hexadecimal: %x\n", f)

	fmt.Printf("%x, #%b\n", f, f)
}
