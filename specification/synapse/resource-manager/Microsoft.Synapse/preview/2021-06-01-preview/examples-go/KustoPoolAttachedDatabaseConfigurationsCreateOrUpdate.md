Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.2.1/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsCreateOrUpdate.json
func ExampleKustoPoolAttachedDatabaseConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolAttachedDatabaseConfigurationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<attached-database-configuration-name>",
		"<resource-group-name>",
		armsynapse.AttachedDatabaseConfiguration{
			Location: to.StringPtr("<location>"),
			Properties: &armsynapse.AttachedDatabaseConfigurationProperties{
				KustoPoolResourceID:               to.StringPtr("<kusto-pool-resource-id>"),
				DatabaseName:                      to.StringPtr("<database-name>"),
				DefaultPrincipalsModificationKind: armsynapse.DefaultPrincipalsModificationKind("Union").ToPtr(),
				TableLevelSharingProperties: &armsynapse.TableLevelSharingProperties{
					ExternalTablesToExclude: []*string{
						to.StringPtr("ExternalTable2")},
					ExternalTablesToInclude: []*string{
						to.StringPtr("ExternalTable1")},
					MaterializedViewsToExclude: []*string{
						to.StringPtr("MaterializedViewTable2")},
					MaterializedViewsToInclude: []*string{
						to.StringPtr("MaterializedViewTable1")},
					TablesToExclude: []*string{
						to.StringPtr("Table2")},
					TablesToInclude: []*string{
						to.StringPtr("Table1")},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.KustoPoolAttachedDatabaseConfigurationsClientCreateOrUpdateResult)
}
```
