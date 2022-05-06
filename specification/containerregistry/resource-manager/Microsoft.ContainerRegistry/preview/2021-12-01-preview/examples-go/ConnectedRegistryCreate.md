Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerregistry%2Farmcontainerregistry%2Fv0.5.0/sdk/resourcemanager/containerregistry/armcontainerregistry/README.md) on how to add the SDK to your project and authenticate.

```go
package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2021-12-01-preview/examples/ConnectedRegistryCreate.json
func ExampleConnectedRegistriesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewConnectedRegistriesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<connected-registry-name>",
		armcontainerregistry.ConnectedRegistry{
			Properties: &armcontainerregistry.ConnectedRegistryProperties{
				ClientTokenIDs: []*string{
					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token")},
				Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
				NotificationsList: []*string{
					to.Ptr("hello-world:*:*"),
					to.Ptr("sample/repo/*:1.0:*")},
				Parent: &armcontainerregistry.ParentProperties{
					SyncProperties: &armcontainerregistry.SyncProperties{
						MessageTTL: to.Ptr("<message-ttl>"),
						Schedule:   to.Ptr("<schedule>"),
						SyncWindow: to.Ptr("<sync-window>"),
						TokenID:    to.Ptr("<token-id>"),
					},
				},
			},
		},
		&armcontainerregistry.ConnectedRegistriesClientBeginCreateOptions{ResumeToken: ""})
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
