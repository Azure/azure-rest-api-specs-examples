package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/JobsList.json
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
		// 			Name: to.Ptr("mnaustest"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/ausresgrpmn/providers/Microsoft.DataBox/jobs/mnaustest"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-30T14:44:39.77401+05:30"); return t}()),
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
		// 			Name: to.Ptr("portalcontractAUS"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/ausresgrpmn/providers/Microsoft.DataBox/jobs/portalcontractAUS"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-10T18:03:22.2578238+05:30"); return t}()),
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
		// 			Name: to.Ptr("testBB-diskAU"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/testBB/providers/Microsoft.DataBox/jobs/testBB-diskAU"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-30T15:17:17.3753642+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
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
		// 			Name: to.Ptr("allXML"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat12/providers/Microsoft.DataBox/jobs/allXML"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-04T11:12:15.5972523+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("bothExportAllxml"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat12/providers/Microsoft.DataBox/jobs/bothExportAllxml"),
		// 			Properties: &armdatabox.JobProperties{
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(true),
		// 				IsCancellableWithoutFee: to.Ptr(true),
		// 				IsDeletable: to.Ptr(false),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(true),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-07T15:51:08.4479315+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("exportAll"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat12/providers/Microsoft.DataBox/jobs/exportAll"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-04T19:44:59.2410723+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("importRegressnTest"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat12/providers/Microsoft.DataBox/jobs/importRegressnTest"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-04T18:31:05.3036028+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
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
		// 			Name: to.Ptr("testBotthXMLAndAll"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat12/providers/Microsoft.DataBox/jobs/testBotthXMLAndAll"),
		// 			Properties: &armdatabox.JobProperties{
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(true),
		// 				IsCancellableWithoutFee: to.Ptr(true),
		// 				IsDeletable: to.Ptr(false),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(true),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T01:57:03.8985885+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("testExportAllOrder"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat12/providers/Microsoft.DataBox/jobs/testExportAllOrder"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-05T15:56:52.6983398+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("testPayload"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat12/providers/Microsoft.DataBox/jobs/testPayload"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-04T17:45:25.464989+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("xmlOnlyOrder"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat12/providers/Microsoft.DataBox/jobs/xmlOnlyOrder"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-04T10:43:04.8775864+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("testbadresourcegroup"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat3198inh0-9)inh(il(h)_fyoin)(upf(yLASD0-FJ(hal-DSNWP0JDS0P3-0SJ93/providers/Microsoft.DataBox/jobs/testbadresourcegroup"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-27T15:17:49.0760408+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
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
		// 			Name: to.Ptr("andipodtest4jan"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat/providers/Microsoft.DataBox/jobs/andipodtest4jan"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-04T09:08:49.9928621+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
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
		// 			Name: to.Ptr("dbtest1"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat/providers/Microsoft.DataBox/jobs/dbtest1"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("NoLongerNeeded null"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-28T15:20:48.1058546+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("degautam14-04-Clone"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat/providers/Microsoft.DataBox/jobs/degautam14-04-Clone"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("IncorrectOrder null"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-14T08:51:18.2067875+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
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
		// 			Name: to.Ptr("degautamtestexport"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat/providers/Microsoft.DataBox/jobs/degautamtestexport"),
		// 			Properties: &armdatabox.JobProperties{
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(true),
		// 				IsCancellableWithoutFee: to.Ptr(true),
		// 				IsDeletable: to.Ptr(false),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(true),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-27T18:08:30.6622356+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("DegautamTestExportOrder"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat/providers/Microsoft.DataBox/jobs/DegautamTestExportOrder"),
		// 			Properties: &armdatabox.JobProperties{
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(true),
		// 				IsCancellableWithoutFee: to.Ptr(true),
		// 				IsDeletable: to.Ptr(false),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(true),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-21T20:41:02.3837388+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("degautamtestorder"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat/providers/Microsoft.DataBox/jobs/degautamtestorder"),
		// 			Properties: &armdatabox.JobProperties{
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(true),
		// 				IsCancellableWithoutFee: to.Ptr(true),
		// 				IsDeletable: to.Ptr(false),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(true),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-24T07:56:23.6839464+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("degautamTestOrder14-04"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat/providers/Microsoft.DataBox/jobs/degautamTestOrder14-04"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("IncorrectOrder null"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-14T08:48:21.2260174+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
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
		// 			Name: to.Ptr("export"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat/providers/Microsoft.DataBox/jobs/export"),
		// 			Properties: &armdatabox.JobProperties{
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(true),
		// 				IsCancellableWithoutFee: to.Ptr(true),
		// 				IsDeletable: to.Ptr(false),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(true),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-12T12:19:19.8627264+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("exportTestResource"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.DataBox/jobs/exportTestResource"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-05T19:50:20.9692355+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("sanakTestImportNew"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.DataBox/jobs/sanakTestImportNew"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-10T15:09:45.7080012+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
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
		// 			Name: to.Ptr("sanakTestImportOld"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.DataBox/jobs/sanakTestImportOld"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-10T15:13:02.5724966+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
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
		// 			Name: to.Ptr("sanakTestImportOld1"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.DataBox/jobs/sanakTestImportOld1"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-20T13:39:28.1940929+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
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
		// 			Name: to.Ptr("sanakTestImportOld2"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.DataBox/jobs/sanakTestImportOld2"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-10T15:14:59.6879599+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
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
		// 			Name: to.Ptr("saranyagorder"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat/providers/Microsoft.DataBox/jobs/saranyagorder"),
		// 			Properties: &armdatabox.JobProperties{
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(true),
		// 				IsCancellableWithoutFee: to.Ptr(true),
		// 				IsDeletable: to.Ptr(false),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(true),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-21T23:43:34.513148+05:30"); return t}()),
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
		// 			Name: to.Ptr("testdeepak04-07"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat/providers/Microsoft.DataBox/jobs/testdeepak04-07"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("IncorrectOrder null"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-04T14:01:50.9232807+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
		// 			},
		// 		},
		// 		{
		// 			Identity: &armdatabox.ResourceIdentity{
		// 				Type: to.Ptr("None"),
		// 			},
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armdatabox.SKU{
		// 				Name: to.Ptr(armdatabox.SKUNameDataBoxHeavy),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Name: to.Ptr("testdurga-heavy"),
		// 			Type: to.Ptr("Microsoft.DataBox/jobs"),
		// 			ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/akvenkat/providers/Microsoft.DataBox/jobs/testdurga-heavy"),
		// 			Properties: &armdatabox.JobProperties{
		// 				CancellationReason: to.Ptr("Old job which is still in ordered state cancelled by the service"),
		// 				DeliveryInfo: &armdatabox.JobDeliveryInfo{
		// 					ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 				DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
		// 				IsCancellable: to.Ptr(false),
		// 				IsCancellableWithoutFee: to.Ptr(false),
		// 				IsDeletable: to.Ptr(true),
		// 				IsPrepareToShipEnabled: to.Ptr(false),
		// 				IsShippingAddressEditable: to.Ptr(false),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-02-11T12:33:16.6231232+05:30"); return t}()),
		// 				Status: to.Ptr(armdatabox.StageNameCancelled),
		// 				TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
		// 			},
		// 	}},
		// }
	}
}
