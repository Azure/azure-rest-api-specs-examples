Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv0.4.0/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
func ExampleAttachedDatabaseConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armkusto.NewAttachedDatabaseConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<attached-database-configuration-name>",
		armkusto.AttachedDatabaseConfiguration{
			Location: to.Ptr("<location>"),
			Properties: &armkusto.AttachedDatabaseConfigurationProperties{
				ClusterResourceID:                 to.Ptr("<cluster-resource-id>"),
				DatabaseName:                      to.Ptr("<database-name>"),
				DefaultPrincipalsModificationKind: to.Ptr(armkusto.DefaultPrincipalsModificationKindUnion),
				TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
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
		&armkusto.AttachedDatabaseConfigurationsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
