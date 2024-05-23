package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ServerDevOpsAuditList.json
func ExampleServerDevOpsAuditSettingsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerDevOpsAuditSettingsClient().NewListByServerPager("devAuditTestRG", "devOpsAuditTestSvr", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ServerDevOpsAuditSettingsListResult = armsql.ServerDevOpsAuditSettingsListResult{
		// 	Value: []*armsql.ServerDevOpsAuditingSettings{
		// 		{
		// 			Name: to.Ptr("Default"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/devOpsAuditingSettings"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/devAuditTestRG/providers/Microsoft.Sql/servers/devOpsAuditTestSvr/devOpsAuditingSettings/default"),
		// 			Properties: &armsql.ServerDevOpsAuditSettingsProperties{
		// 				IsAzureMonitorTargetEnabled: to.Ptr(false),
		// 				IsManagedIdentityInUse: to.Ptr(false),
		// 				State: to.Ptr(armsql.BlobAuditingPolicyStateDisabled),
		// 				StorageAccountSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
		// 			},
		// 	}},
		// }
	}
}
