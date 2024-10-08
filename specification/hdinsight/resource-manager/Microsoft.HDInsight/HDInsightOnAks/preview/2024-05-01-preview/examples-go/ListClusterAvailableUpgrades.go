package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ListClusterAvailableUpgrades.json
func ExampleClusterAvailableUpgradesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterAvailableUpgradesClient().NewListPager("hiloResourcegroup", "clusterpool1", "cluster1", nil)
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
		// page.ClusterAvailableUpgradeList = armhdinsightcontainers.ClusterAvailableUpgradeList{
		// 	Value: []*armhdinsightcontainers.ClusterAvailableUpgrade{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1/clusters/cluster1/availableUpgrades/AKSPatchUpgrade"),
		// 			Properties: &armhdinsightcontainers.ClusterAvailableUpgradeAksPatchUpgradeProperties{
		// 				UpgradeType: to.Ptr(armhdinsightcontainers.ClusterAvailableUpgradeTypeAKSPatchUpgrade),
		// 				CurrentVersion: to.Ptr("1.26.3"),
		// 				CurrentVersionStatus: to.Ptr(armhdinsightcontainers.CurrentClusterAksVersionStatusSupported),
		// 				LatestVersion: to.Ptr("1.26.6"),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1/clusters/cluster1/availableUpgrades/hotfix1"),
		// 			Properties: &armhdinsightcontainers.ClusterAvailableUpgradeHotfixUpgradeProperties{
		// 				UpgradeType: to.Ptr(armhdinsightcontainers.ClusterAvailableUpgradeTypeHotfixUpgrade),
		// 				Description: to.Ptr("Hotfix for historyserver on version 1.16.0-1.0.6.2"),
		// 				ComponentName: to.Ptr("historyserver"),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-29T14:13:12.000Z"); return t}()),
		// 				ExtendedProperties: to.Ptr(""),
		// 				Severity: to.Ptr(armhdinsightcontainers.SeverityLow),
		// 				SourceBuildNumber: to.Ptr("2"),
		// 				SourceClusterVersion: to.Ptr("1.0.6"),
		// 				SourceOssVersion: to.Ptr("1.16.0"),
		// 				TargetBuildNumber: to.Ptr("3"),
		// 				TargetClusterVersion: to.Ptr("1.0.6"),
		// 				TargetOssVersion: to.Ptr("1.16.0"),
		// 			},
		// 	}},
		// }
	}
}
