package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBDatabaseAccountPatch.json
func ExampleDatabaseAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseAccountsClient().BeginUpdate(ctx, "rg1", "ddb1", armcosmos.DatabaseAccountUpdateParameters{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"dept": to.Ptr("finance"),
		},
		Identity: &armcosmos.ManagedServiceIdentity{
			Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armcosmos.ManagedServiceIdentityUserAssignedIdentities{
				"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
			},
		},
		Properties: &armcosmos.DatabaseAccountUpdateProperties{
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
			ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
				DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
				MaxIntervalInSeconds:    to.Ptr[int32](10),
				MaxStalenessPrefix:      to.Ptr[int64](200),
			},
			DefaultIdentity:                      to.Ptr("FirstPartyIdentity"),
			EnableFreeTier:                       to.Ptr(false),
			EnableAnalyticalStorage:              to.Ptr(true),
			EnableBurstCapacity:                  to.Ptr(true),
			EnablePriorityBasedExecution:         to.Ptr(true),
			DefaultPriorityLevel:                 to.Ptr(armcosmos.DefaultPriorityLevelLow),
			EnablePerRegionPerPartitionAutoscale: to.Ptr(true),
			AnalyticalStorageConfiguration: &armcosmos.AnalyticalStorageConfiguration{
				SchemaType: to.Ptr(armcosmos.AnalyticalStorageSchemaTypeWellDefined),
			},
			BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
				Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
				PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
					BackupIntervalInMinutes:        to.Ptr[int32](240),
					BackupRetentionIntervalInHours: to.Ptr[int32](720),
					BackupStorageRedundancy:        to.Ptr(armcosmos.BackupStorageRedundancyLocal),
				},
			},
			NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassAzureServices),
			NetworkACLBypassResourceIDs: []*string{
				to.Ptr("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName"),
			},
			Capacity: &armcosmos.Capacity{
				TotalThroughputLimit: to.Ptr[int32](2000),
			},
			EnablePartitionMerge:                       to.Ptr(true),
			EnforceHierarchicalPartitionKeyIDLastLevel: to.Ptr(false),
			MinimalTLSVersion:                          to.Ptr(armcosmos.MinimalTLSVersionTLS),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.DatabaseAccountsClientUpdateResponse{
	// 	DatabaseAccountGetResults: armcosmos.DatabaseAccountGetResults{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1"),
	// 		Name: to.Ptr("ddb1"),
	// 		Location: to.Ptr("West US"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts"),
	// 		Kind: to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
	// 		Tags: map[string]*string{
	// 			"dept": to.Ptr("finance"),
	// 		},
	// 		Identity: &armcosmos.ManagedServiceIdentity{
	// 			Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
	// 			PrincipalID: to.Ptr("52f4fef3-3c3f-4ff3-b52e-b5c9eeb68656"),
	// 			TenantID: to.Ptr("33e01921-4d64-4f8c-a055-5bdaffd5e33d"),
	// 			UserAssignedIdentities: map[string]*armcosmos.ManagedServiceIdentityUserAssignedIdentities{
	// 				"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armcosmos.ManagedServiceIdentityUserAssignedIdentities{
	// 					ClientID: to.Ptr("fbe75b66-01c5-4f87-a220-233af3270436"),
	// 					PrincipalID: to.Ptr("33e01921-4d64-4f8c-a055-5bdaffd5e33d"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armcosmos.DatabaseAccountGetProperties{
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			DocumentEndpoint: to.Ptr("https://ddb1.documents.azure.com:443/"),
	// 			DatabaseAccountOfferType: to.Ptr("Standard"),
	// 			IPRules: []*armcosmos.IPAddressOrRange{
	// 				{
	// 					IPAddressOrRange: to.Ptr("23.43.230.120"),
	// 				},
	// 				{
	// 					IPAddressOrRange: to.Ptr("110.12.240.0/12"),
	// 				},
	// 			},
	// 			IsVirtualNetworkFilterEnabled: to.Ptr(true),
	// 			DisableKeyBasedMetadataWriteAccess: to.Ptr(false),
	// 			ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
	// 				DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
	// 				MaxIntervalInSeconds: to.Ptr[int32](10),
	// 				MaxStalenessPrefix: to.Ptr[int64](200),
	// 			},
	// 			WriteLocations: []*armcosmos.Location{
	// 				{
	// 					ID: to.Ptr("ddb1-eastus"),
	// 					LocationName: to.Ptr("East US"),
	// 					DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 				},
	// 			},
	// 			ReadLocations: []*armcosmos.Location{
	// 				{
	// 					ID: to.Ptr("ddb1-eastus"),
	// 					LocationName: to.Ptr("East US"),
	// 					DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 				},
	// 			},
	// 			Locations: []*armcosmos.Location{
	// 				{
	// 					ID: to.Ptr("ddb1-eastus"),
	// 					LocationName: to.Ptr("East US"),
	// 					DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 				},
	// 			},
	// 			FailoverPolicies: []*armcosmos.FailoverPolicy{
	// 				{
	// 					ID: to.Ptr("ddb1-eastus"),
	// 					LocationName: to.Ptr("East US"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 				},
	// 			},
	// 			VirtualNetworkRules: []*armcosmos.VirtualNetworkRule{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 					IgnoreMissingVNetServiceEndpoint: to.Ptr(false),
	// 				},
	// 			},
	// 			PrivateEndpointConnections: []*armcosmos.PrivateEndpointConnection{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.DocumentDB/databaseAccounts/account1/privateEndpointConnections/pe1"),
	// 					Properties: &armcosmos.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armcosmos.PrivateEndpointProperty{
	// 							ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/privateEndpoints/pe1"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armcosmos.PrivateLinkServiceConnectionStateProperty{
	// 							Status: to.Ptr("Approved"),
	// 							ActionsRequired: to.Ptr("None"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Cors: []*armcosmos.CorsPolicy{
	// 			},
	// 			DefaultIdentity: to.Ptr("FirstPartyIdentity"),
	// 			EnableFreeTier: to.Ptr(false),
	// 			APIProperties: &armcosmos.APIProperties{
	// 			},
	// 			EnableAnalyticalStorage: to.Ptr(true),
	// 			EnableBurstCapacity: to.Ptr(true),
	// 			EnablePriorityBasedExecution: to.Ptr(true),
	// 			DefaultPriorityLevel: to.Ptr(armcosmos.DefaultPriorityLevelLow),
	// 			AnalyticalStorageConfiguration: &armcosmos.AnalyticalStorageConfiguration{
	// 				SchemaType: to.Ptr(armcosmos.AnalyticalStorageSchemaTypeWellDefined),
	// 			},
	// 			BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
	// 				Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
	// 				PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
	// 					BackupIntervalInMinutes: to.Ptr[int32](240),
	// 					BackupRetentionIntervalInHours: to.Ptr[int32](720),
	// 					BackupStorageRedundancy: to.Ptr(armcosmos.BackupStorageRedundancyLocal),
	// 				},
	// 			},
	// 			NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassAzureServices),
	// 			NetworkACLBypassResourceIDs: []*string{
	// 				to.Ptr("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName"),
	// 			},
	// 			Capacity: &armcosmos.Capacity{
	// 				TotalThroughputLimit: to.Ptr[int32](2000),
	// 			},
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
	// 			EnablePartitionMerge: to.Ptr(true),
	// 			EnforceHierarchicalPartitionKeyIDLastLevel: to.Ptr(false),
	// 			MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTLS),
	// 		},
	// 	},
	// }
}
