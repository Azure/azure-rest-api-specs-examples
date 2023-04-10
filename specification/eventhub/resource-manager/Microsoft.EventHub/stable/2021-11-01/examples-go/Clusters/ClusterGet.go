package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/Clusters/ClusterGet.json
func ExampleClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().Get(ctx, "myResourceGroup", "testCluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armeventhub.Cluster{
	// 	Name: to.Ptr("testCluster"),
	// 	Type: to.Ptr("Microsoft.EventHub/Clusters"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/myResourceGroup/providers/Microsoft.EventHub/clusters/testCluster"),
	// 	Location: to.Ptr("South Central US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armeventhub.ClusterProperties{
	// 		CreatedAt: to.Ptr("2017-05-24T23:23:27.877Z"),
	// 		MetricID: to.Ptr("SN6-008"),
	// 		UpdatedAt: to.Ptr("2017-05-24T23:23:27.877Z"),
	// 	},
	// 	SKU: &armeventhub.ClusterSKU{
	// 		Name: to.Ptr(armeventhub.ClusterSKUNameDedicated),
	// 		Capacity: to.Ptr[int32](4),
	// 	},
	// }
}
