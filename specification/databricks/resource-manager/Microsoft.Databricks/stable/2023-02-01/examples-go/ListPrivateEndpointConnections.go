package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1a87e1a5deb3f986ea1474d233d6680f1e19fc1/specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/ListPrivateEndpointConnections.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("myResourceGroup", "myWorkspace", nil)
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
		// page.PrivateEndpointConnectionsList = armdatabricks.PrivateEndpointConnectionsList{
		// 	Value: []*armdatabricks.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("myWorkspace.23456789-1111-1111-1111-111111111111"),
		// 			Type: to.Ptr("Microsoft.Databricks/workspaces/PrivateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Databricks/workspaces/myWorkspace/PrivateEndpointConnections/myWorkspace.23456789-1111-1111-1111-111111111111"),
		// 			Properties: &armdatabricks.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armdatabricks.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/networkResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armdatabricks.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Approved by johndoe@company.com"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armdatabricks.PrivateLinkServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armdatabricks.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
