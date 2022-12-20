package armcontainerservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-10-02-preview/examples/MaintenanceConfigurationsCreate_Update_MaintenanceWindow.json
func ExampleMaintenanceConfigurationsClient_CreateOrUpdate_createUpdateMaintenanceConfigurationWithMaintenanceWindow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewMaintenanceConfigurationsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "rg1", "clustername1", "aksManagedAutoUpgradeSchedule", armcontainerservice.MaintenanceConfiguration{
		Properties: &armcontainerservice.MaintenanceConfigurationProperties{
			MaintenanceWindow: &armcontainerservice.MaintenanceWindow{
				DurationHours: to.Ptr[int32](10),
				NotAllowedDates: []*armcontainerservice.DateSpan{
					{
						End:   to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-02-25"); return t }()),
						Start: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-02-18"); return t }()),
					},
					{
						End:   to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-01-05"); return t }()),
						Start: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-12-23"); return t }()),
					}},
				Schedule: &armcontainerservice.Schedule{
					RelativeMonthly: &armcontainerservice.RelativeMonthlySchedule{
						DayOfWeek:      to.Ptr(armcontainerservice.WeekDayMonday),
						IntervalMonths: to.Ptr[int32](3),
						WeekIndex:      to.Ptr(armcontainerservice.TypeFirst),
					},
				},
				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-01-01"); return t }()),
				StartTime: to.Ptr("08:30"),
				UTCOffset: to.Ptr("+05:30"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
