package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-01-01-preview/ServersUpdateWithDataEncryptionEnabledAutoUpdate.json
func ExampleServersClient_BeginUpdate_updateAnExistingServerWithDataEncryptionBasedOnCustomerManagedKeyWithAutomaticKeyVersionUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginUpdate(ctx, "exampleresourcegroup", "exampleserver", armpostgresqlflexibleservers.ServerForPatch{
		Identity: &armpostgresqlflexibleservers.UserAssignedIdentity{
			Type: to.Ptr(armpostgresqlflexibleservers.IdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armpostgresqlflexibleservers.UserIdentity{
				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplegeoredundantidentity": {},
				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity":      {},
			},
		},
		Properties: &armpostgresqlflexibleservers.ServerPropertiesForPatch{
			AdministratorLoginPassword: to.Ptr("examplenewpassword"),
			Backup: &armpostgresqlflexibleservers.BackupForPatch{
				BackupRetentionDays: to.Ptr[int32](20),
			},
			CreateMode: to.Ptr(armpostgresqlflexibleservers.CreateModeForPatchUpdate),
			DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
				Type:                            to.Ptr(armpostgresqlflexibleservers.DataEncryptionTypeAzureKeyVault),
				GeoBackupKeyURI:                 to.Ptr("https://examplegeoredundantkeyvault.vault.azure.net/keys/examplekey"),
				GeoBackupUserAssignedIdentityID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplegeoredundantidentity"),
				PrimaryKeyURI:                   to.Ptr("https://exampleprimarykeyvault.vault.azure.net/keys/examplekey"),
				PrimaryUserAssignedIdentityID:   to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity"),
			},
		},
		SKU: &armpostgresqlflexibleservers.SKUForPatch{
			Name: to.Ptr("Standard_D8s_v3"),
			Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
