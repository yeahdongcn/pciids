package pciids

import (
	"github.com/yeahdongcn/pciids/pkg/file"
	"github.com/yeahdongcn/pciids/pkg/ids"
	"github.com/yeahdongcn/pciids/pkg/query"
)

// PCIID type.
type PCIID = ids.PCIID

// All returns all PCI IDs.
var All = query.All

// LatestFile returns the latest PCI IDs database file.
var LatestFile = file.Latest

// Parse will parse a PCI IDs file.
var Parse = ids.Parse

// QueryVendor will query using a specific PCI ID vendor.
var QueryVendor = query.Vendor

// QueryDevice will query using a specific PCI ID device pair.
var QueryDevice = query.Device

// QuerySubDevice will query using a specific PCI ID device and subdevice pair.
var QuerySubDevice = query.SubDevice
