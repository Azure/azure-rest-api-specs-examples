package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/Common/ListJobsWithAllSupportedFilters.json
func ExampleBackupJobsClient_NewListPager_listJobsWithFilters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupJobsClient().NewListPager("NetSDKTestRsVault", "SwaggerTestRg", &armrecoveryservicesbackup.BackupJobsClientListOptions{Filter: to.Ptr("startTime eq '2016-01-01 00:00:00 AM' and endTime eq '2017-11-29 00:00:00 AM' and operation eq 'Backup' and backupManagementType eq 'AzureIaasVM' and status eq 'InProgress'"),
		SkipToken: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.JobResourceList = armrecoveryservicesbackup.JobResourceList{
		// 	Value: []*armrecoveryservicesbackup.JobResource{
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupJobs"),
		// 			ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/NetSDKTestRsVault/backupJobs/00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armrecoveryservicesbackup.AzureIaaSVMJob{
		// 				ActivityID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureIaasVM),
		// 				EntityFriendlyName: to.Ptr("testvm"),
		// 				JobType: to.Ptr("AzureIaaSVMJob"),
		// 				Operation: to.Ptr("Backup"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-03T05:31:07.014Z"); return t}()),
		// 				Status: to.Ptr("InProgress"),
		// 				Duration: to.Ptr("PT12.4272909S"),
		// 				VirtualMachineVersion: to.Ptr("Compute"),
		// 			},
		// 	}},
		// }
	}
}
