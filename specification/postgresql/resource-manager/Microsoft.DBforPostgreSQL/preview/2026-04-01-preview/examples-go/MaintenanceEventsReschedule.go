package armpostgresqlflexibleservers_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-04-01-preview/MaintenanceEventsReschedule.json
func ExampleMaintenanceEventsClient_BeginReschedule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMaintenanceEventsClient().BeginReschedule(ctx, "exampleresourcegroup", "exampleserver", "XXXX-111", armpostgresqlflexibleservers.MaintenanceEventRescheduleRequest{
		PostponeToDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-10T10:00:00+00:00"); return t }()),
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
	// res = armpostgresqlflexibleservers.MaintenanceEventsClientRescheduleResponse{
	// 	MaintenanceEventActionResponse: armpostgresqlflexibleservers.MaintenanceEventActionResponse{
	// 		MaintenanceEventID: to.Ptr("XXXX-111"),
	// 		ServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver"),
	// 		Status: to.Ptr(armpostgresqlflexibleservers.MaintenanceEventStatus("Accepted")),
	// 		PlannedStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-10T10:00:00Z"); return t}()),
	// 		PlannedEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-10T10:05:00Z"); return t}()),
	// 		AppliedNow: to.Ptr(false),
	// 		LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-07T12:00:00Z"); return t}()),
	// 	},
	// }
}
