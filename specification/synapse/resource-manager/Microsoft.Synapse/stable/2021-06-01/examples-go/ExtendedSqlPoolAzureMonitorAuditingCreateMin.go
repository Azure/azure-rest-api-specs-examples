package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ExtendedSqlPoolAzureMonitorAuditingCreateMin.json
func ExampleExtendedSQLPoolBlobAuditingPoliciesClient_CreateOrUpdate_createOrUpdateAnExtendedSqlPoolsAzureMonitorAuditingPolicyWithMinimalParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExtendedSQLPoolBlobAuditingPoliciesClient().CreateOrUpdate(ctx, "blobauditingtest-4799", "blobauditingtest-6440", "testdb", armsynapse.ExtendedSQLPoolBlobAuditingPolicy{
		Properties: &armsynapse.ExtendedSQLPoolBlobAuditingPolicyProperties{
			IsAzureMonitorTargetEnabled: to.Ptr(true),
			State:                       to.Ptr(armsynapse.BlobAuditingPolicyStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtendedSQLPoolBlobAuditingPolicy = armsynapse.ExtendedSQLPoolBlobAuditingPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/extendedAuditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/blobauditingtest-4799/providers/Microsoft.Synapse/workspaces/blobauditingtest-6440/sqlPools/testdb"),
	// 	Properties: &armsynapse.ExtendedSQLPoolBlobAuditingPolicyProperties{
	// 		AuditActionsAndGroups: []*string{
	// 			to.Ptr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("BATCH_COMPLETED_GROUP")},
	// 			IsAzureMonitorTargetEnabled: to.Ptr(true),
	// 			IsStorageSecondaryKeyInUse: to.Ptr(false),
	// 			RetentionDays: to.Ptr[int32](0),
	// 			State: to.Ptr(armsynapse.BlobAuditingPolicyStateEnabled),
	// 			StorageAccountSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 	}
}
