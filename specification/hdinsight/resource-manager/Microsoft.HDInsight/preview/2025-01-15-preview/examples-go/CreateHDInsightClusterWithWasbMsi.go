package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a9dbb28e788355a47dc5bad3ea5f8da212b4bf6/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2025-01-15-preview/examples/CreateHDInsightClusterWithWasbMsi.json
func ExampleClustersClient_BeginCreate_createClusterWithStorageWasbMsi() {
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
		Identity: &armhdinsight.ClusterIdentity{
			Type: to.Ptr(armhdinsight.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armhdinsight.UserAssignedIdentity{
				"/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/msi": {},
			},
		},
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
			ClusterVersion: to.Ptr("5.1"),
			ComputeProfile: &armhdinsight.ComputeProfile{
				Roles: []*armhdinsight.Role{
					{
						Name: to.Ptr("headnode"),
						HardwareProfile: &armhdinsight.HardwareProfile{
							VMSize: to.Ptr("Standard_E8_V3"),
						},
						MinInstanceCount: to.Ptr[int32](1),
						OSProfile: &armhdinsight.OsProfile{
							LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
								Password: to.Ptr("**********"),
								Username: to.Ptr("sshuser"),
							},
						},
						ScriptActions:       []*armhdinsight.ScriptAction{},
						TargetInstanceCount: to.Ptr[int32](2),
						VirtualNetworkProfile: &armhdinsight.VirtualNetworkProfile{
							ID:     to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname"),
							Subnet: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
						},
					},
					{
						Name: to.Ptr("workernode"),
						HardwareProfile: &armhdinsight.HardwareProfile{
							VMSize: to.Ptr("Standard_E8_V3"),
						},
						MinInstanceCount: to.Ptr[int32](1),
						OSProfile: &armhdinsight.OsProfile{
							LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
								Password: to.Ptr("**********"),
								Username: to.Ptr("sshuser"),
							},
						},
						ScriptActions:       []*armhdinsight.ScriptAction{},
						TargetInstanceCount: to.Ptr[int32](3),
						VirtualNetworkProfile: &armhdinsight.VirtualNetworkProfile{
							ID:     to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname"),
							Subnet: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
						},
					},
					{
						Name: to.Ptr("zookeepernode"),
						HardwareProfile: &armhdinsight.HardwareProfile{
							VMSize: to.Ptr("Standard_E8_V3"),
						},
						MinInstanceCount: to.Ptr[int32](1),
						OSProfile: &armhdinsight.OsProfile{
							LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
								Password: to.Ptr("**********"),
								Username: to.Ptr("sshuser"),
							},
						},
						ScriptActions:       []*armhdinsight.ScriptAction{},
						TargetInstanceCount: to.Ptr[int32](3),
						VirtualNetworkProfile: &armhdinsight.VirtualNetworkProfile{
							ID:     to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname"),
							Subnet: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
						},
					}},
			},
			OSType: to.Ptr(armhdinsight.OSTypeLinux),
			StorageProfile: &armhdinsight.StorageProfile{
				Storageaccounts: []*armhdinsight.StorageAccount{
					{
						Name:          to.Ptr("mystorage.blob.core.windows.net"),
						Container:     to.Ptr("containername"),
						IsDefault:     to.Ptr(true),
						MsiResourceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/msi"),
						ResourceID:    to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/mystorage"),
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
	// 	Location: to.Ptr("East Asia"),
	// 	Etag: to.Ptr("ebd7702e-yr32-4968-aba6-f0a3c028b30e"),
	// 	Properties: &armhdinsight.ClusterGetProperties{
	// 		ClusterDefinition: &armhdinsight.ClusterDefinition{
	// 			Blueprint: to.Ptr("https://blueprints.azurehdinsight.net/hadoop-5.1.3000.0.2504170544.json"),
	// 			ComponentVersion: map[string]*string{
	// 				"Hadoop": to.Ptr("3.3"),
	// 			},
	// 			Kind: to.Ptr("Hadoop"),
	// 		},
	// 		ClusterState: to.Ptr("Running"),
	// 		ClusterVersion: to.Ptr("5.1.3000.0"),
	// 		ComputeProfile: &armhdinsight.ComputeProfile{
	// 			Roles: []*armhdinsight.Role{
	// 				{
	// 					Name: to.Ptr("headnode"),
	// 					HardwareProfile: &armhdinsight.HardwareProfile{
	// 						VMSize: to.Ptr("Standard_E8_V3"),
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
	// 						VMSize: to.Ptr("Standard_E8_V3"),
	// 					},
	// 					OSProfile: &armhdinsight.OsProfile{
	// 						LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
	// 							Username: to.Ptr("sshuser"),
	// 						},
	// 					},
	// 					TargetInstanceCount: to.Ptr[int32](3),
	// 				},
	// 				{
	// 					Name: to.Ptr("zookeepernode"),
	// 					HardwareProfile: &armhdinsight.HardwareProfile{
	// 						VMSize: to.Ptr("Standard_E8_V3"),
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
	// 		CreatedDate: to.Ptr("2025-05-08T08:07:09.813"),
	// 		OSType: to.Ptr(armhdinsight.OSTypeLinux),
	// 		ProvisioningState: to.Ptr(armhdinsight.HDInsightClusterProvisioningStateSucceeded),
	// 		QuotaInfo: &armhdinsight.QuotaInfo{
	// 			CoresUsed: to.Ptr[int32](40),
	// 		},
	// 		Tier: to.Ptr(armhdinsight.TierStandard),
	// 	},
	// }
}
