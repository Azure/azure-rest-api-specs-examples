Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.5.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseSecurityAlertCreateMax.json
func ExampleDatabaseSecurityAlertPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsql.NewDatabaseSecurityAlertPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		armsql.SecurityAlertPolicyNameDefault,
		armsql.DatabaseSecurityAlertPolicy{
			Properties: &armsql.SecurityAlertsPolicyProperties{
				DisabledAlerts: []*string{
					to.Ptr("Sql_Injection"),
					to.Ptr("Usage_Anomaly")},
				EmailAccountAdmins: to.Ptr(true),
				EmailAddresses: []*string{
					to.Ptr("test@microsoft.com"),
					to.Ptr("user@microsoft.com")},
				RetentionDays:           to.Ptr[int32](6),
				State:                   to.Ptr(armsql.SecurityAlertsPolicyStateEnabled),
				StorageAccountAccessKey: to.Ptr("<storage-account-access-key>"),
				StorageEndpoint:         to.Ptr("<storage-endpoint>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
