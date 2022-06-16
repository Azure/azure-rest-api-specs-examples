package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ExtendedServerBlobAuditingCreateMax.json
func ExampleExtendedServerBlobAuditingPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsql.NewExtendedServerBlobAuditingPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armsql.ExtendedServerBlobAuditingPolicy{
			Properties: &armsql.ExtendedServerBlobAuditingPolicyProperties{
				AuditActionsAndGroups: []*string{
					to.Ptr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
					to.Ptr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
					to.Ptr("BATCH_COMPLETED_GROUP")},
				IsAzureMonitorTargetEnabled:  to.Ptr(true),
				IsStorageSecondaryKeyInUse:   to.Ptr(false),
				PredicateExpression:          to.Ptr("<predicate-expression>"),
				QueueDelayMs:                 to.Ptr[int32](4000),
				RetentionDays:                to.Ptr[int32](6),
				State:                        to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
				StorageAccountAccessKey:      to.Ptr("<storage-account-access-key>"),
				StorageAccountSubscriptionID: to.Ptr("<storage-account-subscription-id>"),
				StorageEndpoint:              to.Ptr("<storage-endpoint>"),
			},
		},
		&armsql.ExtendedServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
