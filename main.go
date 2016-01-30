package main

import (
	"flag"
	"fmt"
	"os"
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

	_, err := os.Lstat(*ipfsMountPoint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n%s does not exists, can't start..\n Please use ipfs command line to mount it\n", err, *ipfsMountPoint)
		os.Exit(1)
	}

	d := newIPFSDriver(*ipfsMountPoint)
	h := volume.NewHandler(d)
	fmt.Println(h.ServeUnix("root", "ipfs"))
}
