# 🐳 docker-volume-ipfs [![Build Status](https://travis-ci.org/vdemeester/docker-volume-ipfs.svg?branch=master)](https://travis-ci.org/vdemeester/docker-volume-ipfs)

This is an open source volume plugin that allows using an
[ipfs](https://ipfs.io/) filesystem as a volume.

```bash
$ docker-volume-ipfs &
$ docker run -it --rm -v QmPXME1oRtoT627YKaDPDQ3PwA8tdP9rWuAAweLzqSwAWT/readme:/data --volume-driver=ipfs busybox cat data
Hello and Welcome to IPFS!

██╗██████╗ ███████╗███████╗
██║██╔══██╗██╔════╝██╔════╝
██║██████╔╝█████╗  ███████╗
██║██╔═══╝ ██╔══╝  ╚════██║
██║██║     ██║     ███████║
╚═╝╚═╝     ╚═╝     ╚══════╝

If you're seeing this, you have successfully installed
IPFS and are now interfacing with the ipfs merkledag!

 -------------------------------------------------------
| Warning:                                              |
|   This is alpha software. use at your own discretion! |
|   Much is missing or lacking polish. There are bugs.  |
|   Not yet secure. Read the security notes for more.   |
 -------------------------------------------------------

Check out some of the other files in this directory:

  ./about
  ./help
  ./quick-start     <-- usage examples
  ./readme          <-- this file
  ./security-notes
```

## Goals

The main goal is to be able to create docker volumes that are backed by the IPFS filesystem. They could be named volumes,
anonymous volumes. The plugin should also work in a cluster / swarm environment (but by its nature it should be built-in).

First, let's define a "format" for creating volumes. There is two cases: `docker volume create` and `docker run -v`.

For `docker volume create`, it could be an option `ipfs=`, `ipns=`...

```bash
$ docker volume create --driver=ipfs --opt ipfs=QmPXME1oRtoT627YKaDPDQ3PwA8tdP9rWuAAweLzqSwAWT --name ipfs-welcome
$ docker run -v ipfs-welcome:/ipfs/welcome busybox cat /ipfs/welcome/readme
Hello and Welcome to IPFS!

██╗██████╗ ███████╗███████╗
██║██╔══██╗██╔════╝██╔════╝
██║██████╔╝█████╗  ███████╗
██║██╔═══╝ ██╔══╝  ╚════██║
██║██║     ██║     ███████║
╚═╝╚═╝     ╚═╝     ╚══════╝

If you're seeing this, you have successfully installed
# […]
# would also work with ipns
$ docker volume create --driver=ipfs --opt ipns=ipfs.io --name ipfs-io-website
$ docker run -v ipfs-io-website:/var/www/html nginx
```

For `-v` directly in `docker run` (a.k.a. anonymous volume), we could namespace it somehow.

```
# Default to ipfs
$ docker run -it --rm -v QmPXME1oRtoT627YKaDPDQ3PwA8tdP9rWuAAweLzqSwAWT/readme:/data --volume-driver=ipfs busybox cat data
Hello and Welcome to IPFS!
# […]
$ docker run -it --rm -v $(docker volume create --driver=ipfs --opt ipns=ipfs.io):/var/www/html --volume-driver=ipfs nginx
```

## TODO(s)

- [ ] Use `ipfs daemon` API instead of making the assumption that a daemon is running and mounting ipfs using fuse
- [ ] Support ipfs mounts (via fuse)
- [ ] Support ipns mounts
- [ ] Create a volume v2 plugin
