package tar

import (
	"archive/tar"
	"fmt"
	"os"
	"os/exec"
)

func ExampleWriter() {
	var arch, _ = os.Create("arch.tar")
	defer os.RemoveAll("arch.tar")
	defer arch.Close()
	var tw = NewWriter(arch)
	var hdr = &tar.Header{Name: "log.txt", Mode: 0777, Size: 13}
	tw.WriteHeader(hdr)
	tw.Write([]byte("Hello, World!"))
	tw.Close()
	os.MkdirAll("arch", 0777)
	defer os.RemoveAll("arch")
	exec.Command("tar", "-xf", "arch.tar", "--format", "v7", "-C", "arch").Run()
	var data, _ = os.ReadFile("arch/log.txt")
	fmt.Println(string(data))
	// Output:
	// Hello, World!
}
