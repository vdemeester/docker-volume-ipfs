package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/docker/go-plugins-helpers/volume"
	"github.com/vdemeester/docker-volume-ipfs/ipfs"
)

const ipfsID = "_ipfs"

var (
	defaultDir = filepath.Join(volume.DefaultDockerRootDirectory, ipfsID)
	// ipfs configuration
	flDaemonConfig = flag.String("ipfs-config", "", "ipfs configuration to use when starting the daemon and to read mountpoint")
	// start daemon (or not)
	flDaemon = flag.Bool("daemon", true, "start daemon with the volume plugin")
	// init when starting daemon
	flDaemonInit = flag.Bool("daemon-init", true, "init ipfs if not already initialized")
)

func validatePath(path, format string) {
	_, err := os.Lstat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, format, err, path)
		os.Exit(1)
	}
}

func main() {
	var daemon *ipfs.Daemon
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill)

	flag.Parse()

	if *flDaemon {
		daemon = ipfs.NewDaemon(*flDaemonInit, *flDaemonConfig)
		if err := daemon.Setup(); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		go func() {
			daemon.Start()
		}()
	}

	ipfsMountPoint, err := ipfs.ReadConfig(*flDaemonConfig, "Mounts.IPFS")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ipnsMountPoint, err := ipfs.ReadConfig(*flDaemonConfig, "Mounts.IPNS")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func() {
		for sig := range sigs {
			if daemon != nil {
				daemon.Stop(sig)
			}
		}
	}()

	validatePath(ipfsMountPoint, "%v\n%s does not exists, can't start..\n Please use ipfs command line to mount it\n")
	validatePath(ipnsMountPoint, "%v\n%s does not exists, can't start..\n Please use ipfs command line to mount it\n")

	d := newIPFSDriver(ipfsMountPoint, ipnsMountPoint)
	h := volume.NewHandler(d)
	if err := h.ServeUnix("root", "ipfs"); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
