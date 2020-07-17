package global_id

import (
	"fmt"
	"testing"
)

func TestGenerate (t *testing.T) {
	v := Generate()
	fmt.Println("global id =>", v)
}