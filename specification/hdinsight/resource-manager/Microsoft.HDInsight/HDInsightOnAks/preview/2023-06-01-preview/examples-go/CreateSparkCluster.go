package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7f70e351393addbc31d790a908c994c7c8644d9c/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-06-01-preview/examples/CreateSparkCluster.json
func ExampleClustersClient_BeginCreate_hdInsightSparkClusterPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreate(ctx, "hiloResourcegroup", "clusterpool1", "cluster1", armhdinsightcontainers.Cluster{
		Location: to.Ptr("West US 2"),
		Properties: &armhdinsightcontainers.ClusterResourceProperties{
			ClusterProfile: &armhdinsightcontainers.ClusterProfile{
				AuthorizationProfile: &armhdinsightcontainers.AuthorizationProfile{
					UserIDs: []*string{
						to.Ptr("testuser1"),
						to.Ptr("testuser2")},
				},
				ClusterVersion: to.Ptr("0.0.1"),
				IdentityProfile: &armhdinsightcontainers.IdentityProfile{
					MsiClientID:   to.Ptr("de91f1d8-767f-460a-ac11-3cf103f74b34"),
					MsiObjectID:   to.Ptr("40491351-c240-4042-91e0-f644a1d2b441"),
					MsiResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-msi"),
				},
				OssVersion: to.Ptr("2.2.3"),
				ServiceConfigsProfiles: []*armhdinsightcontainers.ClusterServiceConfigsProfile{
					{
						Configs: []*armhdinsightcontainers.ClusterServiceConfig{
							{
								Component: to.Ptr("spark-config"),
								Files: []*armhdinsightcontainers.ClusterConfigFile{
									{
										FileName: to.Ptr("spark-defaults.conf"),
										Values: map[string]*string{
											"spark.eventLog.enabled": to.Ptr("true"),
										},
									}},
							}},
						ServiceName: to.Ptr("spark-service"),
					},
					{
						Configs: []*armhdinsightcontainers.ClusterServiceConfig{
							{
								Component: to.Ptr("yarn-config"),
								Files: []*armhdinsightcontainers.ClusterConfigFile{
									{
										FileName: to.Ptr("core-site.xml"),
										Values: map[string]*string{
											"fs.defaultFS":      to.Ptr("wasb://testcontainer@teststorage.dfs.core.windows.net/"),
											"storage.container": to.Ptr("testcontainer"),
											"storage.key":       to.Ptr("test key"),
											"storage.name":      to.Ptr("teststorage"),
											"storage.protocol":  to.Ptr("wasb"),
										},
									},
									{
										FileName: to.Ptr("yarn-site.xml"),
										Values: map[string]*string{
											"yarn.webapp.ui2.enable": to.Ptr("false"),
										},
									}},
							}},
						ServiceName: to.Ptr("yarn-service"),
					}},
				SparkProfile: &armhdinsightcontainers.SparkProfile{},
				SSHProfile: &armhdinsightcontainers.SSHProfile{
					Count: to.Ptr[int32](2),
				},
			},
			ClusterType: to.Ptr("spark"),
			ComputeProfile: &armhdinsightcontainers.ComputeProfile{
				Nodes: []*armhdinsightcontainers.NodeProfile{
					{
						Type:   to.Ptr("worker"),
						Count:  to.Ptr[int32](4),
						VMSize: to.Ptr("Standard_D3_v2"),
					}},
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
	// res.Cluster = armhdinsightcontainers.Cluster{
	// 	Name: to.Ptr("cluster1"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusterPools/clusters"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1/clusters/cluster1"),
	// 	SystemData: &armhdinsightcontainers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("West US 2"),
	// 	Properties: &armhdinsightcontainers.ClusterResourceProperties{
	// 		ClusterProfile: &armhdinsightcontainers.ClusterProfile{
	// 			AuthorizationProfile: &armhdinsightcontainers.AuthorizationProfile{
	// 				UserIDs: []*string{
	// 					to.Ptr("testuser1"),
	// 					to.Ptr("testuser2")},
	// 				},
	// 				ClusterVersion: to.Ptr("0.0.1"),
	// 				Components: []*armhdinsightcontainers.ClusterComponentsItem{
	// 					{
	// 						Name: to.Ptr("HDFS"),
	// 						Version: to.Ptr("2.2.3"),
	// 				}},
	// 				ConnectivityProfile: &armhdinsightcontainers.ConnectivityProfile{
	// 					SSH: []*armhdinsightcontainers.SSHConnectivityEndpoint{
	// 						{
	// 							Endpoint: to.Ptr("cluster1.clusterpool1.westus2.projecthilo.net/ssh/host/sshnode-0"),
	// 						},
	// 						{
	// 							Endpoint: to.Ptr("cluster1.clusterpool1.westus2.projecthilo.net/ssh/host/sshnode-1"),
	// 					}},
	// 					Web: &armhdinsightcontainers.ConnectivityProfileWeb{
	// 						Fqdn: to.Ptr("cluster1.clusterpool1.westus2.projecthilo.net"),
	// 					},
	// 				},
	// 				IdentityProfile: &armhdinsightcontainers.IdentityProfile{
	// 					MsiClientID: to.Ptr("de91f1d8-767f-460a-ac11-3cf103f74b34"),
	// 					MsiObjectID: to.Ptr("40491351-c240-4042-91e0-f644a1d2b441"),
	// 					MsiResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-msi"),
	// 				},
	// 				OssVersion: to.Ptr("2.2.3"),
	// 				ServiceConfigsProfiles: []*armhdinsightcontainers.ClusterServiceConfigsProfile{
	// 					{
	// 						Configs: []*armhdinsightcontainers.ClusterServiceConfig{
	// 							{
	// 								Component: to.Ptr("spark-config"),
	// 								Files: []*armhdinsightcontainers.ClusterConfigFile{
	// 									{
	// 										FileName: to.Ptr("spark-defaults.conf"),
	// 										Values: map[string]*string{
	// 											"spark.eventLog.enabled": to.Ptr("true"),
	// 										},
	// 								}},
	// 						}},
	// 						ServiceName: to.Ptr("spark-service"),
	// 					},
	// 					{
	// 						Configs: []*armhdinsightcontainers.ClusterServiceConfig{
	// 							{
	// 								Component: to.Ptr("yarn-config"),
	// 								Files: []*armhdinsightcontainers.ClusterConfigFile{
	// 									{
	// 										FileName: to.Ptr("core-site.xml"),
	// 										Values: map[string]*string{
	// 											"fs.defaultFS": to.Ptr("wasb://testcontainer@teststorage.dfs.core.windows.net/"),
	// 											"storage.container": to.Ptr("testcontainer"),
	// 											"storage.key": to.Ptr("test key"),
	// 											"storage.name": to.Ptr("teststorage"),
	// 											"storage.protocol": to.Ptr("wasb"),
	// 										},
	// 									},
	// 									{
	// 										FileName: to.Ptr("yarn-site.xml"),
	// 										Values: map[string]*string{
	// 											"yarn.webapp.ui2.enable": to.Ptr("false"),
	// 										},
	// 								}},
	// 						}},
	// 						ServiceName: to.Ptr("yarn-service"),
	// 				}},
	// 				SparkProfile: &armhdinsightcontainers.SparkProfile{
	// 				},
	// 				SSHProfile: &armhdinsightcontainers.SSHProfile{
	// 					Count: to.Ptr[int32](2),
	// 					PodPrefix: to.Ptr("sshnode"),
	// 				},
	// 			},
	// 			ClusterType: to.Ptr("spark"),
	// 			ComputeProfile: &armhdinsightcontainers.ComputeProfile{
	// 				Nodes: []*armhdinsightcontainers.NodeProfile{
	// 					{
	// 						Type: to.Ptr("worker"),
	// 						Count: to.Ptr[int32](4),
	// 						VMSize: to.Ptr("Standard_D3_v2"),
	// 				}},
	// 			},
	// 			DeploymentID: to.Ptr("45cd32aead6e4a91b079a0cdbfac8c36"),
	// 			ProvisioningState: to.Ptr(armhdinsightcontainers.ProvisioningStatusSucceeded),
	// 		},
	// 	}
}
