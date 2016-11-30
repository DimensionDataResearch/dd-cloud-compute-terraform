package models

import (
	"github.com/DimensionDataResearch/dd-cloud-compute-terraform/maps"
	"github.com/DimensionDataResearch/go-dd-cloud-compute/compute"
)

// DiskAction represents an action to be taken for a Disk.
type DiskAction int

const (
	// DiskActionNone indicates that no action is to be taken.
	DiskActionNone DiskAction = iota

	// DiskActionCreate indicates that the Disk is to be created.
	DiskActionCreate

	// DiskActionUpdate indicates that the Disk is to be updated.
	DiskActionUpdate

	// DiskActionDelete indicates that the Disk is to be deleted.
	DiskActionDelete
)

// Disk represents the Terraform configuration for a ddcloud_server disk.
type Disk struct {
	ID         *string
	SCSIUnitID int
	SizeGB     int
	Speed      string
	Action     DiskAction
}

// ReadMap populates the Disk with values from the specified map.
func (disk *Disk) ReadMap(diskProperties map[string]interface{}) {
	reader := maps.NewReader(diskProperties)

	disk.ID = reader.GetStringPtr("id")
	disk.SCSIUnitID = reader.GetInt("scsi_unit_id")
	disk.SizeGB = reader.GetInt("size_gb")
	disk.Speed = reader.GetString("speed")
}

// ToMap creates a new map using the values from the Disk.
func (disk *Disk) ToMap() map[string]interface{} {
	data := make(map[string]interface{})
	disk.UpdateMap(data)

	return data
}

// UpdateMap updates a map using values from the Disk.
func (disk *Disk) UpdateMap(diskProperties map[string]interface{}) {
	writer := maps.NewWriter(diskProperties)

	writer.SetStringPtr("id", disk.ID)
	writer.SetInt("scsi_unit_id", disk.SCSIUnitID)
	writer.SetInt("size_gb", disk.SizeGB)
	writer.SetString("speed", disk.Speed)
}

// ReadVirtualMachineDisk populates the Disk with values from the specified VirtualMachineDisk.
func (disk *Disk) ReadVirtualMachineDisk(virtualMachineDisk compute.VirtualMachineDisk) {
	disk.ID = virtualMachineDisk.ID
	disk.SCSIUnitID = virtualMachineDisk.SCSIUnitID
	disk.SizeGB = virtualMachineDisk.SizeGB
	disk.Speed = virtualMachineDisk.Speed
}

// ToVirtualMachineDisk updates a map using values from the Disk.
func (disk *Disk) ToVirtualMachineDisk() compute.VirtualMachineDisk {
	virtualMachineDisk := compute.VirtualMachineDisk{}
	disk.UpdateVirtualMachineDisk(&virtualMachineDisk)

	return virtualMachineDisk
}

// UpdateVirtualMachineDisk updates a CloudControl VirtualMachineDisk using values from the Disk.
func (disk *Disk) UpdateVirtualMachineDisk(virtualMachineDisk *compute.VirtualMachineDisk) {
	virtualMachineDisk.ID = disk.ID
	virtualMachineDisk.SCSIUnitID = disk.SCSIUnitID
	virtualMachineDisk.SizeGB = disk.SizeGB
	virtualMachineDisk.Speed = disk.Speed
}

// NewDiskFromMap creates a Disk from the values in the specified map.
func NewDiskFromMap(diskProperties map[string]interface{}) Disk {
	disk := Disk{}
	disk.ReadMap(diskProperties)

	return disk
}

// NewDiskFromVirtualMachineDisk creates a Disk from the values in the specified CloudControl VirtualMachineDisk.
func NewDiskFromVirtualMachineDisk(virtualMachineDisk compute.VirtualMachineDisk) Disk {
	disk := Disk{}
	disk.ReadVirtualMachineDisk(virtualMachineDisk)

	return disk
}

// NewDisksFromStateData creates Disks from an array of Terraform state data.
//
// The values in the diskPropertyList are expected to be map[string]interface{}.
func NewDisksFromStateData(diskPropertyList []interface{}) Disks {
	disks := make(Disks, len(diskPropertyList))
	for index, data := range diskPropertyList {
		diskProperties := data.(map[string]interface{})
		disks[index] = NewDiskFromMap(diskProperties)
	}

	return disks
}

// NewDisksFromMaps creates Disks from an array of Terraform value maps.
func NewDisksFromMaps(diskPropertyList []map[string]interface{}) Disks {
	disks := make(Disks, len(diskPropertyList))
	for index, data := range diskPropertyList {
		disks[index] = NewDiskFromMap(data)
	}

	return disks
}

// NewDisksFromVirtualMachineDisks creates Disks from an array of compute.VirtualMachineDisk.
func NewDisksFromVirtualMachineDisks(virtualMachineDisks []compute.VirtualMachineDisk) Disks {
	disks := make(Disks, len(virtualMachineDisks))
	for index, virtualMachineDisk := range virtualMachineDisks {
		disks[index] = NewDiskFromVirtualMachineDisk(virtualMachineDisk)
	}

	return disks
}
