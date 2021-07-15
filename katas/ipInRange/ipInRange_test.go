package main


import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIpInRange(testing *testing.T){
	var valuesToTest = []struct {
		x, y, result string
	}{
		{"192.168.4.27", "192.168.0.0/16", true},// exercise statement example 1
		{"192.168.4.27", "192.168.1.0/24", false},// exercise statement example 2
	}
	for i := 0; i < len(valuesToTest); i++ {
		assert.Equal(testing, valuesToTest[i].result, ipInRange(valuesToTest[i].x, valuesToTest[i].y), "should be equal")
	}
}