package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2026-06-01/BackupInstanceOperations/PutBackupInstanceWithGenericParameters.json
func ExampleBackupInstancesClient_BeginCreateOrUpdate_createBackupInstanceWithGenericBackupDatasourceParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("97cda027-4279-4cde-b4ff-19afa0021d87", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupInstancesClient().BeginCreateOrUpdate(ctx, "ESAN-ECYBVTRG", "ESANVault", "esan-volgroup-bi", armdataprotection.BackupInstanceResource{
		Tags: map[string]*string{
			"key1": to.Ptr("val1"),
		},
		Properties: &armdataprotection.BackupInstance{
			FriendlyName: to.Ptr("esan-volgroup-bi"),
			DataSourceInfo: &armdataprotection.Datasource{
				ResourceID:       to.Ptr("/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc/volumeGroups/esan-volgroup"),
				ResourceURI:      to.Ptr("SampleresourceUri123"),
				DatasourceType:   to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
				ResourceName:     to.Ptr("esan-volgroup-bi"),
				ResourceType:     to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
				ResourceLocation: to.Ptr("eastus2euap"),
				ObjectType:       to.Ptr("Datasource"),
			},
			DataSourceSetInfo: &armdataprotection.DatasourceSet{
				ResourceID:       to.Ptr("/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc"),
				DatasourceType:   to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
				ResourceType:     to.Ptr("Microsoft.ElasticSan/elasticSans"),
				ResourceLocation: to.Ptr("eastus2euap"),
				ObjectType:       to.Ptr("DatasourceSet"),
			},
			PolicyInfo: &armdataprotection.PolicyInfo{
				PolicyID: to.Ptr("/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.DataProtection/backupVaults/ESANVault/backupPolicies/BVTPolicy"),
				PolicyParameters: &armdataprotection.PolicyParameters{
					DataStoreParametersList: []armdataprotection.DataStoreParametersClassification{
						&armdataprotection.AzureOperationalStoreParameters{
							ResourceGroupID: to.Ptr("/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG"),
							ObjectType:      to.Ptr("AzureOperationalStoreParameters"),
							DataStoreType:   to.Ptr(armdataprotection.DataStoreTypesOperationalStore),
						},
					},
					BackupDatasourceParametersList: []armdataprotection.BackupDatasourceParametersClassification{
						&armdataprotection.GenericBackupDatasourceParameters{
							ObjectType: to.Ptr("GenericBackupDatasourceParameters"),
							ResourceSelectors: []*string{
								to.Ptr("vol1"),
								to.Ptr("vol2"),
								to.Ptr("vol3"),
							},
						},
					},
				},
			},
			ObjectType: to.Ptr("BackupInstance"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataprotection.BackupInstancesClientCreateOrUpdateResponse{
	// 	BackupInstanceResource: armdataprotection.BackupInstanceResource{
	// 		ID: to.Ptr("/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.DataProtection/backupVaults/ESANVault/backupInstances/esan-volgroup-bi"),
	// 		Name: to.Ptr("esan-volgroup-bi"),
	// 		Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances"),
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("val1"),
	// 		},
	// 		Properties: &armdataprotection.BackupInstance{
	// 			FriendlyName: to.Ptr("esan-volgroup-bi"),
	// 			DataSourceInfo: &armdataprotection.Datasource{
	// 				ResourceID: to.Ptr("/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc/volumeGroups/esan-volgroup"),
	// 				ResourceURI: to.Ptr("SampleresourceUri123"),
	// 				DatasourceType: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
	// 				ResourceName: to.Ptr("esan-volgroup-bi"),
	// 				ResourceType: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
	// 				ResourceLocation: to.Ptr("eastus2euap"),
	// 				ObjectType: to.Ptr("Datasource"),
	// 			},
	// 			DataSourceSetInfo: &armdataprotection.DatasourceSet{
	// 				ResourceID: to.Ptr("/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc"),
	// 				DatasourceType: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
	// 				ResourceType: to.Ptr("Microsoft.ElasticSan/elasticSans"),
	// 				ResourceLocation: to.Ptr("eastus2euap"),
	// 				ObjectType: to.Ptr("DatasourceSet"),
	// 			},
	// 			PolicyInfo: &armdataprotection.PolicyInfo{
	// 				PolicyID: to.Ptr("/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.DataProtection/backupVaults/ESANVault/backupPolicies/BVTPolicy"),
	// 				PolicyParameters: &armdataprotection.PolicyParameters{
	// 					DataStoreParametersList: []armdataprotection.DataStoreParametersClassification{
	// 						&armdataprotection.AzureOperationalStoreParameters{
	// 							ResourceGroupID: to.Ptr("/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG"),
	// 							ObjectType: to.Ptr("AzureOperationalStoreParameters"),
	// 							DataStoreType: to.Ptr(armdataprotection.DataStoreTypesOperationalStore),
	// 						},
	// 					},
	// 					BackupDatasourceParametersList: []armdataprotection.BackupDatasourceParametersClassification{
	// 						&armdataprotection.GenericBackupDatasourceParameters{
	// 							ObjectType: to.Ptr("GenericBackupDatasourceParameters"),
	// 							ResourceSelectors: []*string{
	// 								to.Ptr("vol1"),
	// 								to.Ptr("vol2"),
	// 								to.Ptr("vol3"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProtectionStatus: &armdataprotection.ProtectionStatusDetails{
	// 				Status: to.Ptr(armdataprotection.Status("NotProtected")),
	// 			},
	// 			ProvisioningState: to.Ptr("Provisioned"),
	// 			ObjectType: to.Ptr("BackupInstance"),
	// 		},
	// 	},
	// }
}
