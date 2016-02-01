package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/docker/go-plugins-helpers/volume"
)

type ipfsDriver struct {
	mountPoint string
	volumes    map[string]string
	m          *sync.Mutex
}

func newIPFSDriver(ipfsMountPoint string) ipfsDriver {
	d := ipfsDriver{
		mountPoint: ipfsMountPoint,
		volumes:    make(map[string]string),
		m:          &sync.Mutex{},
	}
	return d
}

func (d ipfsDriver) Create(r volume.Request) volume.Response {
	fmt.Printf("Create %v\n", r)

	d.m.Lock()
	defer d.m.Unlock()

	volumeName := r.Name

	if _, ok := d.volumes[volumeName]; ok {
		return volume.Response{}
	}

	volumePath := filepath.Join(d.mountPoint, volumeName)

	_, err := os.Lstat(volumePath)
	if err != nil {
		fmt.Println("Error", volumePath, err.Error())
		return volume.Response{Err: fmt.Sprintf("Error while looking for volumePath %s: %s", volumePath, err.Error())}
	}

	d.volumes[volumeName] = volumePath

	return volume.Response{}
}

func (d ipfsDriver) Path(r volume.Request) volume.Response {
	fmt.Printf("Path %v\n", r)
	fmt.Printf("%v", d.volumes)
	volumeName := r.Name

	if volumePath, ok := d.volumes[volumeName]; ok {
		return volume.Response{Mountpoint: volumePath}
	}

	return volume.Response{}
}

func (d ipfsDriver) Remove(r volume.Request) volume.Response {
	fmt.Printf("Remove %v", r)

	d.m.Lock()
	defer d.m.Unlock()

	volumeName := r.Name

	if _, ok := d.volumes[volumeName]; ok {
		delete(d.volumes, volumeName)
	}

	return volume.Response{}
}

func (d ipfsDriver) Mount(r volume.Request) volume.Response {
	fmt.Printf("Mount %v\n", r)
	volumeName := r.Name

	if volumePath, ok := d.volumes[volumeName]; ok {
		return volume.Response{Mountpoint: volumePath}
	}

	return volume.Response{}
}

func (d ipfsDriver) Unmount(r volume.Request) volume.Response {
	fmt.Printf("Unmount %v: nothing to do\n", r)
	return volume.Response{}
}

func (d ipfsDriver) Get(r volume.Request) volume.Response {
	fmt.Printf("Get %v\n", r)
	volumeName := r.Name

	if volumePath, ok := d.volumes[volumeName]; ok {
		return volume.Response{Volume: &volume.Volume{Name: volumeName, Mountpoint: volumePath}}
	}

	return volume.Response{Err: fmt.Sprintf("volume %s does not exists", volumeName)}
}

func (d ipfsDriver) List(r volume.Request) volume.Response {
	fmt.Printf("List %v\n", r)

	volumes := []*volume.Volume{}

	for name, path := range d.volumes {
		volumes = append(volumes, &volume.Volume{Name: name, Mountpoint: path})
	}

	return volume.Response{Volumes: volumes}
}
