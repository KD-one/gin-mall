package util

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	root, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(root)
	fmt.Println(time.Now().Unix() - 1700000000)
}
