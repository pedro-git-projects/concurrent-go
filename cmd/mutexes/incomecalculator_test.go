package main

import (
	"strings"
	"testing"

	"github.com/pedro-git-projects/concurrent-go/pkg"
)

func TestIncomeCalculator(t *testing.T) {
	result := pkg.PipeStdout(calculateYearlyIncome)
	if !strings.Contains(result, "$34320.00") {
		t.Errorf("Expected $34320.00, got %s", result)
	}
}
