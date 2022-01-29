package testLearning

import (
	"fmt"
	"rebuildServer/utils"
	"testing"
)

func TestPathExists(t *testing.T) {
	fmt.Println(utils.PathExists("testdir"))
	fmt.Println(utils.CreateDir("name", "css"))
}
