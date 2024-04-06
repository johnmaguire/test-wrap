package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	reader, writer := io.Pipe()

	cmdCtx, cmdDone := context.WithCancel(context.Background())

	scannerStopped := make(chan struct{})
	go func() {
		defer close(scannerStopped)

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			fmt.Println(scanner.Bytes())
		}
	}()

	cmd := exec.Command("./test-command.exe")
	cmd.Stdout = writer
	_ = cmd.Start()
	go func() {
		_ = cmd.Wait()
		cmdDone()
		writer.Close()
	}()
	<-cmdCtx.Done()

	<-scannerStopped
}
