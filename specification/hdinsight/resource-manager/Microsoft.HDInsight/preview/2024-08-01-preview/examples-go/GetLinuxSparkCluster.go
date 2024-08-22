package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetLinuxSparkCluster.json
func ExampleClustersClient_Get_getSparkOnLinuxCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().Get(ctx, "rg1", "cluster1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armhdinsight.Cluster{
	// 	Name: to.Ptr("cluster1"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusters"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HDInsight/clusters/cluster1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("val1"),
	// 	},
	// 	Etag: to.Ptr("4cd8b5a9-ca9c-4239-9e5b-3916032e994c"),
	// 	Properties: &armhdinsight.ClusterGetProperties{
	// 		ClusterDefinition: &armhdinsight.ClusterDefinition{
	// 			Blueprint: to.Ptr("https://blueprints.azurehdinsight.net/spark-3.5.1000.0.9625065.json"),
	// 			ComponentVersion: map[string]*string{
	// 				"Spark": to.Ptr("2.0"),
	// 			},
	// 			Kind: to.Ptr("SPARK"),
	// 		},
	// 		ClusterState: to.Ptr("Running"),
	// 		ClusterVersion: to.Ptr("3.5.1000.0"),
	// 		ComputeProfile: &armhdinsight.ComputeProfile{
	// 			Roles: []*armhdinsight.Role{
	// 				{
	// 					Name: to.Ptr("headnode"),
	// 					HardwareProfile: &armhdinsight.HardwareProfile{
	// 						VMSize: to.Ptr("Standard_D12_V2"),
	// 					},
	// 					OSProfile: &armhdinsight.OsProfile{
	// 						LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
	// 							Username: to.Ptr("sshuser"),
	// 						},
	// 					},
	// 					TargetInstanceCount: to.Ptr[int32](2),
	// 				},
	// 				{
	// 					Name: to.Ptr("workernode"),
	// 					HardwareProfile: &armhdinsight.HardwareProfile{
	// 						VMSize: to.Ptr("Standard_D4_V2"),
	// 					},
	// 					OSProfile: &armhdinsight.OsProfile{
	// 						LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
	// 							Username: to.Ptr("sshuser"),
	// 						},
	// 					},
	// 					TargetInstanceCount: to.Ptr[int32](4),
	// 				},
	// 				{
	// 					Name: to.Ptr("zookeepernode"),
	// 					HardwareProfile: &armhdinsight.HardwareProfile{
	// 						VMSize: to.Ptr("Medium"),
	// 					},
	// 					OSProfile: &armhdinsight.OsProfile{
	// 						LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
	// 							Username: to.Ptr("sshuser"),
	// 						},
	// 					},
	// 					TargetInstanceCount: to.Ptr[int32](3),
	// 			}},
	// 		},
	// 		ConnectivityEndpoints: []*armhdinsight.ConnectivityEndpoint{
	// 			{
	// 				Name: to.Ptr("SSH"),
	// 				Location: to.Ptr("cluster1-ssh.azurehdinsight.net"),
	// 				Port: to.Ptr[int32](22),
	// 				Protocol: to.Ptr("TCP"),
	// 			},
	// 			{
	// 				Name: to.Ptr("HTTPS"),
	// 				Location: to.Ptr("cluster1.azurehdinsight.net"),
	// 				Port: to.Ptr[int32](443),
	// 				Protocol: to.Ptr("TCP"),
	// 		}},
	// 		CreatedDate: to.Ptr("2017-03-03T08:35:07.323"),
	// 		OSType: to.Ptr(armhdinsight.OSTypeLinux),
	// 		ProvisioningState: to.Ptr(armhdinsight.HDInsightClusterProvisioningStateSucceeded),
	// 		QuotaInfo: &armhdinsight.QuotaInfo{
	// 			CoresUsed: to.Ptr[int32](40),
	// 		},
	// 		Tier: to.Ptr(armhdinsight.TierStandard),
	// 	},
	// }
}
