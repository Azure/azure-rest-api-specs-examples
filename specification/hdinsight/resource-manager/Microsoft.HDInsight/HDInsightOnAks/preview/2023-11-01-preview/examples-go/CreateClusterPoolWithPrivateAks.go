package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-11-01-preview/examples/CreateClusterPoolWithPrivateAks.json
func ExampleClusterPoolsClient_BeginCreateOrUpdate_clusterPoolPutWithPrivateAks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterPoolsClient().BeginCreateOrUpdate(ctx, "hiloResourcegroup", "clusterpool1", armhdinsightcontainers.ClusterPool{
		Location: to.Ptr("West US 2"),
		Properties: &armhdinsightcontainers.ClusterPoolResourceProperties{
			ClusterPoolProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesClusterPoolProfile{
				ClusterPoolVersion: to.Ptr("1.2"),
			},
			ComputeProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesComputeProfile{
				VMSize: to.Ptr("Standard_D3_v2"),
			},
			NetworkProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesNetworkProfile{
				EnablePrivateAPIServer: to.Ptr(true),
				SubnetID:               to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
			},
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
	// res.ClusterPool = armhdinsightcontainers.ClusterPool{
	// 	Name: to.Ptr("clusterpool1"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusterPools"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1"),
	// 	SystemData: &armhdinsightcontainers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("West US 2"),
	// 	Properties: &armhdinsightcontainers.ClusterPoolResourceProperties{
	// 		AksClusterProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesAksClusterProfile{
	// 			AksClusterAgentPoolIdentityProfile: &armhdinsightcontainers.AksClusterProfileAksClusterAgentPoolIdentityProfile{
	// 				MsiClientID: to.Ptr("a89fb478-2a84-4d9b-8f18-3e8c4d1db3eb"),
	// 				MsiObjectID: to.Ptr("dc7ef861-8b55-4ffb-9003-20885cd895a9"),
	// 				MsiResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ManagedIdentity/userAssignedIdentities/clusterpool1-agentpool"),
	// 			},
	// 			AksClusterResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ContainerService/managedClusters/clusterpool1"),
	// 			AksVersion: to.Ptr("1.24"),
	// 		},
	// 		ClusterPoolProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesClusterPoolProfile{
	// 			ClusterPoolVersion: to.Ptr("1.2"),
	// 		},
	// 		ComputeProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesComputeProfile{
	// 			Count: to.Ptr[int32](3),
	// 			VMSize: to.Ptr("Standard_D3_v2"),
	// 		},
	// 		DeploymentID: to.Ptr("45cd32aead6e4a91b079a0cdbfac8c36"),
	// 		ManagedResourceGroupName: to.Ptr("hdi-45cd32aead6e4a91b079a0cdbfac8c36"),
	// 		NetworkProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesNetworkProfile{
	// 			EnablePrivateAPIServer: to.Ptr(true),
	// 			SubnetID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 		},
	// 		ProvisioningState: to.Ptr(armhdinsightcontainers.ProvisioningStatusSucceeded),
	// 	},
	// }
}
