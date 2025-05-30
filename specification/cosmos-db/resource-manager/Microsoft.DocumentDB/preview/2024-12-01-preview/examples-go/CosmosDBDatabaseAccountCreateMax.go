package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBDatabaseAccountCreateMax.json
func ExampleDatabaseAccountsClient_BeginCreateOrUpdate_cosmosDbDatabaseAccountCreateMax() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseAccountsClient().BeginCreateOrUpdate(ctx, "rg1", "ddb1", armcosmos.DatabaseAccountCreateUpdateParameters{
		Identity: &armcosmos.ManagedServiceIdentity{
			Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armcosmos.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
				"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
			},
		},
		Location: to.Ptr("westus"),
		Tags:     map[string]*string{},
		Kind:     to.Ptr(armcosmos.DatabaseAccountKindMongoDB),
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			AnalyticalStorageConfiguration: &armcosmos.AnalyticalStorageConfiguration{
				SchemaType: to.Ptr(armcosmos.AnalyticalStorageSchemaTypeWellDefined),
			},
			APIProperties: &armcosmos.APIProperties{
				ServerVersion: to.Ptr(armcosmos.ServerVersionThree2),
			},
			BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
				Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
				PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
					BackupIntervalInMinutes:        to.Ptr[int32](240),
					BackupRetentionIntervalInHours: to.Ptr[int32](8),
					BackupStorageRedundancy:        to.Ptr(armcosmos.BackupStorageRedundancyGeo),
				},
			},
			Capacity: &armcosmos.Capacity{
				TotalThroughputLimit: to.Ptr[int32](2000),
			},
			CapacityMode: to.Ptr(armcosmos.CapacityModeProvisioned),
			ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
				DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
				MaxIntervalInSeconds:    to.Ptr[int32](10),
				MaxStalenessPrefix:      to.Ptr[int64](200),
			},
			Cors: []*armcosmos.CorsPolicy{
				{
					AllowedOrigins: to.Ptr("https://test"),
				}},
			CreateMode:                           to.Ptr(armcosmos.CreateModeDefault),
			DatabaseAccountOfferType:             to.Ptr("Standard"),
			DefaultIdentity:                      to.Ptr("FirstPartyIdentity"),
			DefaultPriorityLevel:                 to.Ptr(armcosmos.DefaultPriorityLevelLow),
			EnableAnalyticalStorage:              to.Ptr(true),
			EnableBurstCapacity:                  to.Ptr(true),
			EnableFreeTier:                       to.Ptr(false),
			EnableMaterializedViews:              to.Ptr(false),
			EnablePerRegionPerPartitionAutoscale: to.Ptr(true),
			EnablePriorityBasedExecution:         to.Ptr(true),
			IPRules: []*armcosmos.IPAddressOrRange{
				{
					IPAddressOrRange: to.Ptr("23.43.230.120"),
				},
				{
					IPAddressOrRange: to.Ptr("110.12.240.0/12"),
				}},
			IsVirtualNetworkFilterEnabled: to.Ptr(true),
			KeyVaultKeyURI:                to.Ptr("https://myKeyVault.vault.azure.net"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					IsZoneRedundant:  to.Ptr(false),
					LocationName:     to.Ptr("southcentralus"),
				},
				{
					FailoverPriority: to.Ptr[int32](1),
					IsZoneRedundant:  to.Ptr(false),
					LocationName:     to.Ptr("eastus"),
				}},
			MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTls12),
			NetworkACLBypass:  to.Ptr(armcosmos.NetworkACLBypassAzureServices),
			NetworkACLBypassResourceIDs: []*string{
				to.Ptr("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName")},
			PublicNetworkAccess: to.Ptr(armcosmos.PublicNetworkAccessEnabled),
			VirtualNetworkRules: []*armcosmos.VirtualNetworkRule{
				{
					ID:                               to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
					IgnoreMissingVNetServiceEndpoint: to.Ptr(false),
				}},
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
	// res.DatabaseAccountGetResults = armcosmos.DatabaseAccountGetResults{
	// 	Name: to.Ptr("ddb1"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1"),
	// 	Identity: &armcosmos.ManagedServiceIdentity{
	// 		Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("52f4fef3-3c3f-4ff3-b52e-b5c9eeb68656"),
	// 		TenantID: to.Ptr("33e01921-4d64-4f8c-a055-5bdaffd5e33d"),
	// 		UserAssignedIdentities: map[string]*armcosmos.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 			"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armcosmos.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 				ClientID: to.Ptr("fbe75b66-01c5-4f87-a220-233af3270436"),
	// 				PrincipalID: to.Ptr("075a0ca6-43f6-4434-9abf-c9b1b79f9219"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Kind: to.Ptr(armcosmos.DatabaseAccountKindMongoDB),
	// 	Properties: &armcosmos.DatabaseAccountGetProperties{
	// 		AnalyticalStorageConfiguration: &armcosmos.AnalyticalStorageConfiguration{
	// 			SchemaType: to.Ptr(armcosmos.AnalyticalStorageSchemaTypeWellDefined),
	// 		},
	// 		APIProperties: &armcosmos.APIProperties{
	// 			ServerVersion: to.Ptr(armcosmos.ServerVersionThree2),
	// 		},
	// 		BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
	// 			Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
	// 			PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
	// 				BackupIntervalInMinutes: to.Ptr[int32](240),
	// 				BackupRetentionIntervalInHours: to.Ptr[int32](8),
	// 				BackupStorageRedundancy: to.Ptr(armcosmos.BackupStorageRedundancyGeo),
	// 			},
	// 		},
	// 		Capacity: &armcosmos.Capacity{
	// 			TotalThroughputLimit: to.Ptr[int32](2000),
	// 		},
	// 		CapacityMode: to.Ptr(armcosmos.CapacityModeProvisioned),
	// 		ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
	// 			DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
	// 			MaxIntervalInSeconds: to.Ptr[int32](10),
	// 			MaxStalenessPrefix: to.Ptr[int64](200),
	// 		},
	// 		Cors: []*armcosmos.CorsPolicy{
	// 			{
	// 				AllowedOrigins: to.Ptr("https://test"),
	// 		}},
	// 		CreateMode: to.Ptr(armcosmos.CreateModeDefault),
	// 		DatabaseAccountOfferType: to.Ptr("Standard"),
	// 		DefaultIdentity: to.Ptr("FirstPartyIdentity"),
	// 		DefaultPriorityLevel: to.Ptr(armcosmos.DefaultPriorityLevelLow),
	// 		DisableKeyBasedMetadataWriteAccess: to.Ptr(false),
	// 		EnableAnalyticalStorage: to.Ptr(true),
	// 		EnableBurstCapacity: to.Ptr(true),
	// 		EnableFreeTier: to.Ptr(false),
	// 		EnableMaterializedViews: to.Ptr(false),
	// 		EnablePerRegionPerPartitionAutoscale: to.Ptr(true),
	// 		EnablePriorityBasedExecution: to.Ptr(true),
	// 		FailoverPolicies: []*armcosmos.FailoverPolicy{
	// 			{
	// 				FailoverPriority: to.Ptr[int32](0),
	// 				ID: to.Ptr("ddb1-southcentralus"),
	// 				LocationName: to.Ptr("South Central US"),
	// 			},
	// 			{
	// 				FailoverPriority: to.Ptr[int32](1),
	// 				ID: to.Ptr("ddb1-eastus"),
	// 				LocationName: to.Ptr("East US"),
	// 		}},
	// 		InstanceID: to.Ptr("d9b26648-2f53-4541-b3d8-3044f4f9810d"),
	// 		IPRules: []*armcosmos.IPAddressOrRange{
	// 			{
	// 				IPAddressOrRange: to.Ptr("23.43.230.120"),
	// 			},
	// 			{
	// 				IPAddressOrRange: to.Ptr("110.12.240.0/12"),
	// 		}},
	// 		IsVirtualNetworkFilterEnabled: to.Ptr(true),
	// 		KeyVaultKeyURI: to.Ptr("https://myKeyVault.vault.azure.net"),
	// 		KeysMetadata: &armcosmos.DatabaseAccountKeysMetadata{
	// 			PrimaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
	// 			},
	// 			PrimaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
	// 			},
	// 			SecondaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
	// 			},
	// 			SecondaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
	// 			},
	// 		},
	// 		Locations: []*armcosmos.Location{
	// 			{
	// 				FailoverPriority: to.Ptr[int32](0),
	// 				ID: to.Ptr("ddb1-southcentralus"),
	// 				IsZoneRedundant: to.Ptr(false),
	// 				LocationName: to.Ptr("South Central US"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 			},
	// 			{
	// 				FailoverPriority: to.Ptr[int32](1),
	// 				ID: to.Ptr("ddb1-eastus"),
	// 				IsZoneRedundant: to.Ptr(false),
	// 				LocationName: to.Ptr("East US"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 		}},
	// 		MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTls12),
	// 		NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassAzureServices),
	// 		NetworkACLBypassResourceIDs: []*string{
	// 			to.Ptr("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName")},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			PublicNetworkAccess: to.Ptr(armcosmos.PublicNetworkAccessEnabled),
	// 			ReadLocations: []*armcosmos.Location{
	// 				{
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					IsZoneRedundant: to.Ptr(false),
	// 					LocationName: to.Ptr("South Central US"),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 				},
	// 				{
	// 					FailoverPriority: to.Ptr[int32](1),
	// 					ID: to.Ptr("ddb1-eastus"),
	// 					IsZoneRedundant: to.Ptr(false),
	// 					LocationName: to.Ptr("East US"),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 			}},
	// 			VirtualNetworkRules: []*armcosmos.VirtualNetworkRule{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 			}},
	// 			WriteLocations: []*armcosmos.Location{
	// 				{
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					IsZoneRedundant: to.Ptr(false),
	// 					LocationName: to.Ptr("South Central US"),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 			}},
	// 		},
	// 		SystemData: &armcosmos.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09.000Z"); return t}()),
	// 		},
	// 	}
}
