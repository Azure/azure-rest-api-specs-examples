package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/MaintenanceEvents_List.json
func ExampleMaintenanceEventsClient_NewListPager_listMaintenanceEventsForAHorizonDbPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMaintenanceEventsClient().NewListPager("example-resource-group", "example-cluster", "example-pool-01", nil)
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
		// page = armhorizondb.MaintenanceEventsClientListResponse{
		// 	MaintenanceEventListResult: armhorizondb.MaintenanceEventListResult{
		// 		Value: []*armhorizondb.MaintenanceEvent{
		// 			{
		// 				Properties: &armhorizondb.MaintenanceEventProperties{
		// 					MaintenanceEventID: to.Ptr("ABCD-123"),
		// 					MaintenanceType: to.Ptr(armhorizondb.MaintenanceTypePlannedMaintenance),
		// 					Description: to.Ptr("Scheduled platform maintenance for the HorizonDB pool."),
		// 					OriginalStartTimeUTC: to.Ptr(time.Date(2026, time.August, 12, 7, 0, 0, 0, time.UTC)),
		// 					Status: to.Ptr(armhorizondb.MaintenanceEventStatusPlanned),
		// 					StartTimeUTC: to.Ptr(time.Date(2026, time.August, 12, 7, 0, 0, 0, time.UTC)),
		// 					EndTimeUTC: to.Ptr(time.Date(2026, time.August, 12, 8, 0, 0, 0, time.UTC)),
		// 					EstimatedDowntime: to.Ptr("PT3600S"),
		// 					Deferrable: to.Ptr(true),
		// 					DeferralDeadlineUTC: to.Ptr(time.Date(2026, time.August, 26, 7, 0, 0, 0, time.UTC)),
		// 					LastUpdatedTimeUTC: to.Ptr(time.Date(2026, time.August, 1, 7, 23, 15, 0, time.UTC)),
		// 				},
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/pools/example-pool-01/maintenanceEvents/ABCD-123"),
		// 				Name: to.Ptr("ABCD-123"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/pools/maintenanceEvents"),
		// 			},
		// 			{
		// 				Properties: &armhorizondb.MaintenanceEventProperties{
		// 					MaintenanceEventID: to.Ptr("EFGH-456"),
		// 					MaintenanceType: to.Ptr(armhorizondb.MaintenanceTypePlannedMaintenance),
		// 					Description: to.Ptr("Scheduled platform maintenance for the HorizonDB pool."),
		// 					OriginalStartTimeUTC: to.Ptr(time.Date(2026, time.August, 19, 7, 0, 0, 0, time.UTC)),
		// 					Status: to.Ptr(armhorizondb.MaintenanceEventStatusPlanned),
		// 					StartTimeUTC: to.Ptr(time.Date(2026, time.August, 19, 7, 0, 0, 0, time.UTC)),
		// 					EndTimeUTC: to.Ptr(time.Date(2026, time.August, 19, 8, 0, 0, 0, time.UTC)),
		// 					EstimatedDowntime: to.Ptr("PT3600S"),
		// 					Deferrable: to.Ptr(true),
		// 					DeferralDeadlineUTC: to.Ptr(time.Date(2026, time.September, 2, 7, 0, 0, 0, time.UTC)),
		// 					LastUpdatedTimeUTC: to.Ptr(time.Date(2026, time.August, 2, 7, 23, 15, 0, time.UTC)),
		// 				},
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/pools/example-pool-01/maintenanceEvents/EFGH-456"),
		// 				Name: to.Ptr("EFGH-456"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/pools/maintenanceEvents"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
