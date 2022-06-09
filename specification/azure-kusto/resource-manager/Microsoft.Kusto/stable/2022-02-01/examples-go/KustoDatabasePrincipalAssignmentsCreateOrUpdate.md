```go
package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDatabasePrincipalAssignmentsCreateOrUpdate.json
func ExampleDatabasePrincipalAssignmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkusto.NewDatabasePrincipalAssignmentsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"kustorptest",
		"kustoCluster",
		"Kustodatabase8",
		"kustoprincipal1",
		armkusto.DatabasePrincipalAssignment{
			Properties: &armkusto.DatabasePrincipalProperties{
				PrincipalID:   to.Ptr("87654321-1234-1234-1234-123456789123"),
				PrincipalType: to.Ptr(armkusto.PrincipalTypeApp),
				Role:          to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
				TenantID:      to.Ptr("12345678-1234-1234-1234-123456789123"),
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv1.0.0/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.
