package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/Clusters/ClusterPatch.json
func ExampleClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdate(ctx, "myResourceGroup", "testCluster", armeventhub.Cluster{
		Location: to.Ptr("South Central US"),
		Tags: map[string]*string{
			"tag3": to.Ptr("value3"),
			"tag4": to.Ptr("value4"),
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
	// res.Cluster = armeventhub.Cluster{
	// 	Name: to.Ptr("testCluster"),
	// 	Type: to.Ptr("Microsoft.EventHub/Clusters"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/myResourceGroup/providers/Microsoft.EventHub/clusters/testCluster"),
	// 	Location: to.Ptr("South Central US"),
	// 	Tags: map[string]*string{
	// 		"tag3": to.Ptr("value3"),
	// 		"tag4": to.Ptr("value4"),
	// 	},
	// 	Properties: &armeventhub.ClusterProperties{
	// 		CreatedAt: to.Ptr("2017-06-01T21:37:04.46Z"),
	// 		MetricID: to.Ptr("SN6-008"),
	// 		ProvisioningState: to.Ptr(armeventhub.ProvisioningStateSucceeded),
	// 		UpdatedAt: to.Ptr("2017-06-01T21:37:53.413Z"),
	// 	},
	// 	SKU: &armeventhub.ClusterSKU{
	// 		Name: to.Ptr(armeventhub.ClusterSKUNameDedicated),
	// 		Capacity: to.Ptr[int32](4),
	// 	},
	// }
}
