// package main

// import "fmt"

// type Vehicule struct {
// 	vehiculeBrand string
// 	model         string
// 	topSpeed      int
// }

// type Speed interface {
// 	howFast() string

// }

// func (v Vehicule) getTopSpeed() int {
// 	return v.topSpeed
// }

// // doesn't change the state of our vehicule
// func (v Vehicule) setTopSpeed(newSpeed int) {
// 	v.topSpeed = newSpeed
// }

// // doeas change the state of our vehicule
// // func (v *Vehicule) setTopSpeed2(newSpeed int) {
// // 	v.topSpeed = newSpeed
// // }

// func (v Vehicule) howFast() string {
// 	return "normal speed"
// }

// type Car struct {
// 	Vehicule
// }

// func (c Car) howFast() string {
// 	return "somehow fast speed"
// }

// type Truck struct {
// 	Vehicule
// }

// func (t Truck) howFast() string {
// 	return "too slow in terms of speed"
// }

// func printSpeed(s Speed) {
// 	fmt.Println(s.howFast())
// }

// func main() {

// 	volvo := Truck{
// 		Vehicule: Vehicule{
// 			vehiculeBrand: "volvo",
// 			model:         "ef34",
// 			topSpeed:      550,
// 		},
// 	}

// 	audi := Car{
// 		Vehicule: Vehicule{
// 			vehiculeBrand: "Audi",
// 			model:         "rs5",
// 			topSpeed:      320,
// 		},
// 	}

// 	fmt.Println(audi)

// 	fmt.Println("old speed: ", volvo.getTopSpeed())
// 	volvo.setTopSpeed(220)
// 	fmt.Println("new speed: ", volvo.getTopSpeed())

// 	printSpeed(volvo)
// 	printSpeed(audi)

// }

// package main

// import "fmt"

// type IPAddr [4]byte

// // TODO: Add a "String() string" method to IPAddr.

// func (ip IPAddr) String() string {
// 	var result string
// 	for _, elm := range ip {
// 		result = fmt.Sprintf("%v%v.", result, elm)
// 	}
// 	return result[:len(result)-1]
// }

// func main() {
// 	hosts := map[string]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	for name, ip := range hosts {
// 		fmt.Printf("%v: %v\n", name, ip)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type ErrNegativeSqrt float64

// func (e ErrNegativeSqrt) Error() string {
// 	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
// }

// func Sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, ErrNegativeSqrt(x)
// 	}
// 	return math.Sqrt(x), nil
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// 	fmt.Println(Sqrt(-2))
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	r := strings.NewReader("Hello, Reader!")

// 	b := make([]byte, 8)
// 	for {
// 		n, err := r.Read(b)
// 		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 		fmt.Printf("b[:n] = %q\n", b[:n])
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Read implements io.Reader.
func (r *rot13Reader) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
