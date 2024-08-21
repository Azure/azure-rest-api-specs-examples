package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/CreateHDInsightClusterWithAvailabilityZones.json
func ExampleClustersClient_BeginCreate_createClusterWithAvailabilityZones() {
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
					"ambari-conf": map[string]any{
						"database-name":          "{ambari database name}",
						"database-server":        "{sql server name}.database.windows.net",
						"database-user-name":     "**********",
						"database-user-password": "**********",
					},
					"gateway": map[string]any{
						"restAuthCredential.isEnabled": true,
						"restAuthCredential.password":  "**********",
						"restAuthCredential.username":  "admin",
					},
					"hive-env": map[string]any{
						"hive_database":                       "Existing MSSQL Server database with SQL authentication",
						"hive_database_name":                  "{hive metastore name}",
						"hive_database_type":                  "mssql",
						"hive_existing_mssql_server_database": "{hive metastore name}",
						"hive_existing_mssql_server_host":     "{sql server name}.database.windows.net",
						"hive_hostname":                       "{sql server name}.database.windows.net",
					},
					"hive-site": map[string]any{
						"javax.jdo.option.ConnectionDriverName": "com.microsoft.sqlserver.jdbc.SQLServerDriver",
						"javax.jdo.option.ConnectionPassword":   "**********!",
						"javax.jdo.option.ConnectionURL":        "jdbc:sqlserver://{sql server name}.database.windows.net;database={hive metastore name};encrypt=true;trustServerCertificate=true;create=false;loginTimeout=300;sendStringParametersAsUnicode=true;prepareSQL=0",
						"javax.jdo.option.ConnectionUserName":   "**********",
					},
					"oozie-env": map[string]any{
						"oozie_database":                       "Existing MSSQL Server database with SQL authentication",
						"oozie_database_name":                  "{oozie metastore name}",
						"oozie_database_type":                  "mssql",
						"oozie_existing_mssql_server_database": "{oozie metastore name}",
						"oozie_existing_mssql_server_host":     "{sql server name}.database.windows.net",
						"oozie_hostname":                       "{sql server name}.database.windows.net",
					},
					"oozie-site": map[string]any{
						"oozie.db.schema.name":                   "oozie",
						"oozie.service.JPAService.jdbc.driver":   "com.microsoft.sqlserver.jdbc.SQLServerDriver",
						"oozie.service.JPAService.jdbc.password": "**********",
						"oozie.service.JPAService.jdbc.url":      "jdbc:sqlserver://{sql server name}.database.windows.net;database={oozie metastore name};encrypt=true;trustServerCertificate=true;create=false;loginTimeout=300;sendStringParametersAsUnicode=true;prepareSQL=0",
						"oozie.service.JPAService.jdbc.username": "**********",
					},
				},
				Kind: to.Ptr("hadoop"),
			},
			ClusterVersion: to.Ptr("3.6"),
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
						VirtualNetworkProfile: &armhdinsight.VirtualNetworkProfile{
							ID:     to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname"),
							Subnet: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
						},
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
						VirtualNetworkProfile: &armhdinsight.VirtualNetworkProfile{
							ID:     to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname"),
							Subnet: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
						},
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
		Zones: []*string{
			to.Ptr("1")},
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
	// 					VirtualNetworkProfile: &armhdinsight.VirtualNetworkProfile{
	// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname"),
	// 						Subnet: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
	// 					},
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
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}
