package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/PrivateEndpointConnectionsUpdate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreate(ctx, "myResourceGroup", "myWorkspace", "myWorkspace.23456789-1111-1111-1111-111111111111", armdatabricks.PrivateEndpointConnection{
		Properties: &armdatabricks.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armdatabricks.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Approved by databricksadmin@contoso.com"),
				Status:      to.Ptr(armdatabricks.PrivateLinkServiceConnectionStatusApproved),
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
	// res.PrivateEndpointConnection = armdatabricks.PrivateEndpointConnection{
	// 	Name: to.Ptr("myWorkspace.23456789-1111-1111-1111-111111111111"),
	// 	Type: to.Ptr("Microsoft.Databricks/workspaces/PrivateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Databricks/workspaces/myWorkspace/PrivateEndpointConnections/myWorkspace.23456789-1111-1111-1111-111111111111"),
	// 	Properties: &armdatabricks.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armdatabricks.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/networkResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armdatabricks.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Approved by databricksadmin@contoso.com"),
	// 			ActionRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armdatabricks.PrivateLinkServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armdatabricks.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
