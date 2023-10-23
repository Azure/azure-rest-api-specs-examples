package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1a87e1a5deb3f986ea1474d233d6680f1e19fc1/specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspaceManagedDiskEncryptionCreate.json
func ExampleWorkspacesClient_BeginCreateOrUpdate_createAWorkspaceWithCustomerManagedKeyCmkEncryptionForManagedDisks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginCreateOrUpdate(ctx, "rg", "myWorkspace", armdatabricks.Workspace{
		Location: to.Ptr("westus"),
		Properties: &armdatabricks.WorkspaceProperties{
			Encryption: &armdatabricks.WorkspacePropertiesEncryption{
				Entities: &armdatabricks.EncryptionEntitiesDefinition{
					ManagedDisk: &armdatabricks.ManagedDiskEncryption{
						KeySource: to.Ptr(armdatabricks.EncryptionKeySourceMicrosoftKeyvault),
						KeyVaultProperties: &armdatabricks.ManagedDiskEncryptionKeyVaultProperties{
							KeyName:     to.Ptr("test-cmk-key"),
							KeyVaultURI: to.Ptr("https://test-vault-name.vault.azure.net/"),
							KeyVersion:  to.Ptr("00000000000000000000000000000000"),
						},
						RotationToLatestKeyVersionEnabled: to.Ptr(true),
					},
				},
			},
			ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
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
	// res.Workspace = armdatabricks.Workspace{
	// 	Type: to.Ptr("Microsoft.Databricks/workspaces"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Databricks/workspaces/myWorkspace"),
	// 	Location: to.Ptr("East US 2"),
	// 	Properties: &armdatabricks.WorkspaceProperties{
	// 		Authorizations: []*armdatabricks.WorkspaceProviderAuthorization{
	// 			{
	// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				RoleDefinitionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		}},
	// 		CreatedBy: &armdatabricks.CreatedBy{
	// 			ApplicationID: to.Ptr("44444444-4444-4444-4444-444444444444"),
	// 			Oid: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 			Puid: to.Ptr("33333333"),
	// 		},
	// 		CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-20T00:10:29.2858439Z"); return t}()),
	// 		DiskEncryptionSetID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 		Encryption: &armdatabricks.WorkspacePropertiesEncryption{
	// 			Entities: &armdatabricks.EncryptionEntitiesDefinition{
	// 				ManagedDisk: &armdatabricks.ManagedDiskEncryption{
	// 					KeySource: to.Ptr(armdatabricks.EncryptionKeySourceMicrosoftKeyvault),
	// 					KeyVaultProperties: &armdatabricks.ManagedDiskEncryptionKeyVaultProperties{
	// 						KeyName: to.Ptr("test-cmk-key"),
	// 						KeyVaultURI: to.Ptr("https://test-vault-name.vault.azure.net/"),
	// 						KeyVersion: to.Ptr("00000000000000000000000000000000"),
	// 					},
	// 					RotationToLatestKeyVersionEnabled: to.Ptr(true),
	// 				},
	// 			},
	// 		},
	// 		ManagedDiskIdentity: &armdatabricks.ManagedIdentityConfiguration{
	// 			Type: to.Ptr("SystemAssigned"),
	// 			PrincipalID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 			TenantID: to.Ptr("44444444-4444-4444-4444-444444444444"),
	// 		},
	// 		ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
	// 		ProvisioningState: to.Ptr(armdatabricks.ProvisioningStateSucceeded),
	// 		UIDefinitionURI: to.Ptr("https://path/to/workspaceCreateUiDefinition.json"),
	// 		UpdatedBy: &armdatabricks.CreatedBy{
	// 			ApplicationID: to.Ptr("44444444-4444-4444-4444-444444444444"),
	// 			Oid: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 			Puid: to.Ptr("33333333"),
	// 		},
	// 		WorkspaceID: to.Ptr("5555555555555555"),
	// 		WorkspaceURL: to.Ptr("adb-5555555555555555.19.azuredatabricks.net"),
	// 	},
	// 	SKU: &armdatabricks.SKU{
	// 		Name: to.Ptr("premium"),
	// 	},
	// }
}
