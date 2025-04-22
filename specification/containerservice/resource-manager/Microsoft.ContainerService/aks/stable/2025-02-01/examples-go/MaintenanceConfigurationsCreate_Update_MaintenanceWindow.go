package armcontainerservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/62d827e6c0f471fe0828b517f481f7485a676ffb/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-02-01/examples/MaintenanceConfigurationsCreate_Update_MaintenanceWindow.json
func ExampleMaintenanceConfigurationsClient_CreateOrUpdate_createUpdateMaintenanceConfigurationWithMaintenanceWindow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMaintenanceConfigurationsClient().CreateOrUpdate(ctx, "rg1", "clustername1", "aksManagedAutoUpgradeSchedule", armcontainerservice.MaintenanceConfiguration{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MaintenanceConfiguration = armcontainerservice.MaintenanceConfiguration{
	// 	Name: to.Ptr("aksManagedAutoUpgradeSchedule"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/maintenanceConfigurations/aksManagedAutoUpgradeSchedule"),
	// 	Properties: &armcontainerservice.MaintenanceConfigurationProperties{
	// 		MaintenanceWindow: &armcontainerservice.MaintenanceWindow{
	// 			DurationHours: to.Ptr[int32](10),
	// 			NotAllowedDates: []*armcontainerservice.DateSpan{
	// 				{
	// 					End: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-02-25"); return t}()),
	// 					Start: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-02-18"); return t}()),
	// 				},
	// 				{
	// 					End: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-01-05"); return t}()),
	// 					Start: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-12-23"); return t}()),
	// 			}},
	// 			Schedule: &armcontainerservice.Schedule{
	// 				Weekly: &armcontainerservice.WeeklySchedule{
	// 					DayOfWeek: to.Ptr(armcontainerservice.WeekDayMonday),
	// 					IntervalWeeks: to.Ptr[int32](3),
	// 				},
	// 			},
	// 			StartDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-01-01"); return t}()),
	// 			StartTime: to.Ptr("08:30"),
	// 			UTCOffset: to.Ptr("+05:30"),
	// 		},
	// 	},
	// }
}
