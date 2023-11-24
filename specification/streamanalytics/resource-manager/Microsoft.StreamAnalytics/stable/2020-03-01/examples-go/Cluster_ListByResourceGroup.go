package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Cluster_ListByResourceGroup.json
func ExampleClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListByResourceGroupPager("sjrg", nil)
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
		// page.ClusterListResult = armstreamanalytics.ClusterListResult{
		// 	Value: []*armstreamanalytics.Cluster{
		// 		{
		// 			Name: to.Ptr("An Example Cluster"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/clusters"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/sjrg/providers/Microsoft.StreamAnalytics/clusters/AnExampleStreamingCluster"),
		// 			Location: to.Ptr("Central US"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Etag: to.Ptr("F86B9B70-D5B1-451D-AFC8-0B42D4729B8C"),
		// 			Properties: &armstreamanalytics.ClusterProperties{
		// 				CapacityAllocated: to.Ptr[int32](48),
		// 				CapacityAssigned: to.Ptr[int32](96),
		// 				ClusterID: to.Ptr("B01C67EF-4739-4DDD-9FB2-427EB43DE839"),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-25T01:00:00.000Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armstreamanalytics.ClusterProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armstreamanalytics.ClusterSKU{
		// 				Name: to.Ptr(armstreamanalytics.ClusterSKUNameDefault),
		// 				Capacity: to.Ptr[int32](96),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("A Different Cluster"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/clusters"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/sjrg/providers/Microsoft.StreamAnalytics/clusters/ADifferentStreamingCluster"),
		// 			Location: to.Ptr("Central US"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Etag: to.Ptr("G97C0C81-D5B1-451D-AFC8-0B42D4729B8C"),
		// 			Properties: &armstreamanalytics.ClusterProperties{
		// 				CapacityAllocated: to.Ptr[int32](48),
		// 				CapacityAssigned: to.Ptr[int32](96),
		// 				ClusterID: to.Ptr("B01C67EF-4739-4DDD-9FB2-427EB43DE839"),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-25T01:00:00.000Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armstreamanalytics.ClusterProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armstreamanalytics.ClusterSKU{
		// 				Name: to.Ptr(armstreamanalytics.ClusterSKUNameDefault),
		// 				Capacity: to.Ptr[int32](96),
		// 			},
		// 	}},
		// }
	}
}
