package main

import (
	"fmt"
	"log"
	"os/exec"
	"bytes"
	"io"
	"strings"
)

func main() {
	cmd := exec.Command("ls", "-al")
	stdout, err := cmd.StdoutPipe()

	outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, stdout)
        outC <- buf.String()
    }()

	
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	
	for i := range outC {
		if strings.Contains(i, "oryx") {
			fmt.Printf(i)
			cmd.Wait()
		}
	}



	fmt.Printf("daje")
}