package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/97f789aeb52adfc1e20c386005839f5276874d7d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-05-02-preview/examples/MaintenanceConfigurationsList_MaintenanceWindow.json
func ExampleMaintenanceConfigurationsClient_NewListByManagedClusterPager_listMaintenanceConfigurationsConfiguredWithMaintenanceWindowByManagedCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMaintenanceConfigurationsClient().NewListByManagedClusterPager("rg1", "clustername1", nil)
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
		// page.MaintenanceConfigurationListResult = armcontainerservice.MaintenanceConfigurationListResult{
		// 	Value: []*armcontainerservice.MaintenanceConfiguration{
		// 		{
		// 			Name: to.Ptr("aksManagedNodeOSUpgradeSchedule"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/maintenanceConfigurations/default"),
		// 			Properties: &armcontainerservice.MaintenanceConfigurationProperties{
		// 				MaintenanceWindow: &armcontainerservice.MaintenanceWindow{
		// 					DurationHours: to.Ptr[int32](10),
		// 					Schedule: &armcontainerservice.Schedule{
		// 						Daily: &armcontainerservice.DailySchedule{
		// 							IntervalDays: to.Ptr[int32](5),
		// 						},
		// 					},
		// 					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-01-01"); return t}()),
		// 					StartTime: to.Ptr("13:30"),
		// 					UTCOffset: to.Ptr("-07:00"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("aksManagedAutoUpgradeSchedule"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/maintenanceConfigurations/default"),
		// 			Properties: &armcontainerservice.MaintenanceConfigurationProperties{
		// 				MaintenanceWindow: &armcontainerservice.MaintenanceWindow{
		// 					DurationHours: to.Ptr[int32](5),
		// 					NotAllowedDates: []*armcontainerservice.DateSpan{
		// 						{
		// 							End: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-02-25"); return t}()),
		// 							Start: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-02-18"); return t}()),
		// 						},
		// 						{
		// 							End: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-01-05"); return t}()),
		// 							Start: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-12-23"); return t}()),
		// 					}},
		// 					Schedule: &armcontainerservice.Schedule{
		// 						AbsoluteMonthly: &armcontainerservice.AbsoluteMonthlySchedule{
		// 							DayOfMonth: to.Ptr[int32](15),
		// 							IntervalMonths: to.Ptr[int32](3),
		// 						},
		// 					},
		// 					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-01-01"); return t}()),
		// 					StartTime: to.Ptr("08:30"),
		// 					UTCOffset: to.Ptr("+00:00"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
