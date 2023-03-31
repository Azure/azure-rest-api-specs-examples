package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateEndpointConnections_CreateOrUpdate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "SampleResourceGroup", "account1", "privateEndpointConnection1", armpurview.PrivateEndpointConnection{
		Properties: &armpurview.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armpurview.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Approved by johndoe@company.com"),
				Status:      to.Ptr(armpurview.StatusApproved),
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
	// res.PrivateEndpointConnection = armpurview.PrivateEndpointConnection{
	// 	Name: to.Ptr("privateEndpointConnection1"),
	// 	Type: to.Ptr("Microsoft.Purview/accounts/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-12345678abc/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1/privateEndpointConnections/privateEndpointConnection1"),
	// 	Properties: &armpurview.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armpurview.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-12345678abc/resourceGroups/SampleResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armpurview.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Approved by johndoe@company.com"),
	// 			Status: to.Ptr(armpurview.StatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
