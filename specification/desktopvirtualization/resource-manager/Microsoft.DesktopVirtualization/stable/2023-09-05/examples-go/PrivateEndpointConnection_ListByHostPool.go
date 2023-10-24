package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/PrivateEndpointConnection_ListByHostPool.json
func ExamplePrivateEndpointConnectionsClient_NewListByHostPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByHostPoolPager("resourceGroup1", "hostPool1", &armdesktopvirtualization.PrivateEndpointConnectionsClientListByHostPoolOptions{PageSize: nil,
		IsDescending: nil,
		InitialSkip:  nil,
	})
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
		// page.PrivateEndpointConnectionListResultWithSystemData = armdesktopvirtualization.PrivateEndpointConnectionListResultWithSystemData{
		// 	Value: []*armdesktopvirtualization.PrivateEndpointConnectionWithSystemData{
		// 		{
		// 			Name: to.Ptr("hostPool1.377103f1-5179-4bdf-8556-4cdd3207cc5b"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/hostpools/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostpools/hostPool1/privateEndpointConnections/hostPool1.377103f1-5179-4bdf-8556-4cdd3207cc5b"),
		// 			Properties: &armdesktopvirtualization.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armdesktopvirtualization.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup2/providers/Microsoft.Network/privateEndpoints/endpointName1"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armdesktopvirtualization.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armdesktopvirtualization.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armdesktopvirtualization.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 			SystemData: &armdesktopvirtualization.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
