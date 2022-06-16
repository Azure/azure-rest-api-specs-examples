package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsCreateOrUpdate.json
func ExampleKustoPoolAttachedDatabaseConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolAttachedDatabaseConfigurationsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"kustorptest",
		"kustoclusterrptest4",
		"attachedDatabaseConfigurations1",
		"kustorptest",
		armsynapse.AttachedDatabaseConfiguration{
			Location: to.Ptr("westus"),
			Properties: &armsynapse.AttachedDatabaseConfigurationProperties{
				KustoPoolResourceID:               to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/Workspaces/kustorptest/KustoPools/kustoclusterrptest4"),
				DatabaseName:                      to.Ptr("kustodatabase"),
				DefaultPrincipalsModificationKind: to.Ptr(armsynapse.DefaultPrincipalsModificationKindUnion),
				TableLevelSharingProperties: &armsynapse.TableLevelSharingProperties{
					ExternalTablesToExclude: []*string{
						to.Ptr("ExternalTable2")},
					ExternalTablesToInclude: []*string{
						to.Ptr("ExternalTable1")},
					MaterializedViewsToExclude: []*string{
						to.Ptr("MaterializedViewTable2")},
					MaterializedViewsToInclude: []*string{
						to.Ptr("MaterializedViewTable1")},
					TablesToExclude: []*string{
						to.Ptr("Table2")},
					TablesToInclude: []*string{
						to.Ptr("Table1")},
				},
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
