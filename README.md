# poc-golang-testing

Proof of concept Golang testing unit with `testing package of Golang`

###Run
In terminal from root of repository

- To run all tests
  `$ go test`

- To run specific test naming AddMultiple
  `$ go run test Add`

- To run unit test showing logs when is not used `t.Errorf`
  `$ go test -v`

- To run benchmark tests that starts with keyword bench
 `$ go test -bench=.`

Exercise One

We have a func that have to be sum two numbers:

```
/*
Add two numbers
Return the sum of x and y
 */
func Add (x, y int ) int{
	return x + y
}
```

To test that you can import testing package:

```
package exerciseone

import "testing"

func Test_Add(t *testing.T){
	want := 5
	got := Add(2,3)

	//The result have to be five
	if got != want {
		t.Logf("ERROR: Se esperaba %d, se obtuvo %d", want, got)
		t.Fail()
	}
}
``
```
 Exercise two - Fibonacci benchmark

Method:
```
func TestFibonacciFor(t *testing.T) {
	want := 55
	got := fibonacciRecMem(10)
	if got != want {
		t.Errorf("Se esperaba %d, se obtuvo %d", want, got)
	}
}
```
Benchmark:
```
func BenchmarkFibonacciRecMem(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacciRecMem(30)
	}
}
```
