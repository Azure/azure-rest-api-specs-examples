package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/ServerThreatProtectionSettingsGet.json
func ExampleServerThreatProtectionSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerThreatProtectionSettingsClient().Get(ctx, "threatprotection-6852", "threatprotection-2080", armpostgresqlflexibleservers.ThreatProtectionNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerThreatProtectionSettingsModel = armpostgresqlflexibleservers.ServerThreatProtectionSettingsModel{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/advancedThreatProtectionSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/threatprotection-4799/providers/Microsoft.DBforPostgreSQL/flexibleServers/threatprotection-6440/advancedThreatProtectionSettings/Default"),
	// 	Properties: &armpostgresqlflexibleservers.ServerThreatProtectionProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-03T04:41:33.937Z"); return t}()),
	// 		State: to.Ptr(armpostgresqlflexibleservers.ThreatProtectionStateEnabled),
	// 	},
	// }
}
