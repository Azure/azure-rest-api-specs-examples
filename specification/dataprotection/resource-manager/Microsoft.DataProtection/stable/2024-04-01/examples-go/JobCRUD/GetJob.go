package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/JobCRUD/GetJob.json
func ExampleJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().Get(ctx, "BugBash1", "BugBashVaultForCCYv11", "3c60cb49-63e8-4b21-b9bd-26277b3fdfae", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureBackupJobResource = armdataprotection.AzureBackupJobResource{
	// 	Name: to.Ptr("3c60cb49-63e8-4b21-b9bd-26277b3fdfae"),
	// 	Type: to.Ptr("Microsoft.DataProtection/Backupvaults/backupJobs"),
	// 	ID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/BugBash1/providers/Microsoft.DataProtection/Backupvaults/BugBashVaultForCCYv11/backupJobs/3c60cb49-63e8-4b21-b9bd-26277b3fdfae"),
	// 	Properties: &armdataprotection.AzureBackupJob{
	// 		ActivityID: to.Ptr("c4344fb4-7c11-43a4-8307-7ae7c7fb09b9"),
	// 		BackupInstanceFriendlyName: to.Ptr("mabtestingccybasicv11\\bugbashdb2"),
	// 		BackupInstanceID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/BugBash1/providers/Microsoft.DataProtection/backupVaults/BugBashVaultForCCYv11/backupInstances/28460a9d-707a-45f3-ace6-b16284c2900e"),
	// 		DataSourceID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/DppPostgresTestingCcy/providers/Microsoft.DBforPostgreSQL/servers/mabtestingccybasicv11/databases/bugbashdb2"),
	// 		DataSourceLocation: to.Ptr("centraluseuap"),
	// 		DataSourceName: to.Ptr("bugbashdb2"),
	// 		DataSourceSetName: to.Ptr("mabtestingccybasicv11"),
	// 		DataSourceType: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
	// 		Duration: to.Ptr("00:00:00"),
	// 		ExtendedInfo: &armdataprotection.JobExtendedInfo{
	// 			AdditionalDetails: map[string]*string{
	// 				"PolicyRuleName": to.Ptr("BackupWeekly"),
	// 				"RetentionTag": to.Ptr("Default"),
	// 				"TaskId": to.Ptr("c4344fb4-7c11-43a4-8307-7ae7c7fb09b9"),
	// 			},
	// 			SubTasks: []*armdataprotection.JobSubTask{
	// 				{
	// 					TaskID: to.Ptr[int32](1),
	// 					TaskName: to.Ptr("Trigger Backup"),
	// 					TaskStatus: to.Ptr("Started"),
	// 			}},
	// 		},
	// 		IsUserTriggered: to.Ptr(false),
	// 		Operation: to.Ptr("Backup"),
	// 		OperationCategory: to.Ptr("Backup"),
	// 		PolicyID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/BugBash1/providers/Microsoft.DataProtection/backupVaults/BugBashVaultForCCYv11/backupPolicies/jakavetdailypolicy"),
	// 		PolicyName: to.Ptr("jakavetdailypolicy"),
	// 		ProgressEnabled: to.Ptr(false),
	// 		SourceResourceGroup: to.Ptr("DppPostgresTestingCcy"),
	// 		SourceSubscriptionID: to.Ptr("62b829ee-7936-40c9-a1c9-47a93f9f3965"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-17T03:00:03.760Z"); return t}()),
	// 		Status: to.Ptr("Started"),
	// 		SubscriptionID: to.Ptr("62b829ee-7936-40c9-a1c9-47a93f9f3965"),
	// 		SupportedActions: []*string{
	// 		},
	// 		VaultName: to.Ptr("BugBashVaultForCCYv11"),
	// 	},
	// }
}
