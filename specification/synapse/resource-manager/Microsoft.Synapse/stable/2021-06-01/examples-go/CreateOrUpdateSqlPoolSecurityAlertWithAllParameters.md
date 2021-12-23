Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.1.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolSecurityAlertWithAllParameters.json
func ExampleSQLPoolSecurityAlertPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewSQLPoolSecurityAlertPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		armsynapse.SecurityAlertPolicyNameDefault,
		armsynapse.SQLPoolSecurityAlertPolicy{
			Properties: &armsynapse.SecurityAlertPolicyProperties{
				DisabledAlerts: []*string{
					to.StringPtr("Sql_Injection"),
					to.StringPtr("Usage_Anomaly")},
				EmailAccountAdmins: to.BoolPtr(true),
				EmailAddresses: []*string{
					to.StringPtr("test@microsoft.com"),
					to.StringPtr("user@microsoft.com")},
				RetentionDays:           to.Int32Ptr(6),
				State:                   armsynapse.SecurityAlertPolicyStateEnabled.ToPtr(),
				StorageAccountAccessKey: to.StringPtr("<storage-account-access-key>"),
				StorageEndpoint:         to.StringPtr("<storage-endpoint>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SQLPoolSecurityAlertPolicy.ID: %s\n", *res.ID)
}
```
