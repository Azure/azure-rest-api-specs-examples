package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ecee919199a39cc0d864410f540aa105bf7cdb64/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/ServerThreatProtectionSettingsListByServer.json
func ExampleServerThreatProtectionSettingsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerThreatProtectionSettingsClient().NewListByServerPager("threatprotection-6852", "threatprotection-2080", nil)
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
		// page.ServerThreatProtectionListResult = armpostgresqlflexibleservers.ServerThreatProtectionListResult{
		// 	Value: []*armpostgresqlflexibleservers.ServerThreatProtectionSettingsModel{
		// 		{
		// 			Name: to.Ptr("Default"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/advancedThreatProtectionSettings"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/threatprotection-6852/providers/Microsoft.DBforPostgreSQL/flexibleServers/threatprotection-2080/advancedThreatProtectionSettings/Default"),
		// 			Properties: &armpostgresqlflexibleservers.ServerThreatProtectionProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-03T04:41:33.937Z"); return t}()),
		// 				State: to.Ptr(armpostgresqlflexibleservers.ThreatProtectionStateEnabled),
		// 			},
		// 	}},
		// }
	}
}
