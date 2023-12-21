package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateMaintenanceWindows.json
func ExampleMaintenanceWindowsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMaintenanceWindowsClient().CreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdwdb", "current", armsql.MaintenanceWindows{
		Properties: &armsql.MaintenanceWindowsProperties{
			TimeRanges: []*armsql.MaintenanceWindowTimeRange{
				{
					DayOfWeek: to.Ptr(armsql.DayOfWeekSaturday),
					Duration:  to.Ptr("PT60M"),
					StartTime: to.Ptr("00:00:00"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
