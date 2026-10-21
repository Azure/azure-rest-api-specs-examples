package armhorizondb_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/MaintenanceEvents_Reschedule.json
func ExampleMaintenanceEventsClient_BeginReschedule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMaintenanceEventsClient().BeginReschedule(ctx, "example-resource-group", "example-cluster", "example-pool-01", "ABCD-123", armhorizondb.MaintenanceEventRescheduleRequest{
		PostponeToDateTimeUTC: to.Ptr(time.Date(2026, time.August, 20, 10, 0, 0, 0, time.UTC)),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.MaintenanceEventsClientRescheduleResponse{
	// 	MaintenanceEventActionResponse: armhorizondb.MaintenanceEventActionResponse{
	// 		MaintenanceEventID: to.Ptr("ABCD-123"),
	// 		PoolResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/pools/example-pool-01"),
	// 		Status: to.Ptr(armhorizondb.MaintenanceEventStatusRescheduled),
	// 		PlannedStartTimeUTC: to.Ptr(time.Date(2026, time.August, 20, 10, 0, 0, 0, time.UTC)),
	// 		PlannedEndTimeUTC: to.Ptr(time.Date(2026, time.August, 20, 11, 0, 0, 0, time.UTC)),
	// 		AppliedNow: to.Ptr(false),
	// 		LastUpdatedTimeUTC: to.Ptr(time.Date(2026, time.August, 7, 12, 0, 0, 0, time.UTC)),
	// 	},
	// }
}
