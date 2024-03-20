package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-02-15-preview/examples/CosmosDBDatabaseAccountPatch.json
func ExampleDatabaseAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseAccountsClient().BeginUpdate(ctx, "rg1", "ddb1", armcosmos.DatabaseAccountUpdateParameters{
		Identity: &armcosmos.ManagedServiceIdentity{
			Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armcosmos.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
				"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
			},
		},
		Location: to.Ptr("westus"),
		Properties: &armcosmos.DatabaseAccountUpdateProperties{
			AnalyticalStorageConfiguration: &armcosmos.AnalyticalStorageConfiguration{
				SchemaType: to.Ptr(armcosmos.AnalyticalStorageSchemaTypeWellDefined),
			},
			BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
				Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
				PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
					BackupIntervalInMinutes:        to.Ptr[int32](240),
					BackupRetentionIntervalInHours: to.Ptr[int32](720),
					BackupStorageRedundancy:        to.Ptr(armcosmos.BackupStorageRedundancyGeo),
				},
			},
			Capacity: &armcosmos.Capacity{
				TotalThroughputLimit: to.Ptr[int32](2000),
			},
			ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
				DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
				MaxIntervalInSeconds:    to.Ptr[int32](10),
				MaxStalenessPrefix:      to.Ptr[int64](200),
			},
			DefaultIdentity:      to.Ptr("FirstPartyIdentity"),
			DefaultPriorityLevel: to.Ptr(armcosmos.DefaultPriorityLevelLow),
			DiagnosticLogSettings: &armcosmos.DiagnosticLogSettings{
				EnableFullTextQuery: to.Ptr(armcosmos.EnableFullTextQueryTrue),
			},
			EnableAnalyticalStorage:              to.Ptr(true),
			EnableBurstCapacity:                  to.Ptr(true),
			EnableFreeTier:                       to.Ptr(false),
			EnablePartitionMerge:                 to.Ptr(true),
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
			MinimalTLSVersion:             to.Ptr(armcosmos.MinimalTLSVersionTLS),
			NetworkACLBypass:              to.Ptr(armcosmos.NetworkACLBypassAzureServices),
			NetworkACLBypassResourceIDs: []*string{
				to.Ptr("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName")},
			VirtualNetworkRules: []*armcosmos.VirtualNetworkRule{
				{
					ID:                               to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
					IgnoreMissingVNetServiceEndpoint: to.Ptr(false),
				}},
		},
		Tags: map[string]*string{
			"dept": to.Ptr("finance"),
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
	// 		"dept": to.Ptr("finance"),
	// 	},
	// 	Kind: to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
	// 	Properties: &armcosmos.DatabaseAccountGetProperties{
	// 		AnalyticalStorageConfiguration: &armcosmos.AnalyticalStorageConfiguration{
	// 			SchemaType: to.Ptr(armcosmos.AnalyticalStorageSchemaTypeWellDefined),
	// 		},
	// 		APIProperties: &armcosmos.APIProperties{
	// 		},
	// 		BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
	// 			Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
	// 			PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
	// 				BackupIntervalInMinutes: to.Ptr[int32](240),
	// 				BackupRetentionIntervalInHours: to.Ptr[int32](720),
	// 				BackupStorageRedundancy: to.Ptr(armcosmos.BackupStorageRedundancyGeo),
	// 			},
	// 		},
	// 		Capacity: &armcosmos.Capacity{
	// 			TotalThroughputLimit: to.Ptr[int32](2000),
	// 		},
	// 		ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
	// 			DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
	// 			MaxIntervalInSeconds: to.Ptr[int32](10),
	// 			MaxStalenessPrefix: to.Ptr[int64](200),
	// 		},
	// 		Cors: []*armcosmos.CorsPolicy{
	// 		},
	// 		CreateMode: to.Ptr(armcosmos.CreateModeDefault),
	// 		DatabaseAccountOfferType: to.Ptr("Standard"),
	// 		DefaultIdentity: to.Ptr("FirstPartyIdentity"),
	// 		DefaultPriorityLevel: to.Ptr(armcosmos.DefaultPriorityLevelLow),
	// 		DiagnosticLogSettings: &armcosmos.DiagnosticLogSettings{
	// 			EnableFullTextQuery: to.Ptr(armcosmos.EnableFullTextQueryTrue),
	// 		},
	// 		DisableKeyBasedMetadataWriteAccess: to.Ptr(false),
	// 		DocumentEndpoint: to.Ptr("https://ddb1.documents.azure.com:443/"),
	// 		EnableAnalyticalStorage: to.Ptr(true),
	// 		EnableBurstCapacity: to.Ptr(true),
	// 		EnableFreeTier: to.Ptr(false),
	// 		EnableMaterializedViews: to.Ptr(false),
	// 		EnablePerRegionPerPartitionAutoscale: to.Ptr(true),
	// 		EnablePriorityBasedExecution: to.Ptr(true),
	// 		FailoverPolicies: []*armcosmos.FailoverPolicy{
	// 			{
	// 				FailoverPriority: to.Ptr[int32](0),
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
	// 				DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
	// 				FailoverPriority: to.Ptr[int32](0),
	// 				ID: to.Ptr("ddb1-eastus"),
	// 				LocationName: to.Ptr("East US"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 		}},
	// 		MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTLS),
	// 		NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassAzureServices),
	// 		NetworkACLBypassResourceIDs: []*string{
	// 			to.Ptr("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName")},
	// 			PrivateEndpointConnections: []*armcosmos.PrivateEndpointConnection{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.DocumentDB/databaseAccounts/account1/privateEndpointConnections/pe1"),
	// 					Properties: &armcosmos.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armcosmos.PrivateEndpointProperty{
	// 							ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/privateEndpoints/pe1"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armcosmos.PrivateLinkServiceConnectionStateProperty{
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr("Approved"),
	// 						},
	// 					},
	// 			}},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ReadLocations: []*armcosmos.Location{
	// 				{
	// 					DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					ID: to.Ptr("ddb1-eastus"),
	// 					LocationName: to.Ptr("East US"),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 			}},
	// 			VirtualNetworkRules: []*armcosmos.VirtualNetworkRule{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 					IgnoreMissingVNetServiceEndpoint: to.Ptr(false),
	// 			}},
	// 			WriteLocations: []*armcosmos.Location{
	// 				{
	// 					DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					ID: to.Ptr("ddb1-eastus"),
	// 					LocationName: to.Ptr("East US"),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 			}},
	// 		},
	// 		SystemData: &armcosmos.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09.000Z"); return t}()),
	// 		},
	// 	}
}
