Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fattestation%2Farmattestation%2Fv0.1.0/sdk/resourcemanager/attestation/armattestation/README.md) on how to add the SDK to your project and authenticate.

```go
package armattestation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation"
)

// x-ms-original-file: specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/AttestationProviderPutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armattestation.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<provider-name>",
		"<private-endpoint-connection-name>",
		armattestation.PrivateEndpointConnection{
			Properties: &armattestation.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armattestation.PrivateLinkServiceConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armattestation.PrivateEndpointServiceConnectionStatusApproved.ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateEndpointConnection.ID: %s\n", *res.ID)
}
```
