Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv0.2.1/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDatabasePrincipalAssignmentsCreateOrUpdate.json
func ExampleDatabasePrincipalAssignmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewDatabasePrincipalAssignmentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<database-name>",
		"<principal-assignment-name>",
		armkusto.DatabasePrincipalAssignment{
			Properties: &armkusto.DatabasePrincipalProperties{
				PrincipalID:   to.StringPtr("<principal-id>"),
				PrincipalType: armkusto.PrincipalType("App").ToPtr(),
				Role:          armkusto.DatabasePrincipalRole("Admin").ToPtr(),
				TenantID:      to.StringPtr("<tenant-id>"),
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
	log.Printf("Response result: %#v\n", res.DatabasePrincipalAssignmentsClientCreateOrUpdateResult)
}
```
