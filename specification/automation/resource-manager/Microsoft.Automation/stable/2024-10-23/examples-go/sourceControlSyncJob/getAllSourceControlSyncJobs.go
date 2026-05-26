package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/sourceControlSyncJob/getAllSourceControlSyncJobs.json
func ExampleSourceControlSyncJobClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSourceControlSyncJobClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", "MySourceControl", nil)
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
		// page = armautomation.SourceControlSyncJobClientListByAutomationAccountResponse{
		// 	SourceControlSyncJobListResult: armautomation.SourceControlSyncJobListResult{
		// 		Value: []*armautomation.SourceControlSyncJob{
		// 			{
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/sourceControls/MySourceControl/sourceControlSyncJobs/ce6fe3e3-9db3-4096-a6b4-82bfb4c10a1a"),
		// 				Properties: &armautomation.SourceControlSyncJobProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903+00:00"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:28.903+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armautomation.ProvisioningStateCompleted),
		// 					SourceControlSyncJobID: to.Ptr("ce6fe3e3-9db3-4096-a6b4-82bfb4c10a1a"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:27.903+00:00"); return t}()),
		// 					SyncType: to.Ptr(armautomation.SyncTypePartialSync),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/sourceControls/MySourceControl/sourceControlSyncJobs/ce6fe3e3-9db3-4096-a6b4-82bfb4c10a2a"),
		// 				Properties: &armautomation.SourceControlSyncJobProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903+00:00"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:28.903+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armautomation.ProvisioningStateCompleted),
		// 					SourceControlSyncJobID: to.Ptr("ce6fe3e3-9db3-4096-a6b4-82bfb4c10a2a"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:27.903+00:00"); return t}()),
		// 					SyncType: to.Ptr(armautomation.SyncTypePartialSync),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/sourceControls/MySourceControl/sourceControlSyncJobs/ce6fe3e3-9db3-4096-a6b4-82bfb4c10a3a"),
		// 				Properties: &armautomation.SourceControlSyncJobProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903+00:00"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:28.903+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armautomation.ProvisioningStateCompleted),
		// 					SourceControlSyncJobID: to.Ptr("ce6fe3e3-9db3-4096-a6b4-82bfb4c10a3a"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:27.903+00:00"); return t}()),
		// 					SyncType: to.Ptr(armautomation.SyncTypePartialSync),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/sourceControls/MySourceControl/sourceControlSyncJobs/ce6fe3e3-9db3-4096-a6b4-82bfb4c10a4a"),
		// 				Properties: &armautomation.SourceControlSyncJobProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903+00:00"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:28.903+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armautomation.ProvisioningStateCompleted),
		// 					SourceControlSyncJobID: to.Ptr("ce6fe3e3-9db3-4096-a6b4-82bfb4c10a4a"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:27.903+00:00"); return t}()),
		// 					SyncType: to.Ptr(armautomation.SyncTypePartialSync),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/sourceControls/MySourceControl/sourceControlSyncJobs/ce6fe3e3-9db3-4096-a6b4-82bfb4c10a5a"),
		// 				Properties: &armautomation.SourceControlSyncJobProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903+00:00"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:28.903+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armautomation.ProvisioningStateCompleted),
		// 					SourceControlSyncJobID: to.Ptr("ce6fe3e3-9db3-4096-a6b4-82bfb4c10a5a"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:27.903+00:00"); return t}()),
		// 					SyncType: to.Ptr(armautomation.SyncTypePartialSync),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
