package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ec7ee8842bf615c2f0354bf8b5b8725fdac9454a/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/CosmosDBDatabaseAccountListByResourceGroup.json
func ExampleDatabaseAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabaseAccountsClient().NewListByResourceGroupPager("rg1", nil)
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
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
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
		// 			Kind: to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
		// 			Properties: &armcosmos.DatabaseAccountGetProperties{
		// 				ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
		// 					DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelSession),
		// 					MaxIntervalInSeconds: to.Ptr[int32](5),
		// 					MaxStalenessPrefix: to.Ptr[int64](100),
		// 				},
		// 				Cors: []*armcosmos.CorsPolicy{
		// 				},
		// 				DatabaseAccountOfferType: to.Ptr("Standard"),
		// 				DefaultIdentity: to.Ptr("FirstPartyIdentity"),
		// 				DisableKeyBasedMetadataWriteAccess: to.Ptr(false),
		// 				DocumentEndpoint: to.Ptr("https://ddb1.documents.azure.com:443/"),
		// 				EnableFreeTier: to.Ptr(false),
		// 				FailoverPolicies: []*armcosmos.FailoverPolicy{
		// 					{
		// 						FailoverPriority: to.Ptr[int32](0),
		// 						ID: to.Ptr("ddb1-eastus"),
		// 						LocationName: to.Ptr("East US"),
		// 				}},
		// 				IPRules: []*armcosmos.IPAddressOrRange{
		// 				},
		// 				KeysMetadata: &armcosmos.DatabaseAccountKeysMetadata{
		// 					PrimaryMasterKey: &armcosmos.AccountKeyMetadata{
		// 						GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
		// 					},
		// 					PrimaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
		// 						GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
		// 					},
		// 					SecondaryMasterKey: &armcosmos.AccountKeyMetadata{
		// 						GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
		// 					},
		// 					SecondaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
		// 						GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
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
		// 	}},
		// }
	}
}
