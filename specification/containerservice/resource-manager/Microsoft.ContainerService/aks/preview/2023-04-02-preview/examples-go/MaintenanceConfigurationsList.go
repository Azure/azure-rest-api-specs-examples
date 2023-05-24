package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c183bb012de8e9e1d0d2e67a0994748df4747d2c/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-04-02-preview/examples/MaintenanceConfigurationsList.json
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
		// 			ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/maintenanceConfigurations/default"),
		// 			Properties: &armcontainerservice.MaintenanceConfigurationProperties{
		// 				NotAllowedTime: []*armcontainerservice.TimeSpan{
		// 					{
		// 						End: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-30T12:00:00Z"); return t}()),
		// 						Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-26T03:00:00Z"); return t}()),
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
