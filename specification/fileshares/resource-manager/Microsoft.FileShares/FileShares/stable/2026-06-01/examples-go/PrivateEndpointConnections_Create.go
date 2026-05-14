package armfileshares_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fileshares/armfileshares"
)

// Generated from example definition: 2026-06-01/PrivateEndpointConnections_Create.json
func ExamplePrivateEndpointConnectionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfileshares.NewClientFactory("0681745E-3F9F-4966-80E6-69624A3B29F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreate(ctx, "rgfileshares", "fileshare", "privateEndpointConnection1", armfileshares.PrivateEndpointConnection{
		Properties: &armfileshares.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armfileshares.PrivateLinkServiceConnectionState{
				Status:      to.Ptr(armfileshares.PrivateEndpointServiceConnectionStatusApproved),
				Description: to.Ptr("Approved by admin"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armfileshares.PrivateEndpointConnectionsClientCreateResponse{
	// 	PrivateEndpointConnection: armfileshares.PrivateEndpointConnection{
	// 		Properties: &armfileshares.PrivateEndpointConnectionProperties{
	// 			GroupIDs: []*string{
	// 				to.Ptr("fileshare"),
	// 			},
	// 			PrivateEndpoint: &armfileshares.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/0681745E-3F9F-4966-80E6-69624A3B29F2/resourceGroups/rgfileshares/providers/Microsoft.Network/privateEndpoints/testPrivateEndpoint"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armfileshares.PrivateLinkServiceConnectionState{
	// 				Status: to.Ptr(armfileshares.PrivateEndpointServiceConnectionStatusApproved),
	// 				Description: to.Ptr("Approved by admin"),
	// 				ActionsRequired: to.Ptr("None"),
	// 			},
	// 			ProvisioningState: to.Ptr(armfileshares.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0681745E-3F9F-4966-80E6-69624A3B29F2/resourceGroups/rgfileshares/providers/Microsoft.FileShares/fileShares/fileshare/privateEndpointConnections/privateEndpointConnection1"),
	// 		Name: to.Ptr("privateEndpointConnection1"),
	// 		Type: to.Ptr("Microsoft.FileShares/fileShares/privateEndpointConnections"),
	// 		SystemData: &armfileshares.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armfileshares.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-23T12:37:53.955Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armfileshares.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-23T12:37:53.955Z"); return t}()),
	// 		},
	// 	},
	// }
}
