package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/search/resource-manager/Microsoft.Search/preview/2024-03-01-preview/examples/UpdatePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Update(ctx, "rg1", "mysearchservice", "testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546", armsearch.PrivateEndpointConnection{
		Properties: &armsearch.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armsearch.PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState{
				Description: to.Ptr("Rejected for some reason."),
				Status:      to.Ptr(armsearch.PrivateLinkServiceConnectionStatusRejected),
			},
		},
	}, &armsearch.SearchManagementRequestOptions{ClientRequestID: nil}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armsearch.PrivateEndpointConnection{
	// 	Name: to.Ptr("testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546"),
	// 	Type: to.Ptr("Microsoft.Search/searchServices/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice/privateEndpointConnections/testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546"),
	// 	Properties: &armsearch.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armsearch.PrivateEndpointConnectionPropertiesPrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testEndpoint"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armsearch.PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Rejected for some reason."),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armsearch.PrivateLinkServiceConnectionStatusRejected),
	// 		},
	// 	},
	// }
}
