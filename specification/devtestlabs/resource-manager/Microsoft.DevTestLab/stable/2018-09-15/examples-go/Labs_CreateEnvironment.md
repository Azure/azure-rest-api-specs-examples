Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv0.2.1/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_CreateEnvironment.json
func ExampleLabsClient_BeginCreateEnvironment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateEnvironment(ctx,
		"<resource-group-name>",
		"<name>",
		armdevtestlabs.LabVirtualMachineCreationParameter{
			Name:     to.StringPtr("<name>"),
			Location: to.StringPtr("<location>"),
			Properties: &armdevtestlabs.LabVirtualMachineCreationParameterProperties{
				AllowClaim:              to.BoolPtr(true),
				DisallowPublicIPAddress: to.BoolPtr(true),
				GalleryImageReference: &armdevtestlabs.GalleryImageReference{
					Offer:     to.StringPtr("<offer>"),
					OSType:    to.StringPtr("<ostype>"),
					Publisher: to.StringPtr("<publisher>"),
					SKU:       to.StringPtr("<sku>"),
					Version:   to.StringPtr("<version>"),
				},
				LabSubnetName:       to.StringPtr("<lab-subnet-name>"),
				LabVirtualNetworkID: to.StringPtr("<lab-virtual-network-id>"),
				Password:            to.StringPtr("<password>"),
				Size:                to.StringPtr("<size>"),
				StorageType:         to.StringPtr("<storage-type>"),
				UserName:            to.StringPtr("<user-name>"),
			},
			Tags: map[string]*string{
				"tagName1": to.StringPtr("tagValue1"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
