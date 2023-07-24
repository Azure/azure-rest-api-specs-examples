package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/GetLinuxHadoopAllClustersInResourceGroup.json
func ExampleClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.ClusterListResult = armhdinsight.ClusterListResult{
		// 	Value: []*armhdinsight.Cluster{
		// 		{
		// 			Name: to.Ptr("cluster1"),
		// 			Type: to.Ptr("Microsoft.HDInsight/clusters"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HDInsight/clusters/cluster1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("val1"),
		// 			},
		// 			Etag: to.Ptr("f0212a39-b827-45e0-9ffa-4f5232e58851"),
		// 			Properties: &armhdinsight.ClusterGetProperties{
		// 				ClusterDefinition: &armhdinsight.ClusterDefinition{
		// 					Blueprint: to.Ptr("https://blueprints.azurehdinsight.net/hadoop-3.5.1000.0.9243893.json"),
		// 					Kind: to.Ptr("hadoop"),
		// 				},
		// 				ClusterState: to.Ptr("Running"),
		// 				ClusterVersion: to.Ptr("3.5.1000.0"),
		// 				ComputeProfile: &armhdinsight.ComputeProfile{
		// 					Roles: []*armhdinsight.Role{
		// 						{
		// 							Name: to.Ptr("headnode"),
		// 							HardwareProfile: &armhdinsight.HardwareProfile{
		// 								VMSize: to.Ptr("Standard_D3_V2"),
		// 							},
		// 							OSProfile: &armhdinsight.OsProfile{
		// 								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
		// 									Username: to.Ptr("pulkitssh"),
		// 								},
		// 							},
		// 							TargetInstanceCount: to.Ptr[int32](2),
		// 						},
		// 						{
		// 							Name: to.Ptr("workernode"),
		// 							HardwareProfile: &armhdinsight.HardwareProfile{
		// 								VMSize: to.Ptr("Standard_D3_V2"),
		// 							},
		// 							OSProfile: &armhdinsight.OsProfile{
		// 								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
		// 									Username: to.Ptr("pulkitssh"),
		// 								},
		// 							},
		// 							TargetInstanceCount: to.Ptr[int32](4),
		// 						},
		// 						{
		// 							Name: to.Ptr("zookeepernode"),
		// 							HardwareProfile: &armhdinsight.HardwareProfile{
		// 								VMSize: to.Ptr("Medium"),
		// 							},
		// 							OSProfile: &armhdinsight.OsProfile{
		// 								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
		// 									Username: to.Ptr("pulkitssh"),
		// 								},
		// 							},
		// 							TargetInstanceCount: to.Ptr[int32](3),
		// 					}},
		// 				},
		// 				ConnectivityEndpoints: []*armhdinsight.ConnectivityEndpoint{
		// 					{
		// 						Name: to.Ptr("SSH"),
		// 						Location: to.Ptr("cluster1-ssh.azurehdinsight.net"),
		// 						Port: to.Ptr[int32](22),
		// 						Protocol: to.Ptr("TCP"),
		// 					},
		// 					{
		// 						Name: to.Ptr("HTTPS"),
		// 						Location: to.Ptr("cluster1.azurehdinsight.net"),
		// 						Port: to.Ptr[int32](443),
		// 						Protocol: to.Ptr("TCP"),
		// 				}},
		// 				CreatedDate: to.Ptr("2017-01-11T18:58:26.187"),
		// 				OSType: to.Ptr(armhdinsight.OSTypeLinux),
		// 				ProvisioningState: to.Ptr(armhdinsight.HDInsightClusterProvisioningStateSucceeded),
		// 				QuotaInfo: &armhdinsight.QuotaInfo{
		// 					CoresUsed: to.Ptr[int32](24),
		// 				},
		// 				Tier: to.Ptr(armhdinsight.TierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("cluster2"),
		// 			Type: to.Ptr("Microsoft.HDInsight/clusters"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HDInsight/clusters/cluster2"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("val1"),
		// 			},
		// 			Etag: to.Ptr("f0212a39-b827-45e0-9ffa-4f5232e58851"),
		// 			Properties: &armhdinsight.ClusterGetProperties{
		// 				ClusterDefinition: &armhdinsight.ClusterDefinition{
		// 					Blueprint: to.Ptr("https://blueprints.azurehdinsight.net/hadoop-3.5.1000.0.9243893.json"),
		// 					Kind: to.Ptr("hadoop"),
		// 				},
		// 				ClusterState: to.Ptr("Running"),
		// 				ClusterVersion: to.Ptr("3.5.1000.0"),
		// 				ComputeProfile: &armhdinsight.ComputeProfile{
		// 					Roles: []*armhdinsight.Role{
		// 						{
		// 							Name: to.Ptr("headnode"),
		// 							HardwareProfile: &armhdinsight.HardwareProfile{
		// 								VMSize: to.Ptr("Standard_D3_V2"),
		// 							},
		// 							OSProfile: &armhdinsight.OsProfile{
		// 								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
		// 									Username: to.Ptr("pulkitssh"),
		// 								},
		// 							},
		// 							TargetInstanceCount: to.Ptr[int32](2),
		// 						},
		// 						{
		// 							Name: to.Ptr("workernode"),
		// 							HardwareProfile: &armhdinsight.HardwareProfile{
		// 								VMSize: to.Ptr("Standard_D3_V2"),
		// 							},
		// 							OSProfile: &armhdinsight.OsProfile{
		// 								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
		// 									Username: to.Ptr("pulkitssh"),
		// 								},
		// 							},
		// 							TargetInstanceCount: to.Ptr[int32](4),
		// 						},
		// 						{
		// 							Name: to.Ptr("zookeepernode"),
		// 							HardwareProfile: &armhdinsight.HardwareProfile{
		// 								VMSize: to.Ptr("Medium"),
		// 							},
		// 							OSProfile: &armhdinsight.OsProfile{
		// 								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
		// 									Username: to.Ptr("pulkitssh"),
		// 								},
		// 							},
		// 							TargetInstanceCount: to.Ptr[int32](3),
		// 					}},
		// 				},
		// 				ConnectivityEndpoints: []*armhdinsight.ConnectivityEndpoint{
		// 					{
		// 						Name: to.Ptr("SSH"),
		// 						Location: to.Ptr("cluster2-ssh.azurehdinsight.net"),
		// 						Port: to.Ptr[int32](22),
		// 						Protocol: to.Ptr("TCP"),
		// 					},
		// 					{
		// 						Name: to.Ptr("HTTPS"),
		// 						Location: to.Ptr("cluster2.azurehdinsight.net"),
		// 						Port: to.Ptr[int32](443),
		// 						Protocol: to.Ptr("TCP"),
		// 				}},
		// 				CreatedDate: to.Ptr("2017-01-11T18:58:26.187"),
		// 				OSType: to.Ptr(armhdinsight.OSTypeLinux),
		// 				ProvisioningState: to.Ptr(armhdinsight.HDInsightClusterProvisioningStateSucceeded),
		// 				QuotaInfo: &armhdinsight.QuotaInfo{
		// 					CoresUsed: to.Ptr[int32](24),
		// 				},
		// 				Tier: to.Ptr(armhdinsight.TierStandard),
		// 			},
		// 	}},
		// }
	}
}
