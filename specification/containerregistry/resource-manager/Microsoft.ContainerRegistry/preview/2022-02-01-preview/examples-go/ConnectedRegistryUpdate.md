Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerregistry%2Farmcontainerregistry%2Fv0.6.0/sdk/resourcemanager/containerregistry/armcontainerregistry/README.md) on how to add the SDK to your project and authenticate.

```go
package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ConnectedRegistryUpdate.json
func ExampleConnectedRegistriesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewConnectedRegistriesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myRegistry",
		"myScopeMap",
		armcontainerregistry.ConnectedRegistryUpdateParameters{
			Properties: &armcontainerregistry.ConnectedRegistryUpdateProperties{
				ClientTokenIDs: []*string{
					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),
					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token")},
				Logging: &armcontainerregistry.LoggingProperties{
					AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusEnabled),
					LogLevel:       to.Ptr(armcontainerregistry.LogLevelDebug),
				},
				NotificationsList: []*string{
					to.Ptr("hello-world:*:*"),
					to.Ptr("sample/repo/*:1.0:*")},
				SyncProperties: &armcontainerregistry.SyncUpdateProperties{
					MessageTTL: to.Ptr("P30D"),
					Schedule:   to.Ptr("0 0 */10 * *"),
					SyncWindow: to.Ptr("P2D"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
