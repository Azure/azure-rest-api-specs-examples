package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoManagedPrivateEndpointsList.json
func ExampleManagedPrivateEndpointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedPrivateEndpointsClient().NewListPager("kustorptest", "kustoCluster", nil)
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
		// page.ManagedPrivateEndpointListResult = armkusto.ManagedPrivateEndpointListResult{
		// 	Value: []*armkusto.ManagedPrivateEndpoint{
		// 		{
		// 			Name: to.Ptr("kustoCluster/KustoDatabase8/kustomanagedPrivateEndpoint1"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/ManagedPrivateEndpoints"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/ManagedPrivateEndpoints/kustoManagedPrivateEndpoint1"),
		// 			Properties: &armkusto.ManagedPrivateEndpointProperties{
		// 				GroupID: to.Ptr("blob"),
		// 				PrivateLinkResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/storageAccountTest"),
		// 				RequestMessage: to.Ptr("Please Approve."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kustoCluster/KustoDatabase8/kustomanagedPrivateEndpoint2"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/ManagedPrivateEndpoints"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/ManagedPrivateEndpoints/kustoManagedPrivateEndpoint2"),
		// 			Properties: &armkusto.ManagedPrivateEndpointProperties{
		// 				GroupID: to.Ptr("namespace"),
		// 				PrivateLinkResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHubs/storageAccounts/eventHubTest"),
		// 				RequestMessage: to.Ptr("Please Approve."),
		// 			},
		// 	}},
		// }
	}
}
