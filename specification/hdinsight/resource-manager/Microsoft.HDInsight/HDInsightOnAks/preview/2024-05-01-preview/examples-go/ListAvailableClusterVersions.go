package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ListAvailableClusterVersions.json
func ExampleAvailableClusterVersionsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailableClusterVersionsClient().NewListByLocationPager("westus2", nil)
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
		// page.ClusterVersionsListResult = armhdinsightcontainers.ClusterVersionsListResult{
		// 	Value: []*armhdinsightcontainers.ClusterVersion{
		// 		{
		// 			Name: to.Ptr("flink_1.16.0-1.0.4"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HDInsight/locations/westus2/availableclusterversions/flink_1.16.0-1.0.4"),
		// 			Properties: &armhdinsightcontainers.ClusterVersionProperties{
		// 				ClusterPoolVersion: to.Ptr("1.0"),
		// 				ClusterType: to.Ptr("Flink"),
		// 				ClusterVersion: to.Ptr("1.0.4"),
		// 				Components: []*armhdinsightcontainers.ClusterComponentsItem{
		// 					{
		// 						Name: to.Ptr("TaskManager"),
		// 						Version: to.Ptr("1.16.0"),
		// 					},
		// 					{
		// 						Name: to.Ptr("JobManager"),
		// 						Version: to.Ptr("1.16.0"),
		// 				}},
		// 				IsPreview: to.Ptr(false),
		// 				OssVersion: to.Ptr("1.16.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("spark_3.3.1-1.0.4"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HDInsight/locations/westus2/availableclusterversions/spark_3.3.1-1.0.4"),
		// 			Properties: &armhdinsightcontainers.ClusterVersionProperties{
		// 				ClusterPoolVersion: to.Ptr("1.0"),
		// 				ClusterType: to.Ptr("Spark"),
		// 				ClusterVersion: to.Ptr("1.0.4"),
		// 				Components: []*armhdinsightcontainers.ClusterComponentsItem{
		// 					{
		// 						Name: to.Ptr("Hadoop"),
		// 						Version: to.Ptr("3.2.3"),
		// 					},
		// 					{
		// 						Name: to.Ptr("Hive"),
		// 						Version: to.Ptr("3.1.4"),
		// 				}},
		// 				IsPreview: to.Ptr(false),
		// 				OssVersion: to.Ptr("3.3.1"),
		// 			},
		// 	}},
		// }
	}
}
