package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getQuotaRequestsHistory.json
func ExampleQuotaRequestStatusClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQuotaRequestStatusClient().NewListPager("3f75fdf7-977e-44ad-990d-99f14f0f299f", "Microsoft.Compute", "eastus", &armreservations.QuotaRequestStatusClientListOptions{Filter: nil,
		Top:       nil,
		Skiptoken: nil,
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
		// page.QuotaRequestDetailsList = armreservations.QuotaRequestDetailsList{
		// 	Value: []*armreservations.QuotaRequestDetails{
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000/2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
		// 			Type: to.Ptr("Microsoft.Capacity/serviceLimitsRequests"),
		// 			ID: to.Ptr("2B5C8515-37D8-4B6A-879B-CD641A2CF605"),
		// 			Properties: &armreservations.QuotaRequestProperties{
		// 				Message: to.Ptr("Request completed"),
		// 				ProvisioningState: to.Ptr(armreservations.QuotaRequestStateSucceeded),
		// 				RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-19T19:23:17.904Z"); return t}()),
		// 				Value: []*armreservations.SubRequest{
		// 					{
		// 						Name: &armreservations.ResourceName{
		// 							LocalizedValue: to.Ptr("Standard HCS Family vCPUs"),
		// 							Value: to.Ptr("standardHCSFamily"),
		// 						},
		// 						Limit: to.Ptr[int32](200),
		// 						Message: to.Ptr("Request completed"),
		// 						ProvisioningState: to.Ptr(armreservations.QuotaRequestStateSucceeded),
		// 						SubRequestID: to.Ptr("AD07450A-DE86-4FD3-859B-107BEF218C4C"),
		// 					},
		// 					{
		// 						Name: &armreservations.ResourceName{
		// 							LocalizedValue: to.Ptr("Standard NC Promo Family vCPUs"),
		// 							Value: to.Ptr("standardNCPromoFamily"),
		// 						},
		// 						Limit: to.Ptr[int32](50),
		// 						Message: to.Ptr("Request completed"),
		// 						ProvisioningState: to.Ptr(armreservations.QuotaRequestStateSucceeded),
		// 						SubRequestID: to.Ptr("AD07450A-DE86-4FD3-859B-107BEF218C4C"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000/7E73A85C-83BB-4DE4-903F-076F1A2B91D6"),
		// 			Type: to.Ptr("Microsoft.Capacity/serviceLimitsRequests"),
		// 			ID: to.Ptr("7E73A85C-83BB-4DE4-903F-076F1A2B91D6"),
		// 			Properties: &armreservations.QuotaRequestProperties{
		// 				Message: to.Ptr("Request completed"),
		// 				ProvisioningState: to.Ptr(armreservations.QuotaRequestStateSucceeded),
		// 				RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-18T19:23:17.904Z"); return t}()),
		// 				Value: []*armreservations.SubRequest{
		// 					{
		// 						Name: &armreservations.ResourceName{
		// 							LocalizedValue: to.Ptr("Standard NVSv3 Family vCPUs"),
		// 							Value: to.Ptr("standardNVSv3Family"),
		// 						},
		// 						Limit: to.Ptr[int32](100),
		// 						Message: to.Ptr("Request completed"),
		// 						ProvisioningState: to.Ptr(armreservations.QuotaRequestStateSucceeded),
		// 						SubRequestID: to.Ptr("AD07450A-DE86-4FD3-859B-107BEF218C4C"),
		// 					},
		// 					{
		// 						Name: &armreservations.ResourceName{
		// 							LocalizedValue: to.Ptr("Standard NV Promo Family vCPUs"),
		// 							Value: to.Ptr("standardNVPromoFamily"),
		// 						},
		// 						Limit: to.Ptr[int32](150),
		// 						Message: to.Ptr("Request completed"),
		// 						ProvisioningState: to.Ptr(armreservations.QuotaRequestStateSucceeded),
		// 						SubRequestID: to.Ptr("AD07450A-DE86-4FD3-859B-107BEF218C4C"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000/5E460077-AB53-4802-8997-A6940E0B7649"),
		// 			Type: to.Ptr("Microsoft.Capacity/serviceLimitsRequests"),
		// 			ID: to.Ptr("5E460077-AB53-4802-8997-A6940E0B7649"),
		// 			Properties: &armreservations.QuotaRequestProperties{
		// 				Message: to.Ptr("Request completed"),
		// 				ProvisioningState: to.Ptr(armreservations.QuotaRequestStateSucceeded),
		// 				RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-17T19:23:17.904Z"); return t}()),
		// 				Value: []*armreservations.SubRequest{
		// 					{
		// 						Name: &armreservations.ResourceName{
		// 							LocalizedValue: to.Ptr("Standard NC Promo Family vCPUs"),
		// 							Value: to.Ptr("standardNCPromoFamily"),
		// 						},
		// 						Limit: to.Ptr[int32](70),
		// 						Message: to.Ptr("Request completed"),
		// 						ProvisioningState: to.Ptr(armreservations.QuotaRequestStateSucceeded),
		// 						SubRequestID: to.Ptr("AD07450A-DE86-4FD3-859B-107BEF218C4C"),
		// 					},
		// 					{
		// 						Name: &armreservations.ResourceName{
		// 							LocalizedValue: to.Ptr("Standard HBS Family vCPUs"),
		// 							Value: to.Ptr("standardHBSFamily"),
		// 						},
		// 						Limit: to.Ptr[int32](52),
		// 						Message: to.Ptr("Request completed"),
		// 						ProvisioningState: to.Ptr(armreservations.QuotaRequestStateSucceeded),
		// 						SubRequestID: to.Ptr("AD07450A-DE86-4FD3-859B-107BEF218C4C"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
