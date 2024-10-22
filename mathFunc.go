package main

import "fmt"

type Point struct {
	x, y float32
}

type Polynomy struct {
	coefficient []float32
	bias        float32
}

// a + bx + cx2 + dx3 + ... zxn

func (p *Polynomy) f_p(x float32) float32 {
	result := float32(0)
	power := float32(x)

	for _, c := range p.coefficient {
		result += power * c
		power *= x
	}
	return result + p.bias
}

func CreateArray(size int, bias float32) []float32 {
	a := make([]float32, size)
	for i := 0; i < size; i++ {
		a[i] = float32(i)
	}
	return a
}

func Map(slice []float32, fn func(float32) float32, relativePos int) []float32 {
	result := make([]float32, len(slice))
	lenS := float32(relativePos)
	debug1 := 0

	for i, v := range slice {
		result[i] = fn(v - lenS)

		if (int(v-lenS)%10) == 0 && debug1 == 0 {
			fmt.Printf("x: %2.f | y: %2.f\n", v-lenS, result[i])
		}
	}

	return result
}

// slope change --> movements possible: clockwise, counterclockwise --> change the func
// lets use the other form of describing an N space --> x + y + z = 0, therefore
// if x,z = 0, y=0 as well
/*func rotation(axis []float32, rotation int) int {
	return rotation % 180
}*/
