# poc-golang-testing
Prueba de concepto Golang pruebas unitarias con paquete Test

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
````
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
