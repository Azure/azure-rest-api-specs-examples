Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.5.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/UpdateVMDetachDataDiskUsingToBeDetachedProperty.json
func ExampleVirtualMachinesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachinesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-name>",
		armcompute.VirtualMachineUpdate{
			Properties: &armcompute.VirtualMachineProperties{
				HardwareProfile: &armcompute.HardwareProfile{
					VMSize: armcompute.VirtualMachineSizeTypes("Standard_D2_v2").ToPtr(),
				},
				NetworkProfile: &armcompute.NetworkProfile{
					NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
						{
							ID: to.StringPtr("<id>"),
							Properties: &armcompute.NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &armcompute.OSProfile{
					AdminPassword: to.StringPtr("<admin-password>"),
					AdminUsername: to.StringPtr("<admin-username>"),
					ComputerName:  to.StringPtr("<computer-name>"),
				},
				StorageProfile: &armcompute.StorageProfile{
					DataDisks: []*armcompute.DataDisk{
						{
							CreateOption: armcompute.DiskCreateOptionTypes("Empty").ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(0),
							ToBeDetached: to.BoolPtr(true),
						},
						{
							CreateOption: armcompute.DiskCreateOptionTypes("Empty").ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(1),
							ToBeDetached: to.BoolPtr(false),
						}},
					ImageReference: &armcompute.ImageReference{
						Offer:     to.StringPtr("<offer>"),
						Publisher: to.StringPtr("<publisher>"),
						SKU:       to.StringPtr("<sku>"),
						Version:   to.StringPtr("<version>"),
					},
					OSDisk: &armcompute.OSDisk{
						Name:         to.StringPtr("<name>"),
						Caching:      armcompute.CachingTypesReadWrite.ToPtr(),
						CreateOption: armcompute.DiskCreateOptionTypes("FromImage").ToPtr(),
						ManagedDisk: &armcompute.ManagedDiskParameters{
							StorageAccountType: armcompute.StorageAccountTypes("Standard_LRS").ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachinesClientUpdateResult)
}
```
