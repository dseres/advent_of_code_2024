package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

func TestSolution1(t *testing.T) {
	a := assert.New(t)
	m := new(testInput, 6, 12)
	fmt.Println(m)
	a.Equal(22, m.findPath())
	a.Equal("6,1", m.addMoreBarriers())
}
