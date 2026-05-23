package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub/v2"
)

// Generated from example definition: 2025-05-01-preview/Clusters/ClusterPut.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "testCluster", armeventhub.Cluster{
		Location: to.Ptr("South Central US"),
		SKU: &armeventhub.ClusterSKU{
			Name:     to.Ptr(armeventhub.ClusterSKUNameDedicated),
			Capacity: to.Ptr[int32](1),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armeventhub.ClustersClientCreateOrUpdateResponse{
	// 	Cluster: armeventhub.Cluster{
	// 		Name: to.Ptr("testCluster"),
	// 		Type: to.Ptr("Microsoft.EventHub/Clusters"),
	// 		ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/myResourceGroup/providers/Microsoft.EventHub/clusters/testCluster"),
	// 		Location: to.Ptr("South Central US"),
	// 		Properties: &armeventhub.ClusterProperties{
	// 			CreatedAt: to.Ptr("2017-05-24T23:23:27.877Z"),
	// 			MetricID: to.Ptr("SN6-008"),
	// 			ProvisioningState: to.Ptr(armeventhub.ProvisioningStateSucceeded),
	// 			UpdatedAt: to.Ptr("2017-05-24T23:23:27.877Z"),
	// 		},
	// 		SKU: &armeventhub.ClusterSKU{
	// 			Name: to.Ptr(armeventhub.ClusterSKUNameDedicated),
	// 			Capacity: to.Ptr[int32](1),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
