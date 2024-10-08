package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ListClusterPoolAvailableUpgrades.json
func ExampleClusterPoolAvailableUpgradesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterPoolAvailableUpgradesClient().NewListPager("hiloResourcegroup", "clusterpool1", nil)
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
		// page.ClusterPoolAvailableUpgradeList = armhdinsightcontainers.ClusterPoolAvailableUpgradeList{
		// 	Value: []*armhdinsightcontainers.ClusterPoolAvailableUpgrade{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1/availableUpgrades/AKSPatchUpgrade"),
		// 			Properties: &armhdinsightcontainers.ClusterPoolAvailableUpgradeAksPatchUpgradeProperties{
		// 				UpgradeType: to.Ptr(armhdinsightcontainers.ClusterPoolAvailableUpgradeTypeAKSPatchUpgrade),
		// 				CurrentVersion: to.Ptr("1.26.3"),
		// 				CurrentVersionStatus: to.Ptr(armhdinsightcontainers.CurrentClusterPoolAksVersionStatusDeprecated),
		// 				LatestVersion: to.Ptr("1.26.6"),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1/availableUpgrades/NodeOsUpgrade"),
		// 			Properties: &armhdinsightcontainers.ClusterPoolAvailableUpgradeNodeOsUpgradeProperties{
		// 				UpgradeType: to.Ptr(armhdinsightcontainers.ClusterPoolAvailableUpgradeTypeNodeOsUpgrade),
		// 				LatestVersion: to.Ptr("AKSCBLMariner-V2gen2-202310.09.0"),
		// 			},
		// 	}},
		// }
	}
}
