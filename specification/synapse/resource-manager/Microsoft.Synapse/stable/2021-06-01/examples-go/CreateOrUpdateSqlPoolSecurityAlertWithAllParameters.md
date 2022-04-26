Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.4.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolSecurityAlertWithAllParameters.json
func ExampleSQLPoolSecurityAlertPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolSecurityAlertPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		armsynapse.SecurityAlertPolicyNameDefault,
		armsynapse.SQLPoolSecurityAlertPolicy{
			Properties: &armsynapse.SecurityAlertPolicyProperties{
				DisabledAlerts: []*string{
					to.Ptr("Sql_Injection"),
					to.Ptr("Usage_Anomaly")},
				EmailAccountAdmins: to.Ptr(true),
				EmailAddresses: []*string{
					to.Ptr("test@microsoft.com"),
					to.Ptr("user@microsoft.com")},
				RetentionDays:           to.Ptr[int32](6),
				State:                   to.Ptr(armsynapse.SecurityAlertPolicyStateEnabled),
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
