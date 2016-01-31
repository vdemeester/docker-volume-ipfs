package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
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
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill)

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
	go func() {
		if err := h.ServeUnix("root", "ipfs"); err != nil {
			fmt.Println(err)
		}
	}()
	cmd := startIPFSDaemon()
	cmd.Wait()
}

func startIPFSDaemon() *exec.Cmd {
	cmd := exec.Command("ipfs", "daemon", "--mount")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for IPFS", err)
		os.Exit(1)
	}

	scannerOut := bufio.NewScanner(stdout)
	go func() {
		for scannerOut.Scan() {
			fmt.Printf("IPFS > %s\n", scannerOut.Text())
		}
	}()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StderrPipe for IPFS", err)
		os.Exit(1)
	}

	scannerErr := bufio.NewScanner(stderr)
	go func() {
		for scannerErr.Scan() {
			fmt.Printf("IPFS > %s\n", scannerErr.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting IPFS", err)
		os.Exit(1)
	}

	return cmd
}
