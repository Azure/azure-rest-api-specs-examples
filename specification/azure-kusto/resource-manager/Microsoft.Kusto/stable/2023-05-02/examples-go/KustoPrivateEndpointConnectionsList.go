package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoPrivateEndpointConnectionsList.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("kustorptest", "kustoCluster", nil)
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
		// page.PrivateEndpointConnectionListResult = armkusto.PrivateEndpointConnectionListResult{
		// 	Value: []*armkusto.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("privateEndpointTest"),
		// 			Type: to.Ptr("Microsoft.Kusto/clusters/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/privateEndpointConnections/privateEndpointTest"),
		// 			Properties: &armkusto.PrivateEndpointConnectionProperties{
		// 				GroupID: to.Ptr("cluster"),
		// 				PrivateEndpoint: &armkusto.PrivateEndpointProperty{
		// 					ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/rg1Network/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armkusto.PrivateLinkServiceConnectionStateProperty{
		// 					Description: to.Ptr("Auto-approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr("Approved"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("privateEndpointTest2"),
		// 			Type: to.Ptr("Microsoft.Kusto/clusters/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/privateEndpointConnections/privateEndpointTest2"),
		// 			Properties: &armkusto.PrivateEndpointConnectionProperties{
		// 				GroupID: to.Ptr("cluster"),
		// 				PrivateEndpoint: &armkusto.PrivateEndpointProperty{
		// 					ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/rg1Network/providers/Microsoft.Network/privateEndpoints/privateEndpointName2"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armkusto.PrivateLinkServiceConnectionStateProperty{
		// 					Description: to.Ptr("Auto-approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr("Approved"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}
