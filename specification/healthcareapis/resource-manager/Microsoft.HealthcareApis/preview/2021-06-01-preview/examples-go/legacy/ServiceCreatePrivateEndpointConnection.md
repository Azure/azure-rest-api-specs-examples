Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhealthcareapis%2Farmhealthcareapis%2Fv0.2.1/sdk/resourcemanager/healthcareapis/armhealthcareapis/README.md) on how to add the SDK to your project and authenticate.

```go
package armhealthcareapis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/legacy/ServiceCreatePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<private-endpoint-connection-name>",
		armhealthcareapis.PrivateEndpointConnection{
			Properties: &armhealthcareapis.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armhealthcareapis.PrivateLinkServiceConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armhealthcareapis.PrivateEndpointServiceConnectionStatus("Approved").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionsClientCreateOrUpdateResult)
}
```
