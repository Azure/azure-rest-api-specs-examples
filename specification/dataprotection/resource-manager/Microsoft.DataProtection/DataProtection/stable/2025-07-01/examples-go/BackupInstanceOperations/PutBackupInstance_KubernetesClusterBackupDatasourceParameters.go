package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2025-07-01/BackupInstanceOperations/PutBackupInstance_KubernetesClusterBackupDatasourceParameters.json
func ExampleBackupInstancesClient_BeginCreateOrUpdate_createBackupInstanceWithKubernetesClusterBackupDatasourceParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("62b829ee-7936-40c9-a1c9-47a93f9f3965", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupInstancesClient().BeginCreateOrUpdate(ctx, "aksrg", "aksvault", "aksbi", armdataprotection.BackupInstanceResource{
		Properties: &armdataprotection.BackupInstance{
			DataSourceInfo: &armdataprotection.Datasource{
				DatasourceType:   to.Ptr("Microsoft.ContainerService/managedclusters"),
				ObjectType:       to.Ptr("Datasource"),
				ResourceID:       to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster"),
				ResourceLocation: to.Ptr("eastus2euap"),
				ResourceName:     to.Ptr("akscluster"),
				ResourceType:     to.Ptr("Microsoft.ContainerService/managedclusters"),
				ResourceURI:      to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster"),
			},
			DataSourceSetInfo: &armdataprotection.DatasourceSet{
				DatasourceType:   to.Ptr("Microsoft.ContainerService/managedclusters"),
				ObjectType:       to.Ptr("DatasourceSet"),
				ResourceID:       to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster"),
				ResourceLocation: to.Ptr("eastus2euap"),
				ResourceName:     to.Ptr("akscluster"),
				ResourceType:     to.Ptr("Microsoft.ContainerService/managedclusters"),
				ResourceURI:      to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster"),
			},
			FriendlyName: to.Ptr("aksbi"),
			ObjectType:   to.Ptr("BackupInstance"),
			PolicyInfo: &armdataprotection.PolicyInfo{
				PolicyID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourcegroups/aksrg/providers/Microsoft.DataProtection/BackupVaults/aksvault/backupPolicies/akspolicy"),
				PolicyParameters: &armdataprotection.PolicyParameters{
					BackupDatasourceParametersList: []armdataprotection.BackupDatasourceParametersClassification{
						&armdataprotection.KubernetesClusterBackupDatasourceParameters{
							ExcludedNamespaces: []*string{
								to.Ptr("kube-system"),
							},
							ExcludedResourceTypes: []*string{
								to.Ptr("v1/Secret"),
							},
							IncludeClusterScopeResources: to.Ptr(true),
							IncludedNamespaces: []*string{
								to.Ptr("test"),
							},
							IncludedResourceTypes: []*string{},
							IncludedVolumeTypes: []*armdataprotection.AKSVolumeTypes{
								to.Ptr(armdataprotection.AKSVolumeTypesAzureDisk),
								to.Ptr(armdataprotection.AKSVolumeTypesAzureFileShareSMB),
							},
							LabelSelectors:  []*string{},
							ObjectType:      to.Ptr("KubernetesClusterBackupDatasourceParameters"),
							SnapshotVolumes: to.Ptr(true),
						},
					},
					DataStoreParametersList: []armdataprotection.DataStoreParametersClassification{
						&armdataprotection.AzureOperationalStoreParameters{
							DataStoreType:   to.Ptr(armdataprotection.DataStoreTypesOperationalStore),
							ObjectType:      to.Ptr("AzureOperationalStoreParameters"),
							ResourceGroupID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg"),
						},
					},
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
	// res = armdataprotection.BackupInstancesClientCreateOrUpdateResponse{
	// 	BackupInstanceResource: &armdataprotection.BackupInstanceResource{
	// 		Name: to.Ptr("aksbi"),
	// 		Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances"),
	// 		ID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.DataProtection/backupVaults/aksvault/backupInstances/aksbi"),
	// 		Properties: &armdataprotection.BackupInstance{
	// 			DataSourceInfo: &armdataprotection.Datasource{
	// 				DatasourceType: to.Ptr("Microsoft.ContainerService/managedclusters"),
	// 				ObjectType: to.Ptr("Datasource"),
	// 				ResourceID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster"),
	// 				ResourceLocation: to.Ptr("eastus2euap"),
	// 				ResourceName: to.Ptr("akscluster"),
	// 				ResourceType: to.Ptr("Microsoft.ContainerService/managedclusters"),
	// 				ResourceURI: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster"),
	// 			},
	// 			DataSourceSetInfo: &armdataprotection.DatasourceSet{
	// 				DatasourceType: to.Ptr("Microsoft.ContainerService/managedclusters"),
	// 				ObjectType: to.Ptr("DatasourceSet"),
	// 				ResourceID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster"),
	// 				ResourceLocation: to.Ptr("eastus2euap"),
	// 				ResourceType: to.Ptr("Microsoft.ContainerService/managedclusters"),
	// 			},
	// 			FriendlyName: to.Ptr("aksbi"),
	// 			ObjectType: to.Ptr("BackupInstance"),
	// 			PolicyInfo: &armdataprotection.PolicyInfo{
	// 				PolicyID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.DataProtection/backupVaults/aksvault/backupPolicies/akspolicy"),
	// 				PolicyParameters: &armdataprotection.PolicyParameters{
	// 					BackupDatasourceParametersList: []armdataprotection.BackupDatasourceParametersClassification{
	// 						&armdataprotection.KubernetesClusterBackupDatasourceParameters{
	// 							ExcludedNamespaces: []*string{
	// 								to.Ptr("kube-system"),
	// 							},
	// 							ExcludedResourceTypes: []*string{
	// 								to.Ptr("v1/Secret"),
	// 							},
	// 							IncludeClusterScopeResources: to.Ptr(true),
	// 							IncludedNamespaces: []*string{
	// 								to.Ptr("test"),
	// 							},
	// 							IncludedResourceTypes: []*string{
	// 							},
	// 							IncludedVolumeTypes: []*armdataprotection.AKSVolumeTypes{
	// 								to.Ptr(armdataprotection.AKSVolumeTypesAzureDisk),
	// 								to.Ptr(armdataprotection.AKSVolumeTypesAzureFileShareSMB),
	// 							},
	// 							LabelSelectors: []*string{
	// 							},
	// 							ObjectType: to.Ptr("KubernetesClusterBackupDatasourceParameters"),
	// 							SnapshotVolumes: to.Ptr(true),
	// 						},
	// 					},
	// 					DataStoreParametersList: []armdataprotection.DataStoreParametersClassification{
	// 						&armdataprotection.AzureOperationalStoreParameters{
	// 							DataStoreType: to.Ptr(armdataprotection.DataStoreTypesOperationalStore),
	// 							ObjectType: to.Ptr("AzureOperationalStoreParameters"),
	// 							ResourceGroupID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProtectionStatus: &armdataprotection.ProtectionStatusDetails{
	// 				Status: to.Ptr(armdataprotection.Status("NotProtected")),
	// 			},
	// 			ProvisioningState: to.Ptr("Provisioned"),
	// 		},
	// 	},
	// }
}
