package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBDatabaseAccountList.json
func ExampleDatabaseAccountsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabaseAccountsClient().NewListPager(nil)
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
		// page = armcosmos.DatabaseAccountsClientListResponse{
		// 	DatabaseAccountsListResult: armcosmos.DatabaseAccountsListResult{
		// 		Value: []*armcosmos.DatabaseAccountGetResults{
		// 			{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1"),
		// 				Name: to.Ptr("ddb1"),
		// 				Location: to.Ptr("West US"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts"),
		// 				Kind: to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armcosmos.DatabaseAccountGetProperties{
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					DocumentEndpoint: to.Ptr("https://ddb1.documents.azure.com:443/"),
		// 					IPRules: []*armcosmos.IPAddressOrRange{
		// 					},
		// 					DatabaseAccountOfferType: to.Ptr("Standard"),
		// 					DisableKeyBasedMetadataWriteAccess: to.Ptr(false),
		// 					ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
		// 						DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelSession),
		// 						MaxIntervalInSeconds: to.Ptr[int32](5),
		// 						MaxStalenessPrefix: to.Ptr[int64](100),
		// 					},
		// 					WriteLocations: []*armcosmos.Location{
		// 						{
		// 							ID: to.Ptr("ddb1-eastus"),
		// 							LocationName: to.Ptr("East US"),
		// 							DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
		// 							ProvisioningState: to.Ptr("Succeeded"),
		// 							FailoverPriority: to.Ptr[int32](0),
		// 						},
		// 					},
		// 					ReadLocations: []*armcosmos.Location{
		// 						{
		// 							ID: to.Ptr("ddb1-eastus"),
		// 							LocationName: to.Ptr("East US"),
		// 							DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
		// 							ProvisioningState: to.Ptr("Succeeded"),
		// 							FailoverPriority: to.Ptr[int32](0),
		// 						},
		// 					},
		// 					Locations: []*armcosmos.Location{
		// 						{
		// 							ID: to.Ptr("ddb1-eastus"),
		// 							LocationName: to.Ptr("East US"),
		// 							DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
		// 							ProvisioningState: to.Ptr("Succeeded"),
		// 							FailoverPriority: to.Ptr[int32](0),
		// 						},
		// 					},
		// 					FailoverPolicies: []*armcosmos.FailoverPolicy{
		// 						{
		// 							ID: to.Ptr("ddb1-eastus"),
		// 							LocationName: to.Ptr("East US"),
		// 							FailoverPriority: to.Ptr[int32](0),
		// 						},
		// 					},
		// 					PrivateEndpointConnections: []*armcosmos.PrivateEndpointConnection{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.DocumentDB/databaseAccounts/account1/privateEndpointConnections/pe1"),
		// 							Properties: &armcosmos.PrivateEndpointConnectionProperties{
		// 								PrivateEndpoint: &armcosmos.PrivateEndpointProperty{
		// 									ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/privateEndpoints/pe1"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armcosmos.PrivateLinkServiceConnectionStateProperty{
		// 									Status: to.Ptr("Approved"),
		// 									ActionsRequired: to.Ptr("None"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Cors: []*armcosmos.CorsPolicy{
		// 					},
		// 					DefaultIdentity: to.Ptr("FirstPartyIdentity"),
		// 					EnableFreeTier: to.Ptr(false),
		// 					APIProperties: &armcosmos.APIProperties{
		// 					},
		// 					EnableAnalyticalStorage: to.Ptr(true),
		// 					AnalyticalStorageConfiguration: &armcosmos.AnalyticalStorageConfiguration{
		// 						SchemaType: to.Ptr(armcosmos.AnalyticalStorageSchemaTypeWellDefined),
		// 					},
		// 					InstanceID: to.Ptr("d9b26648-2f53-4541-b3d8-3044f4f9810d"),
		// 					CreateMode: to.Ptr(armcosmos.CreateModeDefault),
		// 					BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
		// 						Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
		// 						PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
		// 							BackupIntervalInMinutes: to.Ptr[int32](240),
		// 							BackupRetentionIntervalInHours: to.Ptr[int32](720),
		// 							BackupStorageRedundancy: to.Ptr(armcosmos.BackupStorageRedundancyGeo),
		// 						},
		// 					},
		// 					NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassNone),
		// 					NetworkACLBypassResourceIDs: []*string{
		// 					},
		// 					EnablePartitionMerge: to.Ptr(true),
		// 					EnableBurstCapacity: to.Ptr(true),
		// 					MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTLS),
		// 					CapacityMode: to.Ptr(armcosmos.CapacityModeProvisioned),
		// 					CapacityModeChangeTransitionState: &armcosmos.CapacityModeChangeTransitionState{
		// 						CapacityModeTransitionBeginTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-24T03:02:16.2747253Z"); return t}()),
		// 						CapacityModeTransitionStatus: to.Ptr(armcosmos.CapacityModeTransitionStatusCompleted),
		// 						CapacityModeLastSuccessfulTransitionEndTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-24T03:02:18.8758329Z"); return t}()),
		// 						CapacityModeTransitionEndTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-24T03:02:18.8758329Z"); return t}()),
		// 						CurrentCapacityMode: to.Ptr(armcosmos.CapacityModeProvisioned),
		// 						PreviousCapacityMode: to.Ptr(armcosmos.CapacityModeServerless),
		// 					},
		// 					EnableMaterializedViews: to.Ptr(false),
		// 					KeysMetadata: &armcosmos.DatabaseAccountKeysMetadata{
		// 						PrimaryMasterKey: &armcosmos.AccountKeyMetadata{
		// 							GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
		// 						},
		// 						SecondaryMasterKey: &armcosmos.AccountKeyMetadata{
		// 							GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
		// 						},
		// 						PrimaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
		// 							GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
		// 						},
		// 						SecondaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
		// 							GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
		// 						},
		// 					},
		// 					EnablePriorityBasedExecution: to.Ptr(true),
		// 					DefaultPriorityLevel: to.Ptr(armcosmos.DefaultPriorityLevelLow),
		// 					EnablePerRegionPerPartitionAutoscale: to.Ptr(true),
		// 					ThroughputPoolDedicatedRUs: to.Ptr[int64](300000),
		// 					ThroughputPoolMaxConsumableRUs: to.Ptr[int64](700000),
		// 				},
		// 				SystemData: &armcosmos.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09Z"); return t}()),
		// 				},
		// 				Identity: &armcosmos.ManagedServiceIdentity{
		// 					Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
		// 					PrincipalID: to.Ptr("52f4fef3-3c3f-4ff3-b52e-b5c9eeb68656"),
		// 					TenantID: to.Ptr("33e01921-4d64-4f8c-a055-5bdaffd5e33d"),
		// 					UserAssignedIdentities: map[string]*armcosmos.ManagedServiceIdentityUserAssignedIdentities{
		// 						"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armcosmos.ManagedServiceIdentityUserAssignedIdentities{
		// 							ClientID: to.Ptr("fbe75b66-01c5-4f87-a220-233af3270436"),
		// 							PrincipalID: to.Ptr("33e01921-4d64-4f8c-a055-5bdaffd5e33d"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
