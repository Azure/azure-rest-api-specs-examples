Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataboxedge%2Farmdataboxedge%2Fv0.1.0/sdk/resourcemanager/databoxedge/armdataboxedge/README.md) on how to add the SDK to your project and authenticate.

```go
package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DataBoxEdgeDevicePut.json
func ExampleDevicesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.DataBoxEdgeDevice{
			Location: to.StringPtr("<location>"),
			SKU: &armdataboxedge.SKU{
				Name: armdataboxedge.SKUNameEdge.ToPtr(),
				Tier: armdataboxedge.SKUTierStandard.ToPtr(),
			},
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataBoxEdgeDevice.ID: %s\n", *res.ID)
}
```
