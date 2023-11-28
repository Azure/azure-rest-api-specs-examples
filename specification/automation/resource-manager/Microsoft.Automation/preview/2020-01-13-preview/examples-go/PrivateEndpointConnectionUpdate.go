package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/PrivateEndpointConnectionUpdate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "rg1", "ddb1", "privateEndpointConnectionName", armautomation.PrivateEndpointConnection{
		Properties: &armautomation.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armautomation.PrivateLinkServiceConnectionStateProperty{
				Description: to.Ptr("Approved by johndoe@contoso.com"),
				Status:      to.Ptr("Approved"),
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
	// res.PrivateEndpointConnection = armautomation.PrivateEndpointConnection{
	// 	Name: to.Ptr("privateEndpointConnectionName"),
	// 	Type: to.Ptr("Microsoft.Automation/automationAccounts/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.Automation/automationAccounts/ddb1/privateEndpointConnections/privateEndpointConnectionName"),
	// 	Properties: &armautomation.PrivateEndpointConnectionProperties{
	// 		GroupIDs: []*string{
	// 			to.Ptr("sql")},
	// 			PrivateEndpoint: &armautomation.PrivateEndpointProperty{
	// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1Network/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armautomation.PrivateLinkServiceConnectionStateProperty{
	// 				Description: to.Ptr("Auto-approved"),
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr("Approved"),
	// 			},
	// 		},
	// 	}
}
