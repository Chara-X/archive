package tar

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func ExampleReader() {
	os.WriteFile("log.txt", []byte("Hello, World!"), 0777)
	defer os.Remove("log.txt")
	exec.Command("tar", "-cf", "arch.tar", "--format", "v7", "log.txt").Run()
	defer os.Remove("arch.tar")
	var arch, _ = os.Open("arch.tar")
	defer arch.Close()
	var tr = NewReader(arch)
	for {
		var hdr, err = tr.Next()
		if err != nil {
			break
		}
		fmt.Println(hdr.Name)
		io.CopyN(os.Stdout, tr, hdr.Size)
	}
	// Output:
	// log.txt
	// Hello, World!
}
