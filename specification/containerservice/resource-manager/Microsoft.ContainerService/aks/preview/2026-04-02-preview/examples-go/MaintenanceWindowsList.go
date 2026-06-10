package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-04-02-preview/MaintenanceWindowsList.json
func ExampleMaintenanceWindowsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMaintenanceWindowsClient().NewListPager("rg-maintenance", nil)
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
		// page = armcontainerservice.MaintenanceWindowsClientListResponse{
		// 	MaintenanceWindowResourceListResult: armcontainerservice.MaintenanceWindowResourceListResult{
		// 		Value: []*armcontainerservice.MaintenanceWindowResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-maintenance/providers/Microsoft.ContainerService/maintenanceWindows/production-weekends"),
		// 				Name: to.Ptr("production-weekends"),
		// 				Type: to.Ptr("Microsoft.ContainerService/maintenanceWindows"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 				},
		// 				Properties: &armcontainerservice.MaintenanceWindowResourceProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservice.ResourceProvisioningStateSucceeded),
		// 					Schedule: &armcontainerservice.Schedule{
		// 						Weekly: &armcontainerservice.WeeklySchedule{
		// 							IntervalWeeks: to.Ptr[int32](1),
		// 							DayOfWeek: to.Ptr(armcontainerservice.WeekDaySaturday),
		// 						},
		// 					},
		// 					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.DateOnly, "2026-04-05"); return t}()),
		// 					StartTime: to.Ptr("02:00"),
		// 					DurationHours: to.Ptr[int32](8),
		// 					UTCOffset: to.Ptr("-07:00"),
		// 					NotAllowedDates: []*armcontainerservice.DateSpan{
		// 					},
		// 				},
		// 				SystemData: &armcontainerservice.SystemData{
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T10:00:00Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T10:00:00Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-maintenance/providers/Microsoft.ContainerService/maintenanceWindows/dev-nightly"),
		// 				Name: to.Ptr("dev-nightly"),
		// 				Type: to.Ptr("Microsoft.ContainerService/maintenanceWindows"),
		// 				Location: to.Ptr("westus2"),
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("development"),
		// 				},
		// 				Properties: &armcontainerservice.MaintenanceWindowResourceProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservice.ResourceProvisioningStateSucceeded),
		// 					Schedule: &armcontainerservice.Schedule{
		// 						Daily: &armcontainerservice.DailySchedule{
		// 							IntervalDays: to.Ptr[int32](1),
		// 						},
		// 					},
		// 					StartTime: to.Ptr("22:00"),
		// 					DurationHours: to.Ptr[int32](6),
		// 					UTCOffset: to.Ptr("+00:00"),
		// 					NotAllowedDates: []*armcontainerservice.DateSpan{
		// 					},
		// 				},
		// 				SystemData: &armcontainerservice.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-02T14:30:00Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-02T14:30:00Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
