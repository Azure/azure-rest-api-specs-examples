Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredisenterprise%2Farmredisenterprise%2Fv0.4.0/sdk/resourcemanager/redisenterprise/armredisenterprise/README.md) on how to add the SDK to your project and authenticate.

```go
package armredisenterprise_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterprisePutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armredisenterprise.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginPut(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<private-endpoint-connection-name>",
		armredisenterprise.PrivateEndpointConnection{
			Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
					Description: to.Ptr("<description>"),
					Status:      to.Ptr(armredisenterprise.PrivateEndpointServiceConnectionStatusApproved),
				},
			},
		},
		&armredisenterprise.PrivateEndpointConnectionsClientBeginPutOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```
