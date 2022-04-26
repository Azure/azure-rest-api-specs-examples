Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvmwarecloudsimple%2Farmvmwarecloudsimple%2Fv0.4.0/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/CreateVirtualMachine.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armvmwarecloudsimple.NewVirtualMachinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<referer>",
		"<virtual-machine-name>",
		armvmwarecloudsimple.VirtualMachine{
			Location: to.Ptr("<location>"),
			Properties: &armvmwarecloudsimple.VirtualMachineProperties{
				AmountOfRAM: to.Ptr[int32](4096),
				Disks: []*armvmwarecloudsimple.VirtualDisk{
					{
						ControllerID:     to.Ptr("<controller-id>"),
						IndependenceMode: to.Ptr(armvmwarecloudsimple.DiskIndependenceModePersistent),
						TotalSize:        to.Ptr[int32](10485760),
						VirtualDiskID:    to.Ptr("<virtual-disk-id>"),
					}},
				Nics: []*armvmwarecloudsimple.VirtualNic{
					{
						Network: &armvmwarecloudsimple.VirtualNetwork{
							ID: to.Ptr("<id>"),
						},
						NicType:      to.Ptr(armvmwarecloudsimple.NICTypeE1000),
						PowerOnBoot:  to.Ptr(true),
						VirtualNicID: to.Ptr("<virtual-nic-id>"),
					}},
				NumberOfCores:  to.Ptr[int32](2),
				PrivateCloudID: to.Ptr("<private-cloud-id>"),
				ResourcePool: &armvmwarecloudsimple.ResourcePool{
					ID: to.Ptr("<id>"),
				},
				TemplateID: to.Ptr("<template-id>"),
			},
		},
		&armvmwarecloudsimple.VirtualMachinesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
