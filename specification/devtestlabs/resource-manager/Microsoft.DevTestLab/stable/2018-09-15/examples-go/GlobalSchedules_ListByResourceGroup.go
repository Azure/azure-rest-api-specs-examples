package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_ListByResourceGroup.json
func ExampleGlobalSchedulesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGlobalSchedulesClient().NewListByResourceGroupPager("resourceGroupName", &armdevtestlabs.GlobalSchedulesClientListByResourceGroupOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
	})
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
		// page.ScheduleList = armdevtestlabs.ScheduleList{
		// 	Value: []*armdevtestlabs.Schedule{
		// 		{
		// 			Name: to.Ptr("LabVmAutoStart"),
		// 			Type: to.Ptr("microsoft.devtestlab/labs/schedules"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/schedules/labvmautostart"),
		// 			Location: to.Ptr("{location}"),
		// 			Properties: &armdevtestlabs.ScheduleProperties{
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-29T22:54:54.933Z"); return t}()),
		// 				NotificationSettings: &armdevtestlabs.NotificationSettings{
		// 					Status: to.Ptr(armdevtestlabs.EnableStatusDisabled),
		// 					TimeInMinutes: to.Ptr[int32](0),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Status: to.Ptr(armdevtestlabs.EnableStatusEnabled),
		// 				TaskType: to.Ptr("LabVmsStartupTask"),
		// 				TimeZoneID: to.Ptr("Hawaiian Standard Time"),
		// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 				WeeklyRecurrence: &armdevtestlabs.WeekDetails{
		// 					Time: to.Ptr("0700"),
		// 					Weekdays: []*string{
		// 						to.Ptr("Monday"),
		// 						to.Ptr("Tuesday"),
		// 						to.Ptr("Wednesday"),
		// 						to.Ptr("Thursday"),
		// 						to.Ptr("Friday")},
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
