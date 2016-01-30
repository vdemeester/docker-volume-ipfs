package main

import (
	"fmt"

	"github.com/docker/go-plugins-helpers/volume"
)

type ipfsDriver struct {
}

func newIPFSDriver() ipfsDriver {
	d := ipfsDriver{}
	return d
}

func (d ipfsDriver) Create(r volume.Request) volume.Response {
	fmt.Printf("Create %v\n", r)
	return volume.Response{}
}

func (d ipfsDriver) Remove(r volume.Request) volume.Response {
	fmt.Printf("Remove %v\n", r)
	return volume.Response{}
}

func (d ipfsDriver) Path(r volume.Request) volume.Response {
	fmt.Printf("Path %v\n", r)
	return volume.Response{}
}

func (d ipfsDriver) Mount(r volume.Request) volume.Response {
	fmt.Printf("Mount %v\n", r)
	return volume.Response{}
}

func (d ipfsDriver) Unmount(r volume.Request) volume.Response {
	fmt.Printf("Unmount %v\n", r)
	return volume.Response{}
}
