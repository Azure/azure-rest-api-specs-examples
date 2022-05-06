Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv0.4.0/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevtestlabs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_CreateOrUpdate.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewVirtualMachinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		armdevtestlabs.LabVirtualMachine{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"tagName1": to.Ptr("tagValue1"),
			},
			Properties: &armdevtestlabs.LabVirtualMachineProperties{
				AllowClaim:              to.Ptr(true),
				DisallowPublicIPAddress: to.Ptr(true),
				GalleryImageReference: &armdevtestlabs.GalleryImageReference{
					Offer:     to.Ptr("<offer>"),
					OSType:    to.Ptr("<ostype>"),
					Publisher: to.Ptr("<publisher>"),
					SKU:       to.Ptr("<sku>"),
					Version:   to.Ptr("<version>"),
				},
				LabSubnetName:       to.Ptr("<lab-subnet-name>"),
				LabVirtualNetworkID: to.Ptr("<lab-virtual-network-id>"),
				Password:            to.Ptr("<password>"),
				Size:                to.Ptr("<size>"),
				StorageType:         to.Ptr("<storage-type>"),
				UserName:            to.Ptr("<user-name>"),
			},
		},
		&armdevtestlabs.VirtualMachinesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
