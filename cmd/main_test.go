package main

import (
	"os"
	"os/exec"
	"testing"
	"time"
)

// TestMainProcess tests if the main function runs and starts without immediate errors
func TestMainProcess(t *testing.T) {
	// If the environment variable is set to 1, execute the actual main function
	if os.Getenv("RUN_MAIN") == "1" {
		main()
		return
	}

	// Create a subprocess that runs this specific test function
	cmd := exec.Command(os.Args[0], "-test.run=TestMainProcess")
	cmd.Env = append(os.Environ(), "RUN_MAIN=1")

	// Start the subprocess (does not wait for completion since the server runs indefinitely)
	err := cmd.Start()
	if err != nil {
		t.Fatalf("Failed to start the process: %v", err)
	}

	// Give the server some time to start up
	time.Sleep(500 * time.Millisecond)

	// Monitor if the process terminates prematurely with an error
	processDone := make(chan error, 1)
	go func() {
		processDone <- cmd.Wait()
	}()

	select {
	case err := <-processDone:
		if err != nil {
			t.Fatalf("The server terminated unexpectedly with error: %v", err)
		}
	case <-time.After(1 * time.Second):
		// The server is still running successfully, kill the process to clean up
		err := cmd.Process.Kill()
		if err != nil {
			t.Fatalf("Failed to terminate the server process: %v", err)
		}
	}
}
