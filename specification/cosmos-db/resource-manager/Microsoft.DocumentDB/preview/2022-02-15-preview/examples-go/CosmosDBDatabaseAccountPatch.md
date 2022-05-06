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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBDatabaseAccountPatch.json
func ExampleDatabaseAccountsClient_BeginUpdate() {
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
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armcosmos.DatabaseAccountUpdateParameters{
			Identity: &armcosmos.ManagedServiceIdentity{
				Type: to.Ptr(armcosmos.ResourceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armcosmos.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
				},
			},
			Location: to.Ptr("<location>"),
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
				DefaultIdentity: to.Ptr("<default-identity>"),
				DiagnosticLogSettings: &armcosmos.DiagnosticLogSettings{
					EnableFullTextQuery: to.Ptr(armcosmos.EnableFullTextQueryTrue),
				},
				EnableAnalyticalStorage: to.Ptr(true),
				EnableFreeTier:          to.Ptr(false),
				IPRules: []*armcosmos.IPAddressOrRange{
					{
						IPAddressOrRange: to.Ptr("<ipaddress-or-range>"),
					},
					{
						IPAddressOrRange: to.Ptr("<ipaddress-or-range>"),
					}},
				IsVirtualNetworkFilterEnabled: to.Ptr(true),
				NetworkACLBypass:              to.Ptr(armcosmos.NetworkACLBypassAzureServices),
				NetworkACLBypassResourceIDs: []*string{
					to.Ptr("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName")},
				VirtualNetworkRules: []*armcosmos.VirtualNetworkRule{
					{
						ID:                               to.Ptr("<id>"),
						IgnoreMissingVNetServiceEndpoint: to.Ptr(false),
					}},
			},
			Tags: map[string]*string{
				"dept": to.Ptr("finance"),
			},
		},
		&armcosmos.DatabaseAccountsClientBeginUpdateOptions{ResumeToken: ""})
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
