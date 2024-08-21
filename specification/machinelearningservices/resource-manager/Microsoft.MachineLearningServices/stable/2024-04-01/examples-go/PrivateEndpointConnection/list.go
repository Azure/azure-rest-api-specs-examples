package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/PrivateEndpointConnection/list.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("rg-1234", "testworkspace", nil)
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
		// page.PrivateEndpointConnectionListResult = armmachinelearning.PrivateEndpointConnectionListResult{
		// 	Value: []*armmachinelearning.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("{privateEndpointConnectionName}"),
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg-1234/providers/Microsoft.MachineLearningServices/workspaces/testworkspace/privateEndpointConnections/{privateEndpointConnectionName}"),
		// 			Properties: &armmachinelearning.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armmachinelearning.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg-1234/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armmachinelearning.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armmachinelearning.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armmachinelearning.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{privateEndpointConnectionName}"),
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg-1234/providers/Microsoft.MachineLearningServices/workspaces/testworkspace/privateEndpointConnections/{privateEndpointConnectionName}"),
		// 			Properties: &armmachinelearning.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armmachinelearning.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg-1234/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armmachinelearning.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armmachinelearning.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armmachinelearning.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
