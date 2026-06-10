package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-04-01-preview/MaintenanceEventsListByServerWithFilter.json
func ExampleMaintenanceEventsClient_NewListPager_listMaintenanceEventsFilteredByStatusForAServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMaintenanceEventsClient().NewListPager("exampleresourcegroup", "exampleserver", &armpostgresqlflexibleservers.MaintenanceEventsClientListOptions{
		MaintenanceStatus: to.Ptr(armpostgresqlflexibleservers.MaintenanceEventStatusFilterUpcoming)})
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
		// page = armpostgresqlflexibleservers.MaintenanceEventsClientListResponse{
		// 	MaintenanceEventResourceListResult: armpostgresqlflexibleservers.MaintenanceEventResourceListResult{
		// 		Value: []*armpostgresqlflexibleservers.MaintenanceEventResource{
		// 			{
		// 				Properties: &armpostgresqlflexibleservers.MaintenanceEventResourceProperties{
		// 					MaintenanceEventID: to.Ptr("XXXX-111"),
		// 					MaintenanceType: to.Ptr(armpostgresqlflexibleservers.MaintenanceTypePlannedMaintenance),
		// 					OriginalStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-02T07:23:00Z"); return t}()),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.MaintenanceEventStatusPlanned),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-02T07:23:00Z"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-02T08:23:00Z"); return t}()),
		// 					EstimatedDowntime: to.Ptr("PT3600S"),
		// 					Deferrable: to.Ptr(true),
		// 					DeferralDeadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-16T07:23:00Z"); return t}()),
		// 					LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T07:23:15.9626227Z"); return t}()),
		// 				},
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/maintenanceEvents/XXXX-111"),
		// 				Name: to.Ptr("XXXX-111"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/maintenanceEvents"),
		// 			},
		// 			{
		// 				Properties: &armpostgresqlflexibleservers.MaintenanceEventResourceProperties{
		// 					MaintenanceEventID: to.Ptr("XXXX-222"),
		// 					MaintenanceType: to.Ptr(armpostgresqlflexibleservers.MaintenanceTypePlannedMaintenance),
		// 					OriginalStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-03T07:23:00Z"); return t}()),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.MaintenanceEventStatusPlanned),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-03T08:23:00Z"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-03T09:23:00Z"); return t}()),
		// 					EstimatedDowntime: to.Ptr("PT3540S"),
		// 					Deferrable: to.Ptr(true),
		// 					DeferralDeadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-18T18:31:00Z"); return t}()),
		// 					LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-02T07:23:15.9626227Z"); return t}()),
		// 				},
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/maintenanceEvents/XXXX-222"),
		// 				Name: to.Ptr("XXXX-222"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/maintenanceEvents"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
