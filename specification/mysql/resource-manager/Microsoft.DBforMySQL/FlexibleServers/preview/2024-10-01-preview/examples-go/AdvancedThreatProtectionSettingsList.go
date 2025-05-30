package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2024-10-01-preview/examples/AdvancedThreatProtectionSettingsList.json
func ExampleAdvancedThreatProtectionSettingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAdvancedThreatProtectionSettingsClient().NewListPager("threatprotection-6852", "threatprotection-2080", nil)
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
		// page.AdvancedThreatProtectionListResult = armmysqlflexibleservers.AdvancedThreatProtectionListResult{
		// 	Value: []*armmysqlflexibleservers.AdvancedThreatProtection{
		// 		{
		// 			Name: to.Ptr("Default"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/advancedThreatProtectionSettings"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/threatprotection-6852/providers/Microsoft.DBforMySQL/flexibleServers/threatprotection-2080/advancedThreatProtectionSettings/Default"),
		// 			SystemData: &armmysqlflexibleservers.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armmysqlflexibleservers.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armmysqlflexibleservers.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmysqlflexibleservers.AdvancedThreatProtectionProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armmysqlflexibleservers.AdvancedThreatProtectionProvisioningStateSucceeded),
		// 				State: to.Ptr(armmysqlflexibleservers.AdvancedThreatProtectionStateEnabled),
		// 			},
		// 	}},
		// }
	}
}
