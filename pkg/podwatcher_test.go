package pkg

import "testing"
import (
	"github.com/stretchr/testify/assert"
)

func TestGetListOfPods(t *testing.T) {

}

func TestDifference(t *testing.T) {
	var setA = []string{"10.10.12.10", "10.10.12.11", "10.10.12.13"}
	var setB = []string{"10.10.12.10", "10.10.12.11", "10.10.12.14"}

	expectedAB := []string{"10.10.12.13"}
	expectedBA := []string{"10.10.12.14"}

	diffAB := difference(setA, setB)
	diffBA := difference(setB, setA)
	diffAA := difference(setA, setA)
	diffBB := difference(setB, setB)

	assert.Equal(t, expectedAB, diffAB)
	assert.Equal(t, expectedBA, diffBA)
	assert.Equal(t, 0, len(diffAA))
	assert.Equal(t, 0, len(diffBB))
}

func TestWatch(t *testing.T) {

}