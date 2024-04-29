package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/mysql/resource-manager/Microsoft.DBforMySQL/Maintenance/preview/2023-10-01-preview/examples/MaintenancesListByServer.json
func ExampleMaintenancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMaintenancesClient().NewListPager("TestGroup", "testserver", nil)
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
		// page.MaintenanceListResult = armmysqlflexibleservers.MaintenanceListResult{
		// 	Value: []*armmysqlflexibleservers.Maintenance{
		// 		{
		// 			Name: to.Ptr("YL4X-3CG"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/maintenances"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/maintenances/YL4X-3CG"),
		// 			Properties: &armmysqlflexibleservers.MaintenanceProperties{
		// 				MaintenanceDescription: to.Ptr("Your Azure Database For MySQL routine maintenance has been successfully completed and your instance is running on the upgraded version. If you have any additional questions or concerns, please contact support."),
		// 				MaintenanceEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-10T21:00:00.000Z"); return t}()),
		// 				MaintenanceExecutionEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-10T20:22:38.000Z"); return t}()),
		// 				MaintenanceExecutionStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-10T20:03:21.000Z"); return t}()),
		// 				MaintenanceStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-10T20:00:00.000Z"); return t}()),
		// 				MaintenanceState: to.Ptr(armmysqlflexibleservers.MaintenanceStateCompleted),
		// 				MaintenanceTitle: to.Ptr("Routine Maintenance: Completed on Mon, 11 Dec 2023 04:22:38 GMT"),
		// 				MaintenanceType: to.Ptr(armmysqlflexibleservers.MaintenanceTypeRoutineMaintenance),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("_T9Q-TS8"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/maintenances"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/maintenances/_T9Q-TS8"),
		// 			Properties: &armmysqlflexibleservers.MaintenanceProperties{
		// 				MaintenanceAvailableScheduleMaxTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T16:00:00.000Z"); return t}()),
		// 				MaintenanceAvailableScheduleMinTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-04T17:23:45.000Z"); return t}()),
		// 				MaintenanceDescription: to.Ptr("Your Azure Database For MySQL routine maintenance will be performed on the scheduled maintenance window between Mon, 15 Jan 2024 05:00:00 GMT and Mon, 15 Jan 2024 06:00:00 GMT. During the maintenance window, your instance may experience temporary service interruption. Our team will minimize any impact and ensure a smooth transition"),
		// 				MaintenanceEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-14T22:00:00.000Z"); return t}()),
		// 				MaintenanceStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-14T21:00:00.000Z"); return t}()),
		// 				MaintenanceState: to.Ptr(armmysqlflexibleservers.MaintenanceStateScheduled),
		// 				MaintenanceTitle: to.Ptr("Routine Maintenance: Scheduled on Mon, 15 Jan 2024 05:00:00 GMT - Mon, 15 Jan 2024 06:00:00 GMT"),
		// 				MaintenanceType: to.Ptr(armmysqlflexibleservers.MaintenanceTypeRoutineMaintenance),
		// 			},
		// 	}},
		// }
	}
}
