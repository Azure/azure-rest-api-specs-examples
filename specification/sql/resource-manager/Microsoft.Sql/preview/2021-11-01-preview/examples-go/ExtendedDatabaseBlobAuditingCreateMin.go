package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8358c7473dfe057d84a6b6a921225063c040b31a/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ExtendedDatabaseBlobAuditingCreateMin.json
func ExampleExtendedDatabaseBlobAuditingPoliciesClient_CreateOrUpdate_createOrUpdateAnExtendedDatabasesBlobAuditingPolicyWithMinimalParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExtendedDatabaseBlobAuditingPoliciesClient().CreateOrUpdate(ctx, "blobauditingtest-4799", "blobauditingtest-6440", "testdb", armsql.ExtendedDatabaseBlobAuditingPolicy{
		Properties: &armsql.ExtendedDatabaseBlobAuditingPolicyProperties{
			State:                   to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
			StorageAccountAccessKey: to.Ptr("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
			StorageEndpoint:         to.Ptr("https://mystorage.blob.core.windows.net"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtendedDatabaseBlobAuditingPolicy = armsql.ExtendedDatabaseBlobAuditingPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/extendedAuditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/blobauditingtest-4799/providers/Microsoft.Sql/servers/blobauditingtest-6440/databases/testdb"),
	// 	Properties: &armsql.ExtendedDatabaseBlobAuditingPolicyProperties{
	// 		AuditActionsAndGroups: []*string{
	// 			to.Ptr("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("FAILED_DATABASE_AUTHENTICATION_GROUP"),
	// 			to.Ptr("BATCH_COMPLETED_GROUP")},
	// 			IsAzureMonitorTargetEnabled: to.Ptr(false),
	// 			IsStorageSecondaryKeyInUse: to.Ptr(false),
	// 			RetentionDays: to.Ptr[int32](0),
	// 			State: to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
	// 			StorageAccountSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 		},
	// 	}
}
