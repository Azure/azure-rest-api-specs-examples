package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ListClusterUpgradeHistory.json
func ExampleClusterUpgradeHistoriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterUpgradeHistoriesClient().NewListPager("hiloResourcegroup", "clusterpool1", "cluster1", nil)
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
		// page.ClusterUpgradeHistoryListResult = armhdinsightcontainers.ClusterUpgradeHistoryListResult{
		// 	Value: []*armhdinsightcontainers.ClusterUpgradeHistory{
		// 		{
		// 			Name: to.Ptr("upgradeDefinitionCR1"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1/clusters/cluster1/upgradeHistories/upgradeDefinitionCR1"),
		// 			Properties: &armhdinsightcontainers.ClusterHotfixUpgradeHistoryProperties{
		// 				UpgradeResult: to.Ptr(armhdinsightcontainers.ClusterUpgradeHistoryUpgradeResultTypeSucceed),
		// 				UpgradeType: to.Ptr(armhdinsightcontainers.ClusterUpgradeHistoryTypeHotfixUpgrade),
		// 				UTCTime: to.Ptr("01/18/2024 03:18:38 AM"),
		// 				ComponentName: to.Ptr("Zookeeper"),
		// 				Severity: to.Ptr(armhdinsightcontainers.ClusterUpgradeHistorySeverityTypeLow),
		// 				SourceBuildNumber: to.Ptr("1"),
		// 				SourceClusterVersion: to.Ptr("1.0.1"),
		// 				SourceOssVersion: to.Ptr("2.4.1"),
		// 				TargetBuildNumber: to.Ptr("2"),
		// 				TargetClusterVersion: to.Ptr("1.0.1"),
		// 				TargetOssVersion: to.Ptr("2.4.1"),
		// 			},
		// 	}},
		// }
	}
}
