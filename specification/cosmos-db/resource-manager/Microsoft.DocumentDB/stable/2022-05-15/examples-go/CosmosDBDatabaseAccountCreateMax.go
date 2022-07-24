package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBDatabaseAccountCreateMax.json
func ExampleDatabaseAccountsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewDatabaseAccountsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"ddb1",
		armcosmos.DatabaseAccountCreateUpdateParameters{
			Location: to.Ptr("westus"),
			Tags:     map[string]*string{},
			Identity: &armcosmos.ManagedServiceIdentity{
				Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armcosmos.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
				},
			},
			Kind: to.Ptr(armcosmos.DatabaseAccountKindMongoDB),
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
				ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
					DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
					MaxIntervalInSeconds:    to.Ptr[int32](10),
					MaxStalenessPrefix:      to.Ptr[int64](200),
				},
				Cors: []*armcosmos.CorsPolicy{
					{
						AllowedOrigins: to.Ptr("https://test"),
					}},
				CreateMode:               to.Ptr(armcosmos.CreateModeDefault),
				DatabaseAccountOfferType: to.Ptr("Standard"),
				DefaultIdentity:          to.Ptr("FirstPartyIdentity"),
				EnableAnalyticalStorage:  to.Ptr(true),
				EnableFreeTier:           to.Ptr(false),
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
				NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassAzureServices),
				NetworkACLBypassResourceIDs: []*string{
					to.Ptr("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName")},
				PublicNetworkAccess: to.Ptr(armcosmos.PublicNetworkAccessEnabled),
				VirtualNetworkRules: []*armcosmos.VirtualNetworkRule{
					{
						ID:                               to.Ptr("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
						IgnoreMissingVNetServiceEndpoint: to.Ptr(false),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
