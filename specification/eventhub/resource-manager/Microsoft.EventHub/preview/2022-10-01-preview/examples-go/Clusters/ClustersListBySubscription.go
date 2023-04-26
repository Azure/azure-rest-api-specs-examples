package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/Clusters/ClustersListBySubscription.json
func ExampleClustersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListBySubscriptionPager(nil)
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
		// page.ClusterListResult = armeventhub.ClusterListResult{
		// 	Value: []*armeventhub.Cluster{
		// 		{
		// 			Name: to.Ptr("testCluster"),
		// 			Type: to.Ptr("Microsoft.EventHub/Clusters"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-EventHub-SouthCentralUS/providers/Microsoft.EventHub/clusters/testCluster"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armeventhub.ClusterProperties{
		// 				CreatedAt: to.Ptr("2016-09-13T23:17:25.24Z"),
		// 				MetricID: to.Ptr("SN6-008"),
		// 				UpdatedAt: to.Ptr("2016-09-13T23:17:28.223Z"),
		// 			},
		// 			SKU: &armeventhub.ClusterSKU{
		// 				Name: to.Ptr(armeventhub.ClusterSKUNameDedicated),
		// 				Capacity: to.Ptr[int32](4),
		// 			},
		// 	}},
		// }
	}
}
