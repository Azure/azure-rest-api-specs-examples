package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ServerBlobAuditingCreateMax.json
func ExampleServerBlobAuditingPoliciesClient_BeginCreateOrUpdate_updateAServersBlobAuditingPolicyWithAllParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerBlobAuditingPoliciesClient().BeginCreateOrUpdate(ctx, "blobauditingtest-4799", "blobauditingtest-6440", armsql.ServerBlobAuditingPolicy{
		Properties: &armsql.ServerBlobAuditingPolicyProperties{
			AuditActionsAndGroups: []*string{
				to.Ptr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
				to.Ptr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
				to.Ptr("BATCH_COMPLETED_GROUP")},
			IsAzureMonitorTargetEnabled:  to.Ptr(true),
			IsStorageSecondaryKeyInUse:   to.Ptr(false),
			QueueDelayMs:                 to.Ptr[int32](4000),
			RetentionDays:                to.Ptr[int32](6),
			State:                        to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
			StorageAccountAccessKey:      to.Ptr("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
			StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
			StorageEndpoint:              to.Ptr("https://mystorage.blob.core.windows.net"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerBlobAuditingPolicy = armsql.ServerBlobAuditingPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/auditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/blobauditingtest-4799/providers/Microsoft.Sql/servers/blobauditingtest-6440/auditingSettings/default"),
	// 	Properties: &armsql.ServerBlobAuditingPolicyProperties{
	// 		AuditActionsAndGroups: []*string{
	// 			to.Ptr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("BATCH_COMPLETED_GROUP")},
	// 			IsAzureMonitorTargetEnabled: to.Ptr(true),
	// 			IsStorageSecondaryKeyInUse: to.Ptr(false),
	// 			QueueDelayMs: to.Ptr[int32](4000),
	// 			RetentionDays: to.Ptr[int32](6),
	// 			State: to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
	// 			StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
	// 			StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 		},
	// 	}
}
