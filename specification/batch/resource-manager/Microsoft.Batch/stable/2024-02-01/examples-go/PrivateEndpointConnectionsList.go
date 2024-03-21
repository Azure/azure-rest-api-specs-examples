package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/PrivateEndpointConnectionsList.json
func ExamplePrivateEndpointConnectionClient_NewListByBatchAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionClient().NewListByBatchAccountPager("default-azurebatch-japaneast", "sampleacct", &armbatch.PrivateEndpointConnectionClientListByBatchAccountOptions{Maxresults: nil})
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
		// page.ListPrivateEndpointConnectionsResult = armbatch.ListPrivateEndpointConnectionsResult{
		// 	Value: []*armbatch.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("testprivateEndpointConnection"),
		// 			Type: to.Ptr("Microsoft.Batch/batchAccounts/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/privateEndpointConnections/testprivateEndpointConnection"),
		// 			Properties: &armbatch.PrivateEndpointConnectionProperties{
		// 				GroupIDs: []*string{
		// 					to.Ptr("batchAccount")},
		// 					PrivateEndpoint: &armbatch.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Network/privateEndpoints/testprivateEndpoint"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armbatch.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Approved by xyz.abc@company.com"),
		// 						Status: to.Ptr(armbatch.PrivateLinkServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armbatch.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}
