package strs

import (
	"fmt"
	"testing"
)

func TestSnakeToCamel(t *testing.T) {
	fmt.Println(SnakeToCamel("ID"))
	fmt.Println(SnakeToCamel("USERNAME"))
	fmt.Println(SnakeToCamel("CREATE_TIME"))
}
