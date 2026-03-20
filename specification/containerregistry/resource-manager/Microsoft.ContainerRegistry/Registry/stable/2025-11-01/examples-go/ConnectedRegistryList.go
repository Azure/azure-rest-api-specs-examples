package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2025-11-01/ConnectedRegistryList.json
func ExampleConnectedRegistriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectedRegistriesClient().NewListPager("myResourceGroup", "myRegistry", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcontainerregistry.ConnectedRegistriesClientListResponse{
		// 	ConnectedRegistryListResult: armcontainerregistry.ConnectedRegistryListResult{
		// 		Value: []*armcontainerregistry.ConnectedRegistry{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/connectedRegistries/myConnectedRegistry"),
		// 				Name: to.Ptr("myConnectedRegistry"),
		// 				Type: to.Ptr("Microsoft.ContainerRegistry/registries/connectedRegistries"),
		// 				Properties: &armcontainerregistry.ConnectedRegistryProperties{
		// 					Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
		// 					Activation: &armcontainerregistry.ActivationProperties{
		// 						Status: to.Ptr(armcontainerregistry.ActivationStatusInactive),
		// 					},
		// 					Parent: &armcontainerregistry.ParentProperties{
		// 						SyncProperties: &armcontainerregistry.SyncProperties{
		// 							TokenID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
		// 							Schedule: to.Ptr("0 9 * * *"),
		// 							MessageTTL: to.Ptr("P2D"),
		// 							SyncWindow: to.Ptr("PT3H"),
		// 						},
		// 					},
		// 					ClientTokenIDs: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),
		// 					},
		// 					Logging: &armcontainerregistry.LoggingProperties{
		// 						LogLevel: to.Ptr(armcontainerregistry.LogLevelInformation),
		// 						AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusDisabled),
		// 					},
		// 					NotificationsList: []*string{
		// 						to.Ptr("hello-world:*:*"),
		// 						to.Ptr("sample/repo/*:1.0:*"),
		// 					},
		// 					GarbageCollection: &armcontainerregistry.GarbageCollectionProperties{
		// 						Enabled: to.Ptr(true),
		// 						Schedule: to.Ptr("0 5 * * *"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
