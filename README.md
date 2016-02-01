# docker-volume-ipfs [![Build Status](https://travis-ci.org/vdemeester/docker-volume-ipfs.svg?branch=master)](https://travis-ci.org/vdemeester/docker-volume-ipfs)

This is an open source volume plugin that allows using an
[ipfs](https://ipfs.io/) filesystem as a volume.

```bash
docker run -it --rm -v QmPXME1oRtoT627YKaDPDQ3PwA8tdP9rWuAAweLzqSwAWT/readme:/data --volume-driver=ipfs busybox cat data
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
