package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/CreateHDInsightClusterWithComputeIsolationProperties.json
func ExampleClustersClient_BeginCreate_createClusterWithComputeIsolationProperties() {
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
				Kind: to.Ptr("hadoop"),
			},
			ClusterVersion: to.Ptr("3.6"),
			ComputeIsolationProperties: &armhdinsight.ComputeIsolationProperties{
				EnableComputeIsolation: to.Ptr(true),
			},
			ComputeProfile: &armhdinsight.ComputeProfile{
				Roles: []*armhdinsight.Role{
					{
						Name: to.Ptr("headnode"),
						HardwareProfile: &armhdinsight.HardwareProfile{
							VMSize: to.Ptr("standard_d3"),
						},
						OSProfile: &armhdinsight.OsProfile{
							LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
								Password: to.Ptr("**********"),
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
							VMSize: to.Ptr("standard_d3"),
						},
						OSProfile: &armhdinsight.OsProfile{
							LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
								Password: to.Ptr("**********"),
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
					}},
			},
			OSType: to.Ptr(armhdinsight.OSTypeLinux),
			StorageProfile: &armhdinsight.StorageProfile{
				Storageaccounts: []*armhdinsight.StorageAccount{
					{
						Name:                to.Ptr("mystorage"),
						Container:           to.Ptr("containername"),
						EnableSecureChannel: to.Ptr(true),
						IsDefault:           to.Ptr(true),
						Key:                 to.Ptr("storage account key"),
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
	// res.Cluster = armhdinsight.Cluster{
	// 	Name: to.Ptr("cluster1"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.HDInsight/clusters/cluster1"),
	// 	Location: to.Ptr("South Central US"),
	// 	Etag: to.Ptr("fdf2a6e8-ce83-42cc-8c2d-0ceb11a370ff"),
	// 	Properties: &armhdinsight.ClusterGetProperties{
	// 		ClusterDefinition: &armhdinsight.ClusterDefinition{
	// 			Blueprint: to.Ptr("https://blueprints.azurehdinsight.net/hadoop-3.6.1000.67.2005040905.json"),
	// 			Kind: to.Ptr("hadoop"),
	// 		},
	// 		ClusterID: to.Ptr("8186508b6234470e9d16c9e8e13bd821"),
	// 		ClusterState: to.Ptr("Running"),
	// 		ClusterVersion: to.Ptr("3.6.1000.67"),
	// 		ComputeIsolationProperties: &armhdinsight.ComputeIsolationProperties{
	// 			EnableComputeIsolation: to.Ptr(true),
	// 		},
	// 		ComputeProfile: &armhdinsight.ComputeProfile{
	// 			Roles: []*armhdinsight.Role{
	// 				{
	// 					Name: to.Ptr("headnode"),
	// 					HardwareProfile: &armhdinsight.HardwareProfile{
	// 						VMSize: to.Ptr("standard_d3"),
	// 					},
	// 					OSProfile: &armhdinsight.OsProfile{
	// 						LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
	// 							Username: to.Ptr("sshuser"),
	// 						},
	// 					},
	// 					TargetInstanceCount: to.Ptr[int32](2),
	// 					VirtualNetworkProfile: &armhdinsight.VirtualNetworkProfile{
	// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname"),
	// 						Subnet: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("workernode"),
	// 					HardwareProfile: &armhdinsight.HardwareProfile{
	// 						VMSize: to.Ptr("standard_d3"),
	// 					},
	// 					OSProfile: &armhdinsight.OsProfile{
	// 						LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
	// 							Username: to.Ptr("sshuser"),
	// 						},
	// 					},
	// 					TargetInstanceCount: to.Ptr[int32](2),
	// 			}},
	// 		},
	// 		CreatedDate: to.Ptr("2020-06-09T12:25:43.48"),
	// 		OSType: to.Ptr(armhdinsight.OSTypeLinux),
	// 		ProvisioningState: to.Ptr(armhdinsight.HDInsightClusterProvisioningStateSucceeded),
	// 		QuotaInfo: &armhdinsight.QuotaInfo{
	// 			CoresUsed: to.Ptr[int32](16),
	// 		},
	// 		Tier: to.Ptr(armhdinsight.TierStandard),
	// 	},
	// }
}
