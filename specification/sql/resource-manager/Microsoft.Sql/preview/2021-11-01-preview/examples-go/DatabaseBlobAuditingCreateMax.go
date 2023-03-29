package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/DatabaseBlobAuditingCreateMax.json
func ExampleDatabaseBlobAuditingPoliciesClient_CreateOrUpdate_createOrUpdateADatabasesBlobAuditingPolicyWithAllParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseBlobAuditingPoliciesClient().CreateOrUpdate(ctx, "blobauditingtest-4799", "blobauditingtest-6440", "testdb", armsql.DatabaseBlobAuditingPolicy{
		Properties: &armsql.DatabaseBlobAuditingPolicyProperties{
			AuditActionsAndGroups: []*string{
				to.Ptr("DATABASE_LOGOUT_GROUP"),
				to.Ptr("DATABASE_ROLE_MEMBER_CHANGE_GROUP"),
				to.Ptr("UPDATE on database::TestDatabaseName by public")},
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseBlobAuditingPolicy = armsql.DatabaseBlobAuditingPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/auditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/blobauditingtest-4799/providers/Microsoft.Sql/servers/blobauditingtest-6440/databases/testdb"),
	// 	Kind: to.Ptr("V12"),
	// 	Properties: &armsql.DatabaseBlobAuditingPolicyProperties{
	// 		AuditActionsAndGroups: []*string{
	// 			to.Ptr("DATABASE_LOGOUT_GROUP"),
	// 			to.Ptr("DATABASE_ROLE_MEMBER_CHANGE_GROUP"),
	// 			to.Ptr("UPDATE on database::TestDatabaseName by public")},
	// 			IsAzureMonitorTargetEnabled: to.Ptr(true),
	// 			IsStorageSecondaryKeyInUse: to.Ptr(false),
	// 			QueueDelayMs: to.Ptr[int32](4000),
	// 			RetentionDays: to.Ptr[int32](0),
	// 			State: to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
	// 			StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
	// 			StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 		},
	// 	}
}
