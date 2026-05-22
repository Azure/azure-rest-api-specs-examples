package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-01-01-preview/ConnectedRegistryGet.json
func ExampleConnectedRegistriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedRegistriesClient().Get(ctx, "myResourceGroup", "myRegistry", "myConnectedRegistry", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerregistry.ConnectedRegistriesClientGetResponse{
	// 	ConnectedRegistry: &armcontainerregistry.ConnectedRegistry{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/connectedRegistries/myConnectedRegistry"),
	// 		Name: to.Ptr("myConnectedRegistry"),
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/connectedRegistries"),
	// 		Properties: &armcontainerregistry.ConnectedRegistryProperties{
	// 			Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
	// 			Activation: &armcontainerregistry.ActivationProperties{
	// 				Status: to.Ptr(armcontainerregistry.ActivationStatusInactive),
	// 			},
	// 			Parent: &armcontainerregistry.ParentProperties{
	// 				SyncProperties: &armcontainerregistry.SyncProperties{
	// 					TokenID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
	// 					Schedule: to.Ptr("0 9 * * *"),
	// 					MessageTTL: to.Ptr("P2D"),
	// 					SyncWindow: to.Ptr("PT3H"),
	// 				},
	// 			},
	// 			ClientTokenIDs: []*string{
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),
	// 			},
	// 			Logging: &armcontainerregistry.LoggingProperties{
	// 				LogLevel: to.Ptr(armcontainerregistry.LogLevelInformation),
	// 				AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusDisabled),
	// 			},
	// 			NotificationsList: []*string{
	// 				to.Ptr("hello-world:*:*"),
	// 				to.Ptr("sample/repo/*:1.0:*"),
	// 			},
	// 			GarbageCollection: &armcontainerregistry.GarbageCollectionProperties{
	// 				Enabled: to.Ptr(true),
	// 				Schedule: to.Ptr("0 5 * * *"),
	// 			},
	// 			RegistrySyncResult: &armcontainerregistry.RegistrySyncResult{
	// 				SyncTrigger: to.Ptr(armcontainerregistry.SyncTriggerInitialSync),
	// 				SyncState: to.Ptr(armcontainerregistry.SyncStateSucceeded),
	// 				LastSyncStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-01T00:00:00.0000000-01:00"); return t}()),
	// 				LastSyncEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-01T00:01:00.0000000-01:00"); return t}()),
	// 				LastSuccessfulSyncEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-01T00:01:00.0000000-01:00"); return t}()),
	// 			},
	// 		},
	// 	},
	// }
}
