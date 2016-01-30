package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/docker/go-plugins-helpers/volume"
)

const ipfsID = "_ipfs"

var (
	defaultDir = filepath.Join(volume.DefaultDockerRootDirectory, ipfsID)
	// ipfs fuse mountpoint
	ipfsMountPoint = flag.String("mount", "/ipfs", "ipfs mount point")
)

func main() {
	// var Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage %s [options]\n", os.Args[0])
	// 	flag.PrintDefaults()
	// }

	flag.Parse()

	d := newIPFSDriver()
	h := volume.NewHandler(d)
	fmt.Println(h.ServeUnix("root", "ipfs"))
}
