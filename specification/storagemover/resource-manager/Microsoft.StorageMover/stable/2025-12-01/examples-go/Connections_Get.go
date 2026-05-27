package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: 2025-12-01/Connections_Get.json
func ExampleConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("60bcfc77-6589-4da2-b7fd-f9ec9322cf95", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionsClient().Get(ctx, "examples-rg", "examples-storageMoverName", "example-connection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragemover.ConnectionsClientGetResponse{
	// 	Connection: &armstoragemover.Connection{
	// 		ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/connections/example-connection"),
	// 		Name: to.Ptr("example-connection"),
	// 		Type: to.Ptr("Microsoft.StorageMover/storageMovers/connections"),
	// 		Properties: &armstoragemover.ConnectionProperties{
	// 			Description: to.Ptr("Example Connection Description"),
	// 			PrivateLinkServiceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Network/privateLinkServices/example-pls"),
	// 			ConnectionStatus: to.Ptr(armstoragemover.ConnectionStatusApproved),
	// 			PrivateEndpointName: to.Ptr("ExamplePrivateEndpointName"),
	// 			PrivateEndpointResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Network/privateEndpoints/example-private-endpoint"),
	// 		},
	// 	},
	// }
}
