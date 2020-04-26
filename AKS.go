package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var n int64 = 1009

	fmt.Println(aks(int64(n)))
}

func aks(n int64) string {
	if perfectPower(n) == true { 					 //step 1
		return ("Composite")
	}

	r := findR(n)                                                    //step 2
	for a := 2; float64(a) < math.Min(float64(r), float64(n)); a++ { //step 3
		if GCD(a, int(n)) > 1 {
			return "Composite"
		}
	}
	if n <= r { //step 4
		return "Prime"
	}

	var x = []*big.Int{} 						 //step 5
	var a float64 = 1
	for a < math.Floor(math.Pow(eulerPhi(r), 1/2)*math.Log2(float64(n))) {
		x = fastPoly([]*big.Int{big.NewInt(1), big.NewInt(1)}, n, r)
		if notZero(x, n) {
			return "Coprime"
		}
		a++
	}
	return "Prime"							 //step 6

}

func notZero(a []*big.Int, n int64) bool {
	for i := 0; i < len(a); i++ {
		if (a[i].Mod(a[i], big.NewInt(n))).Int64() != 0 {
			return true
		}
	}
	return false
}

func perfectPower(n int64) bool {
	var i float64 = 2
	for i < math.Log2(float64(n))+1 {
		var a float64
		a = math.Pow(float64(n), float64(1/i))
		if a == float64(int64(a)) {
			return true
		}
		i++
	}
	return false
}

func findR(n int64) int64 {
	maxK := math.Pow(math.Log2(float64(n)), 2)
	nextR := true
	var r int64 = 1

	for nextR == true {
		r++
		nextR = false
		var k int = 0
		for float64(k) <= maxK && nextR == false {
			k++
			if fastMod(n, k, r) == 0 || fastMod(n, k, r) == 1 {
				nextR = true
			}
		}
	}
	return r
}

func fastMod(base int64, power int, n int64) int {
	var r int = 1
	for power > 0 {
		if int(power)%2 == 1 {
			r = int(r) * int(base) % int(n)
		}
		base = int64(int(math.Pow(float64(base), 2)) % int(n))
		power = power / 2
	}
	return r
}

func fastPoly(base []*big.Int, power int64, r int64) []*big.Int {
	var x = make([]*big.Int, int(r))
	for i := range x {
		x[i] = big.NewInt(0)
	}

	a := base[0]

	x[0] = big.NewInt(1)
	n := power

	for power > 0 {
		if int(power)%2 == 1 {
			x = multi(x, base, n, r)
		}
		base = multi(base, base, n, r)
		power = power / 2
	}

	x[0].Sub(x[0], a)
	x[n%int64(r)].Sub(x[n%int64(r)], big.NewInt(1))

	return x
}

func multi(a []*big.Int, b []*big.Int, n int64, r int64) []*big.Int {
	var temp = big.NewInt(0)
	var x = make([]*big.Int, int(r))
	for i := range x {
		x[i] = big.NewInt(0)
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if (i + j) < int(r) {
				temp.Mul(a[i], b[j])
				x[i+j].Add(x[i+j], temp)
			} else {
				temp.Mul(a[i], b[j])
				x[(i+j)%(int(r))].Add(x[(i+j)%(int(r))], temp)
			}
		}
	}
	return x
}

func eulerPhi(r int64) float64 {
	var x float64 = 0
	for i := 0; i < int(r)+1; i++ {
		if GCD(int(r), i) == 1 {
			x++
		}
	}
	return x
}

// GCD : greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
