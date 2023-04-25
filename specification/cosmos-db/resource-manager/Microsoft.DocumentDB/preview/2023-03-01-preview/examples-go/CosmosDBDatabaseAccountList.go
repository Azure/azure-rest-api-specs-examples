package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/CosmosDBDatabaseAccountList.json
func ExampleDatabaseAccountsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.DatabaseAccountsListResult = armcosmos.DatabaseAccountsListResult{
		// 	Value: []*armcosmos.DatabaseAccountGetResults{
		// 		{
		// 			Name: to.Ptr("ddb1"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1"),
		// 			Identity: &armcosmos.ManagedServiceIdentity{
		// 				Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
		// 				PrincipalID: to.Ptr("52f4fef3-3c3f-4ff3-b52e-b5c9eeb68656"),
		// 				TenantID: to.Ptr("33e01921-4d64-4f8c-a055-5bdaffd5e33d"),
		// 				UserAssignedIdentities: map[string]*armcosmos.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
		// 					"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armcosmos.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
		// 						ClientID: to.Ptr("fbe75b66-01c5-4f87-a220-233af3270436"),
		// 						PrincipalID: to.Ptr("33e01921-4d64-4f8c-a055-5bdaffd5e33d"),
		// 					},
		// 				},
		// 			},
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
		// 			Properties: &armcosmos.DatabaseAccountGetProperties{
		// 				AnalyticalStorageConfiguration: &armcosmos.AnalyticalStorageConfiguration{
		// 					SchemaType: to.Ptr(armcosmos.AnalyticalStorageSchemaTypeWellDefined),
		// 				},
		// 				APIProperties: &armcosmos.APIProperties{
		// 				},
		// 				BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
		// 					Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
		// 					PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
		// 						BackupIntervalInMinutes: to.Ptr[int32](240),
		// 						BackupRetentionIntervalInHours: to.Ptr[int32](720),
		// 						BackupStorageRedundancy: to.Ptr(armcosmos.BackupStorageRedundancyGeo),
		// 					},
		// 				},
		// 				ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
		// 					DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelSession),
		// 					MaxIntervalInSeconds: to.Ptr[int32](5),
		// 					MaxStalenessPrefix: to.Ptr[int64](100),
		// 				},
		// 				Cors: []*armcosmos.CorsPolicy{
		// 				},
		// 				CreateMode: to.Ptr(armcosmos.CreateModeDefault),
		// 				DatabaseAccountOfferType: to.Ptr("Standard"),
		// 				DefaultIdentity: to.Ptr("FirstPartyIdentity"),
		// 				DisableKeyBasedMetadataWriteAccess: to.Ptr(false),
		// 				DocumentEndpoint: to.Ptr("https://ddb1.documents.azure.com:443/"),
		// 				EnableAnalyticalStorage: to.Ptr(true),
		// 				EnableBurstCapacity: to.Ptr(true),
		// 				EnableFreeTier: to.Ptr(false),
		// 				EnableMaterializedViews: to.Ptr(false),
		// 				EnablePartitionMerge: to.Ptr(true),
		// 				FailoverPolicies: []*armcosmos.FailoverPolicy{
		// 					{
		// 						FailoverPriority: to.Ptr[int32](0),
		// 						ID: to.Ptr("ddb1-eastus"),
		// 						LocationName: to.Ptr("East US"),
		// 				}},
		// 				InstanceID: to.Ptr("d9b26648-2f53-4541-b3d8-3044f4f9810d"),
		// 				IPRules: []*armcosmos.IPAddressOrRange{
		// 				},
		// 				KeysMetadata: &armcosmos.DatabaseAccountKeysMetadata{
		// 					PrimaryMasterKey: &armcosmos.AccountKeyMetadata{
		// 						GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
		// 					},
		// 					PrimaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
		// 						GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
		// 					},
		// 					SecondaryMasterKey: &armcosmos.AccountKeyMetadata{
		// 						GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
		// 					},
		// 					SecondaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
		// 						GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
		// 					},
		// 				},
		// 				Locations: []*armcosmos.Location{
		// 					{
		// 						DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
		// 						FailoverPriority: to.Ptr[int32](0),
		// 						ID: to.Ptr("ddb1-eastus"),
		// 						LocationName: to.Ptr("East US"),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 				}},
		// 				MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTLS),
		// 				NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassNone),
		// 				NetworkACLBypassResourceIDs: []*string{
		// 				},
		// 				PrivateEndpointConnections: []*armcosmos.PrivateEndpointConnection{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.DocumentDB/databaseAccounts/account1/privateEndpointConnections/pe1"),
		// 						Properties: &armcosmos.PrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armcosmos.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/privateEndpoints/pe1"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armcosmos.PrivateLinkServiceConnectionStateProperty{
		// 								ActionsRequired: to.Ptr("None"),
		// 								Status: to.Ptr("Approved"),
		// 							},
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ReadLocations: []*armcosmos.Location{
		// 					{
		// 						DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
		// 						FailoverPriority: to.Ptr[int32](0),
		// 						ID: to.Ptr("ddb1-eastus"),
		// 						LocationName: to.Ptr("East US"),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 				}},
		// 				WriteLocations: []*armcosmos.Location{
		// 					{
		// 						DocumentEndpoint: to.Ptr("https://ddb1-eastus.documents.azure.com:443/"),
		// 						FailoverPriority: to.Ptr[int32](0),
		// 						ID: to.Ptr("ddb1-eastus"),
		// 						LocationName: to.Ptr("East US"),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 				}},
		// 			},
		// 			SystemData: &armcosmos.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
