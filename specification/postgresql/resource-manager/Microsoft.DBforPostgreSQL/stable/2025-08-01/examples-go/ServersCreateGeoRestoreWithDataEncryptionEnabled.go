package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ServersCreateGeoRestoreWithDataEncryptionEnabled.json
func ExampleServersClient_BeginCreateOrUpdate_createANewServerUsingARestoreOfAGeographicallyRedundantBackupOfAnExistingServerWithDataEncryptionBasedOnCustomerManagedKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreateOrUpdate(ctx, "exampleresourcegroup", "exampleserver", armpostgresqlflexibleservers.Server{
		Location: to.Ptr("eastus"),
		Identity: &armpostgresqlflexibleservers.UserAssignedIdentity{
			Type: to.Ptr(armpostgresqlflexibleservers.IdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armpostgresqlflexibleservers.UserIdentity{
				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplegeoredundantidentity": {},
				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity":      {},
			},
		},
		Properties: &armpostgresqlflexibleservers.ServerProperties{
			CreateMode: to.Ptr(armpostgresqlflexibleservers.CreateModeGeoRestore),
			DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
				Type:                            to.Ptr(armpostgresqlflexibleservers.DataEncryptionTypeAzureKeyVault),
				GeoBackupKeyURI:                 to.Ptr("https://examplegeoredundantkeyvault.vault.azure.net/keys/examplekey/yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy"),
				GeoBackupUserAssignedIdentityID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplegeoredundantidentity"),
				PrimaryKeyURI:                   to.Ptr("https://exampleprimarykeyvault.vault.azure.net/keys/examplekey/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				PrimaryUserAssignedIdentityID:   to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity"),
			},
			PointInTimeUTC:         to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T18:35:22.123Z"); return t }()),
			SourceServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/examplesourceserver"),
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
