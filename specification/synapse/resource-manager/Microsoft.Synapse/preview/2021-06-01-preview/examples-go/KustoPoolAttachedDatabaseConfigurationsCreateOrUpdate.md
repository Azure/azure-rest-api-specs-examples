Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.4.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsCreateOrUpdate.json
func ExampleKustoPoolAttachedDatabaseConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolAttachedDatabaseConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<attached-database-configuration-name>",
		"<resource-group-name>",
		armsynapse.AttachedDatabaseConfiguration{
			Location: to.Ptr("<location>"),
			Properties: &armsynapse.AttachedDatabaseConfigurationProperties{
				KustoPoolResourceID:               to.Ptr("<kusto-pool-resource-id>"),
				DatabaseName:                      to.Ptr("<database-name>"),
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
		&armsynapse.KustoPoolAttachedDatabaseConfigurationsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
