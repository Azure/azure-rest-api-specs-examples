package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/PrivateEndPointConnectionList.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("SDK-EventHub-4794", "sdk-Namespace-5828", nil)
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
		// page.PrivateEndpointConnectionListResult = armeventhub.PrivateEndpointConnectionListResult{
		// 	Value: []*armeventhub.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("5dc668b3-70e4-437f-b61c-a3c1e594be7a"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/PrivateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-EventHub-7182/providers/Microsoft.EventHub/namespaces/sdk-Namespace-5705-new/privateEndpointConnections/5dc668b3-70e4-437f-b61c-a3c1e594be7a"),
		// 			Properties: &armeventhub.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armeventhub.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-EventHub-7182/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-5705-new"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armeventhub.ConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					Status: to.Ptr(armeventhub.PrivateLinkConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armeventhub.EndPointProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
