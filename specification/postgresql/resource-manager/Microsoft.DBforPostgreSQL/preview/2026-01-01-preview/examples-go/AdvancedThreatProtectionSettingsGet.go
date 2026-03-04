package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-01-01-preview/AdvancedThreatProtectionSettingsGet.json
func ExampleAdvancedThreatProtectionSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdvancedThreatProtectionSettingsClient().Get(ctx, "exampleresourcegroup", "exampleserver", armpostgresqlflexibleservers.ThreatProtectionNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpostgresqlflexibleservers.AdvancedThreatProtectionSettingsClientGetResponse{
	// 	AdvancedThreatProtectionSettingsModel: &armpostgresqlflexibleservers.AdvancedThreatProtectionSettingsModel{
	// 		Name: to.Ptr("Default"),
	// 		Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/advancedThreatProtectionSettings"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/advancedThreatProtectionSettings/Default"),
	// 		Properties: &armpostgresqlflexibleservers.AdvancedThreatProtectionSettingsProperties{
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T18:30:22.123456Z"); return t}()),
	// 			State: to.Ptr(armpostgresqlflexibleservers.ThreatProtectionStateEnabled),
	// 		},
	// 	},
	// }
}
