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

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolBlobAuditingWithAllParameters.json
func ExampleSQLPoolBlobAuditingPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewSQLPoolBlobAuditingPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		armsynapse.Enum11Default,
		armsynapse.SQLPoolBlobAuditingPolicy{
			Properties: &armsynapse.SQLPoolBlobAuditingPolicyProperties{
				AuditActionsAndGroups: []*string{
					to.StringPtr("DATABASE_LOGOUT_GROUP"),
					to.StringPtr("DATABASE_ROLE_MEMBER_CHANGE_GROUP"),
					to.StringPtr("UPDATE on database::TestDatabaseName by public")},
				IsAzureMonitorTargetEnabled:  to.BoolPtr(true),
				IsStorageSecondaryKeyInUse:   to.BoolPtr(false),
				RetentionDays:                to.Int32Ptr(6),
				State:                        armsynapse.BlobAuditingPolicyStateEnabled.ToPtr(),
				StorageAccountAccessKey:      to.StringPtr("<storage-account-access-key>"),
				StorageAccountSubscriptionID: to.StringPtr("<storage-account-subscription-id>"),
				StorageEndpoint:              to.StringPtr("<storage-endpoint>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SQLPoolBlobAuditingPolicy.ID: %s\n", *res.ID)
}
```
