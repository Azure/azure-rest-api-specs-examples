package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-11-01-preview/examples/ListClusterPoolsSubscription.json
func ExampleClusterPoolsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterPoolsClient().NewListBySubscriptionPager(nil)
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
		// page.ClusterPoolListResult = armhdinsightcontainers.ClusterPoolListResult{
		// 	Value: []*armhdinsightcontainers.ClusterPool{
		// 		{
		// 			Name: to.Ptr("clusterpool1"),
		// 			Type: to.Ptr("Microsoft.HDInsight/clusterPools"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1"),
		// 			Location: to.Ptr("West US 2"),
		// 			Tags: map[string]*string{
		// 				"company": to.Ptr("Contoso"),
		// 				"department": to.Ptr("MightyMight"),
		// 			},
		// 			Properties: &armhdinsightcontainers.ClusterPoolResourceProperties{
		// 				AksClusterProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesAksClusterProfile{
		// 					AksClusterAgentPoolIdentityProfile: &armhdinsightcontainers.AksClusterProfileAksClusterAgentPoolIdentityProfile{
		// 						MsiClientID: to.Ptr("a89fb478-2a84-4d9b-8f18-3e8c4d1db3eb"),
		// 						MsiObjectID: to.Ptr("dc7ef861-8b55-4ffb-9003-20885cd895a9"),
		// 						MsiResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ManagedIdentity/userAssignedIdentities/clusterpool1-agentpool"),
		// 					},
		// 					AksClusterResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ContainerService/managedClusters/clusterpool1"),
		// 					AksVersion: to.Ptr("1.24"),
		// 				},
		// 				ClusterPoolProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesClusterPoolProfile{
		// 					ClusterPoolVersion: to.Ptr("1.2"),
		// 				},
		// 				ComputeProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesComputeProfile{
		// 					Count: to.Ptr[int32](3),
		// 					VMSize: to.Ptr("Standard_D3_v2"),
		// 				},
		// 				DeploymentID: to.Ptr("45cd32aead6e4a91b079a0cdbfac8c36"),
		// 				ManagedResourceGroupName: to.Ptr("hdi-45cd32aead6e4a91b079a0cdbfac8c36"),
		// 				ProvisioningState: to.Ptr(armhdinsightcontainers.ProvisioningStatusSucceeded),
		// 			},
		// 	}},
		// }
	}
}
