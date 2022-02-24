Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmariadb%2Farmmariadb%2Fv0.2.1/sdk/resourcemanager/mariadb/armmariadb/README.md) on how to add the SDK to your project and authenticate.

```go
package armmariadb_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerSecurityAlertsCreateMax.json
func ExampleServerSecurityAlertPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmariadb.NewServerSecurityAlertPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armmariadb.SecurityAlertPolicyName("Default"),
		armmariadb.ServerSecurityAlertPolicy{
			Properties: &armmariadb.SecurityAlertPolicyProperties{
				DisabledAlerts: []*string{
					to.StringPtr("Access_Anomaly"),
					to.StringPtr("Usage_Anomaly")},
				EmailAccountAdmins: to.BoolPtr(true),
				EmailAddresses: []*string{
					to.StringPtr("testSecurityAlert@microsoft.com")},
				RetentionDays:           to.Int32Ptr(5),
				State:                   armmariadb.ServerSecurityAlertPolicyStateEnabled.ToPtr(),
				StorageAccountAccessKey: to.StringPtr("<storage-account-access-key>"),
				StorageEndpoint:         to.StringPtr("<storage-endpoint>"),
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
	log.Printf("Response result: %#v\n", res.ServerSecurityAlertPoliciesClientCreateOrUpdateResult)
}
```
