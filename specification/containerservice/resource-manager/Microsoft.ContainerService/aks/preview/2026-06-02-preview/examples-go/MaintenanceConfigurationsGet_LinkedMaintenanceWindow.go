package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-06-02-preview/MaintenanceConfigurationsGet_LinkedMaintenanceWindow.json
func ExampleMaintenanceConfigurationsClient_Get_getALinkedMaintenanceConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMaintenanceConfigurationsClient().Get(ctx, "rg1", "clustername1", "aksManagedAutoUpgradeSchedule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservice.MaintenanceConfigurationsClientGetResponse{
	// 	MaintenanceConfiguration: armcontainerservice.MaintenanceConfiguration{
	// 		Name: to.Ptr("aksManagedAutoUpgradeSchedule"),
	// 		Type: to.Ptr("Microsoft.ContainerService/managedClusters/maintenanceConfigurations"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/maintenanceConfigurations/aksManagedAutoUpgradeSchedule"),
	// 		Properties: &armcontainerservice.MaintenanceConfigurationProperties{
	// 			MaintenanceWindowID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/maintenanceWindows/myMaintenanceWindow"),
	// 			MaintenanceWindow: &armcontainerservice.MaintenanceWindow{
	// 				DurationHours: to.Ptr[int32](8),
	// 				NotAllowedDates: []*armcontainerservice.DateSpan{
	// 					{
	// 						End: to.Ptr(time.Date(2027, time.January, 3, 0, 0, 0, 0, time.UTC)),
	// 						Start: to.Ptr(time.Date(2026, time.December, 23, 0, 0, 0, 0, time.UTC)),
	// 					},
	// 				},
	// 				Schedule: &armcontainerservice.Schedule{
	// 					Weekly: &armcontainerservice.WeeklySchedule{
	// 						DayOfWeek: to.Ptr(armcontainerservice.WeekDaySaturday),
	// 						IntervalWeeks: to.Ptr[int32](1),
	// 					},
	// 				},
	// 				StartTime: to.Ptr("02:00"),
	// 				UTCOffset: to.Ptr("-07:00"),
	// 			},
	// 		},
	// 		SystemData: &armcontainerservice.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2020, time.January, 1, 17, 18, 19, 123456700, time.UTC)),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2020, time.January, 2, 17, 18, 19, 123456700, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
