package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ListClusterInstanceViews.json
func ExampleClustersClient_NewListInstanceViewsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListInstanceViewsPager("rg1", "clusterPool1", "cluster1", nil)
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
		// page.ClusterInstanceViewsResult = armhdinsightcontainers.ClusterInstanceViewsResult{
		// 	Value: []*armhdinsightcontainers.ClusterInstanceViewResult{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Properties: &armhdinsightcontainers.ClusterInstanceViewResultProperties{
		// 				ServiceStatuses: []*armhdinsightcontainers.ServiceStatus{
		// 					{
		// 						Kind: to.Ptr("ZooKeeper"),
		// 						Message: to.Ptr(""),
		// 						Ready: to.Ptr("true"),
		// 				}},
		// 				Status: &armhdinsightcontainers.ClusterInstanceViewPropertiesStatus{
		// 					Message: to.Ptr(""),
		// 					Ready: to.Ptr("True"),
		// 					Reason: to.Ptr(""),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
