//go:generate abigen --abi MathTester.abi --bin MathTester.bin --pkg test --type MathTester --out ../math_tester.go
//go:generate abigen --abi AdderTester.abi --bin AdderTester.bin --pkg test --type AdderTester --out ../adder_tester.go
//go:generate abigen --abi Adder.abi --bin Adder.bin --pkg test --type Adder --out ../adder.go
package lib_math
