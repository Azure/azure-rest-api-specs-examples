Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvmwarecloudsimple%2Farmvmwarecloudsimple%2Fv0.1.0/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple/README.md) on how to add the SDK to your project and authenticate.

```go
package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/CreateVirtualMachine.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvmwarecloudsimple.NewVirtualMachinesClient("<subscription-id>",
		"<referer>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-machine-name>",
		armvmwarecloudsimple.VirtualMachine{
			Location: to.StringPtr("<location>"),
			Properties: &armvmwarecloudsimple.VirtualMachineProperties{
				AmountOfRAM: to.Int32Ptr(4096),
				Disks: []*armvmwarecloudsimple.VirtualDisk{
					{
						ControllerID:     to.StringPtr("<controller-id>"),
						IndependenceMode: armvmwarecloudsimple.DiskIndependenceModePersistent.ToPtr(),
						TotalSize:        to.Int32Ptr(10485760),
						VirtualDiskID:    to.StringPtr("<virtual-disk-id>"),
					}},
				Nics: []*armvmwarecloudsimple.VirtualNic{
					{
						Network: &armvmwarecloudsimple.VirtualNetwork{
							ID: to.StringPtr("<id>"),
						},
						NicType:      armvmwarecloudsimple.NICTypeE1000.ToPtr(),
						PowerOnBoot:  to.BoolPtr(true),
						VirtualNicID: to.StringPtr("<virtual-nic-id>"),
					}},
				NumberOfCores:  to.Int32Ptr(2),
				PrivateCloudID: to.StringPtr("<private-cloud-id>"),
				ResourcePool: &armvmwarecloudsimple.ResourcePool{
					ID: to.StringPtr("<id>"),
				},
				TemplateID: to.StringPtr("<template-id>"),
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
	log.Printf("VirtualMachine.ID: %s\n", *res.ID)
}
```
