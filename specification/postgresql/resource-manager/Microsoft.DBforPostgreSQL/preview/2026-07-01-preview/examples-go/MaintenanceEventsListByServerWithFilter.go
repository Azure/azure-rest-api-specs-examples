package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-07-01-preview/MaintenanceEventsListByServerWithFilter.json
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
		// 					OriginalStartTime: to.Ptr(time.Date(2026, time.April, 2, 7, 23, 0, 0, time.UTC)),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.MaintenanceEventStatusPlanned),
		// 					StartTime: to.Ptr(time.Date(2026, time.April, 2, 7, 23, 0, 0, time.UTC)),
		// 					EndTime: to.Ptr(time.Date(2026, time.April, 2, 8, 23, 0, 0, time.UTC)),
		// 					EstimatedDowntime: to.Ptr("PT3600S"),
		// 					Deferrable: to.Ptr(true),
		// 					DeferralDeadline: to.Ptr(time.Date(2026, time.April, 16, 7, 23, 0, 0, time.UTC)),
		// 					LastUpdatedTime: to.Ptr(time.Date(2026, time.April, 1, 7, 23, 15, 962622700, time.UTC)),
		// 				},
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/maintenanceEvents/XXXX-111"),
		// 				Name: to.Ptr("XXXX-111"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/maintenanceEvents"),
		// 			},
		// 			{
		// 				Properties: &armpostgresqlflexibleservers.MaintenanceEventResourceProperties{
		// 					MaintenanceEventID: to.Ptr("XXXX-222"),
		// 					MaintenanceType: to.Ptr(armpostgresqlflexibleservers.MaintenanceTypePlannedMaintenance),
		// 					OriginalStartTime: to.Ptr(time.Date(2026, time.April, 3, 7, 23, 0, 0, time.UTC)),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.MaintenanceEventStatusPlanned),
		// 					StartTime: to.Ptr(time.Date(2026, time.April, 3, 8, 23, 0, 0, time.UTC)),
		// 					EndTime: to.Ptr(time.Date(2026, time.April, 3, 9, 23, 0, 0, time.UTC)),
		// 					EstimatedDowntime: to.Ptr("PT3540S"),
		// 					Deferrable: to.Ptr(true),
		// 					DeferralDeadline: to.Ptr(time.Date(2026, time.April, 18, 18, 31, 0, 0, time.UTC)),
		// 					LastUpdatedTime: to.Ptr(time.Date(2026, time.April, 2, 7, 23, 15, 962622700, time.UTC)),
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
