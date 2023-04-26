package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/75ece9197dbac70ac0ba651c53a79c1841944be2/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/examples/ServerUpdateWithDataEncryptionEnabled.json
func ExampleServersClient_BeginUpdate_serverUpdateWithDataEncryptionEnabled() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginUpdate(ctx, "TestGroup", "pgtestsvc4", armpostgresqlflexibleservers.ServerForUpdate{
		Identity: &armpostgresqlflexibleservers.UserAssignedIdentity{
			Type: to.Ptr(armpostgresqlflexibleservers.IdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armpostgresqlflexibleservers.UserIdentity{
				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity": {},
			},
		},
		Properties: &armpostgresqlflexibleservers.ServerPropertiesForUpdate{
			AdministratorLoginPassword: to.Ptr("newpassword"),
			Backup: &armpostgresqlflexibleservers.Backup{
				BackupRetentionDays: to.Ptr[int32](20),
			},
			CreateMode: to.Ptr(armpostgresqlflexibleservers.CreateModeForUpdateUpdate),
			DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
				Type:                          to.Ptr(armpostgresqlflexibleservers.ArmServerKeyTypeAzureKeyVault),
				PrimaryKeyURI:                 to.Ptr("https://test-kv.vault.azure.net/keys/test-key1/77f57315bab34b0189daa113fbc78787"),
				PrimaryUserAssignedIdentityID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity"),
			},
			Storage: &armpostgresqlflexibleservers.Storage{
				StorageSizeGB: to.Ptr[int32](1024),
			},
		},
		SKU: &armpostgresqlflexibleservers.SKU{
			Name: to.Ptr("Standard_D8s_v3"),
			Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
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
	// res.Server = armpostgresqlflexibleservers.Server{
	// 	Name: to.Ptr("pgtestsvc4"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc4"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"ElasticServer": to.Ptr("1"),
	// 	},
	// 	Identity: &armpostgresqlflexibleservers.UserAssignedIdentity{
	// 		Type: to.Ptr(armpostgresqlflexibleservers.IdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armpostgresqlflexibleservers.UserIdentity{
	// 			"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity": &armpostgresqlflexibleservers.UserIdentity{
	// 				ClientID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 				PrincipalID: to.Ptr("0a4e0c6e-7751-4078-ae1f-a477306c11e9"),
	// 			},
	// 			"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity-1": &armpostgresqlflexibleservers.UserIdentity{
	// 				ClientID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 				PrincipalID: to.Ptr("90008082-e024-4cc3-8fcf-63bcdb9cf6b6"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armpostgresqlflexibleservers.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
	// 			ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.ActiveDirectoryAuthEnumDisabled),
	// 			PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordAuthEnumEnabled),
	// 		},
	// 		AvailabilityZone: to.Ptr("1"),
	// 		Backup: &armpostgresqlflexibleservers.Backup{
	// 			BackupRetentionDays: to.Ptr[int32](20),
	// 			EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-26T01:16:58.3723361+00:00"); return t}()),
	// 			GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumDisabled),
	// 		},
	// 		DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
	// 			Type: to.Ptr(armpostgresqlflexibleservers.ArmServerKeyTypeAzureKeyVault),
	// 			PrimaryKeyURI: to.Ptr("https://test-kv.vault.azure.net/keys/test-key1/77f57315bab34b0189daa113fbc78787"),
	// 			PrimaryUserAssignedIdentityID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity"),
	// 		},
	// 		FullyQualifiedDomainName: to.Ptr("c7d7483a8ceb.test-private-dns-zone.postgres.database.azure.com"),
	// 		HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
	// 			Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityModeZoneRedundant),
	// 			StandbyAvailabilityZone: to.Ptr("2"),
	// 			State: to.Ptr(armpostgresqlflexibleservers.ServerHAStateHealthy),
	// 		},
	// 		MaintenanceWindow: &armpostgresqlflexibleservers.MaintenanceWindow{
	// 			CustomWindow: to.Ptr("Disabled"),
	// 			DayOfWeek: to.Ptr[int32](0),
	// 			StartHour: to.Ptr[int32](0),
	// 			StartMinute: to.Ptr[int32](0),
	// 		},
	// 		MinorVersion: to.Ptr("6"),
	// 		Network: &armpostgresqlflexibleservers.Network{
	// 			DelegatedSubnetResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-vnet-subnet"),
	// 			PrivateDNSZoneArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/privateDnsZones/test-private-dns-zone.postgres.database.azure.com"),
	// 			PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateDisabled),
	// 		},
	// 		State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
	// 		Storage: &armpostgresqlflexibleservers.Storage{
	// 			StorageSizeGB: to.Ptr[int32](1024),
	// 		},
	// 		Version: to.Ptr(armpostgresqlflexibleservers.ServerVersionTwelve),
	// 	},
	// 	SKU: &armpostgresqlflexibleservers.SKU{
	// 		Name: to.Ptr("Standard_D8s_v3"),
	// 		Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
	// 	},
	// }
}
