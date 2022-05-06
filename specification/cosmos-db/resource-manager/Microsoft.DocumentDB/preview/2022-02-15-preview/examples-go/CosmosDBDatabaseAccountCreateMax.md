Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcosmos%2Farmcosmos%2Fv0.5.0/sdk/resourcemanager/cosmos/armcosmos/README.md) on how to add the SDK to your project and authenticate.

```go
package armcosmos_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBDatabaseAccountCreateMax.json
func ExampleDatabaseAccountsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcosmos.NewDatabaseAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armcosmos.DatabaseAccountCreateUpdateParameters{
			Identity: &armcosmos.ManagedServiceIdentity{
				Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armcosmos.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
				},
			},
			Location: to.Ptr("<location>"),
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
				ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
					DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
					MaxIntervalInSeconds:    to.Ptr[int32](10),
					MaxStalenessPrefix:      to.Ptr[int64](200),
				},
				Cors: []*armcosmos.CorsPolicy{
					{
						AllowedOrigins: to.Ptr("<allowed-origins>"),
					}},
				CreateMode:               to.Ptr(armcosmos.CreateModeDefault),
				DatabaseAccountOfferType: to.Ptr("<database-account-offer-type>"),
				DefaultIdentity:          to.Ptr("<default-identity>"),
				EnableAnalyticalStorage:  to.Ptr(true),
				EnableFreeTier:           to.Ptr(false),
				EnableMaterializedViews:  to.Ptr(false),
				IPRules: []*armcosmos.IPAddressOrRange{
					{
						IPAddressOrRange: to.Ptr("<ipaddress-or-range>"),
					},
					{
						IPAddressOrRange: to.Ptr("<ipaddress-or-range>"),
					}},
				IsVirtualNetworkFilterEnabled: to.Ptr(true),
				KeyVaultKeyURI:                to.Ptr("<key-vault-key-uri>"),
				Locations: []*armcosmos.Location{
					{
						FailoverPriority: to.Ptr[int32](0),
						IsZoneRedundant:  to.Ptr(false),
						LocationName:     to.Ptr("<location-name>"),
					},
					{
						FailoverPriority: to.Ptr[int32](1),
						IsZoneRedundant:  to.Ptr(false),
						LocationName:     to.Ptr("<location-name>"),
					}},
				NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassAzureServices),
				NetworkACLBypassResourceIDs: []*string{
					to.Ptr("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName")},
				PublicNetworkAccess: to.Ptr(armcosmos.PublicNetworkAccessEnabled),
				VirtualNetworkRules: []*armcosmos.VirtualNetworkRule{
					{
						ID:                               to.Ptr("<id>"),
						IgnoreMissingVNetServiceEndpoint: to.Ptr(false),
					}},
			},
		},
		&armcosmos.DatabaseAccountsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
