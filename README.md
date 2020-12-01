# poc-golang-testing

Proof of concept Golang testing unit with `testing package of Golang`

###Run
In terminal from root of repository

- To run all tests
  `$ go test`

- To run specific test naming AddMultiple
  `$ go run test Add`

- To run unit test showing logs `$ go test -v`

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
