package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/AzureIaasVm/BackupPolicies_List.json
func ExampleBackupPoliciesClient_NewListPager_listProtectionPoliciesWithBackupManagementTypeFilterAsAzureIaasVm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupPoliciesClient().NewListPager("NetSDKTestRsVault", "SwaggerTestRg", &armrecoveryservicesbackup.BackupPoliciesClientListOptions{Filter: to.Ptr("backupManagementType eq 'AzureIaasVM'")})
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
		// page.ProtectionPolicyResourceList = armrecoveryservicesbackup.ProtectionPolicyResourceList{
		// 	Value: []*armrecoveryservicesbackup.ProtectionPolicyResource{
		// 		{
		// 			Name: to.Ptr("DefaultPolicy"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupPolicies"),
		// 			ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/NetSDKTestRsVault/backupPolicies/DefaultPolicy"),
		// 			Properties: &armrecoveryservicesbackup.AzureIaaSVMProtectionPolicy{
		// 				BackupManagementType: to.Ptr("AzureIaasVM"),
		// 				ProtectedItemsCount: to.Ptr[int32](0),
		// 				RetentionPolicy: &armrecoveryservicesbackup.LongTermRetentionPolicy{
		// 					RetentionPolicyType: to.Ptr("LongTermRetentionPolicy"),
		// 					DailySchedule: &armrecoveryservicesbackup.DailyRetentionSchedule{
		// 						RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
		// 							Count: to.Ptr[int32](30),
		// 							DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeDays),
		// 						},
		// 						RetentionTimes: []*time.Time{
		// 							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T19:00:00.000Z"); return t}())},
		// 						},
		// 					},
		// 					SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
		// 						SchedulePolicyType: to.Ptr("SimpleSchedulePolicy"),
		// 						ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeDaily),
		// 						ScheduleRunTimes: []*time.Time{
		// 							to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T19:00:00.000Z"); return t}())},
		// 							ScheduleWeeklyFrequency: to.Ptr[int32](0),
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("testPolicy1"),
		// 					Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupPolicies"),
		// 					ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/NetSDKTestRsVault/backupPolicies/testPolicy1"),
		// 					Properties: &armrecoveryservicesbackup.AzureIaaSVMProtectionPolicy{
		// 						BackupManagementType: to.Ptr("AzureIaasVM"),
		// 						ProtectedItemsCount: to.Ptr[int32](0),
		// 						RetentionPolicy: &armrecoveryservicesbackup.LongTermRetentionPolicy{
		// 							RetentionPolicyType: to.Ptr("LongTermRetentionPolicy"),
		// 							DailySchedule: &armrecoveryservicesbackup.DailyRetentionSchedule{
		// 								RetentionDuration: &armrecoveryservicesbackup.RetentionDuration{
		// 									Count: to.Ptr[int32](1),
		// 									DurationType: to.Ptr(armrecoveryservicesbackup.RetentionDurationTypeDays),
		// 								},
		// 								RetentionTimes: []*time.Time{
		// 									to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T02:00:00.000Z"); return t}())},
		// 								},
		// 							},
		// 							SchedulePolicy: &armrecoveryservicesbackup.SimpleSchedulePolicy{
		// 								SchedulePolicyType: to.Ptr("SimpleSchedulePolicy"),
		// 								ScheduleRunFrequency: to.Ptr(armrecoveryservicesbackup.ScheduleRunTypeDaily),
		// 								ScheduleRunTimes: []*time.Time{
		// 									to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-24T02:00:00.000Z"); return t}())},
		// 									ScheduleWeeklyFrequency: to.Ptr[int32](0),
		// 								},
		// 								TimeZone: to.Ptr("Pacific Standard Time"),
		// 							},
		// 					}},
		// 				}
	}
}
