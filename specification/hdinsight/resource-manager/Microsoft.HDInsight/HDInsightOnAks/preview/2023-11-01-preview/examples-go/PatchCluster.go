package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-11-01-preview/examples/PatchCluster.json
func ExampleClustersClient_BeginUpdate_hdInsightClustersPatchTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdate(ctx, "hiloResourcegroup", "clusterpool1", "cluster1", armhdinsightcontainers.ClusterPatch{
		Properties: &armhdinsightcontainers.ClusterPatchProperties{
			ClusterProfile: &armhdinsightcontainers.UpdatableClusterProfile{
				AuthorizationProfile: &armhdinsightcontainers.AuthorizationProfile{
					UserIDs: []*string{
						to.Ptr("Testuser1"),
						to.Ptr("Testuser2")},
				},
				AutoscaleProfile: &armhdinsightcontainers.AutoscaleProfile{
					AutoscaleType:               to.Ptr(armhdinsightcontainers.AutoscaleTypeScheduleBased),
					Enabled:                     to.Ptr(true),
					GracefulDecommissionTimeout: to.Ptr[int32](-1),
					ScheduleBasedConfig: &armhdinsightcontainers.ScheduleBasedConfig{
						DefaultCount: to.Ptr[int32](3),
						Schedules: []*armhdinsightcontainers.Schedule{
							{
								Count: to.Ptr[int32](3),
								Days: []*armhdinsightcontainers.ScheduleDay{
									to.Ptr(armhdinsightcontainers.ScheduleDay("Monday, Tuesday, Wednesday"))},
								EndTime:   to.Ptr("12:00"),
								StartTime: to.Ptr("00:00"),
							},
							{
								Count: to.Ptr[int32](3),
								Days: []*armhdinsightcontainers.ScheduleDay{
									to.Ptr(armhdinsightcontainers.ScheduleDaySunday)},
								EndTime:   to.Ptr("12:00"),
								StartTime: to.Ptr("00:00"),
							}},
						TimeZone: to.Ptr("Cen. Australia Standard Time"),
					},
				},
				LogAnalyticsProfile: &armhdinsightcontainers.ClusterLogAnalyticsProfile{
					ApplicationLogs: &armhdinsightcontainers.ClusterLogAnalyticsApplicationLogs{
						StdErrorEnabled: to.Ptr(true),
						StdOutEnabled:   to.Ptr(true),
					},
					Enabled:        to.Ptr(true),
					MetricsEnabled: to.Ptr(true),
				},
				ServiceConfigsProfiles: []*armhdinsightcontainers.ClusterServiceConfigsProfile{
					{
						Configs: []*armhdinsightcontainers.ClusterServiceConfig{
							{
								Component: to.Ptr("TestComp1"),
								Files: []*armhdinsightcontainers.ClusterConfigFile{
									{
										FileName: to.Ptr("TestFile1"),
										Values: map[string]*string{
											"Test.config.1": to.Ptr("1"),
											"Test.config.2": to.Ptr("2"),
										},
									},
									{
										FileName: to.Ptr("TestFile2"),
										Values: map[string]*string{
											"Test.config.3": to.Ptr("3"),
											"Test.config.4": to.Ptr("4"),
										},
									}},
							},
							{
								Component: to.Ptr("TestComp2"),
								Files: []*armhdinsightcontainers.ClusterConfigFile{
									{
										Path:     to.Ptr("TestPath"),
										Content:  to.Ptr("TestContent"),
										FileName: to.Ptr("TestFile3"),
									},
									{
										FileName: to.Ptr("TestFile4"),
										Values: map[string]*string{
											"Test.config.7": to.Ptr("7"),
											"Test.config.8": to.Ptr("8"),
										},
									}},
							}},
						ServiceName: to.Ptr("TestService1"),
					},
					{
						Configs: []*armhdinsightcontainers.ClusterServiceConfig{
							{
								Component: to.Ptr("TestComp3"),
								Files: []*armhdinsightcontainers.ClusterConfigFile{
									{
										FileName: to.Ptr("TestFile5"),
										Values: map[string]*string{
											"Test.config.9": to.Ptr("9"),
										},
									}},
							}},
						ServiceName: to.Ptr("TestService2"),
					}},
				SSHProfile: &armhdinsightcontainers.SSHProfile{
					Count: to.Ptr[int32](2),
				},
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
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
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
	// 					GracefulDecommissionTimeout: to.Ptr[int32](-1),
	// 					ScheduleBasedConfig: &armhdinsightcontainers.ScheduleBasedConfig{
	// 						DefaultCount: to.Ptr[int32](3),
	// 						Schedules: []*armhdinsightcontainers.Schedule{
	// 							{
	// 								Count: to.Ptr[int32](3),
	// 								Days: []*armhdinsightcontainers.ScheduleDay{
	// 									to.Ptr(armhdinsightcontainers.ScheduleDay("Monday, Tuesday, Wednesday"))},
	// 									EndTime: to.Ptr("12:00"),
	// 									StartTime: to.Ptr("00:00"),
	// 								},
	// 								{
	// 									Count: to.Ptr[int32](3),
	// 									Days: []*armhdinsightcontainers.ScheduleDay{
	// 										to.Ptr(armhdinsightcontainers.ScheduleDaySunday)},
	// 										EndTime: to.Ptr("12:00"),
	// 										StartTime: to.Ptr("00:00"),
	// 								}},
	// 								TimeZone: to.Ptr("Cen. Australia Standard Time"),
	// 							},
	// 						},
	// 						ClusterVersion: to.Ptr("1.0.6"),
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
	// 						LogAnalyticsProfile: &armhdinsightcontainers.ClusterLogAnalyticsProfile{
	// 							ApplicationLogs: &armhdinsightcontainers.ClusterLogAnalyticsApplicationLogs{
	// 								StdErrorEnabled: to.Ptr(true),
	// 								StdOutEnabled: to.Ptr(true),
	// 							},
	// 							Enabled: to.Ptr(true),
	// 							MetricsEnabled: to.Ptr(true),
	// 						},
	// 						OssVersion: to.Ptr("0.410.0"),
	// 						ServiceConfigsProfiles: []*armhdinsightcontainers.ClusterServiceConfigsProfile{
	// 							{
	// 								Configs: []*armhdinsightcontainers.ClusterServiceConfig{
	// 									{
	// 										Component: to.Ptr("TestComp1"),
	// 										Files: []*armhdinsightcontainers.ClusterConfigFile{
	// 											{
	// 												FileName: to.Ptr("TestFile1"),
	// 												Values: map[string]*string{
	// 													"Test.config.1": to.Ptr("1"),
	// 													"Test.config.2": to.Ptr("2"),
	// 												},
	// 											},
	// 											{
	// 												FileName: to.Ptr("TestFile2"),
	// 												Values: map[string]*string{
	// 													"Test.config.3": to.Ptr("3"),
	// 													"Test.config.4": to.Ptr("4"),
	// 												},
	// 										}},
	// 									},
	// 									{
	// 										Component: to.Ptr("TestComp2"),
	// 										Files: []*armhdinsightcontainers.ClusterConfigFile{
	// 											{
	// 												Path: to.Ptr("TestPath"),
	// 												Content: to.Ptr("TestContent"),
	// 												FileName: to.Ptr("TestFile3"),
	// 											},
	// 											{
	// 												FileName: to.Ptr("TestFile4"),
	// 												Values: map[string]*string{
	// 													"Test.config.7": to.Ptr("7"),
	// 													"Test.config.8": to.Ptr("8"),
	// 												},
	// 										}},
	// 								}},
	// 								ServiceName: to.Ptr("TestService1"),
	// 							},
	// 							{
	// 								Configs: []*armhdinsightcontainers.ClusterServiceConfig{
	// 									{
	// 										Component: to.Ptr("TestComp3"),
	// 										Files: []*armhdinsightcontainers.ClusterConfigFile{
	// 											{
	// 												FileName: to.Ptr("TestFile5"),
	// 												Values: map[string]*string{
	// 													"Test.config.9": to.Ptr("9"),
	// 												},
	// 										}},
	// 								}},
	// 								ServiceName: to.Ptr("TestService2"),
	// 						}},
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
	// 			}
}
