package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1a87e1a5deb3f986ea1474d233d6680f1e19fc1/specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/PrivateEndpointConnectionsGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "myResourceGroup", "myWorkspace", "myWorkspace.23456789-1111-1111-1111-111111111111", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armdatabricks.PrivateEndpointConnection{
	// 	Name: to.Ptr("myWorkspace.23456789-1111-1111-1111-111111111111"),
	// 	Type: to.Ptr("Microsoft.Databricks/workspaces/PrivateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Databricks/workspaces/myWorkspace/PrivateEndpointConnections/myWorkspace.23456789-1111-1111-1111-111111111111"),
	// 	Properties: &armdatabricks.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armdatabricks.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/networkResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armdatabricks.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Please approve my request!"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armdatabricks.PrivateLinkServiceConnectionStatusPending),
	// 		},
	// 		ProvisioningState: to.Ptr(armdatabricks.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
