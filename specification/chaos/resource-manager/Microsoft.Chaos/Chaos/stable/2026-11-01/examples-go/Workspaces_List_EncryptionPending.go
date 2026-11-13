package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/Workspaces_List_EncryptionPending.json
func ExampleWorkspacesClient_NewListPager_listCmkCreatedWorkspacesWithCustomerKeyReplacementPending() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListPager("exampleRG", nil)
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
		// page = armchaos.WorkspacesClientListResponse{
		// 	WorkspaceListResult: armchaos.WorkspaceListResult{
		// 		Value: []*armchaos.Workspace{
		// 			{
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace"),
		// 				Name: to.Ptr("exampleWorkspace"),
		// 				Type: to.Ptr("Microsoft.Chaos/workspaces"),
		// 				Location: to.Ptr("westus2"),
		// 				Properties: &armchaos.WorkspaceProperties{
		// 					Scopes: []*string{
		// 						to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG"),
		// 					},
		// 					ProvisioningState: to.Ptr(armchaos.ProvisioningStateUpdating),
		// 					Encryption: &armchaos.WorkspaceEncryption{
		// 						CustomerManagedKeyEncryption: &armchaos.WorkspaceCustomerManagedKeyEncryption{
		// 							KeyEncryptionKeyURL: to.Ptr("https://contoso-vault.vault.azure.net/keys/workspace-key"),
		// 							KeyVaultResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/customer-keys/providers/Microsoft.KeyVault/vaults/contoso-vault"),
		// 						},
		// 						Status: &armchaos.WorkspaceEncryptionStatus{
		// 							State: to.Ptr(armchaos.WorkspaceEncryptionStatePending),
		// 							ObservedProtection: to.Ptr(armchaos.WorkspaceEncryptionProtectionCustomerManaged),
		// 							ObservedKeyEncryptionKeyURL: to.Ptr("https://contoso-vault.vault.azure.net/keys/previous-key/0123456789abcdef0123456789abcdef"),
		// 							LastCheckedAt: to.Ptr(time.Date(2026, time.November, 1, 11, 0, 0, 0, time.UTC)),
		// 						},
		// 						CustomerManagedKeyOnboarding: &armchaos.WorkspaceCustomerManagedKeyOnboarding{
		// 							ApplicationID: to.Ptr("00000000-0000-0000-0000-000000000001"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
