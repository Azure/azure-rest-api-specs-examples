package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoPrivateEndpointConnectionsCreateOrUpdate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "kustorptest", "kustoclusterrptest4", "privateEndpointTest", armkusto.PrivateEndpointConnection{
		Properties: &armkusto.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armkusto.PrivateLinkServiceConnectionStateProperty{
				Description: to.Ptr("Approved by johndoe@contoso.com"),
				Status:      to.Ptr("Approved"),
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
	// res.PrivateEndpointConnection = armkusto.PrivateEndpointConnection{
	// 	Name: to.Ptr("privateEndpointTest"),
	// 	Type: to.Ptr("Microsoft.Kusto/clusters/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/KustoClusterRPTest4/privateEndpointConnections/privateEndpointTest"),
	// 	Properties: &armkusto.PrivateEndpointConnectionProperties{
	// 		PrivateLinkServiceConnectionState: &armkusto.PrivateLinkServiceConnectionStateProperty{
	// 			Description: to.Ptr("Approved by johndoe@contoso.com"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 	},
	// }
}
