package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-11-01-preview/examples/ConnectedRegistryCreate.json
func ExampleConnectedRegistriesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectedRegistriesClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myConnectedRegistry", armcontainerregistry.ConnectedRegistry{
		Properties: &armcontainerregistry.ConnectedRegistryProperties{
			ClientTokenIDs: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token")},
			Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
			NotificationsList: []*string{
				to.Ptr("hello-world:*:*"),
				to.Ptr("sample/repo/*:1.0:*")},
			Parent: &armcontainerregistry.ParentProperties{
				SyncProperties: &armcontainerregistry.SyncProperties{
					MessageTTL: to.Ptr("P2D"),
					Schedule:   to.Ptr("0 9 * * *"),
					SyncWindow: to.Ptr("PT3H"),
					TokenID:    to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
				},
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
	// 		Activation: &armcontainerregistry.ActivationProperties{
	// 			Status: to.Ptr(armcontainerregistry.ActivationStatusInactive),
	// 		},
	// 		ClientTokenIDs: []*string{
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token")},
	// 			Logging: &armcontainerregistry.LoggingProperties{
	// 				AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusDisabled),
	// 				LogLevel: to.Ptr(armcontainerregistry.LogLevelInformation),
	// 			},
	// 			Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
	// 			NotificationsList: []*string{
	// 				to.Ptr("hello-world:*:*"),
	// 				to.Ptr("sample/repo/*:1.0:*")},
	// 				Parent: &armcontainerregistry.ParentProperties{
	// 					SyncProperties: &armcontainerregistry.SyncProperties{
	// 						MessageTTL: to.Ptr("P2D"),
	// 						Schedule: to.Ptr("0 9 * * *"),
	// 						SyncWindow: to.Ptr("PT3H"),
	// 						TokenID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
	// 					},
	// 				},
	// 			},
	// 		}
}
