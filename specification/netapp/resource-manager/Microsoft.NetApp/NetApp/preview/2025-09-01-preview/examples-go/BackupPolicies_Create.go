package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/BackupPolicies_Create.json
func ExampleBackupPoliciesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupPoliciesClient().BeginCreate(ctx, "myRG", "account1", "backupPolicyName", armnetapp.BackupPolicy{
		Location: to.Ptr("westus"),
		Properties: &armnetapp.BackupPolicyProperties{
			DailyBackupsToKeep:   to.Ptr[int32](10),
			Enabled:              to.Ptr(true),
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
	// res = armnetapp.BackupPoliciesClientCreateResponse{
	// 	BackupPolicy: &armnetapp.BackupPolicy{
	// 		Name: to.Ptr("account1/backupPolicyName"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolicies/backupPolicyName"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetapp.BackupPolicyProperties{
	// 			DailyBackupsToKeep: to.Ptr[int32](10),
	// 			Enabled: to.Ptr(true),
	// 			MonthlyBackupsToKeep: to.Ptr[int32](10),
	// 			ProvisioningState: to.Ptr("creating"),
	// 			WeeklyBackupsToKeep: to.Ptr[int32](10),
	// 		},
	// 	},
	// }
}
