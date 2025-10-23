package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-07-01-preview/BackupPolicies_Update.json
func ExampleBackupPoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupPoliciesClient().BeginUpdate(ctx, "myRG", "account1", "backupPolicyName", armnetapp.BackupPolicyPatch{
		Location: to.Ptr("westus"),
		Properties: &armnetapp.BackupPolicyProperties{
			DailyBackupsToKeep:   to.Ptr[int32](5),
			Enabled:              to.Ptr(false),
			MonthlyBackupsToKeep: to.Ptr[int32](10),
			WeeklyBackupsToKeep:  to.Ptr[int32](10),
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
	// res = armnetapp.BackupPoliciesClientUpdateResponse{
	// 	BackupPolicy: &armnetapp.BackupPolicy{
	// 		Name: to.Ptr("account1/backupPolicyName"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolocies/backupPolicyName"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetapp.BackupPolicyProperties{
	// 			DailyBackupsToKeep: to.Ptr[int32](5),
	// 			Enabled: to.Ptr(false),
	// 			MonthlyBackupsToKeep: to.Ptr[int32](10),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			VolumeBackups: []*armnetapp.VolumeBackups{
	// 				{
	// 					BackupsCount: to.Ptr[int32](5),
	// 					PolicyEnabled: to.Ptr(true),
	// 					VolumeName: to.Ptr("volume 1"),
	// 				},
	// 			},
	// 			VolumesAssigned: to.Ptr[int32](1),
	// 			WeeklyBackupsToKeep: to.Ptr[int32](10),
	// 		},
	// 	},
	// }
}
