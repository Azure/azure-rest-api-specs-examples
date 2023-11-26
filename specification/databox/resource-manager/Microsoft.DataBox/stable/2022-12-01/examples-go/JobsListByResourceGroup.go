package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsListByResourceGroup.json
func ExampleJobsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListByResourceGroupPager("YourResourceGroupName", &armdatabox.JobsClientListByResourceGroupOptions{SkipToken: nil})
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
		// page.JobResourceList = armdatabox.JobResourceList{
		// 	Value: []*armdatabox.JobResource{
		// 		{
		// 			Identity: &armdatabox.ResourceIdentity{
		// 				Type: to.Ptr("SystemAssigned"),
		// 				PrincipalID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 				TenantID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			},
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armdatabox.SKU{
		// 				Name: to.Ptr(armdatabox.SKUNameDataBox),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Name: to.Ptr("TestJobName1"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.DataBox/jobs/TestJobName1"),
		// 			Properties: &armdatabox.JobProperties{
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(true),
		// 				IsCancellableWithoutFee: to.Ptr(true),
		// 				IsDeletable: to.Ptr(false),
		// 				IsPrepareToShipEnabled: to.Ptr(true),
		// 				IsShippingAddressEditable: to.Ptr(true),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-03T06:25:54.463Z"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
		// 			},
		// 		},
		// 		{
		// 			Identity: &armdatabox.ResourceIdentity{
		// 				Type: to.Ptr("None"),
		// 			},
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armdatabox.SKU{
		// 				Name: to.Ptr(armdatabox.SKUNameDataBox),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Name: to.Ptr("TestJobName2"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.DataBox/jobs/TestJobName2"),
		// 			Properties: &armdatabox.JobProperties{
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(true),
		// 				IsCancellableWithoutFee: to.Ptr(true),
		// 				IsDeletable: to.Ptr(false),
		// 				IsPrepareToShipEnabled: to.Ptr(true),
		// 				IsShippingAddressEditable: to.Ptr(true),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-07T05:20:36.334Z"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
		// 			},
		// 	}},
		// }
	}
}
