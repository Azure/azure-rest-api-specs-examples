package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresUpdatePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "contoso", "myConnection", armappconfiguration.PrivateEndpointConnection{
		Properties: &armappconfiguration.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armappconfiguration.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Auto-Approved"),
				Status:      to.Ptr(armappconfiguration.ConnectionStatusApproved),
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
	// res.PrivateEndpointConnection = armappconfiguration.PrivateEndpointConnection{
	// 	Name: to.Ptr("myConnection"),
	// 	Type: to.Ptr("Microsoft.AppConfiguration/configurationStores/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso/privateEndpointConnections/myConnection"),
	// 	Properties: &armappconfiguration.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armappconfiguration.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/peexample01"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armappconfiguration.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			ActionsRequired: to.Ptr(armappconfiguration.ActionsRequiredNone),
	// 			Status: to.Ptr(armappconfiguration.ConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armappconfiguration.ProvisioningStateSucceeded),
	// 	},
	// }
}
