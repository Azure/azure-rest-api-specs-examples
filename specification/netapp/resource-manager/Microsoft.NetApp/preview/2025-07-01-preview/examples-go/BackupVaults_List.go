package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-07-01-preview/BackupVaults_List.json
func ExampleBackupVaultsClient_NewListByNetAppAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupVaultsClient().NewListByNetAppAccountPager("myRG", "account1", nil)
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
		// page = armnetapp.BackupVaultsClientListByNetAppAccountResponse{
		// 	BackupVaultsList: armnetapp.BackupVaultsList{
		// 		Value: []*armnetapp.BackupVault{
		// 			{
		// 				Name: to.Ptr("account1/backupVault1"),
		// 				Type: to.Ptr("Microsoft.NetApp/netAppAccounts/backupVaults"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupVaults/backupVault1"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetapp.BackupVaultProperties{
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Tag1": to.Ptr("Value1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
