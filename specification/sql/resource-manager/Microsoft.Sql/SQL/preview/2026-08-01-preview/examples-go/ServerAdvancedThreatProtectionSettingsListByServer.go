package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/ServerAdvancedThreatProtectionSettingsListByServer.json
func ExampleServerAdvancedThreatProtectionSettingsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerAdvancedThreatProtectionSettingsClient().NewListByServerPager("threatprotection-4799", "threatprotection-6440", nil)
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
		// page = armsql.ServerAdvancedThreatProtectionSettingsClientListByServerResponse{
		// 	LogicalServerAdvancedThreatProtectionListResult: armsql.LogicalServerAdvancedThreatProtectionListResult{
		// 		Value: []*armsql.ServerAdvancedThreatProtection{
		// 			{
		// 				Name: to.Ptr("Default"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/advancedThreatProtectionSettings"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/threatprotection-4799/providers/Microsoft.Sql/servers/threatprotection-6440/advancedThreatProtectionSettings/Default"),
		// 				Properties: &armsql.AdvancedThreatProtectionProperties{
		// 					CreationTime: to.Ptr(time.Date(2022, time.April, 3, 4, 41, 33, 937000000, time.UTC)),
		// 					State: to.Ptr(armsql.AdvancedThreatProtectionStateDisabled),
		// 				},
		// 				SystemData: &armsql.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2022, time.April, 3, 4, 41, 33, 937000000, time.UTC)),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2022, time.April, 3, 4, 41, 33, 937000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
