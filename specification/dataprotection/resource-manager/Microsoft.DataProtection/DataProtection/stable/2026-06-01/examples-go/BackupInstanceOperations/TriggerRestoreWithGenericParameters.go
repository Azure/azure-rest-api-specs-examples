package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2026-06-01/BackupInstanceOperations/TriggerRestoreWithGenericParameters.json
func ExampleBackupInstancesClient_BeginTriggerRestore_triggerRestoreWithGenericRestoreDatasourceCriteria() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupInstancesClient().BeginTriggerRestore(ctx, "000pikumar", "PrivatePreviewVault1", "testInstance1", &armdataprotection.AzureBackupRecoveryPointBasedRestoreRequest{
		ObjectType:          to.Ptr("AzureBackupRecoveryPointBasedRestoreRequest"),
		RecoveryPointID:     to.Ptr("hardcodedRP"),
		SourceDataStoreType: to.Ptr(armdataprotection.SourceDataStoreTypeOperationalStore),
		RestoreTargetInfo: &armdataprotection.ItemLevelRestoreTargetInfo{
			RestoreLocation: to.Ptr("southeastasia"),
			RecoveryOption:  to.Ptr(armdataprotection.RecoveryOptionFailIfExists),
			ObjectType:      to.Ptr("ItemLevelRestoreTargetInfo"),
			DatasourceInfo: &armdataprotection.Datasource{
				ResourceID:       to.Ptr("/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc/volumeGroups/target-esan-volgroup"),
				ResourceURI:      to.Ptr("SampleresourceUri123"),
				DatasourceType:   to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
				ResourceName:     to.Ptr("target-esan-volgroup"),
				ResourceType:     to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
				ResourceLocation: to.Ptr("eastus2euap"),
				ObjectType:       to.Ptr("Datasource"),
			},
			DatasourceSetInfo: &armdataprotection.DatasourceSet{
				ResourceID:       to.Ptr("/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc"),
				DatasourceType:   to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
				ResourceType:     to.Ptr("Microsoft.ElasticSan/elasticSans"),
				ResourceLocation: to.Ptr("eastus2euap"),
				ObjectType:       to.Ptr("DatasourceSet"),
			},
			RestoreCriteria: []armdataprotection.ItemLevelRestoreCriteriaClassification{
				&armdataprotection.GenericRestoreDatasourceCriteria{
					ObjectType: to.Ptr("GenericRestoreDatasourceCriteria"),
					ResourceSelectors: &armdataprotection.ResourceListSelectionCriteria{
						ObjectType: to.Ptr("resourceListSelectionCriteria"),
						ResourceIdentifiers: []*string{
							to.Ptr("source-vol1"),
							to.Ptr("source-vol2"),
							to.Ptr("source-vol3"),
						},
						ResourceNameOverrides: map[string]*string{
							"source-vol1": to.Ptr("target-vol1"),
							"source-vol2": to.Ptr("target-vol2"),
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
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataprotection.BackupInstancesClientTriggerRestoreResponse{
	// 	OperationJobExtendedInfo: armdataprotection.OperationJobExtendedInfo{
	// 		JobID: to.Ptr("c60cb49-63e8-4b21-b9bd-26277b3fdfae"),
	// 		ObjectType: to.Ptr("OperationJobExtendedInfo"),
	// 	},
	// }
}
