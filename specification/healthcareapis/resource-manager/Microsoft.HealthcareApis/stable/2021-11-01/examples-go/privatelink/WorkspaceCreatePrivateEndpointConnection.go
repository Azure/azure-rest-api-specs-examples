package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/privatelink/WorkspaceCreatePrivateEndpointConnection.json
func ExampleWorkspacePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacePrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "testRG", "workspace1", "myConnection", armhealthcareapis.PrivateEndpointConnectionDescription{
		Properties: &armhealthcareapis.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armhealthcareapis.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Auto-Approved"),
				Status:      to.Ptr(armhealthcareapis.PrivateEndpointServiceConnectionStatusApproved),
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
	// res.PrivateEndpointConnectionDescription = armhealthcareapis.PrivateEndpointConnectionDescription{
	// 	Name: to.Ptr("myConnection"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/workspaces/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HealthcareApis/workspaces/workspace1/privateEndpointConnections/myConnection"),
	// 	Properties: &armhealthcareapis.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armhealthcareapis.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/peexample01"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armhealthcareapis.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armhealthcareapis.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armhealthcareapis.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
