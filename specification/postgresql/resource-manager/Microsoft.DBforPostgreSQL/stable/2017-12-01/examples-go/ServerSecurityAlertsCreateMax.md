Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresql%2Farmpostgresql%2Fv0.5.0/sdk/resourcemanager/postgresql/armpostgresql/README.md) on how to add the SDK to your project and authenticate.

```go
package armpostgresql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerSecurityAlertsCreateMax.json
func ExampleServerSecurityAlertPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpostgresql.NewServerSecurityAlertPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armpostgresql.SecurityAlertPolicyNameDefault,
		armpostgresql.ServerSecurityAlertPolicy{
			Properties: &armpostgresql.SecurityAlertPolicyProperties{
				DisabledAlerts: []*string{
					to.Ptr("Access_Anomaly"),
					to.Ptr("Usage_Anomaly")},
				EmailAccountAdmins: to.Ptr(true),
				EmailAddresses: []*string{
					to.Ptr("testSecurityAlert@microsoft.com")},
				RetentionDays:           to.Ptr[int32](5),
				State:                   to.Ptr(armpostgresql.ServerSecurityAlertPolicyStateEnabled),
				StorageAccountAccessKey: to.Ptr("<storage-account-access-key>"),
				StorageEndpoint:         to.Ptr("<storage-endpoint>"),
			},
		},
		&armpostgresql.ServerSecurityAlertPoliciesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
