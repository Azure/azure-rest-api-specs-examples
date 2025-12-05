package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/BackupPolicies_List.json
func ExampleBackupPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupPoliciesClient().NewListPager("myRG", "account1", nil)
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
		// page = armnetapp.BackupPoliciesClientListResponse{
		// 	BackupPoliciesList: armnetapp.BackupPoliciesList{
		// 		Value: []*armnetapp.BackupPolicy{
		// 			{
		// 				Name: to.Ptr("account1/backupPolicy1"),
		// 				Type: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolocies/backupPolicy1"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetapp.BackupPolicyProperties{
		// 					DailyBackupsToKeep: to.Ptr[int32](10),
		// 					Enabled: to.Ptr(true),
		// 					MonthlyBackupsToKeep: to.Ptr[int32](10),
		// 					VolumesAssigned: to.Ptr[int32](0),
		// 					WeeklyBackupsToKeep: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
