package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ServerDevOpsAuditCreateMin.json
func ExampleServerDevOpsAuditSettingsClient_BeginCreateOrUpdate_updateAServersDevOpsAuditSettingsWithMinimalInput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerDevOpsAuditSettingsClient().BeginCreateOrUpdate(ctx, "devAuditTestRG", "devOpsAuditTestSvr", armsql.DevOpsAuditingSettingsNameDefault, armsql.ServerDevOpsAuditingSettings{
		Properties: &armsql.ServerDevOpsAuditSettingsProperties{
			State:                   to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
			StorageAccountAccessKey: to.Ptr("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			StorageEndpoint:         to.Ptr("https://mystorage.blob.core.windows.net"),
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
	// res.ServerDevOpsAuditingSettings = armsql.ServerDevOpsAuditingSettings{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/devOpsAuditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/devAuditTestRG/providers/Microsoft.Sql/servers/devOpsAuditTestSvr/devOpsAuditingSettings/default"),
	// 	Properties: &armsql.ServerDevOpsAuditSettingsProperties{
	// 		State: to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
	// 		StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
	// 		StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 	},
	// }
}
