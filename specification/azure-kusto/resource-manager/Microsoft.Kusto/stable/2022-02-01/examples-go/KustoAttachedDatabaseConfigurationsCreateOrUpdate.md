Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv1.0.0/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.

```go
package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
func ExampleAttachedDatabaseConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkusto.NewAttachedDatabaseConfigurationsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"kustorptest",
		"kustoCluster2",
		"attachedDatabaseConfigurationsTest",
		armkusto.AttachedDatabaseConfiguration{
			Location: to.Ptr("westus"),
			Properties: &armkusto.AttachedDatabaseConfigurationProperties{
				ClusterResourceID:                 to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
				DatabaseName:                      to.Ptr("kustodatabase"),
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
```
