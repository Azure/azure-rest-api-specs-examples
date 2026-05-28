package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/CreateOrUpdateMaintenanceWindows.json
func ExampleMaintenanceWindowsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMaintenanceWindowsClient().CreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdwdb", "current", armsql.MaintenanceWindows{
		Properties: &armsql.MaintenanceWindowsProperties{
			TimeRanges: []*armsql.MaintenanceWindowTimeRange{
				{
					DayOfWeek: to.Ptr(armsql.DayOfWeekSaturday),
					Duration:  to.Ptr("PT60M"),
					StartTime: to.Ptr("00:00:00"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.MaintenanceWindowsClientCreateOrUpdateResponse{
	// }
}
