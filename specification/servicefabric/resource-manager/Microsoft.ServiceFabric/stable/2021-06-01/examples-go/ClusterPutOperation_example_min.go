package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterPutOperation_example_min.json
func ExampleClustersClient_BeginCreateOrUpdate_putAClusterWithMinimumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", armservicefabric.Cluster{
		Location: to.Ptr("eastus"),
		Tags:     map[string]*string{},
		Properties: &armservicefabric.ClusterProperties{
			DiagnosticsStorageAccountConfig: &armservicefabric.DiagnosticsStorageAccountConfig{
				BlobEndpoint:            to.Ptr("https://diag.blob.core.windows.net/"),
				ProtectedAccountKeyName: to.Ptr("StorageAccountKey1"),
				QueueEndpoint:           to.Ptr("https://diag.queue.core.windows.net/"),
				StorageAccountName:      to.Ptr("diag"),
				TableEndpoint:           to.Ptr("https://diag.table.core.windows.net/"),
			},
			FabricSettings: []*armservicefabric.SettingsSectionDescription{
				{
					Name: to.Ptr("UpgradeService"),
					Parameters: []*armservicefabric.SettingsParameterDescription{
						{
							Name:  to.Ptr("AppPollIntervalInSeconds"),
							Value: to.Ptr("60"),
						}},
				}},
			ManagementEndpoint: to.Ptr("http://myCluster.eastus.cloudapp.azure.com:19080"),
			NodeTypes: []*armservicefabric.NodeTypeDescription{
				{
					Name: to.Ptr("nt1vm"),
					ApplicationPorts: &armservicefabric.EndpointRangeDescription{
						EndPort:   to.Ptr[int32](30000),
						StartPort: to.Ptr[int32](20000),
					},
					ClientConnectionEndpointPort: to.Ptr[int32](19000),
					DurabilityLevel:              to.Ptr(armservicefabric.DurabilityLevelBronze),
					EphemeralPorts: &armservicefabric.EndpointRangeDescription{
						EndPort:   to.Ptr[int32](64000),
						StartPort: to.Ptr[int32](49000),
					},
					HTTPGatewayEndpointPort: to.Ptr[int32](19007),
					IsPrimary:               to.Ptr(true),
					VMInstanceCount:         to.Ptr[int32](5),
				}},
			ReliabilityLevel: to.Ptr(armservicefabric.ReliabilityLevelSilver),
			UpgradeMode:      to.Ptr(armservicefabric.UpgradeModeAutomatic),
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
	// res.Cluster = armservicefabric.Cluster{
	// 	Name: to.Ptr("myCluster"),
	// 	Type: to.Ptr("Microsoft.ServiceFabric/clusters"),
	// 	Etag: to.Ptr("W/\"636462502169240743\""),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armservicefabric.ClusterProperties{
	// 		AvailableClusterVersions: []*armservicefabric.ClusterVersionDetails{
	// 			{
	// 				CodeVersion: to.Ptr("7.0.470.9590"),
	// 				Environment: to.Ptr(armservicefabric.ClusterEnvironmentWindows),
	// 				SupportExpiryUTC: to.Ptr("2018-06-15T23:59:59.9999999"),
	// 		}},
	// 		ClientCertificateCommonNames: []*armservicefabric.ClientCertificateCommonName{
	// 		},
	// 		ClientCertificateThumbprints: []*armservicefabric.ClientCertificateThumbprint{
	// 		},
	// 		ClusterCodeVersion: to.Ptr("7.0.470.9590"),
	// 		ClusterEndpoint: to.Ptr("https://eastus.servicefabric.azure.com"),
	// 		ClusterID: to.Ptr("92584666-9889-4ae8-8d02-91902923d37f"),
	// 		ClusterState: to.Ptr(armservicefabric.ClusterStateWaitingForNodes),
	// 		DiagnosticsStorageAccountConfig: &armservicefabric.DiagnosticsStorageAccountConfig{
	// 			BlobEndpoint: to.Ptr("https://diag.blob.core.windows.net/"),
	// 			ProtectedAccountKeyName: to.Ptr("StorageAccountKey1"),
	// 			QueueEndpoint: to.Ptr("https://diag.queue.core.windows.net/"),
	// 			StorageAccountName: to.Ptr("diag"),
	// 			TableEndpoint: to.Ptr("https://diag.table.core.windows.net/"),
	// 		},
	// 		FabricSettings: []*armservicefabric.SettingsSectionDescription{
	// 			{
	// 				Name: to.Ptr("UpgradeService"),
	// 				Parameters: []*armservicefabric.SettingsParameterDescription{
	// 					{
	// 						Name: to.Ptr("AppPollIntervalInSeconds"),
	// 						Value: to.Ptr("60"),
	// 				}},
	// 		}},
	// 		ManagementEndpoint: to.Ptr("http://myCluster.eastus.cloudapp.azure.com:19080"),
	// 		NodeTypes: []*armservicefabric.NodeTypeDescription{
	// 			{
	// 				Name: to.Ptr("nt1vm"),
	// 				ApplicationPorts: &armservicefabric.EndpointRangeDescription{
	// 					EndPort: to.Ptr[int32](30000),
	// 					StartPort: to.Ptr[int32](20000),
	// 				},
	// 				ClientConnectionEndpointPort: to.Ptr[int32](19000),
	// 				DurabilityLevel: to.Ptr(armservicefabric.DurabilityLevelBronze),
	// 				EphemeralPorts: &armservicefabric.EndpointRangeDescription{
	// 					EndPort: to.Ptr[int32](64000),
	// 					StartPort: to.Ptr[int32](49000),
	// 				},
	// 				HTTPGatewayEndpointPort: to.Ptr[int32](19007),
	// 				IsPrimary: to.Ptr(true),
	// 				VMInstanceCount: to.Ptr[int32](5),
	// 		}},
	// 		ProvisioningState: to.Ptr(armservicefabric.ProvisioningStateSucceeded),
	// 		ReliabilityLevel: to.Ptr(armservicefabric.ReliabilityLevelSilver),
	// 		UpgradeDescription: &armservicefabric.ClusterUpgradePolicy{
	// 			DeltaHealthPolicy: &armservicefabric.ClusterUpgradeDeltaHealthPolicy{
	// 				MaxPercentDeltaUnhealthyApplications: to.Ptr[int32](0),
	// 				MaxPercentDeltaUnhealthyNodes: to.Ptr[int32](0),
	// 				MaxPercentUpgradeDomainDeltaUnhealthyNodes: to.Ptr[int32](0),
	// 			},
	// 			ForceRestart: to.Ptr(false),
	// 			HealthCheckRetryTimeout: to.Ptr("00:45:00"),
	// 			HealthCheckStableDuration: to.Ptr("00:05:00"),
	// 			HealthCheckWaitDuration: to.Ptr("00:05:00"),
	// 			HealthPolicy: &armservicefabric.ClusterHealthPolicy{
	// 				MaxPercentUnhealthyApplications: to.Ptr[int32](100),
	// 				MaxPercentUnhealthyNodes: to.Ptr[int32](100),
	// 			},
	// 			UpgradeDomainTimeout: to.Ptr("02:00:00"),
	// 			UpgradeReplicaSetCheckTimeout: to.Ptr("10675199.02:48:05.4775807"),
	// 			UpgradeTimeout: to.Ptr("12:00:00"),
	// 		},
	// 		UpgradeMode: to.Ptr(armservicefabric.UpgradeModeAutomatic),
	// 	},
	// }
}
