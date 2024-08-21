package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetAllPrivateEndpointConnectionsInCluster.json
func ExamplePrivateEndpointConnectionsClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByClusterPager("rg1", "cluster1", nil)
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
		// page.PrivateEndpointConnectionListResult = armhdinsight.PrivateEndpointConnectionListResult{
		// 	Value: []*armhdinsight.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("testprivateep.b3bf5fed-9b12-4560-b7d0-2abe1bba07e2"),
		// 			Type: to.Ptr("Microsoft.HDInsight/clusters/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HDInsight/clusters/cluster1/privateEndpointConnections/testprivateep.b3bf5fed-9b12-4560-b7d0-2abe1bba07e2"),
		// 			Properties: &armhdinsight.PrivateEndpointConnectionProperties{
		// 				LinkIdentifier: to.Ptr("620815036"),
		// 				PrivateEndpoint: &armhdinsight.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/testprivateep"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armhdinsight.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr(""),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armhdinsight.PrivateLinkServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armhdinsight.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
