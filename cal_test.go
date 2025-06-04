package cal

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
// cal_test.go
func TestAdd(t *testing.T) {
	assert := assert.New(t)
	result, err := Add(1,2)
	assert.Equal(result, 3)
	assert.NoError(err) 
}

