Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fscvmm%2Farmscvmm%2Fv0.1.0/sdk/resourcemanager/scvmm/armscvmm/README.md) on how to add the SDK to your project and authenticate.

```go
package armscvmm_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/UpdateVirtualMachine.json
func ExampleVirtualMachinesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armscvmm.NewVirtualMachinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<virtual-machine-name>",
		armscvmm.VirtualMachineUpdate{
			Properties: &armscvmm.VirtualMachineUpdateProperties{
				HardwareProfile: &armscvmm.HardwareProfileUpdate{
					CPUCount: to.Ptr[int32](4),
					MemoryMB: to.Ptr[int32](4096),
				},
				NetworkProfile: &armscvmm.NetworkProfileUpdate{
					NetworkInterfaces: []*armscvmm.NetworkInterfacesUpdate{
						{
							Name:            to.Ptr("<name>"),
							IPv4AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
							IPv6AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
							MacAddressType:  to.Ptr(armscvmm.AllocationMethodStatic),
						}},
				},
				StorageProfile: &armscvmm.StorageProfileUpdate{
					Disks: []*armscvmm.VirtualDiskUpdate{
						{
							Name:       to.Ptr("<name>"),
							DiskSizeGB: to.Ptr[int32](10),
						}},
				},
			},
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		&armscvmm.VirtualMachinesClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
