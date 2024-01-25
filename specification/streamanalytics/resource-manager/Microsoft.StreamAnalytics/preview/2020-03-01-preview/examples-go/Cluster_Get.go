package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/Cluster_Get.json
func ExampleClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().Get(ctx, "sjrg", "testcluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armstreamanalytics.Cluster{
	// 	Name: to.Ptr("An Example Cluster"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/clusters"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/sjrg/providers/Microsoft.StreamAnalytics/clusters/AnExampleStreamingCluster"),
	// 	Location: to.Ptr("Central US"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Etag: to.Ptr("F86B9B70-D5B1-451D-AFC8-0B42D4729B8C"),
	// 	Properties: &armstreamanalytics.ClusterProperties{
	// 		CapacityAllocated: to.Ptr[int32](48),
	// 		CapacityAssigned: to.Ptr[int32](96),
	// 		ClusterID: to.Ptr("B01C67EF-4739-4DDD-9FB2-427EB43DE839"),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-25T01:00:00.000Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armstreamanalytics.ClusterProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armstreamanalytics.ClusterSKU{
	// 		Name: to.Ptr(armstreamanalytics.ClusterSKUNameDefault),
	// 		Capacity: to.Ptr[int32](96),
	// 	},
	// }
}
