package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ClusterUpgradeRollback.json
func ExampleClustersClient_BeginUpgradeManualRollback() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpgradeManualRollback(ctx, "hiloResourcegroup", "clusterpool1", "cluster1", armhdinsightcontainers.ClusterUpgradeRollback{
		Properties: &armhdinsightcontainers.ClusterUpgradeRollbackProperties{
			UpgradeHistory: to.Ptr("/subscriptions/10e32bab-26da-4cc4-a441-52b318f824e6/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterpools/clusterpool1/clusters/cluster1/upgradeHistories/01_11_2024_02_35_03_AM-HotfixUpgrade"),
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
	// 				AutoscaleProfile: &armhdinsightcontainers.AutoscaleProfile{
	// 					AutoscaleType: to.Ptr(armhdinsightcontainers.AutoscaleTypeScheduleBased),
	// 					Enabled: to.Ptr(true),
	// 					GracefulDecommissionTimeout: to.Ptr[int32](3600),
	// 					LoadBasedConfig: &armhdinsightcontainers.LoadBasedConfig{
	// 						CooldownPeriod: to.Ptr[int32](300),
	// 						MaxNodes: to.Ptr[int32](20),
	// 						MinNodes: to.Ptr[int32](10),
	// 						PollInterval: to.Ptr[int32](60),
	// 						ScalingRules: []*armhdinsightcontainers.ScalingRule{
	// 							{
	// 								ActionType: to.Ptr(armhdinsightcontainers.ScaleActionTypeScaleup),
	// 								ComparisonRule: &armhdinsightcontainers.ComparisonRule{
	// 									Operator: to.Ptr(armhdinsightcontainers.ComparisonOperatorGreaterThan),
	// 									Threshold: to.Ptr[float32](90),
	// 								},
	// 								EvaluationCount: to.Ptr[int32](3),
	// 								ScalingMetric: to.Ptr("cpu"),
	// 							},
	// 							{
	// 								ActionType: to.Ptr(armhdinsightcontainers.ScaleActionTypeScaledown),
	// 								ComparisonRule: &armhdinsightcontainers.ComparisonRule{
	// 									Operator: to.Ptr(armhdinsightcontainers.ComparisonOperatorLessThan),
	// 									Threshold: to.Ptr[float32](20),
	// 								},
	// 								EvaluationCount: to.Ptr[int32](3),
	// 								ScalingMetric: to.Ptr("cpu"),
	// 						}},
	// 					},
	// 					ScheduleBasedConfig: &armhdinsightcontainers.ScheduleBasedConfig{
	// 						DefaultCount: to.Ptr[int32](10),
	// 						Schedules: []*armhdinsightcontainers.Schedule{
	// 							{
	// 								Count: to.Ptr[int32](20),
	// 								Days: []*armhdinsightcontainers.ScheduleDay{
	// 									to.Ptr(armhdinsightcontainers.ScheduleDayMonday)},
	// 									EndTime: to.Ptr("12:00"),
	// 									StartTime: to.Ptr("00:00"),
	// 								},
	// 								{
	// 									Count: to.Ptr[int32](25),
	// 									Days: []*armhdinsightcontainers.ScheduleDay{
	// 										to.Ptr(armhdinsightcontainers.ScheduleDaySunday)},
	// 										EndTime: to.Ptr("12:00"),
	// 										StartTime: to.Ptr("00:00"),
	// 								}},
	// 								TimeZone: to.Ptr("Cen. Australia Standard Time"),
	// 							},
	// 						},
	// 						ClusterVersion: to.Ptr("1.0.1"),
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
	// 						OssVersion: to.Ptr("2.4.1"),
	// 						SSHProfile: &armhdinsightcontainers.SSHProfile{
	// 							Count: to.Ptr[int32](2),
	// 							PodPrefix: to.Ptr("sshnode"),
	// 						},
	// 						TrinoProfile: &armhdinsightcontainers.TrinoProfile{
	// 						},
	// 					},
	// 					ClusterType: to.Ptr("trino"),
	// 					ComputeProfile: &armhdinsightcontainers.ComputeProfile{
	// 						AvailabilityZones: []*string{
	// 							to.Ptr("1"),
	// 							to.Ptr("2"),
	// 							to.Ptr("3")},
	// 							Nodes: []*armhdinsightcontainers.NodeProfile{
	// 								{
	// 									Type: to.Ptr("worker"),
	// 									Count: to.Ptr[int32](3),
	// 									VMSize: to.Ptr("Standard_D3_v2"),
	// 							}},
	// 						},
	// 						DeploymentID: to.Ptr("45cd32aead6e4a91b079a0cdbfac8c36"),
	// 						ProvisioningState: to.Ptr(armhdinsightcontainers.ProvisioningStatusSucceeded),
	// 					},
	// 				}
}
