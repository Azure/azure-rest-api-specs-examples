package armcosmos_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBRestoreDatabaseAccountCreateUpdate.json
func ExampleDatabaseAccountsClient_BeginCreateOrUpdate_cosmosDbRestoreDatabaseAccountCreateUpdateJson() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewDatabaseAccountsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "ddb1", armcosmos.DatabaseAccountCreateUpdateParameters{
		Location: to.Ptr("westus"),
		Tags:     map[string]*string{},
		Kind:     to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			APIProperties: &armcosmos.APIProperties{
				ServerVersion: to.Ptr(armcosmos.ServerVersionThree2),
			},
			BackupPolicy: &armcosmos.ContinuousModeBackupPolicy{
				Type: to.Ptr(armcosmos.BackupPolicyTypeContinuous),
			},
			ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
				DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
				MaxIntervalInSeconds:    to.Ptr[int32](10),
				MaxStalenessPrefix:      to.Ptr[int64](200),
			},
			CreateMode:               to.Ptr(armcosmos.CreateModeRestore),
			DatabaseAccountOfferType: to.Ptr("Standard"),
			EnableAnalyticalStorage:  to.Ptr(true),
			EnableFreeTier:           to.Ptr(false),
			KeyVaultKeyURI:           to.Ptr("https://myKeyVault.vault.azure.net"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					IsZoneRedundant:  to.Ptr(false),
					LocationName:     to.Ptr("southcentralus"),
				}},
			RestoreParameters: &armcosmos.RestoreParameters{
				DatabasesToRestore: []*armcosmos.DatabaseRestoreResource{
					{
						CollectionNames: []*string{
							to.Ptr("collection1"),
							to.Ptr("collection2")},
						DatabaseName: to.Ptr("db1"),
					},
					{
						CollectionNames: []*string{
							to.Ptr("collection3"),
							to.Ptr("collection4")},
						DatabaseName: to.Ptr("db2"),
					}},
				RestoreMode:           to.Ptr(armcosmos.RestoreModePointInTime),
				RestoreSource:         to.Ptr("/subscriptions/subid/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/1a97b4bb-f6a0-430e-ade1-638d781830cc"),
				RestoreTimestampInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-11T22:05:09Z"); return t }()),
			},
		},
	}, nil)
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
