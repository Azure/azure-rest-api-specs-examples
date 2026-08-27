package armcontainerservice_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-06-01/MaintenanceConfigurationsCreate_Update_MaintenanceWindow.json
func ExampleMaintenanceConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMaintenanceConfigurationsClient().CreateOrUpdate(ctx, "rg1", "clustername1", "aksManagedAutoUpgradeSchedule", armcontainerservice.MaintenanceConfiguration{
		Properties: &armcontainerservice.MaintenanceConfigurationProperties{
			MaintenanceWindow: &armcontainerservice.MaintenanceWindow{
				DurationHours: to.Ptr[int32](10),
				NotAllowedDates: []*armcontainerservice.DateSpan{
					{
						End:   to.Ptr(time.Date(2023, time.February, 25, 0, 0, 0, 0, time.UTC)),
						Start: to.Ptr(time.Date(2023, time.February, 18, 0, 0, 0, 0, time.UTC)),
					},
					{
						End:   to.Ptr(time.Date(2024, time.January, 5, 0, 0, 0, 0, time.UTC)),
						Start: to.Ptr(time.Date(2023, time.December, 23, 0, 0, 0, 0, time.UTC)),
					},
				},
				Schedule: &armcontainerservice.Schedule{
					RelativeMonthly: &armcontainerservice.RelativeMonthlySchedule{
						DayOfWeek:      to.Ptr(armcontainerservice.WeekDayMonday),
						IntervalMonths: to.Ptr[int32](3),
						WeekIndex:      to.Ptr(armcontainerservice.TypeFirst),
					},
				},
				StartDate: to.Ptr(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
				StartTime: to.Ptr("08:30"),
				UTCOffset: to.Ptr("+05:30"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservice.MaintenanceConfigurationsClientCreateOrUpdateResponse{
	// 	MaintenanceConfiguration: armcontainerservice.MaintenanceConfiguration{
	// 		Name: to.Ptr("aksManagedAutoUpgradeSchedule"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/maintenanceConfigurations/aksManagedAutoUpgradeSchedule"),
	// 		Properties: &armcontainerservice.MaintenanceConfigurationProperties{
	// 			MaintenanceWindow: &armcontainerservice.MaintenanceWindow{
	// 				DurationHours: to.Ptr[int32](10),
	// 				NotAllowedDates: []*armcontainerservice.DateSpan{
	// 					{
	// 						End: to.Ptr(time.Date(2023, time.February, 25, 0, 0, 0, 0, time.UTC)),
	// 						Start: to.Ptr(time.Date(2023, time.February, 18, 0, 0, 0, 0, time.UTC)),
	// 					},
	// 					{
	// 						End: to.Ptr(time.Date(2024, time.January, 5, 0, 0, 0, 0, time.UTC)),
	// 						Start: to.Ptr(time.Date(2023, time.December, 23, 0, 0, 0, 0, time.UTC)),
	// 					},
	// 				},
	// 				Schedule: &armcontainerservice.Schedule{
	// 					Weekly: &armcontainerservice.WeeklySchedule{
	// 						DayOfWeek: to.Ptr(armcontainerservice.WeekDayMonday),
	// 						IntervalWeeks: to.Ptr[int32](3),
	// 					},
	// 				},
	// 				StartDate: to.Ptr(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 				StartTime: to.Ptr("08:30"),
	// 				UTCOffset: to.Ptr("+05:30"),
	// 			},
	// 		},
	// 	},
	// }
}
