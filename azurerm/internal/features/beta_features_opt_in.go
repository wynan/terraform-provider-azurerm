package features

import (
	"os"
	"strings"
)

// VMSSExtensionsBeta returns whether or not the beta for VMSS Extensions for Linux and Windows VMSS resources is
// enabled.
//
// Set the Environment Variable `ARM_PROVIDER_VMSS_EXTENSIONS_BETA` to `true`
func VMSSExtensionsBeta() bool {
	return strings.EqualFold(os.Getenv("ARM_PROVIDER_VMSS_EXTENSIONS_BETA"), "true")
}

// VMDataDiskBeta returns whether or not the beta for VM Data Disks for Linux and Windows VM resources is
// enabled.
//
// Set the Environment Variable `ARM_PROVIDER_VM_DATADISKS_BETA` to `true`
func VMDataDiskBeta() bool {
	return strings.EqualFold(os.Getenv("ARM_PROVIDER_VM_DATADISKS_BETA"), "true")
}
