package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-11-01-preview/examples/ListClustersByClusterPoolName.json
func ExampleClustersClient_NewListByClusterPoolNamePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListByClusterPoolNamePager("hiloResourcegroup", "clusterpool1", nil)
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
		// page.ClusterListResult = armhdinsightcontainers.ClusterListResult{
		// 	Value: []*armhdinsightcontainers.Cluster{
		// 		{
		// 			Name: to.Ptr("cluster1"),
		// 			Type: to.Ptr("Microsoft.HDInsight/clusterPools/clusters"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1/clusters/cluster1"),
		// 			Location: to.Ptr("West US 2"),
		// 			Tags: map[string]*string{
		// 				"company": to.Ptr("Contoso"),
		// 				"department": to.Ptr("MightyMight"),
		// 			},
		// 			Properties: &armhdinsightcontainers.ClusterResourceProperties{
		// 				ClusterProfile: &armhdinsightcontainers.ClusterProfile{
		// 					AuthorizationProfile: &armhdinsightcontainers.AuthorizationProfile{
		// 						UserIDs: []*string{
		// 							to.Ptr("testuser1"),
		// 							to.Ptr("testuser2")},
		// 						},
		// 						ClusterVersion: to.Ptr("1.0.6"),
		// 						Components: []*armhdinsightcontainers.ClusterComponentsItem{
		// 							{
		// 								Name: to.Ptr("Trino"),
		// 								Version: to.Ptr("410"),
		// 							},
		// 							{
		// 								Name: to.Ptr("Hive metastore"),
		// 								Version: to.Ptr("3.1.2"),
		// 						}},
		// 						ConnectivityProfile: &armhdinsightcontainers.ConnectivityProfile{
		// 							SSH: []*armhdinsightcontainers.SSHConnectivityEndpoint{
		// 								{
		// 									Endpoint: to.Ptr("cluster1.clusterpool1.westus2.projecthilo.net/ssh/host/sshnode-0"),
		// 								},
		// 								{
		// 									Endpoint: to.Ptr("cluster1.clusterpool1.westus2.projecthilo.net/ssh/host/sshnode-1"),
		// 							}},
		// 							Web: &armhdinsightcontainers.ConnectivityProfileWeb{
		// 								Fqdn: to.Ptr("cluster1.clusterpool1.westus2.projecthilo.net"),
		// 							},
		// 						},
		// 						IdentityProfile: &armhdinsightcontainers.IdentityProfile{
		// 							MsiClientID: to.Ptr("de91f1d8-767f-460a-ac11-3cf103f74b34"),
		// 							MsiObjectID: to.Ptr("40491351-c240-4042-91e0-f644a1d2b441"),
		// 							MsiResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-msi"),
		// 						},
		// 						OssVersion: to.Ptr("0.410.0"),
		// 						SSHProfile: &armhdinsightcontainers.SSHProfile{
		// 							Count: to.Ptr[int32](2),
		// 							PodPrefix: to.Ptr("sshnode"),
		// 						},
		// 						TrinoProfile: &armhdinsightcontainers.TrinoProfile{
		// 						},
		// 					},
		// 					ClusterType: to.Ptr("Trino"),
		// 					ComputeProfile: &armhdinsightcontainers.ComputeProfile{
		// 						Nodes: []*armhdinsightcontainers.NodeProfile{
		// 							{
		// 								Type: to.Ptr("Head"),
		// 								Count: to.Ptr[int32](2),
		// 								VMSize: to.Ptr("Standard_E8as_v5"),
		// 							},
		// 							{
		// 								Type: to.Ptr("Worker"),
		// 								Count: to.Ptr[int32](3),
		// 								VMSize: to.Ptr("Standard_E8as_v5"),
		// 						}},
		// 					},
		// 					DeploymentID: to.Ptr("45cd32aead6e4a91b079a0cdbfac8c36"),
		// 					ProvisioningState: to.Ptr(armhdinsightcontainers.ProvisioningStatusSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}
