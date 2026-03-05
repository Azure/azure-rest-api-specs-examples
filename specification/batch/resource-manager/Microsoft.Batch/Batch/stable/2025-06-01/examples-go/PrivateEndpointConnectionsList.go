package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/PrivateEndpointConnectionsList.json
func ExamplePrivateEndpointConnectionClient_NewListByBatchAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionClient().NewListByBatchAccountPager("default-azurebatch-japaneast", "sampleacct", nil)
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
		// page = armbatch.PrivateEndpointConnectionClientListByBatchAccountResponse{
		// 	ListPrivateEndpointConnectionsResult: armbatch.ListPrivateEndpointConnectionsResult{
		// 		Value: []*armbatch.PrivateEndpointConnection{
		// 			{
		// 				Name: to.Ptr("testprivateEndpointConnection"),
		// 				Type: to.Ptr("Microsoft.Batch/batchAccounts/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/privateEndpointConnections/testprivateEndpointConnection"),
		// 				Properties: &armbatch.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("batchAccount"),
		// 					},
		// 					PrivateEndpoint: &armbatch.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Network/privateEndpoints/testprivateEndpoint"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armbatch.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Approved by xyz.abc@company.com"),
		// 						Status: to.Ptr(armbatch.PrivateLinkServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armbatch.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
