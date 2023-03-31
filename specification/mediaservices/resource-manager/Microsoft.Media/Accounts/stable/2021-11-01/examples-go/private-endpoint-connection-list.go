package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2021-11-01/examples/private-endpoint-connection-list.json
func ExamplePrivateEndpointConnectionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().List(ctx, "contoso", "contososports", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionListResult = armmediaservices.PrivateEndpointConnectionListResult{
	// 	Value: []*armmediaservices.PrivateEndpointConnection{
	// 		{
	// 			Name: to.Ptr("cn1"),
	// 			Type: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnections"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/fabrikam/providers/Microsoft.Media/mediaservices/contososports/privateEndpointConnections/cn1"),
	// 			Properties: &armmediaservices.PrivateEndpointConnectionProperties{
	// 				PrivateEndpoint: &armmediaservices.PrivateEndpoint{
	// 					ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/reosuceGroup1/providers/Microsoft.Network/privateEndpoints/pe1"),
	// 				},
	// 				PrivateLinkServiceConnectionState: &armmediaservices.PrivateLinkServiceConnectionState{
	// 					Description: to.Ptr("Test description"),
	// 					Status: to.Ptr(armmediaservices.PrivateEndpointServiceConnectionStatusApproved),
	// 				},
	// 				ProvisioningState: to.Ptr(armmediaservices.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("cn2"),
	// 			Type: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnections"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/fabrikam/providers/Microsoft.Media/mediaservices/contososports/privateEndpointConnections/cn2"),
	// 			Properties: &armmediaservices.PrivateEndpointConnectionProperties{
	// 				PrivateEndpoint: &armmediaservices.PrivateEndpoint{
	// 					ID: to.Ptr("/subscriptions/22222222-2222-2222-2222-222222222222/resourceGroups/reosuceGroup2/providers/Microsoft.Network/privateEndpoints/pe2"),
	// 				},
	// 				PrivateLinkServiceConnectionState: &armmediaservices.PrivateLinkServiceConnectionState{
	// 					Description: to.Ptr("Test description"),
	// 					Status: to.Ptr(armmediaservices.PrivateEndpointServiceConnectionStatusPending),
	// 				},
	// 				ProvisioningState: to.Ptr(armmediaservices.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 			},
	// 	}},
	// }
}
