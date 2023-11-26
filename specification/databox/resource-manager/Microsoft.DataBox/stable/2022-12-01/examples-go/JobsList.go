package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsList.json
func ExampleJobsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListPager(&armdatabox.JobsClientListOptions{SkipToken: nil})
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
		// 				Type: to.Ptr("None"),
		// 			},
		// 			Location: to.Ptr("australiaeast"),
		// 			SKU: &armdatabox.SKU{
		// 				Name: to.Ptr(armdatabox.SKUNameDataBoxDisk),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Name: to.Ptr("TestJobName1"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/YourSubscriptionId/resourcegroups/TestResourceGroupName/providers/Microsoft.DataBox/jobs/TestJobName1"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-30T09:14:39.774Z"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
		// 			},
		// 		},
		// 		{
		// 			Identity: &armdatabox.ResourceIdentity{
		// 				Type: to.Ptr("None"),
		// 			},
		// 			Location: to.Ptr("australiaeast"),
		// 			SKU: &armdatabox.SKU{
		// 				Name: to.Ptr(armdatabox.SKUNameDataBoxDisk),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Name: to.Ptr("TestJobName2"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/YourSubscriptionId/resourcegroups/TestResourceGroupName/providers/Microsoft.DataBox/jobs/TestJobName2"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-10T12:33:22.257Z"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
		// 			},
		// 	}},
		// }
	}
}
