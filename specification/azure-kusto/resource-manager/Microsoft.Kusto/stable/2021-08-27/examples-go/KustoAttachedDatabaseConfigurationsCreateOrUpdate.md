```go
package armkusto_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
func ExampleAttachedDatabaseConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewAttachedDatabaseConfigurationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<attached-database-configuration-name>",
		armkusto.AttachedDatabaseConfiguration{
			Location: to.StringPtr("<location>"),
			Properties: &armkusto.AttachedDatabaseConfigurationProperties{
				ClusterResourceID:                 to.StringPtr("<cluster-resource-id>"),
				DatabaseName:                      to.StringPtr("<database-name>"),
				DefaultPrincipalsModificationKind: armkusto.DefaultPrincipalsModificationKind("Union").ToPtr(),
				TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
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
	log.Printf("Response result: %#v\n", res.AttachedDatabaseConfigurationsClientCreateOrUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv0.2.1/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.
