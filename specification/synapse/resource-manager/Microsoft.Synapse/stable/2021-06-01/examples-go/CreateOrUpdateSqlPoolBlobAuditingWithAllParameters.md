Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.5.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolBlobAuditingWithAllParameters.json
func ExampleSQLPoolBlobAuditingPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolBlobAuditingPoliciesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"blobauditingtest-4799",
		"blobauditingtest-6440",
		"testdb",
		armsynapse.SQLPoolBlobAuditingPolicy{
			Properties: &armsynapse.SQLPoolBlobAuditingPolicyProperties{
				AuditActionsAndGroups: []*string{
					to.Ptr("DATABASE_LOGOUT_GROUP"),
					to.Ptr("DATABASE_ROLE_MEMBER_CHANGE_GROUP"),
					to.Ptr("UPDATE on database::TestDatabaseName by public")},
				IsAzureMonitorTargetEnabled:  to.Ptr(true),
				IsStorageSecondaryKeyInUse:   to.Ptr(false),
				RetentionDays:                to.Ptr[int32](6),
				State:                        to.Ptr(armsynapse.BlobAuditingPolicyStateEnabled),
				StorageAccountAccessKey:      to.Ptr("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
				StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
				StorageEndpoint:              to.Ptr("https://mystorage.blob.core.windows.net"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
