package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/523ccabf440d8cf1c5b0ea18a8ad1ffedf4902ac/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/ConnectedRegistryUpdate.json
func ExampleConnectedRegistriesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectedRegistriesClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myScopeMap", armcontainerregistry.ConnectedRegistryUpdateParameters{
		Properties: &armcontainerregistry.ConnectedRegistryUpdateProperties{
			ClientTokenIDs: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token")},
			GarbageCollection: &armcontainerregistry.GarbageCollectionProperties{
				Enabled:  to.Ptr(true),
				Schedule: to.Ptr("0 5 * * *"),
			},
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectedRegistry = armcontainerregistry.ConnectedRegistry{
	// 	Name: to.Ptr("myConnectedRegistry"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/connectedRegistries"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/connectedRegistries/myConnectedRegistry"),
	// 	Properties: &armcontainerregistry.ConnectedRegistryProperties{
	// 		ClientTokenIDs: []*string{
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token")},
	// 			GarbageCollection: &armcontainerregistry.GarbageCollectionProperties{
	// 				Enabled: to.Ptr(true),
	// 				Schedule: to.Ptr("0 5 * * *"),
	// 			},
	// 			Logging: &armcontainerregistry.LoggingProperties{
	// 				AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusEnabled),
	// 				LogLevel: to.Ptr(armcontainerregistry.LogLevelDebug),
	// 			},
	// 			Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
	// 			NotificationsList: []*string{
	// 				to.Ptr("hello-world:*:*"),
	// 				to.Ptr("sample/repo/*:1.0:*")},
	// 				Parent: &armcontainerregistry.ParentProperties{
	// 					SyncProperties: &armcontainerregistry.SyncProperties{
	// 						MessageTTL: to.Ptr("P30D"),
	// 						Schedule: to.Ptr("0 0 */10 * *"),
	// 						SyncWindow: to.Ptr("P2D"),
	// 						TokenID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
	// 					},
	// 				},
	// 			},
	// 		}
}
