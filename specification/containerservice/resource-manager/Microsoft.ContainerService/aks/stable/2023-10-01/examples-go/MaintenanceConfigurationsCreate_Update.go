package armcontainerservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c1ba9df47907f9012ae14ca4616aed9e5665f9e5/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-10-01/examples/MaintenanceConfigurationsCreate_Update.json
func ExampleMaintenanceConfigurationsClient_CreateOrUpdate_createUpdateMaintenanceConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMaintenanceConfigurationsClient().CreateOrUpdate(ctx, "rg1", "clustername1", "default", armcontainerservice.MaintenanceConfiguration{
		Properties: &armcontainerservice.MaintenanceConfigurationProperties{
			NotAllowedTime: []*armcontainerservice.TimeSpan{
				{
					End:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-30T12:00:00.000Z"); return t }()),
					Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-26T03:00:00.000Z"); return t }()),
				}},
			TimeInWeek: []*armcontainerservice.TimeInWeek{
				{
					Day: to.Ptr(armcontainerservice.WeekDayMonday),
					HourSlots: []*int32{
						to.Ptr[int32](1),
						to.Ptr[int32](2)},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MaintenanceConfiguration = armcontainerservice.MaintenanceConfiguration{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/maintenanceConfigurations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/maintenanceConfigurations/default"),
	// 	Properties: &armcontainerservice.MaintenanceConfigurationProperties{
	// 		NotAllowedTime: []*armcontainerservice.TimeSpan{
	// 			{
	// 				End: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-30T12:00:00.000Z"); return t}()),
	// 				Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-26T03:00:00.000Z"); return t}()),
	// 		}},
	// 		TimeInWeek: []*armcontainerservice.TimeInWeek{
	// 			{
	// 				Day: to.Ptr(armcontainerservice.WeekDayMonday),
	// 				HourSlots: []*int32{
	// 					to.Ptr[int32](1),
	// 					to.Ptr[int32](2)},
	// 			}},
	// 		},
	// 		SystemData: &armcontainerservice.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
	// 		},
	// 	}
}
