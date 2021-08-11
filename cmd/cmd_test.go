package main

import (
	"fmt"
	"sysutils-go/internal/load"
	"testing"
)

func TestLoad(t *testing.T) {
	avgStat, _ := load.Avg()
	load15 := avgStat.Load15
	fmt.Println("load15: ", load15)
}

func TestCPU(t *testing.T) {

}
