package ipfs

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Daemon represent the ipfs daemon (to run) and options for it.
type Daemon struct {
	init bool
	path string
	cmd  *exec.Cmd
}

// NewDaemon returns a new instance of IPFSDaemon.
func NewDaemon(init bool, path string) *Daemon {
	return &Daemon{
		init: init,
		path: path,
	}
}

// Setup sets the daemon up, initializing ipfs if asked and making sure that
// fuse allow_other is true (otherwise it is not gonna work — permission denied).
func (d *Daemon) Setup() error {
	if d.init {
		cmd := prepareIPFSCommand(d.path, "init")
		out, err := cmd.CombinedOutput()
		if err != nil && !strings.Contains(string(out), "ipfs configuration file already exists!") {
			fmt.Fprintf(os.Stderr, "%s", out)
			return err
		}
	}
	// Make sure fuse config allow others
	// ipfs config --json Mounts.FuseAllowOther true
	cmd := prepareIPFSCommand(d.path, "config", "--json", "Mounts.FuseAllowOther", "true")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", out)
	}
	return err
}

// Start starts the IPFS daemon with mount option.
func (d *Daemon) Start() error {
	cmd := prepareIPFSCommand(d.path, "daemon", "--mount")
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

	d.cmd = cmd
	return nil
}

// Stop kills the daemon
// FIXME(vdemeester): It seems it's not that easy to kill ipfs
func (d *Daemon) Stop(sig os.Signal) error {
	if d.cmd != nil {
		// return d.cmd.Process.Signal(sig)
		return d.cmd.Process.Kill()
	}
	return nil
}

func prepareIPFSCommand(path string, args ...string) *exec.Cmd {
	cmd := exec.Command("ipfs", args...)
	if path != "" {
		cmd.Env = append(cmd.Env, fmt.Sprintf("IPFS_PATH=%s", path))
	}
	return cmd
}

// ReadConfig reads configuration key for the specified config path.
func ReadConfig(configPath, key string) (string, error) {
	cmd := prepareIPFSCommand(configPath, "config", key)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", out)
	}
	return string(out), err
}
