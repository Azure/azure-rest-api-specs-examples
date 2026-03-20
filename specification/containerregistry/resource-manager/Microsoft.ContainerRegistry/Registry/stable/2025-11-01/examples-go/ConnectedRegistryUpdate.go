package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2025-11-01/ConnectedRegistryUpdate.json
func ExampleConnectedRegistriesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectedRegistriesClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myScopeMap", armcontainerregistry.ConnectedRegistryUpdateParameters{
		Properties: &armcontainerregistry.ConnectedRegistryUpdateProperties{
			SyncProperties: &armcontainerregistry.SyncUpdateProperties{
				Schedule:   to.Ptr("0 0 */10 * *"),
				MessageTTL: to.Ptr("P30D"),
				SyncWindow: to.Ptr("P2D"),
			},
			Logging: &armcontainerregistry.LoggingProperties{
				LogLevel:       to.Ptr(armcontainerregistry.LogLevelDebug),
				AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusEnabled),
			},
			ClientTokenIDs: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token"),
			},
			NotificationsList: []*string{
				to.Ptr("hello-world:*:*"),
				to.Ptr("sample/repo/*:1.0:*"),
			},
			GarbageCollection: &armcontainerregistry.GarbageCollectionProperties{
				Enabled:  to.Ptr(true),
				Schedule: to.Ptr("0 5 * * *"),
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
	// res = armcontainerregistry.ConnectedRegistriesClientUpdateResponse{
	// 	ConnectedRegistry: &armcontainerregistry.ConnectedRegistry{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/connectedRegistries/myConnectedRegistry"),
	// 		Name: to.Ptr("myConnectedRegistry"),
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/connectedRegistries"),
	// 		Properties: &armcontainerregistry.ConnectedRegistryProperties{
	// 			Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
	// 			Parent: &armcontainerregistry.ParentProperties{
	// 				SyncProperties: &armcontainerregistry.SyncProperties{
	// 					TokenID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
	// 					Schedule: to.Ptr("0 0 */10 * *"),
	// 					MessageTTL: to.Ptr("P30D"),
	// 					SyncWindow: to.Ptr("P2D"),
	// 				},
	// 			},
	// 			ClientTokenIDs: []*string{
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token"),
	// 			},
	// 			Logging: &armcontainerregistry.LoggingProperties{
	// 				LogLevel: to.Ptr(armcontainerregistry.LogLevelDebug),
	// 				AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusEnabled),
	// 			},
	// 			NotificationsList: []*string{
	// 				to.Ptr("hello-world:*:*"),
	// 				to.Ptr("sample/repo/*:1.0:*"),
	// 			},
	// 			GarbageCollection: &armcontainerregistry.GarbageCollectionProperties{
	// 				Enabled: to.Ptr(true),
	// 				Schedule: to.Ptr("0 5 * * *"),
	// 			},
	// 		},
	// 	},
	// }
}
