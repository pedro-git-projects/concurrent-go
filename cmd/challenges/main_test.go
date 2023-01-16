package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestUptadeMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("lambda", &wg)
	wg.Wait()

	if msg != "lambda" {
		t.Errorf("Expected %s, but it is not there", msg)
	}
}

func TestPrintMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "lambda"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "lambda") {
		t.Errorf("Expected %s, but it is not there", msg)
	}
}

func TestMain(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected Hello, universe! but it is not there")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected Hello, cosmos! but it is not there")
	}

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected Hello, universe! but it is not there")
	}
}
