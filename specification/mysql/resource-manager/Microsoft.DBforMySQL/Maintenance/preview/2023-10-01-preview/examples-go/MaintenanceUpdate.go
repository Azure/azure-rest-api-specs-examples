package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/mysql/resource-manager/Microsoft.DBforMySQL/Maintenance/preview/2023-10-01-preview/examples/MaintenanceUpdate.json
func ExampleMaintenancesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMaintenancesClient().BeginUpdate(ctx, "TestGroup", "testserver", "_T9Q-TS8", &armmysqlflexibleservers.MaintenancesClientBeginUpdateOptions{Parameters: &armmysqlflexibleservers.MaintenanceUpdate{
		Properties: &armmysqlflexibleservers.MaintenancePropertiesForUpdate{
			MaintenanceStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-19T16:00:00.000Z"); return t }()),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Maintenance = armmysqlflexibleservers.Maintenance{
	// 	Name: to.Ptr("_T9Q-TS8"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/maintenances"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/maintenances/_T9Q-TS8"),
	// 	Properties: &armmysqlflexibleservers.MaintenanceProperties{
	// 		MaintenanceAvailableScheduleMaxTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T16:00:00.000Z"); return t}()),
	// 		MaintenanceAvailableScheduleMinTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-04T17:23:45.000Z"); return t}()),
	// 		MaintenanceDescription: to.Ptr("Your Azure Database For MySQL routine maintenance will be performed on the scheduled maintenance window between Sat, 20 Jan 2024 00:00:00 GMT and Sat, 20 Jan 2024 01:00:00 GMT. During the maintenance window, your instance may experience temporary service interruption. Our team will minimize any impact and ensure a smooth transition"),
	// 		MaintenanceEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-19T17:00:00.000Z"); return t}()),
	// 		MaintenanceStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-19T16:00:00.000Z"); return t}()),
	// 		MaintenanceState: to.Ptr(armmysqlflexibleservers.MaintenanceStateReScheduled),
	// 		MaintenanceTitle: to.Ptr("Routine Maintenance: Rescheduled to Sat, 20 Jan 2024 00:00:00 GMT"),
	// 		MaintenanceType: to.Ptr(armmysqlflexibleservers.MaintenanceTypeRoutineMaintenance),
	// 	},
	// }
}
