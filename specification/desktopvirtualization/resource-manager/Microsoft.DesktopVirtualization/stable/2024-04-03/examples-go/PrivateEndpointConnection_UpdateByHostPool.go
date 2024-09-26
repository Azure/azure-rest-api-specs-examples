package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/PrivateEndpointConnection_UpdateByHostPool.json
func ExamplePrivateEndpointConnectionsClient_UpdateByHostPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().UpdateByHostPool(ctx, "resourceGroup1", "hostPool1", "hostPool1.377103f1-5179-4bdf-8556-4cdd3207cc5b", armdesktopvirtualization.PrivateEndpointConnection{
		Properties: &armdesktopvirtualization.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armdesktopvirtualization.PrivateLinkServiceConnectionState{
				Description:     to.Ptr("Approved by admin@consoto.com"),
				ActionsRequired: to.Ptr("None"),
				Status:          to.Ptr(armdesktopvirtualization.PrivateEndpointServiceConnectionStatusApproved),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionWithSystemData = armdesktopvirtualization.PrivateEndpointConnectionWithSystemData{
	// 	Name: to.Ptr("hostPool1.377103f1-5179-4bdf-8556-4cdd3207cc5b"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/hostpools/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostpools/hostPool1/privateEndpointConnections/hostPool1.377103f1-5179-4bdf-8556-4cdd3207cc5b"),
	// 	SystemData: &armdesktopvirtualization.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 	},
	// 	Properties: &armdesktopvirtualization.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armdesktopvirtualization.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup2/providers/Microsoft.Network/privateEndpoints/endpointName1"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armdesktopvirtualization.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Approved by admin@consoto.com"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armdesktopvirtualization.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armdesktopvirtualization.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
