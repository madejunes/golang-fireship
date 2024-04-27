package main

import "fmt"

func main() {
	/* type-nya explisit ketik sendiri */
	const a int32 = 32 					// 32-bit intger
	const b float32 = 33.5 			// 32-bit float
	var c complex128 = 1 + 4i	// 12b-bit complex number
	var d uint = 14 						
	// TIL const boleh declared and not used, var dan := ga boleh declared and not used
	
	/* type-nya otomatis sesuai value */
	n := 42 		// int
	pi := 3.14	// float64
	x, y := true, false // bool
	z := "anu" // string

	fmt.Printf("Type ketik sendiri:\n %T %T %T %T\n", a, b, c, d)
	fmt.Printf("Type otomatis:\n %T %T %T %T %T\n", n, pi, x, y, z)
	// ooh ternyata %T ngeprint type

	fmt.Println("coba pake println", a, b)
}