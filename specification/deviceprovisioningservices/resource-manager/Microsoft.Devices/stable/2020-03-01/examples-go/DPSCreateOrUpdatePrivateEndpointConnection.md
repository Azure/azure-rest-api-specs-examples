Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdeviceprovisioningservices%2Farmdeviceprovisioningservices%2Fv0.1.0/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2020-03-01/examples/DPSCreateOrUpdatePrivateEndpointConnection.json
func ExampleIotDpsResourceClient_BeginCreateOrUpdatePrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeviceprovisioningservices.NewIotDpsResourceClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdatePrivateEndpointConnection(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<private-endpoint-connection-name>",
		armdeviceprovisioningservices.PrivateEndpointConnection{
			Properties: &armdeviceprovisioningservices.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armdeviceprovisioningservices.PrivateLinkServiceConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armdeviceprovisioningservices.PrivateLinkServiceConnectionStatusApproved.ToPtr(),
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
	log.Printf("PrivateEndpointConnection.ID: %s\n", *res.ID)
}
```
