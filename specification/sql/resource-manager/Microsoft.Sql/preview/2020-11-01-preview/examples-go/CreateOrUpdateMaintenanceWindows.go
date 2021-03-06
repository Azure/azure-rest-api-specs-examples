package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateMaintenanceWindows.json
func ExampleMaintenanceWindowsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewMaintenanceWindowsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CreateOrUpdate(ctx,
		"Default-SQL-SouthEastAsia",
		"testsvr",
		"testdwdb",
		"current",
		armsql.MaintenanceWindows{
			Properties: &armsql.MaintenanceWindowsProperties{
				TimeRanges: []*armsql.MaintenanceWindowTimeRange{
					{
						DayOfWeek: to.Ptr(armsql.DayOfWeekSaturday),
						Duration:  to.Ptr("PT60M"),
						StartTime: to.Ptr("00:00:00"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
