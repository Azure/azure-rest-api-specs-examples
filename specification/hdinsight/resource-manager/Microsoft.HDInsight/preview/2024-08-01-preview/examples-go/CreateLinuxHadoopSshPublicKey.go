package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/CreateLinuxHadoopSshPublicKey.json
func ExampleClustersClient_BeginCreate_createHadoopOnLinuxClusterWithSshPublicKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreate(ctx, "rg1", "cluster1", armhdinsight.ClusterCreateParametersExtended{
		Properties: &armhdinsight.ClusterCreateProperties{
			ClusterDefinition: &armhdinsight.ClusterDefinition{
				Configurations: map[string]any{
					"gateway": map[string]any{
						"restAuthCredential.isEnabled": true,
						"restAuthCredential.password":  "**********",
						"restAuthCredential.username":  "admin",
					},
				},
				Kind: to.Ptr("Hadoop"),
			},
			ClusterVersion: to.Ptr("3.5"),
			ComputeProfile: &armhdinsight.ComputeProfile{
				Roles: []*armhdinsight.Role{
					{
						Name: to.Ptr("headnode"),
						HardwareProfile: &armhdinsight.HardwareProfile{
							VMSize: to.Ptr("Standard_D3_V2"),
						},
						MinInstanceCount: to.Ptr[int32](1),
						OSProfile: &armhdinsight.OsProfile{
							LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
								SSHProfile: &armhdinsight.SSHProfile{
									PublicKeys: []*armhdinsight.SSHPublicKey{
										{
											CertificateData: to.Ptr("**********"),
										}},
								},
								Username: to.Ptr("sshuser"),
							},
						},
						TargetInstanceCount: to.Ptr[int32](2),
					},
					{
						Name: to.Ptr("workernode"),
						HardwareProfile: &armhdinsight.HardwareProfile{
							VMSize: to.Ptr("Standard_D3_V2"),
						},
						MinInstanceCount: to.Ptr[int32](1),
						OSProfile: &armhdinsight.OsProfile{
							LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
								Password: to.Ptr("**********"),
								Username: to.Ptr("sshuser"),
							},
						},
						TargetInstanceCount: to.Ptr[int32](4),
					},
					{
						Name: to.Ptr("zookeepernode"),
						HardwareProfile: &armhdinsight.HardwareProfile{
							VMSize: to.Ptr("Small"),
						},
						MinInstanceCount: to.Ptr[int32](1),
						OSProfile: &armhdinsight.OsProfile{
							LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
								Password: to.Ptr("**********"),
								Username: to.Ptr("sshuser"),
							},
						},
						TargetInstanceCount: to.Ptr[int32](3),
					}},
			},
			OSType: to.Ptr(armhdinsight.OSTypeLinux),
			StorageProfile: &armhdinsight.StorageProfile{
				Storageaccounts: []*armhdinsight.StorageAccount{
					{
						Name:                to.Ptr("mystorage.blob.core.windows.net"),
						Container:           to.Ptr("containername"),
						EnableSecureChannel: to.Ptr(true),
						IsDefault:           to.Ptr(true),
						Key:                 to.Ptr("storagekey"),
					}},
			},
			Tier: to.Ptr(armhdinsight.TierStandard),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("val1"),
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
	// res.Cluster = armhdinsight.Cluster{
	// 	Name: to.Ptr("cluster1"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.HDInsight/clusters/cluster1"),
	// 	Location: to.Ptr("East US"),
	// 	Etag: to.Ptr("fdf2a6e8-ce83-42cc-8c2d-0ceb11a370ff"),
	// 	Properties: &armhdinsight.ClusterGetProperties{
	// 		ClusterDefinition: &armhdinsight.ClusterDefinition{
	// 			Blueprint: to.Ptr("https://blueprints.azurehdinsight.net/hadoop-4.0.1000.1.1910270459.json"),
	// 			ComponentVersion: map[string]*string{
	// 				"Hadoop": to.Ptr("3.1"),
	// 			},
	// 			Kind: to.Ptr("HADOOP"),
	// 		},
	// 		ClusterState: to.Ptr("Running"),
	// 		ClusterVersion: to.Ptr("4.0.1000.1"),
	// 		ComputeProfile: &armhdinsight.ComputeProfile{
	// 			Roles: []*armhdinsight.Role{
	// 				{
	// 					Name: to.Ptr("headnode"),
	// 					HardwareProfile: &armhdinsight.HardwareProfile{
	// 						VMSize: to.Ptr("standard_d12_v2"),
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
	// 						VMSize: to.Ptr("standard_d4_v2"),
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
	// 						VMSize: to.Ptr("standard_a2_v2"),
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
	// 		CreatedDate: to.Ptr("2019-11-18T12:25:43.48"),
	// 		OSType: to.Ptr(armhdinsight.OSTypeLinux),
	// 		ProvisioningState: to.Ptr(armhdinsight.HDInsightClusterProvisioningStateSucceeded),
	// 		QuotaInfo: &armhdinsight.QuotaInfo{
	// 			CoresUsed: to.Ptr[int32](40),
	// 		},
	// 		Tier: to.Ptr(armhdinsight.TierStandard),
	// 	},
	// }
}
