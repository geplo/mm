package main

import (
	"fmt"
	"math"
)

type operator struct {
	kind string
	node *node
}

func (op *operator) dump(lvl int) string {
	if op == nil {
		return ""
	}
	return " " + op.kind + " " + op.node.dump(lvl)
}

// a leaf as a value but no subnode.
type node struct {
	value   complex128
	subnode *node
	op      *operator
}

// Scalar .
func Scalar(value complex128) *node {
	return &node{value: value}
}

// ScalarSqrt .
func ScalarSqrt(value complex128) *node {
	return Scalar(value).Sqrt()
}

func isReal(n complex128) bool {
	return complex(real(n), 0) == n
}

func getReal(n complex128) float64 {
	return real(n)
}

func getImaginary(n complex128) float64 {
	// Remove the real part if any.
	n = (n - complex(real(n), 0))

	imm := math.Sqrt(-1 * real(n*n))
	if imm != 0 && complex(0, imm) == n {
		return imm
	}
	return -imm
}

func (n *node) Plus(n1 *node) *node {
	return &node{subnode: n, op: &operator{kind: "+", node: n1}}
}

func (n *node) Minus(n1 *node) *node {
	return &node{subnode: n, op: &operator{kind: "-", node: n1}}
}

func (n *node) Mul(n1 *node) *node {
	return &node{subnode: n, op: &operator{kind: "*", node: n1}}
}

func (n *node) Div(n1 *node) *node {
	return &node{subnode: n, op: &operator{kind: "/", node: n1}}
}

func (n *node) Sqrt() *node {
	return &node{subnode: n, op: &operator{kind: "root", node: &node{value: 2}}}
}

func (n *node) Square() *node {
	return &node{subnode: n, op: &operator{kind: "^", node: &node{value: 2}}}
}

func (n node) dump(lvl int) string {
	var ret string
	if n.subnode == nil {
		ret += fmt.Sprintf("%v", n.value)
	} else {
		ret += n.subnode.dump(lvl + 1)
	}
	if n.op != nil {
		ret += n.op.dump(lvl + 1)
	}
	return "(" + ret + ")"
}

func (n *node) String() string {
	if n == nil {
		return ""
	}
	return n.dump(0)
}

func main() {
	//   1 + √2
	// ( ────── )²
	//     2
	c := (1 + 0i) + 1.5 - 0i - 2.2i
	c = 5e-56 - 1e-55 + -0.0000000000001i
	imm := getImaginary(c)
	op := '+'
	if imm < 0 {
		op = '-'
		imm *= -1
	}
	fmt.Printf("c = (%g%c%gi)\n", getReal(c), op, imm)
	fmt.Printf("c = %v\n", c)

	imm = math.Sqrt(-1 * real(c*c))
	fmt.Printf("%v\n", imm)
	fmt.Printf("%v\n", complex(0, imm) == c)
	fmt.Printf("--> %v\n", c)
	fmt.Printf("--> %v\n", real(c))
	fmt.Printf("--> %v\n", (c - complex(real(c), 0)))
	fmt.Printf("--> %v\n", complex(real(c), 0) == c)
	operation := (Scalar(1).Plus(ScalarSqrt(2)).Div(Scalar(2))).Square()
	fmt.Printf("%s\n", operation)
}
