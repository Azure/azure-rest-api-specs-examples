package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/28c5000054bf2e8112b5543025a519fa60902503/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2024-07-01/examples/MaintenanceConfigurationsList.json
func ExampleMaintenanceConfigurationsClient_NewListByManagedClusterPager_listMaintenanceConfigurationsByManagedCluster() {
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
		// 			Name: to.Ptr("default"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/maintenanceConfigurations/default"),
		// 			Properties: &armcontainerservice.MaintenanceConfigurationProperties{
		// 				NotAllowedTime: []*armcontainerservice.TimeSpan{
		// 					{
		// 						End: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-30T12:00:00.000Z"); return t}()),
		// 						Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-26T03:00:00.000Z"); return t}()),
		// 				}},
		// 				TimeInWeek: []*armcontainerservice.TimeInWeek{
		// 					{
		// 						Day: to.Ptr(armcontainerservice.WeekDayMonday),
		// 						HourSlots: []*int32{
		// 							to.Ptr[int32](1),
		// 							to.Ptr[int32](2)},
		// 					}},
		// 				},
		// 		}},
		// 	}
	}
}
