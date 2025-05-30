package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/Common/BackupProtectionContainers_UsageSummary_Get.json
func ExampleBackupUsageSummariesClient_NewListPager_getProtectedContainersUsagesSummary() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupUsageSummariesClient().NewListPager("testVault", "testRG", &armrecoveryservicesbackup.BackupUsageSummariesClientListOptions{Filter: to.Ptr("type eq 'BackupProtectionContainerCountSummary'"),
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
		// page.BackupManagementUsageList = armrecoveryservicesbackup.BackupManagementUsageList{
		// 	Value: []*armrecoveryservicesbackup.BackupManagementUsage{
		// 		{
		// 			Name: &armrecoveryservicesbackup.NameInfo{
		// 				LocalizedValue: to.Ptr("Azure Backup Server"),
		// 				Value: to.Ptr("AzureBackupServer"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](2),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservicesbackup.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservicesbackup.NameInfo{
		// 				LocalizedValue: to.Ptr("Azure Backup Agent"),
		// 				Value: to.Ptr("MAB"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](3),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservicesbackup.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservicesbackup.NameInfo{
		// 				LocalizedValue: to.Ptr("SQL in Azure VM"),
		// 				Value: to.Ptr("AzureWorkload"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](1),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservicesbackup.UsagesUnitCount),
		// 	}},
		// }
	}
}
