Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.4.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateWorkspaceManagedSqlServerBlobAuditingSettingsWithAllParameters.json
func ExampleWorkspaceManagedSQLServerBlobAuditingPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armsynapse.BlobAuditingPolicyNameDefault,
		armsynapse.ServerBlobAuditingPolicy{
			Properties: &armsynapse.ServerBlobAuditingPolicyProperties{
				AuditActionsAndGroups: []*string{
					to.Ptr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
					to.Ptr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
					to.Ptr("BATCH_COMPLETED_GROUP")},
				IsAzureMonitorTargetEnabled:  to.Ptr(true),
				IsStorageSecondaryKeyInUse:   to.Ptr(false),
				QueueDelayMs:                 to.Ptr[int32](4000),
				RetentionDays:                to.Ptr[int32](6),
				State:                        to.Ptr(armsynapse.BlobAuditingPolicyStateEnabled),
				StorageAccountAccessKey:      to.Ptr("<storage-account-access-key>"),
				StorageAccountSubscriptionID: to.Ptr("<storage-account-subscription-id>"),
				StorageEndpoint:              to.Ptr("<storage-endpoint>"),
			},
		},
		&armsynapse.WorkspaceManagedSQLServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
