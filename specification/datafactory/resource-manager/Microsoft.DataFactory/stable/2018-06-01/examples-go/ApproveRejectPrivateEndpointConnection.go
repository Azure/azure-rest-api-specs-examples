package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ApproveRejectPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPrivateEndpointConnectionClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"exampleResourceGroup",
		"exampleFactoryName",
		"connection",
		armdatafactory.PrivateLinkConnectionApprovalRequestResource{
			Properties: &armdatafactory.PrivateLinkConnectionApprovalRequest{
				PrivateEndpoint: &armdatafactory.PrivateEndpoint{
					ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/privateEndpoints/myPrivateEndpoint"),
				},
				PrivateLinkServiceConnectionState: &armdatafactory.PrivateLinkConnectionState{
					Description:     to.Ptr("Approved by admin."),
					ActionsRequired: to.Ptr(""),
					Status:          to.Ptr("Approved"),
				},
			},
		},
		&armdatafactory.PrivateEndpointConnectionClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
