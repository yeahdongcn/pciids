# PCIIDs

*Lookup vendor and device names using PCI IDs!*

[![Build Status](https://travis-ci.org/powersj/pciids.svg?branch=master)](https://travis-ci.org/powersj/pciids/) [![Go Report Card](https://goreportcard.com/badge/github.com/powersj/pciids)](https://goreportcard.com/report/github.com/powersj/pciids) [![Go Reference](https://pkg.go.dev/badge/github.com/powersj/pciids.svg)](https://pkg.go.dev/github.com/powersj/pciids)

## CLI

The CLI expects either two or four PCI IDs to look up:

```text
$ pciids 1d0f efa1
1d0f:efa1           - Amazon.com, Inc. Elastic Fabric Adapter (EFA)
$ pciids 10de 2206 10de 1467
10de:2206 10de:1467 - NVIDIA Corporation GA102 [GeForce RTX 3080]
```

If there are multiple matches all will be listed.

### JSON output

The command can take the `--json` flag to produce JSON output:

```json
$ pciids 121a 0009 121a 0009 --json
[
    {
        "vendorID": "121a",
        "deviceID": "0009",
        "vendorName": "3Dfx Interactive, Inc.",
        "deviceName": "Voodoo 4 / Voodoo 5",
        "subVendorID": "121a",
        "subDeviceID": "0009",
        "subVendorName": "3Dfx Interactive, Inc.",
        "subDeviceName": "Voodoo5 AGP 5500/6000"
    }
]
```

## Install

Below outlines the various ways to obtain and install pciids.

### From binary

Download the [latest release](https://github.com/powersj/pciids/releases/latest)
of pciids for your platform and extract the tarball:

```shell
wget pciids<version>_<os>_<arch>.tar.gz
tar zxvf pciids<version>_<os>_<arch>.tar.gz
```

The tarball will extract the readme, license, and the pre-compiled binary.

### From source

To build and install directly from source run:

```shell
git clone https://github.com/powersj/pciids
cd pciids
make
```

The default make command will run `go build` and produce a binary in the root
directory.

### From go

To download using the `go get` command run:

```shell
go get github.com/powersj/pciids
```

The executable object file location will exist at `${GOPATH}/bin/pciids`