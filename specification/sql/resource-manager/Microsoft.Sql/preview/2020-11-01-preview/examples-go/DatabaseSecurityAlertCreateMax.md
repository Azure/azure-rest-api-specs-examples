Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.3.1/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseSecurityAlertCreateMax.json
func ExampleDatabaseSecurityAlertPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewDatabaseSecurityAlertPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		armsql.SecurityAlertPolicyName("Default"),
		armsql.DatabaseSecurityAlertPolicy{
			Properties: &armsql.SecurityAlertsPolicyProperties{
				DisabledAlerts: []*string{
					to.StringPtr("Sql_Injection"),
					to.StringPtr("Usage_Anomaly")},
				EmailAccountAdmins: to.BoolPtr(true),
				EmailAddresses: []*string{
					to.StringPtr("test@microsoft.com"),
					to.StringPtr("user@microsoft.com")},
				RetentionDays:           to.Int32Ptr(6),
				State:                   armsql.SecurityAlertsPolicyStateEnabled.ToPtr(),
				StorageAccountAccessKey: to.StringPtr("<storage-account-access-key>"),
				StorageEndpoint:         to.StringPtr("<storage-endpoint>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DatabaseSecurityAlertPoliciesClientCreateOrUpdateResult)
}
```
