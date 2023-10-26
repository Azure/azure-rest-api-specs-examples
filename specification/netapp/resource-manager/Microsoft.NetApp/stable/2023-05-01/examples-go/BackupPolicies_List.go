package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/BackupPolicies_List.json
func ExampleBackupPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.BackupPoliciesList = armnetapp.BackupPoliciesList{
		// 	Value: []*armnetapp.BackupPolicy{
		// 		{
		// 			Name: to.Ptr("account1/backupPolicy1"),
		// 			Type: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolocies/backupPolicy1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetapp.BackupPolicyProperties{
		// 				DailyBackupsToKeep: to.Ptr[int32](10),
		// 				Enabled: to.Ptr(true),
		// 				MonthlyBackupsToKeep: to.Ptr[int32](10),
		// 				VolumesAssigned: to.Ptr[int32](0),
		// 				WeeklyBackupsToKeep: to.Ptr[int32](10),
		// 			},
		// 	}},
		// }
	}
}
