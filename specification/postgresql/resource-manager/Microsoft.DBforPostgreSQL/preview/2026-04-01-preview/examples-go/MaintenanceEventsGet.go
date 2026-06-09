package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-04-01-preview/MaintenanceEventsGet.json
func ExampleMaintenanceEventsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMaintenanceEventsClient().Get(ctx, "exampleresourcegroup", "exampleserver", "XXXX-111", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpostgresqlflexibleservers.MaintenanceEventsClientGetResponse{
	// 	MaintenanceEventResource: armpostgresqlflexibleservers.MaintenanceEventResource{
	// 		Properties: &armpostgresqlflexibleservers.MaintenanceEventResourceProperties{
	// 			MaintenanceEventID: to.Ptr("XXXX-111"),
	// 			MaintenanceType: to.Ptr(armpostgresqlflexibleservers.MaintenanceTypePlannedMaintenance),
	// 			OriginalStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-02T07:23:00Z"); return t}()),
	// 			Status: to.Ptr(armpostgresqlflexibleservers.MaintenanceEventStatusInProgress),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-02T07:23:00Z"); return t}()),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-02T08:23:00Z"); return t}()),
	// 			EstimatedDowntime: to.Ptr("PT3600S"),
	// 			Deferrable: to.Ptr(true),
	// 			DeferralDeadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-16T07:23:00Z"); return t}()),
	// 			LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T07:23:15.9626227Z"); return t}()),
	// 		},
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/maintenanceEvents/XXXX-111"),
	// 		Name: to.Ptr("XXXX-111"),
	// 		Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/maintenanceEvents"),
	// 	},
	// }
}
