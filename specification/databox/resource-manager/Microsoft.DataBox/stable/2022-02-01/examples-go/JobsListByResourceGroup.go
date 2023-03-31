package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/JobsListByResourceGroup.json
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
	pager := clientFactory.NewJobsClient().NewListByResourceGroupPager("SdkRg5154", &armdatabox.JobsClientListByResourceGroupOptions{SkipToken: nil})
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
		// 				PrincipalID: to.Ptr("fac84c35-5490-4b11-81b9-770053ccbe3b"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armdatabox.SKU{
		// 				Name: to.Ptr(armdatabox.SKUNameDataBox),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Name: to.Ptr("SdkJob5928"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/SdkRg5154/providers/Microsoft.DataBox/jobs/SdkJob5928"),
		// 			Properties: &armdatabox.JobProperties{
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(true),
		// 				IsCancellableWithoutFee: to.Ptr(true),
		// 				IsDeletable: to.Ptr(false),
		// 				IsPrepareToShipEnabled: to.Ptr(true),
		// 				IsShippingAddressEditable: to.Ptr(true),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-03T11:55:54.463792+05:30"); return t}()),
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
		// 			Name: to.Ptr("SdkJob952"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/SdkRg5154/providers/Microsoft.DataBox/jobs/SdkJob952"),
		// 			Properties: &armdatabox.JobProperties{
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(true),
		// 				IsCancellableWithoutFee: to.Ptr(true),
		// 				IsDeletable: to.Ptr(false),
		// 				IsPrepareToShipEnabled: to.Ptr(true),
		// 				IsShippingAddressEditable: to.Ptr(true),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-07T10:50:36.3341513+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
		// 			},
		// 	}},
		// }
	}
}
