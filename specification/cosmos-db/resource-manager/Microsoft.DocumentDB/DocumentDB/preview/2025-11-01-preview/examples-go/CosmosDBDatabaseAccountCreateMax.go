package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBDatabaseAccountCreateMax.json
func ExampleDatabaseAccountsClient_BeginCreateOrUpdate_cosmosDbDatabaseAccountCreateMax() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseAccountsClient().BeginCreateOrUpdate(ctx, "rg1", "ddb1", armcosmos.DatabaseAccountCreateUpdateParameters{
		Location: to.Ptr("westus"),
		Tags:     map[string]*string{},
		Kind:     to.Ptr(armcosmos.DatabaseAccountKindMongoDB),
		Identity: &armcosmos.ManagedServiceIdentity{
			Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armcosmos.ManagedServiceIdentityUserAssignedIdentities{
				"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
			},
		},
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			DatabaseAccountOfferType: to.Ptr("Standard"),
			IPRules: []*armcosmos.IPAddressOrRange{
				{
					IPAddressOrRange: to.Ptr("23.43.230.120"),
				},
				{
					IPAddressOrRange: to.Ptr("110.12.240.0/12"),
				},
			},
			IsVirtualNetworkFilterEnabled: to.Ptr(true),
			VirtualNetworkRules: []*armcosmos.VirtualNetworkRule{
				{
					ID:                               to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
					IgnoreMissingVNetServiceEndpoint: to.Ptr(false),
				},
			},
			PublicNetworkAccess: to.Ptr(armcosmos.PublicNetworkAccessEnabled),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					LocationName:     to.Ptr("southcentralus"),
					IsZoneRedundant:  to.Ptr(false),
				},
				{
					FailoverPriority: to.Ptr[int32](1),
					LocationName:     to.Ptr("eastus"),
					IsZoneRedundant:  to.Ptr(false),
				},
			},
			CreateMode: to.Ptr(armcosmos.CreateModeDefault),
			ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
				DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
				MaxIntervalInSeconds:    to.Ptr[int32](10),
				MaxStalenessPrefix:      to.Ptr[int64](200),
			},
			KeyVaultKeyURI:  to.Ptr("https://myKeyVault.vault.azure.net"),
			DefaultIdentity: to.Ptr("FirstPartyIdentity"),
			EnableFreeTier:  to.Ptr(false),
			APIProperties: &armcosmos.APIProperties{
				ServerVersion: to.Ptr(armcosmos.ServerVersionThree2),
			},
			EnableAnalyticalStorage: to.Ptr(true),
			AnalyticalStorageConfiguration: &armcosmos.AnalyticalStorageConfiguration{
				SchemaType: to.Ptr(armcosmos.AnalyticalStorageSchemaTypeWellDefined),
			},
			BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
				Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
				PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
					BackupIntervalInMinutes:        to.Ptr[int32](240),
					BackupRetentionIntervalInHours: to.Ptr[int32](8),
					BackupStorageRedundancy:        to.Ptr(armcosmos.BackupStorageRedundancyGeo),
				},
			},
			Cors: []*armcosmos.CorsPolicy{
				{
					AllowedOrigins: to.Ptr("https://test"),
				},
			},
			NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassAzureServices),
			NetworkACLBypassResourceIDs: []*string{
				to.Ptr("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName"),
			},
			Capacity: &armcosmos.Capacity{
				TotalThroughputLimit: to.Ptr[int32](2000),
			},
			CapacityMode:                          to.Ptr(armcosmos.CapacityModeProvisioned),
			EnableMaterializedViews:               to.Ptr(false),
			EnableBurstCapacity:                   to.Ptr(true),
			MinimalTLSVersion:                     to.Ptr(armcosmos.MinimalTLSVersionTls12),
			EnablePriorityBasedExecution:          to.Ptr(true),
			DefaultPriorityLevel:                  to.Ptr(armcosmos.DefaultPriorityLevelLow),
			EnablePerRegionPerPartitionAutoscale:  to.Ptr(true),
			EnableAllVersionsAndDeletesChangeFeed: to.Ptr(false),
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
	// res = armcosmos.DatabaseAccountsClientCreateOrUpdateResponse{
	// 	DatabaseAccountGetResults: &armcosmos.DatabaseAccountGetResults{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1"),
	// 		Name: to.Ptr("ddb1"),
	// 		Location: to.Ptr("West US"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts"),
	// 		Kind: to.Ptr(armcosmos.DatabaseAccountKindMongoDB),
	// 		Tags: map[string]*string{
	// 		},
	// 		Identity: &armcosmos.ManagedServiceIdentity{
	// 			Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
	// 			PrincipalID: to.Ptr("52f4fef3-3c3f-4ff3-b52e-b5c9eeb68656"),
	// 			TenantID: to.Ptr("33e01921-4d64-4f8c-a055-5bdaffd5e33d"),
	// 			UserAssignedIdentities: map[string]*armcosmos.ManagedServiceIdentityUserAssignedIdentities{
	// 				"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armcosmos.ManagedServiceIdentityUserAssignedIdentities{
	// 					ClientID: to.Ptr("fbe75b66-01c5-4f87-a220-233af3270436"),
	// 					PrincipalID: to.Ptr("075a0ca6-43f6-4434-9abf-c9b1b79f9219"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armcosmos.DatabaseAccountGetProperties{
	// 			ProvisioningState: to.Ptr("Initializing"),
	// 			IsVirtualNetworkFilterEnabled: to.Ptr(true),
	// 			DatabaseAccountOfferType: to.Ptr("Standard"),
	// 			DisableKeyBasedMetadataWriteAccess: to.Ptr(false),
	// 			ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
	// 				DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
	// 				MaxIntervalInSeconds: to.Ptr[int32](10),
	// 				MaxStalenessPrefix: to.Ptr[int64](200),
	// 			},
	// 			WriteLocations: []*armcosmos.Location{
	// 				{
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					LocationName: to.Ptr("South Central US"),
	// 					ProvisioningState: to.Ptr("Initializing"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					IsZoneRedundant: to.Ptr(false),
	// 				},
	// 			},
	// 			ReadLocations: []*armcosmos.Location{
	// 				{
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					LocationName: to.Ptr("South Central US"),
	// 					ProvisioningState: to.Ptr("Initializing"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					IsZoneRedundant: to.Ptr(false),
	// 				},
	// 				{
	// 					ID: to.Ptr("ddb1-eastus"),
	// 					LocationName: to.Ptr("East US"),
	// 					ProvisioningState: to.Ptr("Initializing"),
	// 					FailoverPriority: to.Ptr[int32](1),
	// 					IsZoneRedundant: to.Ptr(false),
	// 				},
	// 			},
	// 			Locations: []*armcosmos.Location{
	// 				{
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					LocationName: to.Ptr("South Central US"),
	// 					ProvisioningState: to.Ptr("Initializing"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					IsZoneRedundant: to.Ptr(false),
	// 				},
	// 				{
	// 					ID: to.Ptr("ddb1-eastus"),
	// 					LocationName: to.Ptr("East US"),
	// 					ProvisioningState: to.Ptr("Initializing"),
	// 					FailoverPriority: to.Ptr[int32](1),
	// 					IsZoneRedundant: to.Ptr(false),
	// 				},
	// 			},
	// 			FailoverPolicies: []*armcosmos.FailoverPolicy{
	// 				{
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					LocationName: to.Ptr("South Central US"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 				},
	// 				{
	// 					ID: to.Ptr("ddb1-eastus"),
	// 					LocationName: to.Ptr("East US"),
	// 					FailoverPriority: to.Ptr[int32](1),
	// 				},
	// 			},
	// 			VirtualNetworkRules: []*armcosmos.VirtualNetworkRule{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 				},
	// 			},
	// 			PublicNetworkAccess: to.Ptr(armcosmos.PublicNetworkAccessEnabled),
	// 			KeyVaultKeyURI: to.Ptr("https://myKeyVault.vault.azure.net"),
	// 			KeyVaultKeyURIVersion: to.Ptr("009d400efbef459bac31fb86fcce8884"),
	// 			DefaultIdentity: to.Ptr("FirstPartyIdentity"),
	// 			EnableFreeTier: to.Ptr(false),
	// 			IPRules: []*armcosmos.IPAddressOrRange{
	// 				{
	// 					IPAddressOrRange: to.Ptr("23.43.230.120"),
	// 				},
	// 				{
	// 					IPAddressOrRange: to.Ptr("110.12.240.0/12"),
	// 				},
	// 			},
	// 			APIProperties: &armcosmos.APIProperties{
	// 				ServerVersion: to.Ptr(armcosmos.ServerVersionThree2),
	// 			},
	// 			EnableAnalyticalStorage: to.Ptr(true),
	// 			AnalyticalStorageConfiguration: &armcosmos.AnalyticalStorageConfiguration{
	// 				SchemaType: to.Ptr(armcosmos.AnalyticalStorageSchemaTypeWellDefined),
	// 			},
	// 			CreateMode: to.Ptr(armcosmos.CreateModeDefault),
	// 			InstanceID: to.Ptr("d9b26648-2f53-4541-b3d8-3044f4f9810d"),
	// 			BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
	// 				Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
	// 				PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
	// 					BackupIntervalInMinutes: to.Ptr[int32](240),
	// 					BackupRetentionIntervalInHours: to.Ptr[int32](8),
	// 					BackupStorageRedundancy: to.Ptr(armcosmos.BackupStorageRedundancyGeo),
	// 				},
	// 			},
	// 			Cors: []*armcosmos.CorsPolicy{
	// 				{
	// 					AllowedOrigins: to.Ptr("https://test"),
	// 				},
	// 			},
	// 			NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassAzureServices),
	// 			NetworkACLBypassResourceIDs: []*string{
	// 				to.Ptr("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName"),
	// 			},
	// 			Capacity: &armcosmos.Capacity{
	// 				TotalThroughputLimit: to.Ptr[int32](2000),
	// 			},
	// 			CapacityMode: to.Ptr(armcosmos.CapacityModeProvisioned),
	// 			EnableMaterializedViews: to.Ptr(false),
	// 			EnableBurstCapacity: to.Ptr(true),
	// 			MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTls12),
	// 			KeysMetadata: &armcosmos.DatabaseAccountKeysMetadata{
	// 				PrimaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 				},
	// 				SecondaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 				},
	// 				PrimaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 				},
	// 				SecondaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 				},
	// 			},
	// 			EnablePriorityBasedExecution: to.Ptr(true),
	// 			DefaultPriorityLevel: to.Ptr(armcosmos.DefaultPriorityLevelLow),
	// 			EnablePerRegionPerPartitionAutoscale: to.Ptr(true),
	// 			ThroughputPoolDedicatedRUs: to.Ptr[int64](0),
	// 			ThroughputPoolMaxConsumableRUs: to.Ptr[int64](0),
	// 		},
	// 		SystemData: &armcosmos.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09Z"); return t}()),
	// 		},
	// 	},
	// }
}
